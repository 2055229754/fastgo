package main

import (
	"fmt"
	"log"

	_ "github.com/2055229754/fastgo/app"
	"github.com/2055229754/fastgo/fast"
)

func main() {
	// fa := fast.New()
	// for _, v := range fa.Route.Routes {
	// 	fmt.Println("route", v)
	// }
	// fa.Start()
	//获取数据库控制器DB
	Db, err := fast.GetDb("127.0.0.1", 3306, "root", "root")
	if err != nil {
		log.Printf("获取数据库控制器出错:%s\n", err.Error())
	}
	defer Db.Close() //延迟关闭数据库控制器,释放数据库连接
	//单行数据查询
	showMasterStatus := fast.QueryAndParse(Db, "select id from dd_admin limit 1")
	fmt.Println(showMasterStatus["id"])
	for k, v := range showMasterStatus {
		fmt.Println(k, "=>", v)

	}
	log.Printf("单行数据-数据库状态:%v\n", fast.Data2Json(showMasterStatus))
	log.Printf("单行数据-数据库状态-File:%v\n", showMasterStatus["File"])

	//多行数据查询
	showProcessList := fast.QueryAndParseRows(Db, "show processlist")
	log.Printf("多行数据-进程信息:%v\n", fast.Data2Json(showProcessList))
	log.Printf("多行数据-进程信息-Host:%v\n", showProcessList[0]["Host"])
}
