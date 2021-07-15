package test

import (
	"fmt"
	"github.com/vikadata/vika.go/lib/common"
	vkerror "github.com/vikadata/vika.go/lib/common/error"
	"github.com/vikadata/vika.go/lib/common/profile"
	"github.com/vikadata/vika.go/lib/common/util"
	vika "github.com/vikadata/vika.go/lib/datasheet"
	"os"
	"testing"
)

func TestCreateRecords(t *testing.T) {
	// VIKA_HOST 可以不用设置，默认使用生产的host
	credential := common.NewCredential(os.Getenv("VIKA_TOKEN"))
	cpf := profile.NewClientProfile()
	datasheet, _ := vika.NewDatasheet(credential, os.Getenv("VIKA_DATASHEET_ID"), cpf)
	request := vika.NewCreateRecordsRequest()
	request.Records = []*vika.Fields{
		{
			Fields: &vika.Field{
				os.Getenv("VIKA_FIELD"): vika.NumberFieldValue(900),
			},
		},
	}
	records, err := datasheet.CreateRecords(request)
	if _, ok := err.(*vkerror.VikaSDKError); ok {
		t.Errorf("An API error has returned: %s", err)
	}
	// 非SDK异常，直接失败。实际代码中可以加入其他的处理。
	if err != nil {
		t.Errorf("An unexcepted error has returned: %s", err)
		panic(err)
	}
	t.Log(len(records))
}

func TestDescribeAllRecords(t *testing.T) {
	// VIKA_HOST 可以不用设置，默认使用生产的host
	credential := common.NewCredential(os.Getenv("VIKA_TOKEN"))
	cpf := profile.NewClientProfile()
	datasheet, _ := vika.NewDatasheet(credential, os.Getenv("VIKA_DATASHEET_ID"), cpf)
	request := vika.NewDescribeRecordRequest()
	request.Sort = []*vika.Sort{
		{
			Field: common.StringPtr(os.Getenv("VIKA_FIELD")),
			Order: common.StringPtr("desc"),
		},
	}
	request.Fields = common.StringPtrs([]string{os.Getenv("VIKA_FIELD")})
	records, err := datasheet.DescribeAllRecords(request)
	if _, ok := err.(*vkerror.VikaSDKError); ok {
		t.Errorf("An API error has returned: %s", err)
	}
	// 非SDK异常，直接失败。实际代码中可以加入其他的处理。
	if err != nil {
		t.Errorf("An unexcepted error has returned: %s", err)
		panic(err)
	}
	t.Log(len(records))
}

func TestDescribeRecords(t *testing.T) {
	// VIKA_HOST 可以不用设置，默认使用生产的host
	credential := common.NewCredential(os.Getenv("VIKA_TOKEN"))
	cpf := profile.NewClientProfile()
	datasheet, _ := vika.NewDatasheet(credential, os.Getenv("VIKA_DATASHEET_ID"), cpf)
	request := vika.NewDescribeRecordRequest()
	request.Sort = []*vika.Sort{
		{
			Field: common.StringPtr(os.Getenv("VIKA_FIELD")),
			Order: common.StringPtr("desc"),
		},
	}
	request.Fields = common.StringPtrs([]string{os.Getenv("VIKA_FIELD")})
	records, err := datasheet.DescribeRecords(request)
	if _, ok := err.(*vkerror.VikaSDKError); ok {
		t.Errorf("An API error has returned: %s", err)
	}
	// 非SDK异常，直接失败。实际代码中可以加入其他的处理。
	if err != nil {
		t.Errorf("An unexcepted error has returned: %s", err)
		panic(err)
	}
	t.Log(len(records.Records))
}

