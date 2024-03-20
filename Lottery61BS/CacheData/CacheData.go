package CacheData

import (
	"github.com/TtMyth123/DengG/Lottery61BS/CacheData/TtMapC"
	"github.com/TtMyth123/kit/CacheDataKit"
)

var mCache CacheDataKit.ICache

func Init() {
	CacheDataKit.Register("TtMapCache", TtMapC.NewTtMapCache)
	mpConfig := make(map[string]interface{})

	var err error
	mCache, err = CacheDataKit.NewCache("TtMapCache", mpConfig)
	if err != nil {
		print(err)
	}
}

func GetICache() CacheDataKit.ICache {
	return mCache
}
