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
	err = c.Set("da", "weisd", 300)
	if err != nil {
		t.Fatal(err)
	}

	res := ""
	c.Get("da", &res)

	if res != "weisd" {
		t.Fatal("base put faield")
	}

	t.Log("ok")

	// use tags/namespace
	err = c.Tags([]string{"dd"}).Set("da", "weisd", 300)
	if err != nil {
		t.Fatal(err)
	}
	res = ""
	c.Tags([]string{"dd"}).Get("da", &res)

	if res != "weisd" {
		t.Fatal("tags put faield")
	}

	t.Log("ok")

	err = c.Tags([]string{"aa"}).Set("aa", "aaa", 300)
	if err != nil {
		t.Fatal(err)
	}

	res = ""
	c.Tags([]string{"aa"}).Get("aa", &res)

	if res != "aaa" {
		t.Fatal("not aaa")
	}

	t.Log("ok")

	// flush namespace
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

	// still store in
	res = ""
	c.Tags([]string{"dd"}).Get("da", &res)
	if res != "weisd" {
		t.Fatal("where")
	}

	t.Log("ok")

}
