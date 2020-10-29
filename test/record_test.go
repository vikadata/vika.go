/*
 *
 * MIT License
 *
 * Copyright (c) 2020 VIKADATA
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */
package test

import (
	vika "github.com/vikadata/vika.go/lib"
	"os"
	"reflect"
	_sort "sort"
	"strings"
	"testing"
)

/* 必要步骤：
 * 实例化一个Vika对象，入参需要传入VIKA账户开发者token\n
 * 这里采用的是从环境变量读取的方式，需要在环境变量中先设置这两个值。go env -u -
 * 你也可以直接在代码中写死token，但是小心不要将代码复制、上传或者分享给他人，
 * 以免泄露token危及你的财产安全。
 * token获取: https://vika.cn */
func TestDatasheetAll(t *testing.T) {
	// VIKA_HOST 可以不用设置，默认使用生产的host
	vikaClient := vika.New(vika.VikaConfig{Token: os.Getenv("VIKA_TOKEN"), Host: os.Getenv("VIKA_HOST")})
	params := vika.RecordQueryParam{}
	_, err := vikaClient.Datasheet(os.Getenv("VIKA_DATASHEET_ID")).All(params)
	if err != nil {
		t.Errorf("TestDatasheetAll:Datasheet.All(%#v) err = %#v; expected nil", params, err)
	}
}

func TestPage(t *testing.T) {
	vikaClient := vika.New(vika.VikaConfig{Token: os.Getenv("VIKA_TOKEN"), Host: os.Getenv("VIKA_HOST")})
	params := vika.RecordQueryParam{}
	records, err := vikaClient.Datasheet(os.Getenv("VIKA_DATASHEET_ID")).Get(params)
	if err != nil {
		t.Errorf("TestPage:Datasheet.All(%#v) err = %#v; expected nil", params, err)
	}
	if reflect.ValueOf(records.Data.PageNum).IsZero() {
		t.Errorf("TestPage.All(%#v) err = %#v; expected page", params, err)
	}
}

func TestOrderParam(t *testing.T) {
	// VIKA_HOST 可以不用设置，默认使用生产的host
	vikaClient := vika.New(vika.VikaConfig{Token: os.Getenv("VIKA_TOKEN"), Host: os.Getenv("VIKA_HOST")})
	params := vika.RecordQueryParam{}
	// 获取全部的数据, VIKA_FIELD需要为float64类型
	sort := []vika.Sort{{Field: os.Getenv("VIKA_FIELD"), Order: "asc"}}
	v := reflect.ValueOf(sort)
	num := v.Len()
	querySort := make([]interface{}, num)
	for i := 0; i < num; i++ {
		querySort[i] = v.Index(i).Interface()
	}
	params.Sort = querySort
	records, err := vikaClient.Datasheet(os.Getenv("VIKA_DATASHEET_ID")).All(params)
	if err != nil {
		t.Errorf("TestOrderParam:Datasheet.All(%#v) err = %#v; expected nil", params, err)
	}
	sortFieldValue := make([]float64, len(records.Data.Records))
	for i := range records.Data.Records {
		sortFieldValue[i] = reflect.ValueOf(records.Data.Records[i].Fields[os.Getenv("VIKA_FIELD")]).Float()
	}
	if !_sort.Float64sAreSorted(sortFieldValue) {
		t.Errorf("TestDatasheetOrder:Datasheet.All(%#v) err = %#v; expected querySorted", params, err)
	}
}

func TestCellFormatStringParam(t *testing.T) {
	vikaClient := vika.New(vika.VikaConfig{Token: os.Getenv("VIKA_TOKEN"), Host: os.Getenv("VIKA_HOST")})
	params := vika.RecordQueryParam{}
	params.CellFormat = vika.CellFormatString
	records, err := vikaClient.Datasheet(os.Getenv("VIKA_DATASHEET_ID")).All(params)
	if err != nil {
		t.Errorf("TestCellFormatStringParam:Datasheet.All(%#v) err = %#v; expected nil", params, err)
	}
	if reflect.TypeOf(records.Data.Records[0].Fields[os.Getenv("VIKA_FIELD")]).Name() != "string" {
		t.Errorf("TestCellFormatStringParam:Datasheet.All(%#v) err = %#v; expected CellFormatString", params, err)
	}
}

func TestFieldsParam(t *testing.T) {
	vikaClient := vika.New(vika.VikaConfig{Token: os.Getenv("VIKA_TOKEN"), Host: os.Getenv("VIKA_HOST")})
	params := vika.RecordQueryParam{}
	params.Fields = []string{os.Getenv("VIKA_FIELD")}
	records, err := vikaClient.Datasheet(os.Getenv("VIKA_DATASHEET_ID")).All(params)
	if err != nil {
		t.Errorf("TestFieldsParam:Datasheet.All(%#v) err = %#v; expected nil", params, err)
	}
	if len(records.Data.Records[0].Fields) > 1 {
		t.Errorf("TestFieldsParam:Datasheet.All(%#v) err = %#v; expected fielda", params, err)
	}
}

