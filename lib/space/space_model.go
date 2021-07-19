package space

import (
	"github.com/vikadata/vika.go/lib/common"
	vkhttp "github.com/vikadata/vika.go/lib/common/http"
	"github.com/vikadata/vika.go/lib/common/profile"
)

type Space struct {
	common.Client
	SpaceId string
}

// SpaceBaseInfo describe the property of the vika space
type SpaceBaseInfo struct {
	// Id spaceId
	Id *string `json:"id,omitempty" name:"id"`
	// Name space name
	Name *string `json:"name,omitempty" name:"name"`
	// IsAdmin the manager of the space
	IsAdmin *bool `json:"isAdmin,omitempty" name:"isAdmin"`
}

type DescribeSpacesRequest struct {
	*vkhttp.BaseRequest
}

type SpaceResponse struct {
	Spaces []*SpaceBaseInfo `json:"spaces"`
}

type DescribeSpacesResponse struct {
	*vkhttp.BaseResponse
	// api返回数据
	Data *SpaceResponse `json:"data"`
}

// NewSpace init space instance
func NewSpace(credential *common.Credential, spaceId string, clientProfile *profile.ClientProfile) (space *Space, err error) {
	space = &Space{}
	if spaceId != "" {
		space.SpaceId = spaceId
	}
	space.Init().WithCredential(credential).WithProfile(clientProfile)
	return
}


