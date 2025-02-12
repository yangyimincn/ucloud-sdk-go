//Code is generated by ucloud code generator, don't modify it by hand, it will cause undefined behaviors.
//go:generate ucloud-gen-go-api UMem ResizeUDredisSpace

package umem

import (
	"github.com/yangyimincn/ucloud-sdk-go/ucloud/request"
	"github.com/yangyimincn/ucloud-sdk-go/ucloud/response"
)

// ResizeUDredisSpaceRequest is request schema for ResizeUDredisSpace action
type ResizeUDredisSpaceRequest struct {
	request.CommonBase

	// [公共参数] 地域。 参见 [地域和可用区列表](../summary/regionlist.html)
	// Region *string `required:"true"`

	// [公共参数] 可用区。参见 [可用区列表](../summary/regionlist.html)
	// Zone *string `required:"false"`

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](../summary/get_project_list.html)
	// ProjectId *string `required:"false"`

	// 高性能UMem 内存空间Id
	SpaceId *string `required:"true"`

	// 内存大小, 单位:GB (需要大于原size,<= 1024)
	Size *int `required:"true"`

	// 使用的代金券Id
	CouponId *string `required:"false"`
}

// ResizeUDredisSpaceResponse is response schema for ResizeUDredisSpace action
type ResizeUDredisSpaceResponse struct {
	response.CommonBase
}

// NewResizeUDredisSpaceRequest will create request of ResizeUDredisSpace action.
func (c *UMemClient) NewResizeUDredisSpaceRequest() *ResizeUDredisSpaceRequest {
	req := &ResizeUDredisSpaceRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

// ResizeUDredisSpace - 调整内存空间容量
func (c *UMemClient) ResizeUDredisSpace(req *ResizeUDredisSpaceRequest) (*ResizeUDredisSpaceResponse, error) {
	var err error
	var res ResizeUDredisSpaceResponse

	err = c.Client.InvokeAction("ResizeUDredisSpace", req, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}
