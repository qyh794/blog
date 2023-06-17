package redis

import (
	"blog/models"
	"encoding/json"
	"github.com/go-redis/redis"
	"time"
)

const (
	oneWeekExpiration = 7 * 24 * time.Hour
)

func GetPostIDsInOrder(p *models.ParamPostList) ([]string, error) {
	key := getRedisKey(KeyPostTimeZset)
	if p.Order == models.OrderScore {
		key = getRedisKey(KeyPostScoreZset)
	}
	return getIDsFromKey(p.Page, p.Size, key)
}

func getIDsFromKey(page int64, size int64, key string) ([]string, error) {
	start := (page - 1) * size
	end := start + size - 1
	return client.ZRange(key, start, end).Result()
}

func GetPostVoteData(ids []string) (data []int64, err error) {
	pipeline := client.Pipeline()
	for _, id := range ids {
		key := getRedisKey(KeyPostVotedZset + id)
		pipeline.ZCount(key, "1", "1")
	}
	exec, err := pipeline.Exec()
	if err != nil {
		return nil, err
	}
	data = make([]int64, 0, len(exec))
	for _, cmder := range exec {
		val := cmder.(*redis.IntCmd).Val()
		data = append(data, val)
	}
	return
}

func GetPostInOrderFromCache() (data []string, err error) {
	key := getRedisKey(KeyPostCacheZset)
	result, err := client.ZRange(key, 0, -1).Result()
	if err != nil {
		return nil, err
	}
	return result, nil
}

func AddPostCache(data []*models.PostDetail) error {
	// 存储博客信息的zSet缓存
	key := getRedisKey(KeyPostCacheZset)
	postCache := make([]redis.Z, len(data))
	for i, post := range data {
		score := float64(post.VoteNum)
		member, err := json.Marshal(post)
		if err != nil {
			return err
		}
		postCache[i] = redis.Z{
			Score:  score,
			Member: member,
		}
	}
	_, err := client.ZAdd(key, postCache...).Result()
	if err != nil {
		return err
	}
	// 设置string类型key判断缓存是否过期
	isTimeoutKey := getRedisKey(KeyPostCacheIsTimeOutString)
	err = client.Set(isTimeoutKey, "0", oneWeekExpiration).Err()
	if err != nil {
		return err
	}
	return nil
}

func IsCacheExpiration() (string, error) {
	key := getRedisKey(KeyPostCacheIsTimeOutString)
	result, err := client.Get(key).Result()
	if err != nil {
		// 其他错误
		return "", err
	} else if err == redis.Nil {
		// key过期
		return "", nil
	}
	return result, nil
}
