# Vika

[Vika](https://vika.cn) Golang SDK 是对维格表 Fusion API 的官方封装。

## 快速开始
### 环境要求

go 1.15 +  

### 安装

```shell
go get github.com/vikadata/vika.go
```



## 获取 API TOKEN

访问维格表的工作台，点击左下角的个人头像，进入「用户中心 > 开发者配置」。点击生成Token(首次使用需要绑定邮箱)。

### 使用
```go
package main

import (
	"fmt"
	"github.com/vikadata/vika.go"
	"reflect"
)

func main() {
	datasheet := vika.New(vika.VikaConfig{Token: "YOUR_API_TOKEN"}).Datasheet("datasheetId")
	// 获取全部的数据
	sort := []vika.Sort{{Field: "field1", Order: "desc"}, {Field: "field2", Order: "asc"}}

	v := reflect.ValueOf(sort)
	num := v.Len()
	querySort := make([]interface{}, num)
	for i := 0; i < num; i++ {
		querySort[i] = v.Index(i).Interface()
	}
	result, err := datasheet.All(vika.RecordQueryParam{FieldKey: vika.FieldKeyName, Sort: querySort})
	fmt.Printf("Body: %#v\n error: %s", result, err)
	// 分页获取数据
	page, pageErr := datasheet.Get(vika.RecordQueryParam{FieldKey: vika.FieldKeyId})
	fmt.Printf("Body: %#v\n error: %s", page, pageErr)
	// 上传文件
	upload, uploadErr := datasheet.Upload("image.png")
	fmt.Printf("Body: %#v\n error: %s", upload, uploadErr)
	// 添加记录
	var addForm = []vika.NewRecord{
		{
			Fields: map[string]interface{}{
				"field1": "value",
			},
		},
		{
			Fields: map[string]interface{}{
				"field1": "value",
			},
		},
	}
	add, addErr := datasheet.Add(addForm, vika.FieldKeyName)
	fmt.Printf("Body: %#v\n error: %s", add, addErr)
	// 修改记录
	var updateForm = []vika.Record{
		{
			Fields: map[string]interface{}{
				"field1": "value",
			},
			RecordId: add.Data.Records[0].RecordId,
		},
		{
			Fields: map[string]interface{}{
				"field1": "value",
			},
			RecordId: add.Data.Records[1].RecordId,
		},
	}
	update, updateErr := datasheet.Update(updateForm, vika.FieldKeyName)
	fmt.Printf("Body: %#v\n error: %s", update, updateErr)
	// 删除记录
	deleteArr := []string{add.Data.Records[0].RecordId, add.Data.Records[1].RecordId}
	deleteRes, deleteErr := datasheet.Del(deleteArr)
	fmt.Printf("Body: %#v\n error: %s", deleteRes, deleteErr)
}

```