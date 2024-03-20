package models

import (
	"fmt"
	"github.com/TtMyth123/DengG/Lottery61BS/models/mconst"
	"github.com/TtMyth123/kit/ttLog"
	"github.com/astaxie/beego/orm"
)

type LoNumInfo struct {
	BaseInfo
	Num           int
	Color         string `orm:"size(20)"`
	ColorTitle    string `orm:"size(20)"`
	ZodiacCn      string `orm:"size(20)"`
	ZodiacCnTitle string `orm:"size(20)"`
	DS            string `orm:"size(20)"`
	DSTitle       string `orm:"size(20)"`
}

const (
	DS_Single = "Single"
	DS_Double = "Double"
)

const (
	DS_T_Single = "单"
	DS_T_Double = "双"
)
const (
	Color_red   = "red"
	Color_blue  = "blue"
	Color_green = "green"
)
const (
	Color_T_red   = "红"
	Color_T_blue  = "蓝"
	Color_T_green = "绿"
)

const (
	ZodiacCn_Shu  = "Shu"
	ZodiacCn_Niu  = "Niu"
	ZodiacCn_Hu   = "Hu"
	ZodiacCn_Tu   = "Tu"
	ZodiacCn_Long = "Long"
	ZodiacCn_She  = "She"
	ZodiacCn_Ma   = "Ma"
	ZodiacCn_Yang = "Yang"
	ZodiacCn_Hou  = "Hou"
	ZodiacCn_Ji   = "Ji"
	ZodiacCn_Gou  = "Gou"
	ZodiacCn_Zhu  = "Zhu"
)

const (
	ZodiacCn_T_Shu  = "鼠"
	ZodiacCn_T_Niu  = "牛"
	ZodiacCn_T_Hu   = "虎"
	ZodiacCn_T_Tu   = "兔"
	ZodiacCn_T_Long = "龙"
	ZodiacCn_T_She  = "蛇"
	ZodiacCn_T_Ma   = "马"
	ZodiacCn_T_Yang = "羊"
	ZodiacCn_T_Hou  = "猴"
	ZodiacCn_T_Ji   = "鸡"
	ZodiacCn_T_Gou  = "狗"
	ZodiacCn_T_Zhu  = "猪"
)

func (a *LoNumInfo) TableName() string {
	return mconst.TableName_LoNumInfo
}
func setMp(mp map[int]string, strV string, nums ...int) {
	for _, num := range nums {
		mp[num] = strV
	}
}
func getNum2ZodiacCnMp() (map[int]string, map[int]string) {
	mpT := make(map[int]string)
	mp := make(map[int]string)
	setMp(mp, ZodiacCn_Shu, 5, 17, 29, 41)
	setMp(mp, ZodiacCn_Niu, 4, 16, 28, 40)
	setMp(mp, ZodiacCn_Hu, 3, 15, 27, 39)
	setMp(mp, ZodiacCn_Tu, 2, 14, 26, 38)
	setMp(mp, ZodiacCn_Long, 1, 13, 25, 37, 49)
	setMp(mp, ZodiacCn_She, 12, 24, 36, 48)
	setMp(mp, ZodiacCn_Ma, 11, 23, 35, 47)
	setMp(mp, ZodiacCn_Yang, 10, 22, 34, 46)
	setMp(mp, ZodiacCn_Hou, 9, 21, 33, 45)
	setMp(mp, ZodiacCn_Ji, 8, 20, 32, 44)
	setMp(mp, ZodiacCn_Gou, 7, 19, 31, 43)
	setMp(mp, ZodiacCn_Zhu, 6, 18, 30, 42)

	setMp(mpT, ZodiacCn_T_Shu, 5, 17, 29, 41)
	setMp(mpT, ZodiacCn_T_Niu, 4, 16, 28, 40)
	setMp(mpT, ZodiacCn_T_Hu, 3, 15, 27, 39)
	setMp(mpT, ZodiacCn_T_Tu, 2, 14, 26, 38)
	setMp(mpT, ZodiacCn_T_Long, 1, 13, 26, 37, 49)
	setMp(mpT, ZodiacCn_T_She, 12, 24, 36, 48)
	setMp(mpT, ZodiacCn_T_Ma, 11, 23, 35, 47)
	setMp(mpT, ZodiacCn_T_Yang, 10, 22, 34, 46)
	setMp(mpT, ZodiacCn_T_Hou, 9, 21, 33, 45)
	setMp(mpT, ZodiacCn_T_Ji, 8, 20, 32, 44)
	setMp(mpT, ZodiacCn_T_Gou, 7, 19, 31, 43)
	setMp(mpT, ZodiacCn_T_Zhu, 6, 18, 30, 42)

	return mp, mpT
}
func getNum2ColorMp() (map[int]string, map[int]string) {
	mpT := make(map[int]string)
	mp := make(map[int]string)

	setMp(mp, Color_red, 1, 2, 7, 8, 12, 13, 18, 19, 23, 24, 29, 30, 34, 35, 40, 45, 46)
	setMp(mp, Color_blue, 3, 4, 9, 10, 14, 15, 20, 25, 26, 31, 36, 37, 41, 42, 47, 48)
	setMp(mp, Color_green, 5, 6, 11, 16, 17, 21, 22, 27, 28, 32, 33, 38, 39, 43, 44, 49)

	setMp(mpT, Color_T_red, 1, 2, 7, 8, 12, 13, 18, 19, 23, 24, 29, 30, 34, 35, 40, 45, 46)
	setMp(mpT, Color_T_blue, 3, 4, 9, 10, 14, 15, 20, 25, 26, 31, 36, 37, 41, 42, 47, 48)
	setMp(mpT, Color_T_green, 5, 6, 11, 16, 17, 21, 22, 27, 28, 32, 33, 38, 39, 43, 44, 49)

	return nil, nil
}

func InitLoNumInfo(o orm.Ormer) {
	c, _ := o.QueryTable(mconst.TableName_LoNumInfo).Count()

	if c == 0 {
		mpColor, mpColorT := getNum2ColorMp()
		mpZodiacCn, mpZodiacCnT := getNum2ZodiacCnMp()
		arrData := make([]LoNumInfo, 0)
		for i := int64(1); i <= 49; i++ {
			ii := int(i)
			DS := DS_Single
			DSTitle := DS_T_Single
			if i%2 == 0 {
				DS = DS_Double
				DSTitle = DS_T_Double
			}
			Color := mpColor[ii]
			ColorTitle := mpColorT[ii]
			ZodiacCn := mpZodiacCn[ii]
			ZodiacCnTitle := mpZodiacCnT[ii]
			arrData = append(arrData, LoNumInfo{BaseInfo: BaseInfo{Id: i}, Num: int(i),
				Color: Color, ColorTitle: ColorTitle,
				DS: DS, DSTitle: DSTitle,
				ZodiacCn: ZodiacCn, ZodiacCnTitle: ZodiacCnTitle})
		}

		_, e := o.InsertMulti(len(arrData), &arrData)
		if e != nil {
			ttLog.LogError("站点类型 出错 e:", e)
		}
	}
}

type SNum struct {
	Id    int64
	Color string
}

func GetAllLoNumInfos() []SNum {
	o := orm.NewOrm()

	arrData := make([]SNum, 0)
	sql := fmt.Sprintf(`select * from %s`, mconst.TableName_LoNumInfo)
	_, e := o.Raw(sql).QueryRows(&arrData)
	if e != nil {
		fmt.Println(e)
	}
	return arrData
}
