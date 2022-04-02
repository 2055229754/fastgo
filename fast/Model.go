package fast

import (
	"database/sql"
	"reflect"

	"github.com/2055229754/fastgo/Debug"
	_ "github.com/go-sql-driver/mysql"
)

type Model struct {
	DbInfo *DbInfo
	Db     *sql.DB
	err    error
	Table  string
}

func (M Model) Ins(model interface{}) Model {
	flectVal := reflect.ValueOf(model)
	method := flectVal.MethodByName("TableName")
	if method.IsValid() {
		res := method.Call(nil)
		table := res[0].String()
		if table == "" {
			Debug.Error("请设置模型数据表")
		}
		var mysqlinfo = FastApp.Config.Mysql
		if mysqlinfo.Prefix != "" {
			table = mysqlinfo.Prefix + table
		}
		M.DbInfo = NewDb(table)
		M.Table = table
		return M
	}
	Debug.Error("请设置模型数据表")
	return M
}

//设置fields
func (M Model) Fields(fields string) Model {
	M.DbInfo.SetFields(fields)
	return M
}
func (M Model) WhereSql(condition string) Model {
	M.DbInfo.SetWhereSql(condition)
	return M
}
func (M Model) WhereMap(condition [][]string) Model {
	M.DbInfo.SetWhereMap(condition)
	return M
}
