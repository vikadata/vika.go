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

type NodeType string

const (
	NodeType_Datasheet NodeType = "Datasheet"
	NodeType_Folder    NodeType = "Folder"
	NodeType_Form      NodeType = "Form"
	NodeType_Dashboard NodeType = "Dashboard"
)

// SpaceBaseInfo describe the property of the space
type SpaceBaseInfo struct {
	// Id spaceId
	Id *string `json:"id,omitempty" name:"id"`
	// Name space name
	Name *string `json:"name,omitempty" name:"name"`
	// IsAdmin the manager of the space
	IsAdmin *bool `json:"isAdmin,omitempty" name:"isAdmin"`
}

// NodeBaseInfo node base info
type NodeBaseInfo struct {
	// Id node id
	Id *string `json:"id,omitempty" name:"id"`
	// Name node name
	Name *string `json:"name,omitempty" name:"name"`
	// Type type of the node
	Type *NodeType `json:"type,omitempty" name:"type"`
	// Icon node icon
	Icon *string `json:"icon,omitempty" name:"icon"`
	// IsFav is the node in the favorite folder
	IsFav *bool `json:"isFav,omitempty" name:"isFav"`
}

// NodeDetail node detail
type NodeDetail struct {
	NodeBaseInfo
	Children []*NodeBaseInfo `json:"children,omitempty" name:"children"`
}

// DescribeSpacesRequest space list request
type DescribeSpacesRequest struct {
	*vkhttp.BaseRequest
}

// DescribeNodesRequest node list request
type DescribeNodesRequest struct {
	*vkhttp.BaseRequest
}

// DescribeNodeRequest node list request
type DescribeNodeRequest struct {
	*vkhttp.BaseRequest
	NodeId *string
}

// SpaceResponse space list response
type SpaceResponse struct {
	Spaces []*SpaceBaseInfo `json:"spaces"`
}

// NodeResponse node list response
type NodeResponse struct {
	Nodes []*NodeBaseInfo `json:"nodes"`
}

// DescribeSpacesResponse space list response data
type DescribeSpacesResponse struct {
	*vkhttp.BaseResponse
	// api返回数据
	Data *SpaceResponse `json:"data"`
}

// DescribeNodesResponse node list response data
type DescribeNodesResponse struct {
	*vkhttp.BaseResponse
	// api返回数据
	Data *NodeResponse `json:"data"`
}

// DescribeNodeResponse node detail response
type DescribeNodeResponse struct {
	*vkhttp.BaseResponse
	// api返回数据
	Data *NodeDetail `json:"data"`
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
