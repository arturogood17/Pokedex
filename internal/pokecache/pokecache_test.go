package pokecache

import (
	"fmt"
	"testing"
	"time"
)

func TestAddGet(t *testing.T) {
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
	for i, v := range cases {
		t.Run(fmt.Sprintf("Test case add: %d", i), func(t *testing.T) {
			ch := NewCache(5 * time.Second)
			ch.Add(v.key, v.val)
			val, ok := ch.Get(v.key)
			if !ok {
				t.Errorf("expected to find key")
				return
			}
			if string(val) != string(v.val) {
				t.Errorf("expected to find value")
				return
			}
		})
	}
}

func TestReapLoop(t *testing.T) {
	const baseTime = 5 * time.Millisecond
	const waitTime = baseTime + 5*time.Millisecond
	ch := NewCache(baseTime)
	ch.Add("https://example.com", []byte("testdata"))
	_, ok := ch.Get("https://example.com")
	if !ok {
		t.Errorf("expected to find key")
		return
	}
	time.Sleep(waitTime)

	_, ok = ch.Get("https://example.com")
	if ok {
		t.Errorf("expected to not find key")
		return
	}
}
