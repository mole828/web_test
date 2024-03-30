package cache

func Default() WebCache {
	return New_WebCache_RedisImpl("redis-master-service")
}
