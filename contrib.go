package cache

import (
	"github.com/insionng/vodka"
)

const VodkaCacheStoreKey = "VodkaCacheStore"

func Store(value interface{}) Cache {
	var cacher Cache
	switch v := value.(type) {
	case vodka.Context:
		cacher = v.Get(VodkaCacheStoreKey).(Cache)
		if cacher == nil {
			panic("VodkaStore not found, forget to Use Middleware ?")
		}
	default:

		panic("unknown Context")
	}

	if cacher == nil {
		panic("cache context not found")
	}

	return cacher
}

func VodkaCacher(opt Options) vodka.MiddlewareFunc {
	return func(next vodka.HandlerFunc) vodka.HandlerFunc {
		return func(self vodka.Context) error {
			tagcache, err := New(opt)
			if err != nil {
				return err
			}

			self.Set(VodkaCacheStoreKey, tagcache)

			if err = next(self); err != nil {
				return err
			}

			return nil
		}
	}
}
