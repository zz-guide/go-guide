package sql

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"path/filepath"
	"regexp"
	"src/asset"
	"src/config"
	"time"
)

//常量配置
var (
	defaultDriver   = "mysql"
	defaultDatabase string
	sqlBuf          bytes.Buffer
)

const EXTRA = "?charset=utf8"
const DOC_NAME = "doc"

//全局变量
var DB *sql.DB
var err error

type H map[string]interface{}

type DataInfo struct {
	Name            string
	Engine          string
	Version         string
	Row_format      string
	Rows            string
	Avg_row_length  string
	Data_length     string
	Max_data_length string
	Index_length    string
	Data_free       string
	Auto_increment  string
	Create_time     string
	Update_time     string
	Check_time      string
	Collation       string
	Checksum        string
	Create_options  string
	Comment         string
	Create_sql      string
	Columns         []map[string]string
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func InitDb() {
	defaultDatabase = config.Database
	DB, _ = sql.Open(defaultDriver, config.Username+":"+config.Password+"@tcp("+config.Host+":"+config.Port+")/"+defaultDatabase+EXTRA)
	checkErr(err)
}

/**
用来动态映射数据库字段和结构体
*/
func TransFormColumn(str string) []map[string]string {
	rows, err := DB.Query(str)
	checkErr(err)

	defer rows.Close()
	columns, _ := rows.Columns()

	values := make([]sql.RawBytes, len(columns))

	scans := make([]interface{}, len(columns))

	for i := range values {
		scans[i] = &values[i]
	}

	var result []map[string]string

	for rows.Next() {
		//这个地方传引用以后，values也会受影响
		_ = rows.Scan(scans...)
		each := make(map[string]string)

		for i, col := range values {
			each[columns[i]] = string(col)
		}

		result = append(result, each)
	}

	return result
}

/**
将字符串str中的find替换为replace
*/
func StrReplace(str string, pattern string, replace string) string {
	re := regexp.MustCompile(pattern)
	return re.ReplaceAllString(str, replace)
}

func BuildDatabaseData() []DataInfo {
	//查询每张表的信息
	originData := TransFormColumn("SHOW TABLE STATUS from " + defaultDatabase)

	slice := make([]DataInfo, 0)

	for _, value := range originData {

		tableInfo := new(DataInfo)

		//返回给前端的json数据
		tableInfo.Name = value["Name"]
		tableInfo.Engine = value["Engine"]
		tableInfo.Version = value["Version"]
		tableInfo.Row_format = value["Row_format"]
		tableInfo.Rows = value["Rows"]
		tableInfo.Avg_row_length = value["Avg_row_length"]
		tableInfo.Data_length = value["Data_length"]
		tableInfo.Max_data_length = value["Max_data_length"]
		tableInfo.Index_length = value["Index_length"]
		tableInfo.Data_free = value["Data_free"]
		tableInfo.Auto_increment = value["Auto_increment"]
		tableInfo.Create_time = value["Create_time"]
		tableInfo.Update_time = value["Update_time"]
		tableInfo.Check_time = value["Check_time"]
		tableInfo.Collation = value["Collation"]
		tableInfo.Checksum = value["Checksum"]
		tableInfo.Create_options = value["Create_options"]
		tableInfo.Comment = value["Comment"]

		//查询表的全字段信息
		tableInfo.Columns = TransFormColumn(`SHOW FULL FIELDS FROM ` + value["Name"])

		//查询建表语句
		ct := TransFormColumn("SHOW CREATE TABLE " + value["Name"])
		var createTableSql string
		for _, value1 := range ct {
			createTableSql = value1["Create Table"]
		}

		//正则
		tableInfo.Create_sql = StrReplace(createTableSql, `AUTO_INCREMENT=\d+ `, "") + ";"

		sqlBuf.WriteString(tableInfo.Create_sql + "\n\n")

		slice = append(slice, *tableInfo)

		fmt.Println("Build table " + value["Name"] + " ... SUCCESS!!! ")
	}

	return slice
}

//默认将db.sql生成到当前目录
func GetSqlPath() string {
	abs, _ := filepath.Abs("./")
	return abs
}

func CreateDbJSonFile(d string) {
	f, err := os.Create("./" + DOC_NAME + "/db.json")
	checkErr(err)
	defer f.Close()
	f.WriteString(d)
}

func CreateDbSqlFile(d string) {
	f, err := os.Create("./" + DOC_NAME + "/db.sql")
	checkErr(err)
	defer f.Close()
	f.WriteString(d)
}

func toJson() ([]byte, error) {
	return json.Marshal(
		H{
			"errcode": 200,
			"errmsg":  "success",
			"errtime": time.Now().Unix(),
			"data":    BuildDatabaseData(),
		})
}

//生成json文件
func StartServer() []byte {
	InitDb()
	b, _ := toJson()

	if err != nil {
		fmt.Println("json 格式化失败:", err)
	}

	return b
}

func StartClient() {
	InitDb()
	b, _ := toJson()

	if err != nil {
		fmt.Println("json 格式化失败:", err)
	}

	CreateTemplate()
	CreateDbJSonFile(string(b[:]))
	CreateDbSqlFile(sqlBuf.String())

	fmt.Println("数据库文档生成成功！")
}

/*func CopyTemplate(){
	cmd := exec.Command("cp", "-rp", "./static", "./" + DOC_NAME)

	err := cmd.Run()
	if err != nil {
		fmt.Println("Execute Command failed:" + err.Error())
		return
	}

	fmt.Println("Execute Command finished.")
}*/

func CreateTemplate() {
	path := "./" + DOC_NAME

	os.RemoveAll(path)

	err := os.MkdirAll(path, 0777)
	if err != nil {
		fmt.Printf("%s", err)
	}

	by, _ := asset.Asset("static/axios.min.js")
	by1, _ := asset.Asset("static/vue.min.js")
	by2, _ := asset.Asset("static/index.html")
	by3, _ := asset.Asset("static/index.css")

	//生成相应的模板文件
	generatorFile("./"+DOC_NAME+"/axios.min.js", string(by))
	generatorFile("./"+DOC_NAME+"/vue.min.js", string(by1))
	generatorFile("./"+DOC_NAME+"/index.html", string(by2))
	generatorFile("./"+DOC_NAME+"/index.css", string(by3))
}

func generatorFile(filePath string, d string) {
	f, err := os.Create(filePath)
	checkErr(err)
	defer f.Close()
	f.WriteString(d)
}
