//Code is generated by ucloud code generator, don't modify it by hand, it will cause undefined behaviors.
//go:generate ucloud-gen-go-api UNet UpdateFirewall

package unet

import (
	"github.com/yangyimincn/ucloud-sdk-go/ucloud/request"
	"github.com/yangyimincn/ucloud-sdk-go/ucloud/response"
)

// UpdateFirewallRequest is request schema for UpdateFirewall action
type UpdateFirewallRequest struct {
	request.CommonBase

	// [公共参数] 地域。 参见 [地域和可用区列表](../summary/regionlist.html)
	// Region *string `required:"true"`

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](../summary/get_project_list.html)
	// ProjectId *string `required:"false"`

	// 防火墙资源ID
	FWId *string `required:"true"`

	// 防火墙规则，例如：TCP|22|192.168.1.1/22|DROP|LOW|禁用22端口，第一个参数代表协议：第二个参数代表端口号，第三个参数为ip，第四个参数为ACCEPT（接受）和DROP（拒绝），第五个参数优先级：HIGH（高），MEDIUM（中），LOW（低），第六个参数为该条规则的自定义备注
	Rule []string `required:"true"`
}

// UpdateFirewallResponse is response schema for UpdateFirewall action
type UpdateFirewallResponse struct {
	response.CommonBase

	// 防火墙id
	FWId string
}

// NewUpdateFirewallRequest will create request of UpdateFirewall action.
func (c *UNetClient) NewUpdateFirewallRequest() *UpdateFirewallRequest {
	req := &UpdateFirewallRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

// UpdateFirewall - 更新防火墙规则
func (c *UNetClient) UpdateFirewall(req *UpdateFirewallRequest) (*UpdateFirewallResponse, error) {
	var err error
	var res UpdateFirewallResponse

	err = c.Client.InvokeAction("UpdateFirewall", req, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}
