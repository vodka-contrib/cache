# cache

    Middleware cache provides cache management for Vodka v2+. It can use many cache adapters, including memory, file, Redis.


## import
```go
    "github.com/vodka-contrib/cache"
	_ "github.com/vodka-contrib/cache/redis"
```

## Documentation
```
package cache

import (
	"testing"
)

func Test_TagCache(t *testing.T) {

	c, err := New(Options{Adapter: "memory"})
	if err != nil {
		t.Fatal(err)
	}

	// base use
	err = c.Put("da", "vodka", 300)
	if err != nil {
		t.Fatal(err)
	}

	res := c.Get("da")

	if res != "vodka" {
		t.Fatal("base put faield")
	}

	t.Log("ok")

	// use tags/namespace
	err = c.Tags([]string{"dd"}).Put("da", "vodka", 300)
	if err != nil {
		t.Fatal(err)
	}
	res = c.Tags([]string{"dd"}).Get("da")

	if res != "vodka" {
		t.Fatal("tags put faield")
	}

	t.Log("ok")

	err = c.Tags([]string{"aa"}).Put("aa", "aaa", 300)
	if err != nil {
		t.Fatal(err)
	}

	res = c.Tags([]string{"aa"}).Get("aa")

	if res != "aaa" {
		t.Fatal("not aaa")
	}

	t.Log("ok")

	// flush namespace
	err = c.Tags([]string{"aa"}).Flush()
	if err != nil {
		t.Fatal(err)
	}

	res = c.Tags([]string{"aa"}).Get("aa")
	if res != "" {
		t.Fatal("flush faield")
	}

	res = c.Tags([]string{"aa"}).Get("bb")
	if res != "" {
		t.Fatal("flush faield")
	}

	// still store in
	res = c.Tags([]string{"dd"}).Get("da")
	if res != "vodka" {
		t.Fatal("where ")
	}

	t.Log("ok")

}
```


## vodka Middleware
```go
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

```
