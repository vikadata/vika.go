package datasheet

import (
	"encoding/json"
	vkhttp "github.com/vikadata/vika.go/lib/common/http"
)

func (r *DescribeRecordRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRecordRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type (
	FieldValue       interface{}
	NumberFieldValue int64
	TextFieldValue   string
	UnitFieldValue   struct {
		UnitName string `json:"unitName,omitempty" name:"unitName"`
		UnitType string `json:"unitType,omitempty" name:"unitType"`
		UnitId   string `json:"unitId,omitempty" name:"unitId"`
	}
	AttachmentValue struct {
		Token string `json:"token,omitempty" name:"token"`
		Name  string `json:"name,omitempty" name:"name"`
	}
)
type (
	Field map[string]FieldValue
)

// 需要排序的字段
type Sort struct {

	// 需要排序的字段名称
	Field *string `json:"Field,omitempty" name:"field"`

	// 排序顺序 desc/asc
	Order *string `json:"Order,omitempty" name:"order"`
}

type DescribeRecordRequest struct {
	*vkhttp.BaseRequest

	// 按照一个或者多个recordId查询。recordId形如：`rec*****`。（此参数的具体格式可参考API[开发者文档](https://vika.cn/help/api-get-records/)的`输入参数说明`一节）。每次请求的实例的上限为100。参数不支持同时指定`RecordIds`和`Filters`。
	RecordIds []*string `json:"recordIds,omitempty" name:"recordIds" list`

	// viewId 按照【视图】进行过滤。形如：viw*****。类型：String 必选：否
	ViewId *string `json:"viewId,omitempty" name:"viewId"`

	// fields 按照【字段】进行过滤，可通过登录[空间站](https://vika.cn)进入数表进行查看。形如：******。类型：String 必选：否
	Fields []*string `json:"fields,omitempty" name:"fields"`

	// filterByFormula 按照【公式】进行过滤, 公式使用方式详见[一分钟上手公式](https://vika.cn/help/tutorial-getting-started-with-formulas/)。形如：max({field})。类型：String 必选：否
	FilterByFormula *string `json:"filterByFormula,omitempty" name:"filterByFormula"`

	// cellFormat 按照【单元格值类型】进行过滤,默认为 json，指定为 string 时所有值都将被自动转换为 string 格式。形如: json。类型：String 必选：否
	CellFormat *string `json:"cellFormat,omitempty" name:"cellFormat"`

	// fieldKey 按照【列形式】进行过滤,默认使用列名 ‘name’ 。形如：name。类型：String 必选：否
	FieldKey *string `json:"fieldKey,omitempty" name:"fieldKey"`

	// 参数不支持同时指定`RecordIds`和`Filters`。
	// sort 按照【排序】进行过滤。 形如：{field: ‘fieldname’, order: ‘asc/desc’}。类型：String 必选：否
	Sort []*Sort `json:"sort,omitempty" name:"sort" list`

	// 指定分页的页码，默认为 1。与参数pageSize配合使用。 进一步介绍请参考 API[开发者文档](https://vika.cn/help/api-get-records/)中的相关小节。
	PageNum *int64 `json:"pageNum,omitempty" name:"pageNum"`

	// 返回数量，默认为100，最大值为1000。关于`PageSize`的更进一步介绍请参考 API[开发者文档](https://vika.cn/help/api-get-records/)中的相关小节。
	PageSize *int64 `json:"pageSize,omitempty" name:"pageSize"`

	// 限制返回记录的总数量, 进一步介绍请参考 API[开发者文档](https://vika.cn/help/api-get-records/)中的相关小节。
	MaxRecords *int64 `json:"maxRecords,omitempty" name:"maxRecords"`
}

type Fields struct {
	Fields *Field `json:"fields,omitempty" name:"fields" map`
}

type BaseRecord struct {
	// 记录的ID 形如：`rec*****`
	RecordId *string `json:"recordId,omitempty" name:"recordId"`
	// 记录的列对应的key/value
	Fields *Field `json:"fields,omitempty" name:"fields" map`
}

type Record struct {
	*BaseRecord
	// 记录创建时间 形如：时间戳
	CreatedAt *int64 `json:"createdAt,omitempty" name:"createdAt"`
}

type CreateRecordsRequest struct {
	*vkhttp.BaseRequest
	// 记录的列对应的key/value
	Records []*Fields `json:"records,omitempty" name:"records" map`
}

type ModifyRecordsRequest struct {
	*vkhttp.BaseRequest
	// 记录的列对应的key/value
	Records []*BaseRecord `json:"records,omitempty" name:"records" map`
}

type DeleteRecordsRequest struct {
	*vkhttp.BaseRequest
	// 记录的列对应的key/value
	RecordIds []*string `json:"recordIds,omitempty" name:"recordIds" list`
}

type UploadRequest struct {
	*vkhttp.BaseRequest
	// 文件路径
	FilePath string `json:"filePath,omitempty" name:"filePath" string`
}

type RecordPagination struct {
	// 当前页数
	PageNum *int64 `json:"pageNum,omitempty" name:"pageNum"`

	PageSize *int64 `json:"pageSize,omitempty" name:"pageSize"`

	Total *int64 `json:"total,omitempty" name:"total"`

	Records []*Record `json:"records"`
}

type Attachment struct {
	// 附件唯一标识
	Token *string `json:"token,omitempty" name:"token"`
	//附件原始名称
	Name *string `json:"name,omitempty" name:"name"`
	//附件大小
	Size *int64 `json:"size,omitempty" name:"size"`
	// 附件宽 图片才返回
	Width *int64 `json:"width,omitempty" name:"width"`
	// 附件高 图片才返回
	Height *int64 `json:"height,omitempty" name:"height"`
	// 附件类型 形如：image/jpeg
	MimeType *string `json:"mimeType,omitempty" name:"mimeType"`
	// pdf预览图,只有pdf格式才会返回
	Preview *string `json:"preview,omitempty" name:"preview"`
	// 附件访问路径
	Url *string `json:"url,omitempty" name:"url"`
}

type DescribeRecordResponse struct {
	*vkhttp.BaseResponse
	// api返回数据
	Data *RecordPagination `json:"Data"`
}

type UploadResponse struct {
	*vkhttp.BaseResponse
	// api返回数据
	Data *Attachment `json:"Data"`
}
