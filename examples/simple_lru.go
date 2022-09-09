package examples

import (
	"fmt"
	"github.com/tarruin/LRUCache"
)

func SimpleLRU() {
	lru := LRUCache.NewStorage(3)
	db := map[string]string{
		"key1": "value1",
		"key2": "value2",
	}
	var (
		inputKey    string
		outputValue string
		ok          bool
	)
	_, _ = fmt.Scanln(&inputKey)
	if outputValue, ok = lru.Get(inputKey); !ok {
		outputValue = db[inputKey]
		if !lru.Add(inputKey, outputValue) {
			panic("error on add key to lru")
		}
	}
}
