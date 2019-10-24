//Code is generated by ucloud code generator, don't modify it by hand, it will cause undefined behaviors.
//go:generate ucloud-gen-go-api UMem DeleteUMemSpace

package umem

import (
	"github.com/yangyimincn/ucloud-sdk-go/ucloud/request"
	"github.com/yangyimincn/ucloud-sdk-go/ucloud/response"
)

// DeleteUMemSpaceRequest is request schema for DeleteUMemSpace action
type DeleteUMemSpaceRequest struct {
	request.CommonBase

	// [公共参数] 地域。 参见 [地域和可用区列表](../summary/regionlist.html)
	// Region *string `required:"true"`

	// [公共参数] 可用区。参见 [可用区列表](../summary/regionlist.html)
	// Zone *string `required:"false"`

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](../summary/get_project_list.html)
	// ProjectId *string `required:"false"`

	// UMem内存空间ID
	SpaceId *string `required:"true"`
}

// DeleteUMemSpaceResponse is response schema for DeleteUMemSpace action
type DeleteUMemSpaceResponse struct {
	response.CommonBase
}

// NewDeleteUMemSpaceRequest will create request of DeleteUMemSpace action.
func (c *UMemClient) NewDeleteUMemSpaceRequest() *DeleteUMemSpaceRequest {
	req := &DeleteUMemSpaceRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

// DeleteUMemSpace - 删除UMem内存空间
func (c *UMemClient) DeleteUMemSpace(req *DeleteUMemSpaceRequest) (*DeleteUMemSpaceResponse, error) {
	var err error
	var res DeleteUMemSpaceResponse

	err = c.Client.InvokeAction("DeleteUMemSpace", req, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}
