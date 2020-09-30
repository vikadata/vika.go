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
	"encoding/json"
	"github.com/danhper/structomap"
	"reflect"
	"strconv"
)

type vikaRequestSerializer struct {
	*structomap.Base
}

// 实例化
func newVikaRequestSerializer() *vikaRequestSerializer {
	serializer := &vikaRequestSerializer{structomap.New()}
	return serializer
}

func (serializer *vikaRequestSerializer) withPage() *vikaRequestSerializer {
	serializer.PickFunc(func(pageSize interface{}) interface{} {
		if reflect.ValueOf(pageSize).IsZero() {
			return DefaultPageSize
		}
		return pageSize
	}, "PageSize").PickFunc(func(pageNum interface{}) interface{} {
		if reflect.ValueOf(pageNum).IsZero() {
			return DefaultPageNum
		}
		return pageNum
	}, "PageNum")
	return serializer
}

func (serializer *vikaRequestSerializer) withFieldKey() *vikaRequestSerializer {
	serializer.Pick("FieldKey")
	return serializer
}

func (serializer *vikaRequestSerializer) withQuery() *vikaRequestSerializer {
	serializer.Pick("Sort", "RecordIds", "ViewId", "Fields", "FilterByFormula", "CellFormat", "MaxRecords")
	return serializer
}

func (serializer *vikaRequestSerializer) withFields() *vikaRequestSerializer {
	serializer.Pick("Fields")
	return serializer
}

func (serializer *vikaRequestSerializer) withRecords() *vikaRequestSerializer {
	serializer.Pick("Fields", "RecordId")
	return serializer
}

func (serializer *vikaRequestSerializer) withRecordIds() *vikaRequestSerializer {
	serializer.Pick("RecordIds")
	return serializer
}

func bodyByte(records []map[string]interface{}, fieldKey string) []byte {
	form := make(map[string]interface{})
	form["records"] = records
	form["fieldKey"] = fieldKey
	return mapToBytes(form)
}

func recordIdQueryMap(ids []string) map[string]interface{} {
	form := make(map[string]interface{})
	for k, v := range ids {
		form["recordIds["+strconv.Itoa(k)+"]"] = v
	}
	return form
}

func recordIdByte(ids []string) []byte {
	form := make(map[string]interface{})
	for k, v := range ids {
		form["recordIds["+strconv.Itoa(k)+"]"] = v
	}
	return mapToBytes(form)
}

func mapToBytes(imap map[string]interface{}) []byte {
	buf, err := json.Marshal(imap)
	if err != nil {
		panic(err)
	}
	return buf
}
