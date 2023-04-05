package dao

import (
	"fmt"
	"fwwb/library"
)

type Danger struct {
	ID      int    `json:"id" form:"id"`                   // id
	Ctime   int    `json:"ctime" form:"ctime"`             // 创建时间
	Address string `json:"address" form:"address"`         // 地址
	Desc    string `json:"description" form:"description"` // 描述
	Imgs    string `json:"imgs" form:"imgs"`               // 图片
	Status  int    `json:"status" from:"status"`           // 状态 0 未处理 1 已处理
}

func InsertDanger(d Danger) error {
	_, err := library.MysqlDB.Exec("insert INTO dangers(ctime,address,desc,imgs,status) values(?,?,?,?,?)",
		d.Ctime, d.Address, d.Desc, d.Imgs, d.Status)
	if err != nil {
		return err
	}
	return nil
}

func SelectDangers() ([]Danger, error) {
	res := make([]Danger, 0)
	rows, err := library.MysqlDB.Query("select * from dangers")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		d := Danger{}
		err = rows.Scan(&d.ID, &d.Ctime, &d.Address, &d.Desc, &d.Imgs, &d.Status)
		if err != nil {
			fmt.Printf("Scan failed,err:%v\n", err)
			return nil, err
		}
		res = append(res, d)
	}
	return res, nil
}
