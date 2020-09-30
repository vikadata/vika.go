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

type Record struct {
	RecordId string                 `json:"recordId"`
	Fields   map[string]interface{} `json:"fields"`
}

type NewRecord struct {
	Fields map[string]interface{} `json:"fields"`
}

type RecordDeleteParam struct {
	RecordIds []string `url:"recordIds"`
}

type RecordPage struct {
	Total    int      `json:"total"`
	PageSize int      `json:"pageSize"`
	PageNum  int      `json:"pageNum"`
	Records  []Record `json:"records"`
}

type Records struct {
	Records []Record `json:"records"`
}

type RecordResponse struct {
	Success bool    `json:"success"`
	Code    int     `json:"code"`
	Message string  `json:"message"`
	Data    Records `json:"data"`
}

type PageResponse struct {
	Success bool       `json:"success"`
	Code    int        `json:"code"`
	Message string     `json:"message"`
	Data    RecordPage `json:"data"`
}

type AttachmentResponse struct {
	Success bool       `json:"success"`
	Code    int        `json:"code"`
	Message string     `json:"message"`
	Data    Attachment `json:"data"`
}

type Attachment struct {
	Token    string `json:"token,omitempty"`
	Name     string `json:"name,omitempty"`
	Size     int    `json:"size,omitempty"`
	Width    int    `json:"width,omitempty"`
	Height   int    `json:"height,omitempty"`
	MimeType string `json:"mimeType,omitempty"`
	Preview  string `json:"preview,omitempty"`
	Url      string `json:"url,omitempty"`
}

type DeleteResponse struct {
	Success bool   `json:"success"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Datasheet struct {
	datasheetId string
	request     *vikaRequest
	fieldKey    string
}

type Sort struct {
	Field string `json:"field,omitempty"`
	Order string `json:"order,omitempty"`
}

type RecordQueryParam struct {
	PageSize        int           `url:"pageSize,omitempty"`
	MaxRecords      int           `url:"maxRecord,omitempty"`
	PageNum         int           `url:"pageNum,omitempty"`
	Sort            []interface{} `url:"sort,omitempty"`
	RecordIds       []string      `url:"recordIds,omitempty"`
	ViewId          string        `url:"viewId,omitempty"`
	Fields          []string      `url:"fields,omitempty"`
	FilterByFormula string        `url:"filterByFormula,omitempty"`
	CellFormat      string        `url:"cellFormat,omitempty"`
	FieldKey        string        `url:"fieldKey,omitempty"`
}
