package test

import (
	"fmt"
	"github.com/vikadata/vika.go/lib/common"
	vkerror "github.com/vikadata/vika.go/lib/common/error"
	"github.com/vikadata/vika.go/lib/common/profile"
	"github.com/vikadata/vika.go/lib/common/util"
	vika "github.com/vikadata/vika.go/lib/datasheet"
	"github.com/vikadata/vika.go/lib/space"
	"os"
	"testing"
)

func TestCreateRecords(t *testing.T) {
	// VIKA_HOST 可以不用设置，默认使用生产的host
	credential := common.NewCredential(os.Getenv("TOKEN"))
	cpf := profile.NewClientProfile()
	datasheet, _ := vika.NewDatasheet(credential, os.Getenv("DATASHEET_ID"), cpf)
	request := vika.NewCreateRecordsRequest()
	request.Records = []*vika.Fields{
		{
			Fields: &vika.Field{
				os.Getenv("FIELD"): vika.NumberFieldValue(900),
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
	// HOST 可以不用设置，默认使用生产的host
	credential := common.NewCredential(os.Getenv("TOKEN"))
	cpf := profile.NewClientProfile()
	datasheet, _ := vika.NewDatasheet(credential, os.Getenv("DATASHEET_ID"), cpf)
	request := vika.NewDescribeRecordRequest()
	request.Sort = []*vika.Sort{
		{
			Field: common.StringPtr(os.Getenv("FIELD")),
			Order: common.StringPtr("desc"),
		},
	}
	request.Fields = common.StringPtrs([]string{os.Getenv("FIELD")})
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
	// HOST 可以不用设置，默认使用生产的host
	credential := common.NewCredential(os.Getenv("TOKEN"))
	cpf := profile.NewClientProfile()
	datasheet, _ := vika.NewDatasheet(credential, os.Getenv("DATASHEET_ID"), cpf)
	request := vika.NewDescribeRecordRequest()
	request.Sort = []*vika.Sort{
		{
			Field: common.StringPtr(os.Getenv("FIELD")),
			Order: common.StringPtr("desc"),
		},
	}
	request.Fields = common.StringPtrs([]string{os.Getenv("FIELD")})
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
	// HOST 可以不用设置，默认使用生产的host
	credential := common.NewCredential(os.Getenv("TOKEN"))
	cpf := profile.NewClientProfile()
	datasheet, _ := vika.NewDatasheet(credential, os.Getenv("DATASHEET_ID"), cpf)
	describeRequest := vika.NewDescribeRecordRequest()
	describeRequest.FilterByFormula = common.StringPtr("{" + os.Getenv("FIELD") + "}=900")
	record, _ := datasheet.DescribeRecord(describeRequest)
	request := vika.NewModifyRecordsRequest()
	request.Records = []*vika.BaseRecord{
		{
			Fields: &vika.Field{
				os.Getenv("FIELD"): vika.NumberFieldValue(1000),
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
	// HOST 可以不用设置，默认使用生产的host
	credential := common.NewCredential(os.Getenv("TOKEN"))
	cpf := profile.NewClientProfile()
	datasheet, _ := vika.NewDatasheet(credential, os.Getenv("DATASHEET_ID"), cpf)
	describeRequest := vika.NewDescribeRecordRequest()
	describeRequest.FilterByFormula = common.StringPtr("{" + os.Getenv("FIELD") + "}=1000")
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
	// HOST 可以不用设置，默认使用生产的host
	credential := common.NewCredential(os.Getenv("TOKEN"))
	cpf := profile.NewClientProfile()
	cpf.Upload = true
	datasheet, _ := vika.NewDatasheet(credential, os.Getenv("DATASHEET_ID"), cpf)
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
	// HOST 可以不用设置，默认使用生产的host
	credential := common.NewCredential(os.Getenv("TOKEN"))
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Domain = os.Getenv("HOST")
	datasheet, _ := vika.NewDatasheet(credential, os.Getenv("DATASHEET_ID"), cpf)
	describeRequest := vika.NewDescribeFieldsRequest()
	describeRequest.ViewId = common.StringPtr(os.Getenv("DATASHEET_VIEW_ID"))
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

func TestDescribeViews(t *testing.T) {
	// HOST 可以不用设置，默认使用生产的host
	credential := common.NewCredential(os.Getenv("TOKEN"))
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Domain = os.Getenv("HOST")
	datasheet, _ := vika.NewDatasheet(credential, os.Getenv("DATASHEET_ID"), cpf)
	describeRequest := vika.NewDescribeViewsRequest()
	views, err := datasheet.DescribeViews(describeRequest)
	if _, ok := err.(*vkerror.VikaSDKError); ok {
		t.Errorf("An API error has returned: %s", err)
	}
	// 非SDK异常，直接失败。实际代码中可以加入其他的处理。
	if err != nil {
		t.Errorf("An unexcepted error has returned: %s", err)
		panic(err)
	}
	util.Dd(views)
	t.Log(len(views))
}

func TestDescribeSpaces(t *testing.T) {
	// HOST 可以不用设置，默认使用生产的host
	credential := common.NewCredential(os.Getenv("TOKEN"))
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Domain = os.Getenv("HOST")
	spaceClient, _ := space.NewSpace(credential, "", cpf)
	describeRequest := space.NewDescribeSpacesRequest()
	spaces, err := spaceClient.DescribeSpaces(describeRequest)
	if _, ok := err.(*vkerror.VikaSDKError); ok {
		t.Errorf("An API error has returned: %s", err)
	}
	// 非SDK异常，直接失败。实际代码中可以加入其他的处理。
	if err != nil {
		t.Errorf("An unexcepted error has returned: %s", err)
		panic(err)
	}
	util.Dd(spaces)
	t.Log(len(spaces))
}

func TestDescribeNodes(t *testing.T) {
	// HOST 可以不用设置，默认使用生产的host
	credential := common.NewCredential(os.Getenv("TOKEN"))
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Domain = os.Getenv("HOST")
	spaceClient, _ := space.NewSpace(credential, os.Getenv("SPACE_ID"), cpf)
	describeRequest := space.NewDescribeNodesRequest()
	nodes, err := spaceClient.DescribeNodes(describeRequest)
	if _, ok := err.(*vkerror.VikaSDKError); ok {
		t.Errorf("An API error has returned: %s", err)
	}
	// 非SDK异常，直接失败。实际代码中可以加入其他的处理。
	if err != nil {
		t.Errorf("An unexcepted error has returned: %s", err)
		panic(err)
	}
	util.Dd(nodes)
	t.Log(len(nodes))
}

func TestDescribeNode(t *testing.T) {
	// HOST 可以不用设置，默认使用生产的host
	credential := common.NewCredential(os.Getenv("TOKEN"))
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Domain = os.Getenv("HOST")
	spaceClient, _ := space.NewSpace(credential, os.Getenv("SPACE_ID"), cpf)
	describeRequest := space.NewDescribeNodeRequest()
	describeRequest.NodeId = common.StringPtr(os.Getenv("DATASHEET_ID"))
	node, err := spaceClient.DescribeNode(describeRequest)
	if _, ok := err.(*vkerror.VikaSDKError); ok {
		t.Errorf("An API error has returned: %s", err)
	}
	// 非SDK异常，直接失败。实际代码中可以加入其他的处理。
	if err != nil {
		t.Errorf("An unexcepted error has returned: %s", err)
		panic(err)
	}
	util.Dd(node)
	t.Log(node)
}