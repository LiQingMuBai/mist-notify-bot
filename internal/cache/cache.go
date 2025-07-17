package cache

import (
	"context"
	"time"
)

// Cache 定义缓存接口
type Cache interface {
	// Get 获取缓存值
	Get(key string) (string, error)

	// Set 设置缓存值
	Set(key string, value string, expiration time.Duration) error

	// Delete 删除缓存值
	Delete(key string) error

	// Exists 检查key是否存在
	Exists(key string) (bool, error)

	// Clear 清空所有缓存
	Clear(ctx context.Context) error

	// Close 关闭缓存连接
	Close() error
}