func TestModifyRecords(t *testing.T) {
	// VIKA_HOST 可以不用设置，默认使用生产的host
	credential := common.NewCredential(os.Getenv("VIKA_TOKEN"))
	cpf := profile.NewClientProfile()
	datasheet, _ := vika.NewDatasheet(credential, os.Getenv("VIKA_DATASHEET_ID"), cpf)
	describeRequest := vika.NewDescribeRecordRequest()
	describeRequest.FilterByFormula = common.StringPtr("{" + os.Getenv("VIKA_FIELD") + "}=900")
	record, _ := datasheet.DescribeRecord(describeRequest)
	request := vika.NewModifyRecordsRequest()
	request.Records = []*vika.BaseRecord{
		{
			Fields: &vika.Field{
				os.Getenv("VIKA_FIELD"): vika.NumberFieldValue(1000),
			},
			RecordId: record.RecordId,
		},
	}
	records, err := datasheet.ModifyRecords(request)
	if _, ok := err.(*vkerror.VikaSDKError); ok {
		t.Errorf("An API error has returned: %s", err)
	}
	// 非SDK异常，直接失败。实际代码中可以加入其他的处理。
	if err != nil {
		t.Errorf("An unexcepted error has returned: %s", err)
		panic(err)
	}
	t.Log(len(records))
}

func TestDeleteRecords(t *testing.T) {
	// VIKA_HOST 可以不用设置，默认使用生产的host
	credential := common.NewCredential(os.Getenv("VIKA_TOKEN"))
	cpf := profile.NewClientProfile()
	datasheet, _ := vika.NewDatasheet(credential, os.Getenv("VIKA_DATASHEET_ID"), cpf)
	describeRequest := vika.NewDescribeRecordRequest()
	describeRequest.FilterByFormula = common.StringPtr("{" + os.Getenv("VIKA_FIELD") + "}=1000")
	record, _ := datasheet.DescribeRecord(describeRequest)
	request := vika.NewDeleteRecordsRequest()
	request.RecordIds = []*string{record.RecordId}
	err := datasheet.DeleteRecords(request)
	if _, ok := err.(*vkerror.VikaSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	// 非SDK异常，直接失败。实际代码中可以加入其他的处理。
	if err != nil {
		t.Errorf("An unexcepted error has returned: %s", err)
		panic(err)
	}
}

func TestUpload(t *testing.T) {
	// VIKA_HOST 可以不用设置，默认使用生产的host
	credential := common.NewCredential(os.Getenv("VIKA_TOKEN"))
	cpf := profile.NewClientProfile()
	cpf.Upload = true
	datasheet, _ := vika.NewDatasheet(credential, os.Getenv("VIKA_DATASHEET_ID"), cpf)
	request := vika.NewUploadRequest()
	request.FilePath = "image.png"
	attachment, err := datasheet.UploadFile(request)
	if _, ok := err.(*vkerror.VikaSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	// 非SDK异常，直接失败。实际代码中可以加入其他的处理。
	if err != nil {
		t.Errorf("An unexcepted error has returned: %s", err)
		panic(err)
	}
	t.Log(attachment)
}

func TestDescribeFields(t *testing.T) {
	// VIKA_HOST 可以不用设置，默认使用生产的host
	credential := common.NewCredential(os.Getenv("VIKA_TOKEN"))
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Domain = os.Getenv("VIKA_HOST")
	datasheet, _ := vika.NewDatasheet(credential, os.Getenv("VIKA_DATASHEET_ID"), cpf)
	describeRequest := vika.NewDescribeFieldsRequest()
	describeRequest.ViewId = common.StringPtr(os.Getenv("VIKA_DATASHEET_VIEW_ID"))
	fields, err := datasheet.DescribeFields(describeRequest)
	if _, ok := err.(*vkerror.VikaSDKError); ok {
		t.Errorf("An API error has returned: %s", err)
	}
	// 非SDK异常，直接失败。实际代码中可以加入其他的处理。
	if err != nil {
		t.Errorf("An unexcepted error has returned: %s", err)
		panic(err)
	}
	for _, value := range fields {
		property := value.SelectFieldProperty()
		util.Dd(property)
	}
	t.Log(len(fields))
}
