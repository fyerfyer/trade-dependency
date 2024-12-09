package cache

type Cache interface {
	Set(key string, data interface{}, expire int) error
	Exists(key string) bool
	Get(key string) ([]byte, error)
	Delete(key string) (bool, error)
}
