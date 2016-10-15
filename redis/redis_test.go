package redis

import (
	"testing"

	"github.com/vodka-contrib/cache"
)

func TestRedisCache(t *testing.T) {
	var err error
	c, err := cache.New(cache.Options{Adapter: "redis", AdapterConfig: `{"Addr":":6379"}`, Section: "test"})
	if err != nil {
		t.Fatal(err)
	}

	err = c.Set("da", "weisd", 300)
	if err != nil {
		t.Fatal(err)
	}

	res := ""
	err = c.Get("da", &res)
	if err != nil {
		t.Fatal(err)
	}

	if res != "weisd" {
		t.Fatal(res)
	}

	t.Log("ok")
	t.Log("test", res)

	err = c.Tags([]string{"dd"}).Set("da", "weisd", 300)
	if err != nil {
		t.Fatal(err)
	}
	res = ""
	err = c.Tags([]string{"dd"}).Get("da", &res)
	if err != nil {
		t.Fatal(err)
	}

	if res != "weisd" {
		t.Fatal("not weisd")
	}

	t.Log("ok")
	t.Log("dd", res)

	err = c.Tags([]string{"aa"}).Set("aa", "aaa", 300)
	if err != nil {
		t.Fatal(err)
	}

	err = c.Tags([]string{"aa"}).Set("bb", "bbb", 300)
	if err != nil {
		t.Fatal(err)
	}

	res = ""
	err = c.Tags([]string{"aa"}).Get("aa", &res)
	if err != nil {
		t.Fatal(err)
	}

	if res != "aaa" {
		t.Fatal("not aaa")
	}

	t.Log("ok")
	t.Log("aa", res)

	err = c.Tags([]string{"aa"}).Flush()
	if err != nil {
		t.Fatal(err)
	}

	res = ""
	c.Tags([]string{"aa"}).Get("aa", &res)
	if res != "" {
		t.Fatal("flush faield")
	}

	res = ""
	c.Tags([]string{"aa"}).Get("bb", &res)
	if res != "" {
		t.Fatal("flush faield")
	}

	res = ""
	err = c.Tags([]string{"dd"}).Get("da", &res)
	if err != nil {
		t.Fatal(err)
	}

	if res != "weisd" {
		t.Fatal("not weisd")
	}

	t.Log("ok")

	err = c.Flush()
	if err != nil {
		t.Fatal(err)
	}

	res = ""
	c.Get("da", &res)
	if res != "" {
		t.Fatal("flush failed")
	}

	t.Log("get dd da", res)

}
