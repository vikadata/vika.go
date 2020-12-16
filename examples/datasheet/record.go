package main

import (
	"encoding/json"
	"fmt"
	"github.com/vikadata/vika.go/lib/common"
	vkerror "github.com/vikadata/vika.go/lib/common/error"
	"github.com/vikadata/vika.go/lib/common/profile"
	vika "github.com/vikadata/vika.go/lib/datasheet"
	"os"
)

func main() {
	// 必要步骤：
	// 实例化一个认证对象，入参需要传入vika开发者token。
	// 这里采用的是从环境变量读取的方式，需要在环境变量中先设置这个值。
	// 你也可以直接在代码中写死token，但是小心不要将代码复制、上传或者分享给他人，
	// 以免泄露token危及你的财产安全。
	credential := common.NewCredential(
		os.Getenv("VIKA_TOKEN"),
	)
	cpf := profile.NewClientProfile()
	// 上传图片
	attachment1, err := uploadImage(credential, cpf)
	if _, ok := err.(*vkerror.VikaSDKError); ok {
		fmt.Printf("uploadImage:An API error has returned: %s", err)
		return
	}
	// 非SDK异常，直接失败。实际代码中可以加入其他的处理。
	if err != nil {
		panic(err)
	}
	json1, _ := json.Marshal(attachment1)
	fmt.Printf("上传图片结果：%s\n", json1)
}

func uploadImage(credential *common.Credential, cpf *profile.ClientProfile) (*vika.Attachment, error) {
	cpf.Upload = true
	datasheet, _ := vika.NewDatasheet(credential, os.Getenv("VIKA_DATASHEET_ID"), cpf)
	request := vika.NewUploadRequest()
	// 如果不设置Domain使用默认的域名
	request.FilePath = "图片完整路径"
	return datasheet.UploadFile(request)
}
