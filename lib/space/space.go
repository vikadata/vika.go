package space

import (
	vkhttp "github.com/vikadata/vika.go/lib/common/http"
)

const spacePath = "/fusion/v1/spaces"

// NewDescribeSpacesRequest init get spaces request instance
func NewDescribeSpacesRequest() (request *DescribeSpacesRequest) {
	request = &DescribeSpacesRequest{
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
// DescribeSpaces get all user's spaces list
func (c *Space) DescribeSpaces(request *DescribeSpacesRequest) (views []*SpaceBaseInfo, err error) {
	if request == nil {
		request = NewDescribeSpacesRequest()
	}
	request.Init().SetPath(spacePath)
	request.SetHttpMethod(vkhttp.GET)
	response := newDescribeSpacesResponse()
	err = c.Send(request, response)
	if err != nil {
		return nil, err
	}
	return response.Data.Spaces, nil
}