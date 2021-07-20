package datasheet

import (
	"encoding/json"
	"fmt"
	vkhttp "github.com/vikadata/vika.go/lib/common/http"
)

const fieldPath = "/fusion/v1/datasheets/%s/fields"

// FieldType the vika datasheet field type
type FieldType string

// all vika datasheet field types
const (
	FieldType_SingleText       FieldType = "SingleText"
	FieldType_Text             FieldType = "Text"
	FieldType_SingleSelect     FieldType = "SingleSelect"
	FieldType_MultiSelect      FieldType = "MultiSelect"
	FieldType_Number           FieldType = "Number"
	FieldType_Currency         FieldType = "Currency"
	FieldType_Percent          FieldType = "Percent"
	FieldType_DateTime         FieldType = "DateTime"
	FieldType_Attachment       FieldType = "Attachment"
	FieldType_Member           FieldType = "Member"
	FieldType_Checkbox         FieldType = "Checkbox"
	FieldType_Rating           FieldType = "Rating"
	FieldType_URL              FieldType = "URL"
	FieldType_Phone            FieldType = "Phone"
	FieldType_MagicLink        FieldType = "MagicLink"
	FieldType_MagicLookUp      FieldType = "MagicLookUp"
	FieldType_Formula          FieldType = "Formula"
	FieldType_AutoNumber       FieldType = "AutoNumber"
	FieldType_CreatedTime      FieldType = "CreatedTime"
	FieldType_LastModifiedTime FieldType = "LastModifiedTime"
	FieldType_CreatedBy        FieldType = "CreatedBy"
	FieldType_LastModifiedBy   FieldType = "LastModifiedBy"
)

// IDatasheetField define how to obtain the field properties
type IDatasheetField interface {
	SingleTextFieldProperty() *SingleTextFieldProperty
	SelectFieldProperty() *SelectFieldProperty
	PercentFieldProperty() *PercentFieldProperty
	FormulaFieldProperty() *FormulaFieldProperty
	UserFieldProperty() *UserFieldProperty
	MagicLinkFieldProperty() *MagicLinkFieldProperty
	MagicLookUpFieldProperty() *MagicLookUpFieldProperty
	RatingFieldProperty() *RatingFieldProperty
	CheckboxFieldProperty() *CheckboxFieldProperty
	DateTimeFieldProperty() *DateTimeFieldProperty
	MemberFieldProperty() *MemberFieldProperty
	CurrencyFieldProperty() *CurrencyFieldProperty
	NumberFieldProperty() *NumberFieldProperty
}

// DatasheetField describe the fields of table
type DatasheetField struct {
	// 字段 ID
	Id *string `json:"id,omitempty" name:"id"`
	// 字段名称
	Name *string `json:"name,omitempty" name:"name"`
	// 字段类型
	Type *FieldType `json:"type,omitempty" name:"type"`
	// 字段权限，即列权限，true 为可编辑，false 为只读
	Editable *bool `json:"editable,omitempty" name:"editable"`
	// 字段属性。不同的字段有不同的属性
	Property *json.RawMessage `json:"property,omitempty" name:"property"`
}

func (field *DatasheetField) SingleTextFieldProperty() *SingleTextFieldProperty {
	if *field.Type == FieldType_SingleText {
		property := &SingleTextFieldProperty{}
		err := json.Unmarshal(*field.Property, &property)
		if err == nil {
			return property
		}
	}
	return nil
}

func (field *DatasheetField) SelectFieldProperty() *SelectFieldProperty {
	if *field.Type == FieldType_SingleSelect || *field.Type == FieldType_MultiSelect {
		property := &SelectFieldProperty{}
		err := json.Unmarshal(*field.Property, &property)
		if err == nil {
			return property
		}
	}
	return nil
}

func (field *DatasheetField) PercentFieldProperty() *PercentFieldProperty {
	if *field.Type == FieldType_Percent {
		property := &PercentFieldProperty{}
		err := json.Unmarshal(*field.Property, &property)
		if err == nil {
			return property
		}
	}
	return nil
}

