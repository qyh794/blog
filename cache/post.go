package cache

import (
	"blog/dao/redis"
	"blog/models"
	"blog/pkg/convert"
	"go.uber.org/zap"
)

func GetPostListFromCache() (data []*models.PostDetail, err error) {
	// 判断缓存是否过期
	expiration, err := redis.IsCacheExpiration()
	// 过期
	if expiration == "" {
		return nil, err
	}
	// 未过期，查询缓存
	cache, err := redis.GetPostInOrderFromCache()
	// 缓存命中
	if len(cache) > 0 {
		zap.L().Info("缓存命中")
		return convert.ConvertToPostDetailList(cache), nil
	}
	zap.L().Info("缓存未命中")
	// 缓存未命中
	return nil, err
}
