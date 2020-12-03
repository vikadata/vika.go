package common

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/h2non/filetype"
	vkerror "github.com/vikadata/vika.go/lib/common/error"
	vkhttp "github.com/vikadata/vika.go/lib/common/http"
	"github.com/vikadata/vika.go/lib/common/profile"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"net/http/httputil"
	"net/textproto"
	"os"
	"path"
	"strings"
	"time"
)

type Client struct {
	httpClient  *http.Client
	httpProfile *profile.HttpProfile
	profile     *profile.ClientProfile
	credential  *Credential
	debug       bool
}

func (c *Client) Init() *Client {
	c.httpClient = &http.Client{}
	c.debug = false
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	return c
}

func (c *Client) WithCredential(cred *Credential) *Client {
	c.credential = cred
	return c
}

func (c *Client) WithProfile(clientProfile *profile.ClientProfile) *Client {
	c.profile = clientProfile
	c.httpProfile = clientProfile.HttpProfile
	c.httpClient.Timeout = time.Duration(c.httpProfile.ReqTimeout) * time.Second
	c.debug = clientProfile.Debug
	return c
}

func (c *Client) WithDebug(flag bool) *Client {
	c.debug = flag
	return c
}

func FileBuffer(filePath string) ([]byte, string, error) {
	//新建一个缓冲，用于存放文件内容
	bodyBuffer := &bytes.Buffer{}
	//创建一个multipart文件写入器，方便按照http规定格式写入内容
	bodyWriter := multipart.NewWriter(bodyBuffer)
	fileWriter, err := createFormFile(bodyWriter, "file", filePath)
	if err != nil {
		msg := fmt.Sprintf("Fail to get response because %s", err)
		return nil, "", vkerror.NewVikaSDKError(500, msg, "ClientError.MultipartError")
	}
	file, err := os.Open(filePath)
	if err != nil {
		msg := fmt.Sprintf("Fail to get response because %s", err)
		return nil, "", vkerror.NewVikaSDKError(500, msg, "ClientError.FileReadError")
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

func (c *Client) Send(request vkhttp.Request, response vkhttp.Response) (err error) {
	if request.GetScheme() == "" {
		request.SetScheme(c.httpProfile.Scheme)
	}

	if request.GetDomain() == "" {
		domain := c.httpProfile.Domain
		if domain == "" {
			domain = vkhttp.Domain
		}
		request.SetDomain(domain)
	}

	if request.GetHttpMethod() == "" {
		request.SetHttpMethod(c.httpProfile.ReqMethod)
	}
	return c.sendWithToken(request, response)
}

func (c *Client) sendWithToken(request vkhttp.Request, response vkhttp.Response) (err error) {
	headers := map[string]string{
		"User-Agent": "lib-go",
	}
	if c.credential.Token != "" {
		headers["Authorization"] = "Bearer " + c.credential.Token
	}
	headers["Content-Type"] = request.GetContentType()

	// start process

	// build canonical request string
	httpRequestMethod := request.GetHttpMethod()
	canonicalQueryString := ""
	if httpRequestMethod == vkhttp.GET || httpRequestMethod == vkhttp.DELETE {
		err = vkhttp.ConstructParams(request)
		if err != nil {
			return err
		}
		params := make(map[string]string)
		for key, value := range request.GetParams() {
			params[key] = value
		}
		canonicalQueryString = vkhttp.GetUrlQueriesEncoded(params)
	}
	requestPayload := ""
	if httpRequestMethod == vkhttp.POST || httpRequestMethod == vkhttp.PATCH {
		if c.profile.Upload {
			requestPayload = string(request.GetFile())
		} else {
			b, err := json.Marshal(request)
			if err != nil {
				return err
			}
			requestPayload = string(b)
		}
	}

	url := request.GetScheme() + "://" + request.GetDomain() + request.GetPath()
	if canonicalQueryString != "" {
		url = url + "?" + canonicalQueryString
	}
	httpRequest, err := http.NewRequest(httpRequestMethod, url, strings.NewReader(requestPayload))
	if err != nil {
		return err
	}
	for k, v := range headers {
		httpRequest.Header[k] = []string{v}
	}
	if c.debug {
		outbytes, err := httputil.DumpRequest(httpRequest, true)
		if err != nil {
			log.Printf("[ERROR] dump request failed because %s", err)
			return err
		}
		log.Printf("[DEBUG] http request = %s", outbytes)
	}
	httpResponse, err := c.httpClient.Do(httpRequest)
	if err != nil {
		msg := fmt.Sprintf("Fail to get response because %s", err)
		return vkerror.NewVikaSDKError(500, msg, "ClientError.NetworkError")
	}
	err = vkhttp.ParseFromHttpResponse(httpResponse, response)
	return err
}

// 重写文件类型
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
