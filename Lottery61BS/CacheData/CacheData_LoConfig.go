package CacheData

import "github.com/TtMyth123/DengG/Lottery61BS/models"

const key_LoConfig = "LoConfig"

func LoadConfigFormDB() error {
	aLoConfig := &models.LoConfig{}
	aLoConfig.Id = 1
	e := aLoConfig.Read(nil)
	if e != nil {
		return e
	}
	GetICache().SetCache(key_LoConfig, aLoConfig, -1)
	return nil
}

func LoadConfig2DB() error {
	aLoConfig := GetLoConfig()
	e := aLoConfig.Update(nil)
	if e != nil {
		return e
	}
	return nil
}
func GetLoConfig() models.LoConfig {
	aLoConfig := &models.LoConfig{}
	config, e := GetICache().Get(key_LoConfig, aLoConfig)
	if e != nil {
		LoadConfigFormDB()
		config, e = GetICache().Get(key_LoConfig, aLoConfig)

		aLoConfig = config.(*models.LoConfig)
	}
	aLoConfig = config.(*models.LoConfig)
	return *aLoConfig
}

func SetLoConfigAnd2DB(aLoConfig models.LoConfig) error {
	e := GetICache().SetCache(key_LoConfig, &aLoConfig, -1)
	if e != nil {
		return e
	}
	e = LoadConfig2DB()
	return e
}
