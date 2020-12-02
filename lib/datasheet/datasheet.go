// Package datasheet provides the operations about datasheet
package datasheet

import (
	"fmt"
	"github.com/vikadata/vika.go/lib/common"
	vkhttp "github.com/vikadata/vika.go/lib/common/http"
	"github.com/vikadata/vika.go/lib/common/profile"
	"math"
)

const maxPageSize = 1000
const recordPath = "/fusion/v1/datasheets/%s/records"
const attachPath = "/fusion/v1/datasheets/%s/attachments"

type Datasheet struct {
	common.Client
	DatasheetId string
}

// init datasheet instance
func NewDatasheet(credential *common.Credential, datasheetId string, clientProfile *profile.ClientProfile) (datasheet *Datasheet, err error) {
	datasheet = &Datasheet{}
	datasheet.DatasheetId = datasheetId
	datasheet.Init().WithCredential(credential).WithProfile(clientProfile)
	return
}

// init datasheet record request instance
func NewDescribeRecordRequest() (request *DescribeRecordRequest) {
	request = &DescribeRecordRequest{
		BaseRequest: &vkhttp.BaseRequest{},
	}

	return
}

func NewCreateRecordsRequest() (request *CreateRecordsRequest) {
	request = &CreateRecordsRequest{
		BaseRequest: &vkhttp.BaseRequest{},
	}
	return
}

func NewModifyRecordsRequest() (request *ModifyRecordsRequest) {
	request = &ModifyRecordsRequest{
		BaseRequest: &vkhttp.BaseRequest{},
	}
	return
}

func NewDeleteRecordsRequest() (request *DeleteRecordsRequest) {
	request = &DeleteRecordsRequest{
		BaseRequest: &vkhttp.BaseRequest{},
	}
	return
}

func NewUploadRequest() (request *UploadRequest) {
	request = &UploadRequest{
		BaseRequest: &vkhttp.BaseRequest{},
	}
	return
}

func NewDescribeRecordResponse() (response *DescribeRecordResponse) {
	response = &DescribeRecordResponse{
		BaseResponse: &vkhttp.BaseResponse{},
	}
	return
}

func NewUploadResponse() (response *UploadResponse) {
	response = &UploadResponse{
		BaseResponse: &vkhttp.BaseResponse{},
	}
	return
}

// 本接口 (DescribeAllRecord) 用于查询所有记录的详细信息。
//
// * 可以根据视图`ViewId`、列名称或者列ID等信息来查询记录的详细信息。过滤信息详细请见`RecordRequest`。
// * 如果参数为空，返回当前数表的所有记录。
func (c *Datasheet) DescribeAllRecords(request *DescribeRecordRequest) (records []*Record, err error) {
	if request == nil {
		request = NewDescribeRecordRequest()
	}
	request.Init().SetPath(fmt.Sprintf(recordPath, c.DatasheetId))
	request.SetHttpMethod(vkhttp.GET)
	request.PageSize = common.Int64Ptr(maxPageSize)
	request.PageNum = common.Int64Ptr(1)
	response := NewDescribeRecordResponse()
	err = c.Send(request, response)
	if err != nil {
		return nil, err
	}
	total := response.Data.Total
	// 计算循环总次数
	if *total > maxPageSize {
		times := int(math.Ceil(float64(*total / maxPageSize)))
		for i := 1; i <= times; i++ {
			request.PageNum = common.Int64Ptr(int64(i + 1))
			tmp := NewDescribeRecordResponse()
			err = c.Send(request, tmp)
			if err != nil {
				// 其中任何一次失败 都失败
				return nil, err
			}
			response.Data.Records = append(response.Data.Records, tmp.Data.Records...)
		}
	}
	response.Data.PageNum = common.Int64Ptr(0)
	return response.Data.Records, nil
}

// 本接口 (DescribeRecords) 用于查询分页记录的详细信息。
//
// * 可以根据视图`ViewId`、列名称或者列ID等信息来查询记录的详细信息。过滤信息详细请见`RecordRequest`。
// * 如果参数为空，返回根据默认分页的查询的数据。默认每页100条
func (c *Datasheet) DescribeRecords(request *DescribeRecordRequest) (pagination *RecordPagination, err error) {
	if request == nil {
		request = NewDescribeRecordRequest()
	}
	request.Init().SetPath(fmt.Sprintf(recordPath, c.DatasheetId))
	request.SetHttpMethod(vkhttp.GET)
	response := NewDescribeRecordResponse()
	err = c.Send(request, response)
	if err != nil {
		return nil, err
	}
	return response.Data, nil
}

// 本接口 (DescribeRecord) 用于获取单条记录
//
// * 可以根据视图`ViewId`、列名称或者列ID等信息来查询记录的详细信息。过滤信息详细请见`RecordRequest`。
// * 返回查询到的第一条记录
func (c *Datasheet) DescribeRecord(request *DescribeRecordRequest) (record *Record, err error) {
	if request == nil {
		request = NewDescribeRecordRequest()
	}
	request.Init().SetPath(fmt.Sprintf(recordPath, c.DatasheetId))
	request.SetHttpMethod(vkhttp.GET)
	response := NewDescribeRecordResponse()
	err = c.Send(request, response)
	if err != nil {
		return nil, err
	}
	return response.Data.Records[0], nil
}

// 本接口 (CreateRecords) 用于创建多条记录
func (c *Datasheet) CreateRecords(request *CreateRecordsRequest) (records []*Record, err error) {
	if request == nil {
		request = NewCreateRecordsRequest()
	}
	request.Init().SetPath(fmt.Sprintf(recordPath, c.DatasheetId))
	request.SetContentType(vkhttp.JsonContent)
	response := NewDescribeRecordResponse()
	err = c.Send(request, response)
	if err != nil {
		return nil, err
	}
	return response.Data.Records, nil
}

// 本接口 (DeleteRecords) 用于修改多条记录
func (c *Datasheet) ModifyRecords(request *ModifyRecordsRequest) (records []*Record, err error) {
	if request == nil {
		request = NewModifyRecordsRequest()
	}
	request.Init().SetPath(fmt.Sprintf(recordPath, c.DatasheetId))
	request.SetContentType(vkhttp.JsonContent)
	request.SetHttpMethod(vkhttp.PATCH)
	response := NewDescribeRecordResponse()
	err = c.Send(request, response)
	if err != nil {
		return nil, err
	}
	return response.Data.Records, nil
}

// 本接口 (DeleteRecords) 用于删除多条记录
func (c *Datasheet) DeleteRecords(request *DeleteRecordsRequest) (err error) {
	if request == nil {
		request = NewDeleteRecordsRequest()
	}
	request.Init().SetPath(fmt.Sprintf(recordPath, c.DatasheetId))
	request.SetHttpMethod(vkhttp.DELETE)
	response := NewDescribeRecordResponse()
	err = c.Send(request, response)
	return
}

// 本接口 (UploadFile) 用于上传附件
func (c *Datasheet) UploadFile(request *UploadRequest) (attachment *Attachment, err error) {
	if request == nil {
		request = NewUploadRequest()
	}
	body, contentType, err := common.FileBuffer(request.FilePath)
	if err != nil {
		return
	}
	request.Init()
	request.SetPath(fmt.Sprintf(attachPath, c.DatasheetId))
	request.SetHttpMethod(vkhttp.POST)
	request.SetFile(body)
	request.SetContentType(contentType)
	response := NewUploadResponse()
	err = c.Send(request, response)
	if err != nil {
		return nil, err
	}
	return response.Data, nil
}