func (field *DatasheetField) FormulaFieldProperty() *FormulaFieldProperty {
	if *field.Type == FieldType_Formula {
		property := &FormulaFieldProperty{}
		err := json.Unmarshal(*field.Property, &property)
		if err == nil {
			return property
		}
	}
	return nil
}

func (field *DatasheetField) UserFieldProperty() *UserFieldProperty {
	if *field.Type == FieldType_LastModifiedBy || *field.Type == FieldType_CreatedBy {
		property := &UserFieldProperty{}
		err := json.Unmarshal(*field.Property, &property)
		if err == nil {
			return property
		}
	}
	return nil
}

func (field *DatasheetField) MagicLinkFieldProperty() *MagicLinkFieldProperty {
	if *field.Type == FieldType_MagicLink {
		property := &MagicLinkFieldProperty{}
		err := json.Unmarshal(*field.Property, &property)
		if err == nil {
			return property
		}
	}
	return nil
}

func (field *DatasheetField) MagicLookUpFieldProperty() *MagicLookUpFieldProperty {
	if *field.Type == FieldType_MagicLookUp {
		property := &MagicLookUpFieldProperty{}
		err := json.Unmarshal(*field.Property, &property)
		if err == nil {
			return property
		}
	}
	return nil
}

func (field *DatasheetField) RatingFieldProperty() *RatingFieldProperty {
	if *field.Type == FieldType_Rating {
		property := &RatingFieldProperty{}
		err := json.Unmarshal(*field.Property, &property)
		if err == nil {
			return property
		}
	}
	return nil
}

func (field *DatasheetField) CheckboxFieldProperty() *CheckboxFieldProperty {
	if *field.Type == FieldType_Checkbox {
		property := &CheckboxFieldProperty{}
		err := json.Unmarshal(*field.Property, &property)
		if err == nil {
			return property
		}
	}
	return nil
}

func (field *DatasheetField) DateTimeFieldProperty() *DateTimeFieldProperty {
	if *field.Type == FieldType_DateTime || *field.Type == FieldType_LastModifiedTime || *field.Type == FieldType_CreatedTime {
		property := &DateTimeFieldProperty{}
		err := json.Unmarshal(*field.Property, &property)
		if err == nil {
			return property
		}
	}
	return nil
}

func (field *DatasheetField) MemberFieldProperty() *MemberFieldProperty {
	if *field.Type == FieldType_Member {
		property := &MemberFieldProperty{}
		err := json.Unmarshal(*field.Property, &property)
		if err == nil {
			return property
		}
	}
	return nil
}

func (field *DatasheetField) CurrencyFieldProperty() *CurrencyFieldProperty {
	if *field.Type == FieldType_Currency {
		property := &CurrencyFieldProperty{}
		err := json.Unmarshal(*field.Property, &property)
		if err == nil {
			return property
		}
	}
	return nil
}

func (field *DatasheetField) NumberFieldProperty() *NumberFieldProperty {
	if *field.Type == FieldType_Number {
		property := &NumberFieldProperty{}
		err := json.Unmarshal(*field.Property, &property)
		if err == nil {
			return property
		}
	}
	return nil
}

func NewDescribeFieldsRequest() (request *DescribeFieldsRequest) {
	request = &DescribeFieldsRequest{
		BaseRequest: &vkhttp.BaseRequest{},
	}
	return
}

func newDescribeFieldsResponse() (response *DescribeFieldsResponse) {
	response = &DescribeFieldsResponse{
		BaseResponse: &vkhttp.BaseResponse{},
	}
	return
}

func (c *Datasheet) DescribeFields(request *DescribeFieldsRequest) (fields []*DatasheetField, err error) {
	if request == nil {
		request = NewDescribeFieldsRequest()
	}
	request.Init().SetPath(fmt.Sprintf(fieldPath, c.DatasheetId))
	request.SetHttpMethod(vkhttp.GET)
	response := newDescribeFieldsResponse()
	err = c.Send(request, response)
	if err != nil {
		return nil, err
	}
	return response.Data.Fields, nil
}
