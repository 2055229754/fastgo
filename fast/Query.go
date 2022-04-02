package fast

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/2055229754/fastgo/Debug"
	_ "github.com/go-sql-driver/mysql"
)

type DbInfo struct {
	Table    string
	Host     string
	Port     int
	Prefix   string
	Username string
	Password string
	Charset  string
	Db       *sql.DB
	Quetybuilder
	err error
}
type Quetybuilder struct {
	Where  string
	Order  string
	Limit  int
	Offset int
	Fields string
}

var (
	Db  *sql.DB
	err error
)

func NewDb(table string) *DbInfo {
	q := &DbInfo{}
	q.Table = table
	return q
}
func (db *DbInfo) SetFields(fields string) {
	db.Fields = fields
}
func (db *DbInfo) SetWhereSql(sql string) {
	db.BuildWhere(sql)
}
func (db *DbInfo) SetWhereMap(condition [][]string) {
	var sql string
	for _, con := range condition {
		var conlen = len(con)
		if conlen == 2 {
			sql += fmt.Sprintf(" %s='%s' AND", con[0], con[1])
		}
		if conlen == 3 {
			operate := strings.ToUpper(con[1])

			if operate == "IN" {
				sql += fmt.Sprintf(" %s IN (%s) AND", con[0], con[2])
			} else if operate == "LIKE" {
				sql += fmt.Sprintf(" %s LIKE '%s' AND", con[0], con[2])
			} else {
				sql += fmt.Sprintf(" %s %s '%s' AND", con[0], con[1], con[2])
			}
		}
		if conlen != 2 && conlen != 3 {
			Debug.Error("SetWhereMap 格式错误!")
		}
	}
	fmt.Println("原始Sql", sql)
	if sql != "" {
		sql = sql[1 : len(sql)-4]
	}
	fmt.Println("修正Sql", sql)
}
func (db *DbInfo) BuildWhere(sql string) {
	if db.Where == "" {
		db.Where = sql
	} else {
		db.Where += " AND " + sql
	}
}

func (db *DbInfo) StringToWhere(where string) {
	db.Where = where
}
func (db *DbInfo) ArrayToWhere(where map[string]interface{}) {

}

//获取数据库控制器
func GetDb(host string, port int, username string, password string) (*sql.DB, error) {
	//DSN (Data Source Name)数据源连接格式:[username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]
	//这里我们可以不选择数据库,或者增加可选参数,比如timeout(建立连接超时时间)
	//mysqlConnStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/mysql?&charset=utf8&parseTime=True&loc=Local&timeout=5s", username, password, host, port)
	dsn := fmt.Sprintf("%s:%s@(%s:%d)/ddjz?timeout=5s", username, password, host, port)
	Db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Printf("配置连接出错:%s\n", err.Error())
		return Db, err
	}
	// 设置连接池中空闲连接的最大数量。
	Db.SetMaxIdleConns(1)
	// 设置打开数据库连接的最大数量。
	Db.SetMaxOpenConns(1)
	// 设置连接可复用的最大时间。
	Db.SetConnMaxLifetime(time.Second * 30)
	//设置连接最大空闲时间
	Db.SetConnMaxIdleTime(time.Second * 30)

	//检查连通性
	err = Db.Ping()
	if err != nil {
		log.Printf("数据库连接出错:%s\n", err.Error())
		return Db, err
	}

	return Db, err
}

//单行数据解析 查询数据库，解析查询结果，支持动态行数解析
func QueryAndParse(Db *sql.DB, queryStr string) map[string]string {
	rows, err := Db.Query(queryStr)
	defer rows.Close()

	if err != nil {
		log.Printf("查询出错,SQL语句:%s\n错误详情:%s\n", queryStr, err.Error())
		return nil
	}

	//获取列名cols
	cols, _ := rows.Columns()
	if len(cols) > 0 {
		buff := make([]interface{}, len(cols))       // 创建临时切片buff
		data := make([][]byte, len(cols))            // 创建存储数据的字节切片2维数组data
		dataKv := make(map[string]string, len(cols)) //创建dataKv, 键值对的map对象
		for i, _ := range buff {
			buff[i] = &data[i] //将字节切片地址赋值给临时切片,这样data才是真正存放数据
		}

		for rows.Next() {
			rows.Scan(buff...) // ...是必须的,表示切片
		}

		for k, col := range data {
			dataKv[cols[k]] = string(col)
			//fmt.Printf("%30s:\t%s\n", cols[k], col)
		}
		return dataKv
	} else {
		return nil
	}
}

//多行数据解析
func QueryAndParseRows(Db *sql.DB, queryStr string) []map[string]string {
	rows, err := Db.Query(queryStr)
	defer rows.Close()
	if err != nil {
		fmt.Printf("查询出错:\nSQL:\n%s, 错误详情:%s\n", queryStr, err.Error())
		return nil
	}
	//获取列名cols
	cols, _ := rows.Columns()
	if len(cols) > 0 {
		var ret []map[string]string
		for rows.Next() {
			buff := make([]interface{}, len(cols))
			data := make([][]byte, len(cols)) //数据库中的NULL值可以扫描到字节中
			for i, _ := range buff {
				buff[i] = &data[i]
			}
			rows.Scan(buff...) //扫描到buff接口中，实际是字符串类型data中
			//将每一行数据存放到数组中
			dataKv := make(map[string]string, len(cols))
			for k, col := range data { //k是index，col是对应的值
				//fmt.Printf("%30s:\t%s\n", cols[k], col)
				dataKv[cols[k]] = string(col)
			}
			ret = append(ret, dataKv)
		}
		return ret
	} else {
		return nil
	}
}

//任意可序列化数据转为Json,便于查看
func Data2Json(anyData interface{}) string {
	JsonByte, err := json.Marshal(anyData)
	if err != nil {
		log.Printf("数据序列化为json出错:\n%s\n", err.Error())
	}
	return string(JsonByte)
}
