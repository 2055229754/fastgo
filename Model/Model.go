package Model

import (
	"reflect"

	"github.com/2055229754/fastgo/Debug"
	"github.com/2055229754/fastgo/fast"
	_ "github.com/go-sql-driver/mysql"
)

type Model struct {
	DbInfo *fast.DbInfo
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
		var mysqlinfo = fast.FastApp.Config.Mysql
		if mysqlinfo.Prefix != "" {
			table = mysqlinfo.Prefix + table
		}
		M.DbInfo = fast.NewDb(table)
		M.Table = table
		return M
	}
	Debug.Error("请设置模型数据表")
	return M
}

//设置fields
func (M Model) Insert(data map[string]string) (interface{}, error) {
	res, err := M.DbInfo.Insert(data)
	if err != nil {
		return 0, err
	}
	return res, nil
}

//设置fields
func (M Model) Delete() (interface{}, error) {
	res, err := M.DbInfo.Delete()
	if err != nil {
		return 0, err
	}
	return res, nil
}

func (M Model) Update(data map[string]string) (interface{}, error) {
	result, err := M.DbInfo.Update(data)
	return result, err
}

//设置fields
func (M Model) Fields(fields string) Model {
	M.DbInfo.SetFields(fields)
	return M
}
func (M Model) Where(condition string) Model {
	M.DbInfo.SetWhereSql(condition)
	return M
}
func (M Model) WhereMap(condition [][]string) Model {
	M.DbInfo.SetWhereMap(condition)
	return M
}

func (M Model) OrWhere(condition [][]string) Model {
	M.DbInfo.SetOrWhere(condition)
	return M
}

func (M Model) Limit(num int) Model {
	M.DbInfo.SetLimit(num)
	return M
}

func (M Model) Offset(num int) Model {
	M.DbInfo.SetOffset(num)
	return M
}
func (M Model) Order(order string) Model {
	M.DbInfo.SetOrder(order)
	return M
}
func (M Model) FetchSql(flag bool) Model {
	M.DbInfo.SetFetch(flag)
	return M
}
func (M Model) Find() interface{} {
	result := M.DbInfo.QueryAndParseOne()
	return result
}
func (M Model) Select() interface{} {
	result := M.DbInfo.QueryAndParseRows()
	return result
}
