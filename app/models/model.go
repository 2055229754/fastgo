package models

import (
	"fmt"

	"github.com/2055229754/fastgo/fast"
)

type Admin struct {
	fast.Model
}

func (admin *Admin) TableName() string {
	return "admin"
}

func (admin *Admin) Show() {
	// res := admin.Ins(admin).Fields("*").WhereSql("id=1")
	res := admin.Ins(admin).Fields("*").WhereMap([][]string{{"id", "1"}, {"name", "sss"}, {"realname", "like", "%adad%"}})
	// res := admin.Ins(admin).Fields("*").Where([]string{"id", "1"})
	fmt.Println("模型：", res.DbInfo.Quetybuilder)
}
