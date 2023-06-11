package redis

import (
	"errors"
	"time"

	"github.com/go-redis/redis"
)

const (
	oneWeekInSeconds = 7 * 24 * 3600
	scorePerVote     = 432
)

var (
	ErrVoteTimeExpire = errors.New("投票时间已过")
	ErrVoteRepeated   = errors.New("不能重复投票")
)

func CreatePost(postID int64) error {
	pipeline := client.TxPipeline()
	// 创建帖子创建时间的key
	pipeline.ZAdd(getRedisKey(KeyPostTimeZset), redis.Z{
		Score: float64(time.Now().Unix()),
		Member: postID,
	})
	// 创建帖子分数的key
	pipeline.ZAdd(getRedisKey(KeyPostScoreZset), redis.Z{
		Score: float64(time.Now().Unix()),
		Member: postID,
	})
	_, err := pipeline.Exec()
	return err
}

func VoteForPost(userID, postID string, direction float64) error {
	// 获取帖子创建时间，规定了帖子创建一周后不能投票
	postCreateTime := client.ZScore(getRedisKey(KeyPostTimeZset), postID).Val()
	// 判断帖子是否超出了投票时间
	if float64(time.Now().Unix())-postCreateTime > oneWeekInSeconds {
		return ErrVoteTimeExpire
	}
	// 拿到原来用户给帖子的投票数据
	oldVote := client.ZScore(getRedisKey(KeyPostVotedZset+postID), userID).Val()
	if oldVote == direction {
		return ErrVoteRepeated
	}
	// 计算分数过后给帖子修改分数
	pipeline := client.TxPipeline()
	pipeline.ZIncrBy(getRedisKey(KeyPostScoreZset), (direction-oldVote) * scorePerVote, postID)
	// 如果用户取消投票则删除redis中对应的key
	if direction == 0 {
		pipeline.ZRem(getRedisKey(KeyPostVotedZset + postID), userID)
	} else {
		// 否则更新原有的投票数据
		pipeline.ZAdd(getRedisKey(KeyPostVotedZset + postID), redis.Z{
			Score: direction,
			Member: userID,
		})
	}
	// 操作应该放入pipeline中执行
	_, err := pipeline.Exec()
	return err
}
