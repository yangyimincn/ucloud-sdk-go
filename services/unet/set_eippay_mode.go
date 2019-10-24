//Code is generated by ucloud code generator, don't modify it by hand, it will cause undefined behaviors.
//go:generate ucloud-gen-go-api UNet SetEIPPayMode

package unet

import (
	"github.com/yangyimincn/ucloud-sdk-go/ucloud/request"
	"github.com/yangyimincn/ucloud-sdk-go/ucloud/response"
)

// SetEIPPayModeRequest is request schema for SetEIPPayMode action
type SetEIPPayModeRequest struct {
	request.CommonBase

	// [公共参数] 地域。 参见 [地域和可用区列表](../summary/regionlist.html)
	// Region *string `required:"true"`

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](../summary/get_project_list.html)
	// ProjectId *string `required:"false"`

	// 弹性IP的资源Id
	EIPId *string `required:"true"`

	// 计费模式. 枚举值："Traffic", 流量计费模式; "Bandwidth", 带宽计费模式
	PayMode *string `required:"true"`

	// 调整的目标带宽值, 单位Mbps. 各地域的带宽值范围如下: 流量计费[1-200],其余情况[1-800]
	Bandwidth *int `required:"true"`
}

// SetEIPPayModeResponse is response schema for SetEIPPayMode action
type SetEIPPayModeResponse struct {
	response.CommonBase
}

// NewSetEIPPayModeRequest will create request of SetEIPPayMode action.
func (c *UNetClient) NewSetEIPPayModeRequest() *SetEIPPayModeRequest {
	req := &SetEIPPayModeRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

// SetEIPPayMode - 设置弹性IP计费模式, 切换时会涉及付费/退费.
func (c *UNetClient) SetEIPPayMode(req *SetEIPPayModeRequest) (*SetEIPPayModeResponse, error) {
	var err error
	var res SetEIPPayModeResponse

	err = c.Client.InvokeAction("SetEIPPayMode", req, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}
