package main

import (
	"fmt"
	"github.com/tencentyun/scf-go-lib/cloudfunction"
	"github.com/vikadata/vika.go/lib/common"
	"github.com/vikadata/vika.go/lib/common/profile"
	vika "github.com/vikadata/vika.go/lib/datasheet"
	"math"
	"os"
)

func getRecords(credential *common.Credential, cpf *profile.ClientProfile) []*vika.Record {
	datasheet, _ := vika.NewDatasheet(credential, os.Getenv("DATASHEET_ID"), cpf)
	request := vika.NewDescribeRecordRequest()
	request.Sort = []*vika.Sort{
		{
			Field: common.StringPtr(os.Getenv("SORT_FIELD_NAME")),
			Order: common.StringPtr("asc"),
		},
	}
	request.Fields = common.StringPtrs([]string{os.Getenv("SORT_FIELD_NAME")})
	request.PageNum = common.Int64Ptr(1)
	request.ViewId = common.StringPtr(os.Getenv("VIEW_ID"))
	request.PageSize = common.Int64Ptr(100)
	records, err := datasheet.DescribeRecords(request)
	if err != nil {
		fmt.Println("获取记录失败")
		panic(err)
	}
	return records.Records
}

func hello() (string, error) {
	// 必要步骤：
	// 实例化一个认证对象，入参需要传入vika开发者token。
	// 这里采用的是从环境变量读取的方式，需要在环境变量中先设置这个值。
	// 你也可以直接在代码中写死token，但是小心不要将代码复制、上传或者分享给他人，
	// 以免泄露token危及你的财产安全。
	credential := common.NewCredential(
		os.Getenv("VIKA_TOKEN"),
	)
	fmt.Println("开始执行")
	cpf := profile.NewClientProfile()
	records := getRecords(credential, cpf)
	var recordIds []*string
	for i := range records {
		recordIds = append(recordIds, records[i].RecordId)
	}
	fmt.Println("获取到的记录条数：", len(recordIds))
	return deleteRecords(credential, cpf, recordIds)
}

func deleteRecords(credential *common.Credential, cpf *profile.ClientProfile, recordIds []*string) (string, error) {
	datasheet, _ := vika.NewDatasheet(credential, os.Getenv("DATASHEET_ID"), cpf)
	request := vika.NewDeleteRecordsRequest()
	for i := 0; i < int(math.Ceil(float64(len(recordIds)/10))); i++ {
		request.RecordIds = recordIds[i*10 : (i+1)*10]
		err := datasheet.DeleteRecords(request)
		if err != nil {
			fmt.Println("删除失败", i)
		}
		fmt.Println("删除成功", i)
	}
	return "", nil
}
func main() {
	// Make the handler available for Remote Procedure Call by Cloud Function
	cloudfunction.Start(hello)
}
