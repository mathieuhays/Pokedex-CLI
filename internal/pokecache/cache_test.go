package pokecache

import (
	"fmt"
	"testing"
	"time"
)

func TestAddGet(t *testing.T) {
	const interval = 5 * time.Second
	cases := []struct {
		key string
		val []byte
	}{
		{
			key: "https://example.com",
			val: []byte("testdata"),
		},
		{
			key: "https://example.com/path",
			val: []byte("moretestdata"),
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			cache := NewCache(interval, interval)
			cache.Add(c.key, c.val)
			val, ok := cache.Get(c.key)
			if !ok {
				t.Errorf("expected to find key")
				return
			}
			if string(val) != string(c.val) {
				t.Errorf("expected to find value")
				return
			}
		})
	}

	t.Run("expired cache not returned even without realloop execution", func(t *testing.T) {
		expireTime := time.Millisecond * 5
		purgeInterval := time.Second * 10
		cache := NewCache(expireTime, purgeInterval)
		testKey := "test"
		cache.Add(testKey, []byte("test data"))

		_, ok := cache.Get(testKey)
		if !ok {
			t.Errorf("expected to find key")
			return
		}

		time.Sleep(expireTime * 2)

		_, ok = cache.Get(testKey)
		if ok {
			t.Errorf("expected to not find key")
		}
	})
}

func TestRealLoop(t *testing.T) {
	const baseTime = 5 * time.Millisecond
	const waitTime = baseTime + 5*time.Millisecond
	const testKey = "https://example.com"
	cache := NewCache(baseTime, baseTime)
	cache.Add(testKey, []byte("testdata"))

	_, ok := cache.Get(testKey)
	if !ok {
		t.Errorf("expected to find key")
		return
	}

	time.Sleep(waitTime)

	_, ok = cache.Get(testKey)
	if ok {
		t.Errorf("expected to not find key")
		return
	}
}
