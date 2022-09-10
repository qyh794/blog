package redis

const (
	KeyPrefix        = "blog:"
	KeyPostTimeZset  = "post:time"
	KeyPostScoreZset = "post:score"
	KeyPostVotedZset = "post:voted:"
)

func getRedisKey(key string) string {
	return KeyPrefix + key
}