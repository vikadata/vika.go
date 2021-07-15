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

// SingleTextFieldProperty describe the single text field property
type SingleTextFieldProperty struct {
	DefaultValue string `json:"defaultValue,omitempty" name:"defaultValue"`
}

// SelectFieldProperty describe the single select field and multi select field property
type SelectFieldProperty struct {
	Options []*SelectFieldOption `json:"options,omitempty" name:"options"`
}

// NumberFieldProperty describe number field property
type NumberFieldProperty struct {
	DefaultValue *string `json:"defaultValue,omitempty" name:"defaultValue"`
	Precision    *int    `json:"precision,omitempty" name:"precision"`
}

// CurrencyFieldProperty describe currency field property
type CurrencyFieldProperty struct {
	DefaultValue *string `json:"defaultValue,omitempty" name:"defaultValue"`
	Precision    *int    `json:"precision,omitempty" name:"precision"`
	Symbol       *string `json:"symbol,omitempty" name:"symbol"`
}

// PercentFieldProperty describe percent field property
type PercentFieldProperty struct {
	DefaultValue *string `json:"defaultValue,omitempty" name:"defaultValue"`
	Precision    *int    `json:"precision,omitempty" name:"precision"`
}

// DateTimeFieldProperty describe date time field property
type DateTimeFieldProperty struct {
	Format      *string `json:"format,omitempty" name:"format"`
	IncludeTime *bool   `json:"includeTime,omitempty" name:"includeTime"`
	AutoFill    *bool   `json:"autoFill,omitempty" name:"autoFill"`
}

// MemberFieldProperty describe member field property
type MemberFieldProperty struct {
	Options       []*MemberFieldOption `json:"options,omitempty" name:"options"`
	IsMulti       *bool                `json:"isMulti,omitempty" name:"isMulti"`
	ShouldSendMsg *bool                `json:"shouldSendMsg,omitempty" name:"shouldSendMsg"`
}

// CheckboxFieldProperty describe checkbox field property
type CheckboxFieldProperty struct {
	Icon *string `json:"icon,omitempty" name:"icon"`
}

// RatingFieldProperty describe rating field property
type RatingFieldProperty struct {
	Icon *string `json:"icon,omitempty" name:"icon"`
	Max  *int    `json:"max,omitempty" name:"max"`
}

// MagicLinkFieldProperty describe magic link field property
type MagicLinkFieldProperty struct {
	ForeignDatasheetId *string `json:"foreignDatasheetId,omitempty" name:"foreignDatasheetId"`
	BrotherFieldId     *string `json:"brotherFieldId,omitempty" name:"brotherFieldId"`
	LimitToViewId      *string `json:"limitToViewId,omitempty" name:"limitToViewId"`
	LimitSingleRecord  *bool   `json:"limitSingleRecord,omitempty" name:"limitSingleRecord"`
}

// MagicLookUpFieldProperty describe magic lookup field property
type MagicLookUpFieldProperty struct {
	RelatedLinkFieldId *string            `json:"relatedLinkFieldId,omitempty" name:"relatedLinkFieldId"`
	TargetFieldId      *string            `json:"targetFieldId,omitempty" name:"targetFieldId"`
	RollupFunction     *RollupFunction    `json:"rollupFunction,omitempty" name:"rollupFunction"`
	ValueType          *ValueType         `json:"valueType,omitempty" name:"valueType"`
	EntityField        *LookUpFieldEntity `json:"entityField,omitempty" name:"entityField"`
	Format             *FieldFormat       `json:"format,omitempty" name:"format"`
}

// FormulaFieldProperty describe formula field property
type FormulaFieldProperty struct {
	Expression *string      `json:"expression,omitempty" name:"expression"`
	ValueType  *ValueType   `json:"valueType,omitempty" name:"valueType"`
	HasError   *bool        `json:"hasError,omitempty" name:"hasError"`
	Format     *FieldFormat `json:"format,omitempty" name:"format"`
}

// UserFieldProperty describe createdBy,lastModifiedBy field property
type UserFieldProperty struct {
	Options []*UserInfo `json:"options,omitempty" name:"options"`
}

// SelectFieldOption describe the single select field and multi select field option property
type SelectFieldOption struct {
	Id    *string                 `json:"id,omitempty" name:"id"`
	Name  *string                 `json:"name,omitempty" name:"name"`
	Color *SelectFieldOptionColor `json:"color,omitempty"`
}

