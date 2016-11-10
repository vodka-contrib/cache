package cache

import (
	"github.com/insionng/vodka"
)

const VodkaCacheStoreKey = "VodkaCacheStore"

func Store(value interface{}) Cache {
	var cacher Cache
	var okay bool
	switch v := value.(type) {
	case vodka.Context:
		if cacher, okay = v.Get(VodkaCacheStoreKey).(Cache); !okay {
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

func Cacher(opt Options) vodka.MiddlewareFunc {
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