func TestViewIdParam(t *testing.T) {
	vikaClient := vika.New(vika.VikaConfig{Token: os.Getenv("VIKA_TOKEN"), Host: os.Getenv("VIKA_HOST")})
	params := vika.RecordQueryParam{}
	params.ViewId = os.Getenv("VIKA_VIEW_ID")
	records, err := vikaClient.Datasheet(os.Getenv("VIKA_DATASHEET_ID")).All(params)
	if err != nil {
		t.Errorf("TestViewIdParam:Datasheet.All(%#v) err = %#v; expected nil", params, err)
	}
	if len(records.Data.Records) == 0 {
		t.Errorf("TestViewIdParam:Datasheet.All(%#v) err = %#v; expected view data", params, err)
	}
}

func TestFieldKeyParam(t *testing.T) {
	vikaClient := vika.New(vika.VikaConfig{Token: os.Getenv("VIKA_TOKEN"), Host: os.Getenv("VIKA_HOST"), FieldKey: vika.FieldKeyId})
	params := vika.RecordQueryParam{}
	records, err := vikaClient.Datasheet(os.Getenv("VIKA_DATASHEET_ID")).All(params)
	if err != nil {
		t.Errorf("TestViewIdParam:Datasheet.All(%#v) err = %#v; expected nil", params, err)
	}
	for k := range records.Data.Records[0].Fields {
		if !strings.Contains(k, "fld") {
			t.Errorf("TestViewIdParam:Datasheet.All(%#v) err = %#v; expected fieldIddata ", params, err)
		}
	}
}

func TestRecordIdsParam(t *testing.T) {
	vikaClient := vika.New(vika.VikaConfig{Token: os.Getenv("VIKA_TOKEN"), Host: os.Getenv("VIKA_HOST"), FieldKey: vika.FieldKeyId})
	datasheet := vikaClient.Datasheet(os.Getenv("VIKA_DATASHEET_ID"))
	page, _ := datasheet.Get(vika.RecordQueryParam{PageNum: 1, PageSize: 2})
	recordIds := []string{page.Data.Records[0].RecordId, page.Data.Records[1].RecordId}
	records, err := datasheet.Find(recordIds)
	if err != nil {
		t.Errorf("TestRecordIdsParam:Datasheet.All(%#v) err = %#v; expected nil", recordIds, err)
	}
	if len(records.Data.Records) < len(recordIds) {
		t.Errorf("TestRecordIdsParam:Datasheet.All(%#v) err = %#v; expected fieldIddata ", recordIds, err)
	}
}

func TestAddRecords(t *testing.T) {
	vikaClient := vika.New(vika.VikaConfig{Token: os.Getenv("VIKA_TOKEN"), Host: os.Getenv("VIKA_HOST"), FieldKey: vika.FieldKeyName})
	datasheet := vikaClient.Datasheet(os.Getenv("VIKA_DATASHEET_ID"))
	var addForm = []vika.NewRecord{
		{
			Fields: map[string]interface{}{
				os.Getenv("VIKA_FIELD"): 88,
			},
		},
		{
			Fields: map[string]interface{}{
				os.Getenv("VIKA_FIELD"): 99,
			},
		},
	}
	records, err := datasheet.Add(addForm)
	if err != nil {
		t.Errorf("TestAddRecords:Datasheet.All(%#v) err = %#v; expected nil", addForm, err)
	}
	if len(records.Data.Records) < len(addForm) {
		t.Errorf("TestAddRecords:Datasheet.All(%#v) err = %#v; expected add ", len(addForm), err)
	}
}

func TestUpdate(t *testing.T) {
	vikaClient := vika.New(vika.VikaConfig{Token: os.Getenv("VIKA_TOKEN"), Host: os.Getenv("VIKA_HOST"), FieldKey: vika.FieldKeyName})
	datasheet := vikaClient.Datasheet(os.Getenv("VIKA_DATASHEET_ID"))
	page, _ := datasheet.Get(vika.RecordQueryParam{PageNum: 1, PageSize: 2, FilterByFormula: "OR({数字ID}=88, {数字ID}=99)"})
	var updateForm = []vika.Record{
		{
			Fields: map[string]interface{}{
				"数字ID": 100,
			},
			RecordId: page.Data.Records[0].RecordId,
		},
		{
			Fields: map[string]interface{}{
				"数字ID": 101,
			},
			RecordId: page.Data.Records[1].RecordId,
		},
	}
	records, err := datasheet.Update(updateForm)
	if err != nil {
		t.Errorf("TestUpdate:Datasheet.All(%#v) err = %#v; expected nil", updateForm, err)
	}
	if len(records.Data.Records) < len(updateForm) {
		t.Errorf("TestUpdate:Datasheet.All(%#v) err = %#v; expected add ", len(updateForm), err)
	}
}

func TestDelete(t *testing.T) {
	vikaClient := vika.New(vika.VikaConfig{Token: os.Getenv("VIKA_TOKEN"), Host: os.Getenv("VIKA_HOST"), FieldKey: vika.FieldKeyName})
	datasheet := vikaClient.Datasheet(os.Getenv("VIKA_DATASHEET_ID"))
	page, _ := datasheet.Get(vika.RecordQueryParam{PageNum: 1, PageSize: 2, FilterByFormula: "OR({数字ID}=100, {数字ID}=101)"})
	recordIds := []string{page.Data.Records[0].RecordId, page.Data.Records[1].RecordId}
	_, err := datasheet.Del(recordIds)
	if err != nil {
		t.Errorf("TestAddRecords:Datasheet.All(%#v) err = %#v; expected nil", err, err)
	}
}