// MemberFieldOption describe the member field option property
type MemberFieldOption struct {
	Id     *string     `json:"id,omitempty" name:"id"`
	Name   *string     `json:"name,omitempty" name:"name"`
	Type   *MemberType `json:"type,omitempty" name:"type"`
	Avatar *string     `json:"avatar,omitempty" name:"avatar"`
}

// UserInfo describe the user base info
type UserInfo struct {
	Id     *string `json:"id,omitempty" name:"id"`
	Name   *string `json:"name,omitempty" name:"name"`
	Avatar *string `json:"avatar,omitempty" name:"avatar"`
}

// SelectFieldOptionColor describe the single select field and multi select field option's color property
type SelectFieldOptionColor struct {
	Name  *string `json:"name,omitempty" name:"name"`
	Value *string `json:"value,omitempty" name:"value"`
}

type LookUpFieldEntity struct {
	DatasheetId *string         `json:"datasheetId,omitempty" name:"datasheetId"`
	Field       *DatasheetField `json:"field,omitempty" name:"field"`
}

type LookUpFieldFormat struct {
	Type   *string         `json:"type,omitempty" name:"type"`
	Format *DatasheetField `json:"field,omitempty" name:"field"`
}

// FieldFormat the format of the field set for the record value to show
type FieldFormat struct {
	DateTimeFieldFormat
	NumberFieldFormat
	CurrencyFieldFormat
}

// DateTimeFieldFormat the format for make the value just like the datetime filed shows
type DateTimeFieldFormat struct {
	DateFormat  *string `json:"dateFormat,omitempty" name:"dateFormat"`
	TimeFormat  *string `json:"timeFormat,omitempty" name:"timeFormat"`
	IncludeTime *bool   `json:"includeTime,omitempty" name:"includeTime"`
}

// NumberFieldFormat the format for make the value just like the number filed shows
type NumberFieldFormat struct {
	Precision *int `json:"precision,omitempty" name:"precision"`
}

// CurrencyFieldFormat the format for make the value just like the currency filed shows
type CurrencyFieldFormat struct {
	Precision *int `json:"precision,omitempty" name:"precision"`
	Symbol    *int `json:"symbol,omitempty" name:"symbol"`
}

// MemberType the vika datasheet member field type
type MemberType string

// ValueType the vika datasheet basic value type
type ValueType string

// RollupFunction the vika datasheet supported customize function
type RollupFunction string

// all datasheet member field types
const (
	MemberType_Member MemberType = "Member"
	MemberType_Team   MemberType = "Team"
)

// datasheet supported basic value type
const (
	ValueType_String   ValueType = "String"
	ValueType_Boolean  ValueType = "Boolean"
	ValueType_Number   ValueType = "Number"
	ValueType_DateTime ValueType = "DateTime"
	ValueType_Array    ValueType = "Array"
)

// datasheet supported all rollup functions
const (
	RollupFunction_VALUES       RollupFunction = "VALUES"
	RollupFunction_AVERAGE      RollupFunction = "AVERAGE"
	RollupFunction_COUNT        RollupFunction = "COUNT"
	RollupFunction_COUNTA       RollupFunction = "COUNTA"
	RollupFunction_COUNTALL     RollupFunction = "COUNTALL"
	RollupFunction_SUM          RollupFunction = "SUM"
	RollupFunction_MIN          RollupFunction = "MIN"
	RollupFunction_MAX          RollupFunction = "MAX"
	RollupFunction_AND          RollupFunction = "AND"
	RollupFunction_OR           RollupFunction = "OR"
	RollupFunction_XOR          RollupFunction = "XOR"
	RollupFunction_CONCATENATE  RollupFunction = "CONCATENATE"
	RollupFunction_ARRAYJOIN    RollupFunction = "ARRAYJOIN"
	RollupFunction_ARRAYUNIQUE  RollupFunction = "ARRAYUNIQUE"
	RollupFunction_ARRAYCOMPACT RollupFunction = "ARRAYCOMPACT"
)

// Sort 需要排序的字段
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

type DescribeFieldsRequest struct {
	*vkhttp.BaseRequest
	// viewId 按照【视图】进行过滤。形如：viw*****。类型：String 必选：否
	ViewId *string `json:"viewId,omitempty" name:"viewId"`
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
	Data *RecordPagination `json:"data"`
}

type UploadResponse struct {
	*vkhttp.BaseResponse
	// api返回数据
	Data *Attachment `json:"data"`
}

type FieldsResponse struct {
	Fields []*DatasheetField `json:"fields"`
}

type DescribeFieldsResponse struct {
	*vkhttp.BaseResponse
	// api返回数据
	Data *FieldsResponse `json:"data"`
}
