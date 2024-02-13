package cache

import (
	"github.com/rosedblabs/rosedb/v2"
	"videoteque/fs"
)

var cache *rosedb.DB

func init() {
	options := rosedb.DefaultOptions
	options.DirPath = fs.CacheDir()

	db, err := rosedb.Open(options)
	if err != nil {
		panic(err)
	}

	cache = db
}

func Read(key string) string {
	val, _ := cache.Get([]byte(key))

	return string(val)
}

func Write(key string, val []byte) {
	cache.Put([]byte(key), val)
}

func Delete(key string) {
	cache.Delete([]byte(key))
}
