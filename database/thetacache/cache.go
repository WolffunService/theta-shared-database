package thetacache

import (
	"crypto"
	"fmt"
	"github.com/WolffunGame/theta-shared-common/thetalog"
	"github.com/WolffunGame/theta-shared-database/database/thetacache/cachestore"
	"github.com/dgraph-io/ristretto"
	"reflect"
	"time"
)

func NewCacheService() *Cache {
	cacheService.Name = "Cache"

	ristrettoCache, err := ristretto.NewCache(&ristretto.Config{
		NumCounters: 1e7,
		MaxCost:     1 << 30,
		BufferItems: 64,
		Metrics:     false,
	})

	ristrettoStore := cachestore.NewRistretto(ristrettoCache, nil)
	cacheService.store = ristrettoStore
	cacheService.initErr = err

	if cacheService.initErr != nil {
		thetalog.Err(cacheService.initErr).Msg("cannot init theta cache")
	}

	return &cacheService
}

var cacheService Cache

// Cache represents the configuration needed by a cache
type Cache struct {
	Name    string
	store   StoreInterface
	initErr error
}

// Get returns the object stored in cache if it exists
func (c *Cache) Get(key interface{}) (interface{}, error) {
	cacheKey := c.getCacheKey(key)
	return c.store.Get(cacheKey)
}

// GetTo Get the object stored in cache then set to destination object
//func (c *Cache) GetTo(key interface{}, dest interface{}) (interface{}, error) {
//	cacheKey := c.getCacheKey(key)
//	cached, err := c.store.Get(cacheKey)
//	if err == nil {
//
//	}
//	return dest, err
//}

// GetWithTTL returns the object stored in cache and its corresponding TTL
func (c *Cache) GetWithTTL(key interface{}) (interface{}, time.Duration, error) {
	cacheKey := c.getCacheKey(key)
	return c.store.GetWithTTL(cacheKey)
}

// Set populates the cache item using the given key
func (c *Cache) Set(key, object interface{}, options *cachestore.Options) error {
	cacheKey := c.getCacheKey(key)
	return c.store.Set(cacheKey, object, options)
}

// SetSimple Set populates the cache item using the given key
func (c *Cache) SetSimple(key, object interface{}, tags ...string) error {
	cacheKey := c.getCacheKey(key)

	return c.store.Set(cacheKey, object, &cachestore.Options{
		Cost:       1,
		Expiration: 24 * time.Hour,
		Tags:       tags,
	})
}

// Delete removes the cache item using the given key
func (c *Cache) Delete(key interface{}) error {
	cacheKey := c.getCacheKey(key)
	return c.store.Delete(cacheKey)
}

// Invalidate invalidates cache item from given options
func (c *Cache) Invalidate(options cachestore.InvalidateOptions) error {
	return c.store.Invalidate(options)
}

// Clear resets all cache data
func (c *Cache) Clear() error {
	return c.store.Clear()
}

// GetType returns the cache type
func (c *Cache) GetType() string {
	return "cache"
}

// getCacheKey returns the cache key for the given key object by returning
// the key if type is string or by computing a checksum of key structure
// if its type is other than string
func (c *Cache) getCacheKey(key interface{}) string {
	switch key.(type) {
	case string:
		return key.(string)
	default:
		return checksum(key)
	}
}

// checksum hashes a given object into a string
func checksum(object interface{}) string {
	digester := crypto.MD5.New()
	_, _ = fmt.Fprint(digester, reflect.TypeOf(object))
	_, _ = fmt.Fprint(digester, object)
	hash := digester.Sum(nil)

	return fmt.Sprintf("%x", hash)
}

// StoreInterface is the interface for all available cache stores
// Ai cần thêm store ví dụ redis các thứ thì tự implement interface này
type StoreInterface interface {
	Get(key interface{}) (interface{}, error)
	GetWithTTL(key interface{}) (interface{}, time.Duration, error)
	Set(key interface{}, value interface{}, options *cachestore.Options) error
	Delete(key interface{}) error
	Invalidate(options cachestore.InvalidateOptions) error
	Clear() error
	GetType() string
}
