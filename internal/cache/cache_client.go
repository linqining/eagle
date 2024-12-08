package cache

import (
	"context"

	"github.com/linqining/eagle/internal/model"
	"github.com/linqining/eagle/pkg/cache"
	"github.com/linqining/eagle/pkg/encoding"
	"github.com/linqining/eagle/pkg/redis"
)

func getCacheClient(ctx context.Context) cache.Cache {
	jsonEncoding := encoding.JSONEncoding{}
	cachePrefix := ""
	client := cache.NewRedisCache(redis.RedisClient, cachePrefix, jsonEncoding, func() interface{} {
		return &model.UserBaseModel{}
	})

	return client
}
