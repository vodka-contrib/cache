package cache

import (
	"github.com/insionng/vodka"
)

const EchoCacheStoreKey = "EchoCacheStore"

func Store(value interface{}) Cache {
	var cacher Cache
	switch v := value.(type) {
	case *vodka.Context:
		cacher = v.Get(EchoCacheStoreKey).(Cache)
		if cacher == nil {
			panic("EchoStore not found, forget to Use Middleware ?")
		}
	default:

		panic("unknown Context")
	}

	if cacher == nil {
		panic("cache context not found")
	}

	return cacher
}

func EchoCacher(opt Options) vodka.MiddlewareFunc {
	return func(h vodka.HandlerFunc) vodka.HandlerFunc {
		return func(c *vodka.Context) error {
			tagcache, err := New(opt)
			if err != nil {
				return err
			}

			c.Set(EchoCacheStoreKey, tagcache)

			if err = h(c); err != nil {
				return err
			}

			return nil
		}
	}
}
