# PLEASE NOTE, THIS PROJECT IS NO LONGER BEING MAINTAINED

---

# Vika

[Vika](https://vika.cn) Golang SDK 是对维格表 API 的封装。

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
    "github.com/vikadata/vika.go/lib/common"
    vkerror "github.com/vikadata/vika.go/lib/common/error"
    "github.com/vikadata/vika.go/lib/common/profile"
    vika "github.com/vikadata/vika.go/lib/datasheet"
)

func main() {
    credential := common.NewCredential("YOUR_API_TOKEN")
    cpf := profile.NewClientProfile()
    datasheet, _ := vika.NewDatasheet(credential, "datasheetId", cpf)
    // 获取全部的数据
    request := vika.NewDescribeRecordRequest()
    request.Sort = []*vika.Sort{
        {
            Field: common.StringPtr("number_field"),
            Order: common.StringPtr("desc"),
        },
    }
    request.Fields = common.StringPtrs([]string{"number_field"})
    records, err := datasheet.DescribeAllRecords(request)
    if _, ok := err.(*vkerror.SDKError); ok {
       fmt.Printf("An API error has returned: %s", err)
       return
    }
    // 非SDK异常，直接失败。实际代码中可以加入其他的处理。
    if err != nil {
        panic(err)
    }
    // 打印返回的数据
    fmt.Printf("%#v\n", records)
    // 分页获取数据
    page, err := datasheet.DescribeRecords(request)
	if _, ok := err.(*vkerror.SDKError); ok {
       fmt.Printf("An API error has returned: %s", err)
       return
    }
    // 非SDK异常，直接失败。实际代码中可以加入其他的处理。
    if err != nil {
        panic(err)
    }
    // 打印返回的数据
    fmt.Printf("%#v\n", page)
    // 添加记录
    createRequest := vika.NewCreateRecordsRequest()
    createRequest.Records = []*vika.Fields{
        {
            Fields: &vika.Field{
                "number_field": vika.NumberFieldValue(900),
            },
        },
    }
    createRecords, err := datasheet.CreateRecords(createRequest)
    if _, ok := err.(*vkerror.SDKError); ok {
       fmt.Printf("An API error has returned: %s", err)
       return
    }
    // 非SDK异常，直接失败。实际代码中可以加入其他的处理。
    if err != nil {
        panic(err)
    }
    // 打印返回的数据
    fmt.Printf("%#v\n", createRecords)
	// 修改记录
    modifyRequest := vika.NewModifyRecordsRequest()
    modifyRequest.Records = []*vika.BaseRecord{
        {
            Fields: &vika.Field{
                "number_field": vika.NumberFieldValue(1000),
            },
            RecordId: common.StringPtr("recordId"),
        },
    }
    modifyRecords, err := datasheet.ModifyRecords(modifyRequest)
    if _, ok := err.(*vkerror.SDKError); ok {
       fmt.Printf("An API error has returned: %s", err)
       return
    }
    // 非SDK异常，直接失败。实际代码中可以加入其他的处理。
    if err != nil {
        panic(err)
    }
    // 打印返回的数据
    fmt.Printf("%#v\n", modifyRecords)
	// 删除记录
    deleteRequest := vika.NewDeleteRecordsRequest()
    request.RecordIds =	common.StringPtrs([]string{"recordId1", "recordId2"})
    err = datasheet.DeleteRecords(deleteRequest)
    if _, ok := err.(*vkerror.SDKError); ok {
       fmt.Printf("An API error has returned: %s", err)
       return
    }
    // 非SDK异常，直接失败。实际代码中可以加入其他的处理。
    if err != nil {
        panic(err)
    }
    // 上传文件
    cpf.Upload = true
    uploadRequest := vika.NewUploadRequest()
    request.FilePath = "image.png"
    attachment, err := datasheet.UploadFile(request)
    if _, ok := err.(*vkerror.SDKError); ok {
       fmt.Printf("An API error has returned: %s", err)
       return
    }
    // 非SDK异常，直接失败。实际代码中可以加入其他的处理。
    if err != nil {
        panic(err)
    }
    // 打印返回的数据
    fmt.Printf("%#v\n", attachment)
}

```
