package space

import (
	"fmt"
	vkhttp "github.com/vikadata/vika.go/lib/common/http"
)

const spaceListPath = "/fusion/v1/spaces"
const nodeListPath = "/fusion/v1/spaces/%s/nodes"
const nodeDetailPath = "/fusion/v1/spaces/%s/nodes/%s"

// NewDescribeSpacesRequest init get spaces request instance
func NewDescribeSpacesRequest() (request *DescribeSpacesRequest) {
	request = &DescribeSpacesRequest{
		BaseRequest: &vkhttp.BaseRequest{},
	}
	return
}

func NewDescribeNodesRequest() (request *DescribeNodesRequest) {
	request = &DescribeNodesRequest{
		BaseRequest: &vkhttp.BaseRequest{},
	}
	return
}

func NewDescribeNodeRequest() (request *DescribeNodeRequest) {
	request = &DescribeNodeRequest{
		BaseRequest: &vkhttp.BaseRequest{},
	}
	return
}

func newDescribeSpacesResponse() (response *DescribeSpacesResponse) {
	response = &DescribeSpacesResponse{
		BaseResponse: &vkhttp.BaseResponse{},
	}
	return
}

func newDescribeNodesResponse() (response *DescribeNodesResponse) {
	response = &DescribeNodesResponse{
		BaseResponse: &vkhttp.BaseResponse{},
	}
	return
}

func newDescribeNodeResponse() (response *DescribeNodeResponse) {
	response = &DescribeNodeResponse{
		BaseResponse: &vkhttp.BaseResponse{},
	}
	return
}

// DescribeSpaces get all user's spaces list
func (c *Space) DescribeSpaces(request *DescribeSpacesRequest) (spaces []*SpaceBaseInfo, err error) {
	if request == nil {
		request = NewDescribeSpacesRequest()
	}
	request.Init().SetPath(spaceListPath)
	request.SetHttpMethod(vkhttp.GET)
	response := newDescribeSpacesResponse()
	err = c.Send(request, response)
	if err != nil {
		return nil, err
	}
	return response.Data.Spaces, nil
}

// DescribeNodes get all space nodes
func (c *Space) DescribeNodes(request *DescribeNodesRequest) (node []*NodeBaseInfo, err error) {
	if request == nil {
		request = NewDescribeNodesRequest()
	}
	request.Init().SetPath(fmt.Sprintf(nodeListPath, c.SpaceId))
	request.SetHttpMethod(vkhttp.GET)
	response := newDescribeNodesResponse()
	err = c.Send(request, response)
	if err != nil {
		return nil, err
	}
	return response.Data.Nodes, nil
}

// DescribeNode get node detail
func (c *Space) DescribeNode(request *DescribeNodeRequest) (node *NodeDetail, err error) {
	if request == nil {
		request = NewDescribeNodeRequest()
	}
	request.Init().SetPath(fmt.Sprintf(nodeDetailPath, c.SpaceId, *request.NodeId))
	request.SetHttpMethod(vkhttp.GET)
	response := newDescribeNodeResponse()
	err = c.Send(request, response)
	if err != nil {
		return nil, err
	}
	return response.Data, nil
}
