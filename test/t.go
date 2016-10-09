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
	v.Use(cache.VodkaCacher(cache.Options{Adapter: "redis", AdapterConfig: `{"Addr":":6379"}`, Section: "test", Interval: 5}))

	v.GET("/cache/put/", func(self vodka.Context) error {
		err := cache.Store(self).Put("name", "vodka", 60)
		if err != nil {
			return err
		}

		return self.String(http.StatusOK, "store okay")
	})

	v.GET("/cache/get/", func(self vodka.Context) error {
		name := cache.Store(self).Get("name")

		return self.String(http.StatusOK, fmt.Sprintf("get name %s", name))
	})

	v.Run(fasthttp.New(":7891"))
}
