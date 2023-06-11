package redis

import (
	"blog/models"
	"github.com/go-redis/redis"
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

func UpdateCache(data []*models.PostDetail) error {

}
