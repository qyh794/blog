package redis

const (
	KeyPrefix                   = "blog:"
	KeyPostTimeZset             = "post:time"
	KeyPostScoreZset            = "post:score"
	KeyPostVotedZset            = "post:voted:"
	KeyPostCacheZset            = "post:cache"
	KeyPostCacheIsTimeOutString = "post:cacheIsTimeout"
)

func getRedisKey(key string) string {
	return KeyPrefix + key
}
