package CacheData

import (
	"fmt"
	"github.com/TtMyth123/DengG/Lottery61BS/models"
	"github.com/TtMyth123/kit"
	"sync"
)

var userId2UUIDMap sync.Map

func init() {

}

func GetTokenKey(strToken string) string {
	return fmt.Sprintf("token_%s", strToken)
}

func Get401TokenKey(strToken string) string {
	return fmt.Sprintf("token_old_%s", strToken)
}

func Get401Ip(strToken string) string {
	key401Token := Get401TokenKey(strToken)
	strIp := ""
	GetICache().GetCache(key401Token, &strIp)
	return strIp
}

func DelToken(strToken string) error {
	aSimpleUser, _ := GetToken(strToken)
	key := GetTokenKey(strToken)
	GetICache().DelCache(key)
	userId2UUIDMap.Delete(aSimpleUser.Id)

	return nil
}

func SetToken(aSimpleUser models.SysUser) (string, error) {
	strUUID := kit.GetGuid()
	key := GetTokenKey(strUUID)
	e := GetICache().SetCache(key, aSimpleUser, 0)
	if e != nil {
		return "", e
	}
	if vOldUUID, ok := userId2UUIDMap.Load(aSimpleUser.Id); ok {
		if strOldUUID, ok := vOldUUID.(string); ok {
			if oldSimpleUser, e := GetToken(strOldUUID); e == nil {
				if oldSimpleUser.LoginIp != aSimpleUser.LoginIp || true {
					key401Token := Get401TokenKey(strOldUUID)
					GetICache().SetCache(key401Token, aSimpleUser.LoginIp, 60)
				}
			}

			DelToken(strOldUUID)
		}
	}
	userId2UUIDMap.Store(aSimpleUser.Id, strUUID)
	return strUUID, e
}

func GetToken(strToken string) (models.SysUser, error) {
	SimpleUser := models.SysUser{}
	key := GetTokenKey(strToken)
	e := GetICache().GetCache(key, &SimpleUser)
	if e != nil {
		return SimpleUser, e
	}

	return SimpleUser, nil
}
