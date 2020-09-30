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
	"bytes"
	json2 "encoding/json"
	"errors"
	"fmt"
	"github.com/google/go-querystring/query"
	"github.com/h2non/filetype"
	"github.com/valyala/fasthttp"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/textproto"
	"os"
	"path"
	"reflect"
	"strings"
	"time"
)

const recordPath = "/fusion/v1/datasheets/%s/records"
const attachPath = "/fusion/v1/datasheets/%s/attachments"

type vikaRequest struct {
	host        string
	timeout     time.Duration
	datasheetId string
	token       string
}
type vikaHeader struct {
	method      string
	token       string
	contentType string
	timeout     time.Duration
}

type vikaUri struct {
	scheme string
	path   string
	host   string
}

func newVikaRequest(vika Vika, datasheetId string) *vikaRequest {
	request := &vikaRequest{}
	request.host = vika.Host
	request.timeout = vika.RequestTimeout
	request.datasheetId = datasheetId
	request.token = vika.Token
	return request
}

func (cli *vikaRequest) getRecords(params RecordQueryParam) (*PageResponse, error) {
	if params.Sort != nil {
		for key, sort := range params.Sort {
			buf, _ := json2.Marshal(sort)
			params.Sort[key] = string(buf)
		}
	}
	values, _ := query.Values(params)
	res, err := request(vikaHeader{
		fasthttp.MethodGet, cli.token, "application/x-www-form-urlencoded", cli.timeout},
		recordUri(cli.host, cli.datasheetId, recordPath), values.Encode(), nil)
	if err != nil {
		return nil, err
	}
	result := &PageResponse{}
	jsonErr := json2.Unmarshal(res, result)
	if jsonErr != nil {
		return nil, jsonErr
	}
	if result.Success != true {
		return nil, errors.New(result.Message)
	}
	return result, nil
}

func (cli *vikaRequest) addRecords(records []NewRecord, fieldKey string) (*RecordResponse, error) {
	form, _ := newVikaRequestSerializer().withFields().UseCamelCase().TransformArray(records)
	res, err := request(vikaHeader{
		fasthttp.MethodPost, cli.token, "application/json", cli.timeout},
		recordUri(cli.host, cli.datasheetId, recordPath), "", bodyByte(form, fieldKey))
	if err != nil {
		return nil, err
	}
	result := &RecordResponse{}
	jsonErr := json2.Unmarshal(res, result)
	if jsonErr != nil {
		return nil, jsonErr
	}
	if result.Success != true {
		return nil, errors.New(result.Message)
	}
	return result, nil
}

func (cli *vikaRequest) updateRecords(records []Record, fieldKey string) (*RecordResponse, error) {
	form, _ := newVikaRequestSerializer().withRecords().UseCamelCase().TransformArray(records)
	res, err := request(vikaHeader{
		fasthttp.MethodPatch, cli.token, "application/json", cli.timeout},
		recordUri(cli.host, cli.datasheetId, recordPath), "", bodyByte(form, fieldKey))
	if err != nil {
		return nil, err
	}
	result := &RecordResponse{}
	jsonErr := json2.Unmarshal(res, result)
	if jsonErr != nil {
		return nil, jsonErr
	}
	if result.Success != true {
		return nil, errors.New(result.Message)
	}
	return result, nil
}

func (cli *vikaRequest) delRecords(ids []string) (*DeleteResponse, error) {
	values, _ := query.Values(RecordDeleteParam{RecordIds: ids})
	res, err := request(vikaHeader{
		fasthttp.MethodDelete, cli.token, "application/json", cli.timeout},
		recordUri(cli.host, cli.datasheetId, recordPath), values.Encode(), recordIdByte(ids))
	if err != nil {
		return nil, err
	}
	result := &DeleteResponse{}
	jsonErr := json2.Unmarshal(res, result)
	if jsonErr != nil {
		return nil, jsonErr
	}
	if result.Success != true {
		return nil, errors.New(result.Message)
	}
	return result, nil
}

func (cli *vikaRequest) uploadAsset(filePath string) (*AttachmentResponse, error) {
	body, contentType, err := fileBuffer(filePath)
	if err != nil {
		return nil, err
	}
	res, err := request(vikaHeader{fasthttp.MethodPost, cli.token, contentType, cli.timeout},
		recordUri(cli.host, cli.datasheetId, attachPath), "", body)
	if err != nil {
		return nil, err
	}
	result := &AttachmentResponse{}
	jsonErr := json2.Unmarshal(res, result)
	if jsonErr != nil {
		return nil, jsonErr
	}
	if result.Success != true {
		return nil, errors.New(result.Message)
	}
	return result, nil
}

func request(header vikaHeader, uri vikaUri, query string, body []byte) ([]byte, error) {
	req, res := fasthttp.AcquireRequest(), fasthttp.AcquireResponse()
	defer func() {
		fasthttp.ReleaseRequest(req)
		fasthttp.ReleaseResponse(res)
	}()
	// header
	setHeader(req, header)
	// uri
	setUri(req, uri)
	// query
	if !reflect.ValueOf(query).IsZero() {
		req.URI().SetQueryString(query)
	}
	// body
	if body != nil {
		req.SetBody(body)
	}
	err := fasthttp.DoTimeout(req, res, header.timeout)
	if err != nil {
		return nil, err
	}
	return res.Body(), err
}

func setHeader(req *fasthttp.Request, header vikaHeader) {
	// header
	req.Header.SetUserAgent("vika-go")
	req.Header.Set(fasthttp.HeaderAuthorization, "Bearer "+header.token)
	req.Header.SetMethod(header.method)
	req.Header.SetContentType(header.contentType)
}

func setUri(req *fasthttp.Request, uri vikaUri) {
	// uri
	req.URI().SetHost(uri.host)
	req.URI().SetPath(uri.path)
	req.URI().SetScheme(uri.scheme)
}

func fileBuffer(filePath string) ([]byte, string, error) {
	//新建一个缓冲，用于存放文件内容
	bodyBuffer := &bytes.Buffer{}
	//创建一个multipart文件写入器，方便按照http规定格式写入内容
	bodyWriter := multipart.NewWriter(bodyBuffer)
	fileWriter, err := createFormFile(bodyWriter, "file", filePath)
	if err != nil {
		return nil, "", err
	}
	file, err := os.Open(filePath)
	if err != nil {
		return nil, "", err
	}
	//不要忘记关闭打开的文件
	defer file.Close()
	_, err = io.Copy(fileWriter, file)
	if err != nil {
		return nil, "", nil
	}
	//关闭bodyWriter停止写入数据
	_ = bodyWriter.Close()
	contentType := bodyWriter.FormDataContentType()
	return bodyBuffer.Bytes(), contentType, nil
}

func recordUri(host string, datasheetId string, path string) vikaUri {
	arr := strings.Split(host, "://")
	return vikaUri{
		host:   arr[1],
		scheme: arr[0],
		path:   fmt.Sprintf(path, datasheetId),
	}
}

func createFormFile(bodyWriter *multipart.Writer, fieldname, filePath string) (io.Writer, error) {
	buf, _ := ioutil.ReadFile(filePath)
	kind, _ := filetype.Match(buf)
	h := make(textproto.MIMEHeader)
	h.Set("Content-Disposition",
		fmt.Sprintf(`form-data; name="%s"; filename="%s"`,
			fieldname, path.Base(filePath)))
	h.Set("Content-Type", kind.MIME.Value)
	return bodyWriter.CreatePart(h)
}
