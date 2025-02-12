//Code is generated by ucloud code generator, don't modify it by hand, it will cause undefined behaviors.
//go:generate ucloud-gen-go-api UNet DisassociateEIPWithShareBandwidth

package unet

import (
	"github.com/yangyimincn/ucloud-sdk-go/ucloud/request"
	"github.com/yangyimincn/ucloud-sdk-go/ucloud/response"
)

// DisassociateEIPWithShareBandwidthRequest is request schema for DisassociateEIPWithShareBandwidth action
type DisassociateEIPWithShareBandwidthRequest struct {
	request.CommonBase

	// [公共参数] 地域。 参见 [地域和可用区列表](../summary/regionlist.html)
	// Region *string `required:"true"`

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](../summary/get_project_list.html)
	// ProjectId *string `required:"false"`

	// 共享带宽ID
	ShareBandwidthId *string `required:"true"`

	// 移出共享带宽后，EIP的外网带宽, 单位为Mbps. 各地域带宽范围如下：  流量计费[1-200],带宽计费[1-800]
	Bandwidth *int `required:"true"`

	// EIP的资源Id；默认移出该共享带宽下所有的EIP
	EIPIds []string `required:"false"`

	// 移出共享带宽后，EIP的计费模式. 枚举值: "Traffic", 流量计费; "Bandwidth", 带宽计费;  默认为 "Bandwidth".
	PayMode *string `required:"false"`
}

// DisassociateEIPWithShareBandwidthResponse is response schema for DisassociateEIPWithShareBandwidth action
type DisassociateEIPWithShareBandwidthResponse struct {
	response.CommonBase
}

// NewDisassociateEIPWithShareBandwidthRequest will create request of DisassociateEIPWithShareBandwidth action.
func (c *UNetClient) NewDisassociateEIPWithShareBandwidthRequest() *DisassociateEIPWithShareBandwidthRequest {
	req := &DisassociateEIPWithShareBandwidthRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

// DisassociateEIPWithShareBandwidth - 将EIP移出共享带宽
func (c *UNetClient) DisassociateEIPWithShareBandwidth(req *DisassociateEIPWithShareBandwidthRequest) (*DisassociateEIPWithShareBandwidthResponse, error) {
	var err error
	var res DisassociateEIPWithShareBandwidthResponse

	err = c.Client.InvokeAction("DisassociateEIPWithShareBandwidth", req, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}
