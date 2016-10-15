package main

import (
	"fmt"
	"net/http"

	"github.com/insionng/vodka"
	"github.com/insionng/vodka/engine/fasthttp"
	"github.com/vodka-contrib/cache"
	_ "github.com/vodka-contrib/cache/redis"
)

func main() {

	v := vodka.New()
	v.Use(cache.Cacher(cache.Options{Adapter: "redis", AdapterConfig: `{"Addr":":6379"}`, Section: "test", Interval: 5}))

	v.GET("/cache/put/", func(self vodka.Context) error {
		err := cache.Store(self).Set("name", "vodka", 60)
		if err != nil {
			return err
		}

		return self.String(http.StatusOK, "store okay")
	})

	v.GET("/cache/get/", func(self vodka.Context) error {
		var name string
		cache.Store(self).Get("name", &name)

		return self.String(http.StatusOK, fmt.Sprintf("get name %s", name))
	})

	v.Run(fasthttp.New(":7891"))
}
