package main

import (
	"encoding/json"
	"fmt"

	. "github.com/admpub/mgodb"
	"gopkg.in/mgo.v2/bson"
)

func main() {
	var table = `mag_list`
	c := &Config{
		User: `root`,
		Pass: `root`,
		Host: `localhost:27017`,
		Name: `admin`,
	}
	Setup(c.Dsn(), c.Name)

	// ======================
	// 获取单行数据
	// ======================
	result := ResultSet{}
	FindAll(table, bson.M{}, result)
	b, err := json.MarshalIndent(result, ``, `  `)

	// 更新数据
	//result["article_url"] = "http://abc"
	//Update(c.Table, result)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(b))

	// ======================
	// 获取多行数据
	// ======================
	results := []ResultSet{}
	FindAll(table, bson.M{}, &results)
	b, err = json.MarshalIndent(results, ``, `  `)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(b))
}
