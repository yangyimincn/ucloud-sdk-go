//Code is generated by ucloud code generator, don't modify it by hand, it will cause undefined behaviors.
//go:generate ucloud-gen-go-api UNet ResizeShareBandwidth

package unet

import (
	"github.com/yangyimincn/ucloud-sdk-go/ucloud/request"
	"github.com/yangyimincn/ucloud-sdk-go/ucloud/response"
)

// ResizeShareBandwidthRequest is request schema for ResizeShareBandwidth action
type ResizeShareBandwidthRequest struct {
	request.CommonBase

	// [公共参数] 地域。 参见 [地域和可用区列表](../summary/regionlist.html)
	// Region *string `required:"true"`

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](../summary/get_project_list.html)
	// ProjectId *string `required:"false"`

	// 带宽值，单位为Mb，范围 [20-5000] (最大值受地域限制)
	ShareBandwidth *int `required:"true"`

	// 共享带宽的Id
	ShareBandwidthId *string `required:"true"`
}

// ResizeShareBandwidthResponse is response schema for ResizeShareBandwidth action
type ResizeShareBandwidthResponse struct {
	response.CommonBase
}

// NewResizeShareBandwidthRequest will create request of ResizeShareBandwidth action.
func (c *UNetClient) NewResizeShareBandwidthRequest() *ResizeShareBandwidthRequest {
	req := &ResizeShareBandwidthRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

// ResizeShareBandwidth - 调整共享带宽的带宽值
func (c *UNetClient) ResizeShareBandwidth(req *ResizeShareBandwidthRequest) (*ResizeShareBandwidthResponse, error) {
	var err error
	var res ResizeShareBandwidthResponse

	err = c.Client.InvokeAction("ResizeShareBandwidth", req, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}
