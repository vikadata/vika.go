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

package vika

import (
	"math"
	"reflect"
)

type IDatasheet interface {
	All(params RecordQueryParam) (*RecordResponse, error)
	Get(params RecordQueryParam) (*PageResponse, error)
	Add(records []NewRecord, fieldKey string) (*RecordResponse, error)
	Update(records []Record, fieldKey string) (*RecordResponse, error)
	find(recordIds []string, fieldKey string)
	Del(ids []string) (*RecordResponse, error)
	Upload()
}

func newDatasheet(datasheetId string, vika Vika) *Datasheet {
	datasheet := &Datasheet{}
	datasheet.datasheetId = datasheetId
	// 默认全局的request配置
	datasheet.request = newVikaRequest(vika, datasheetId)
	datasheet.fieldKey = vika.FieldKey
	return datasheet
}

func (datasheet Datasheet) All(params RecordQueryParam) (*RecordResponse, error) {
	if reflect.ValueOf(params.FieldKey).IsZero() {
		params.FieldKey = datasheet.fieldKey
	}
	params.PageSize = MaxPageSize
	result, err := datasheet.request.getRecords(params)
	if err != nil {
		return nil, err
	}
	response := &RecordResponse{Success: true, Message: "SUCCESS", Code: 200, Data: Records{
		result.Data.Records,
	}}
	total := result.Data.Total
	// 计算循环总次数
	if total > MaxPageSize {
		times := int(math.Ceil(float64(total / MaxPageSize)))
		for i := 1; i <= times; i++ {
			params.PageNum = i + 1
			tmp, tmpErr := datasheet.request.getRecords(params)
			if tmpErr != nil {
				// 其中任何一次失败 都失败
				return nil, err
			}
			response.Data.Records = append(response.Data.Records, tmp.Data.Records...)
		}
	}
	return response, err
}

func (datasheet Datasheet) Get(params RecordQueryParam) (*PageResponse, error) {
	if reflect.ValueOf(params.FieldKey).IsZero() {
		params.FieldKey = datasheet.fieldKey
	}
	return datasheet.request.getRecords(params)

}

func (datasheet Datasheet) Add(records []NewRecord, fieldKey ...string) (*RecordResponse, error) {
	fieldKeyParam := datasheet.fieldKey
	if fieldKey != nil {
		fieldKeyParam = fieldKey[0]
	}
	return datasheet.request.addRecords(records, fieldKeyParam)
}

func (datasheet Datasheet) Update(records []Record, fieldKey ...string) (*RecordResponse, error) {
	fieldKeyParam := datasheet.fieldKey
	if fieldKey != nil {
		fieldKeyParam = fieldKey[0]
	}
	return datasheet.request.updateRecords(records, fieldKeyParam)
}

func (datasheet Datasheet) Del(ids []string) (*DeleteResponse, error) {
	return datasheet.request.delRecords(ids)
}

func (datasheet Datasheet) Upload(filePath string) (*AttachmentResponse, error) {
	return datasheet.request.uploadAsset(filePath)
}

func (datasheet Datasheet) Find(recordIds []string, fieldKey ...string) (*RecordResponse, error) {
	fieldKeyParam := datasheet.fieldKey
	if fieldKey != nil {
		fieldKeyParam = fieldKey[0]
	}
	params := RecordQueryParam{RecordIds: recordIds, FieldKey: fieldKeyParam}
	if reflect.ValueOf(params.FieldKey).IsZero() {
		params.FieldKey = datasheet.fieldKey
	}
	res, err := datasheet.request.getRecords(RecordQueryParam{RecordIds: recordIds})
	if err != nil {
		return nil, err
	}
	return &RecordResponse{Success: true, Message: "SUCCESS", Code: 200, Data: Records{
		res.Data.Records,
	}}, nil
}
