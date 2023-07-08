package cache

// Define interface for caches
type Cacher interface {
	Get(key string) (any, bool)
	Set(key string, value any) error
	Delete(key string)
	Clear()
	Count() int
}
