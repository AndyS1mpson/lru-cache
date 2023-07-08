package cache

// Interface for caches
type Cacher interface {
	Get(key string) (any, bool)
	Set(key string, value any) error
	Delete(key string)
	Clear()
	Count() int
}
