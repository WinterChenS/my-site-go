package initialize

import (
	"github.com/coocood/freecache"
	"winterchen.com/my-site-go/src/global"
)

func InitCache() {
	global.Cache = freecache.NewCache(100 * 1024 * 1024)
}
