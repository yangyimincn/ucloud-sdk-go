//Code is generated by ucloud code generator, don't modify it by hand, it will cause undefined behaviors.
//go:generate ucloud-gen-go-api UNet UpdateFirewallAttribute

package unet

import (
	"github.com/yangyimincn/ucloud-sdk-go/ucloud/request"
	"github.com/yangyimincn/ucloud-sdk-go/ucloud/response"
)

// UpdateFirewallAttributeRequest is request schema for UpdateFirewallAttribute action
type UpdateFirewallAttributeRequest struct {
	request.CommonBase

	// [公共参数] 地域。 参见 [地域和可用区列表](../summary/regionlist.html)
	// Region *string `required:"true"`

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](../summary/get_project_list.html)
	// ProjectId *string `required:"false"`

	// 防火墙资源ID
	FWId *string `required:"true"`

	// 防火墙名称，默认为空，为空则不做修改。Name,Tag,Remark必须填写1个及以上
	Name *string `required:"false"`

	// 防火墙业务组，默认为空，为空则不做修改。Name,Tag,Remark必须填写1个及以上
	Tag *string `required:"false"`

	// 防火墙备注，默认为空，为空则不做修改。Name,Tag,Remark必须填写1个及以上
	Remark *string `required:"false"`
}

// UpdateFirewallAttributeResponse is response schema for UpdateFirewallAttribute action
type UpdateFirewallAttributeResponse struct {
	response.CommonBase
}

// NewUpdateFirewallAttributeRequest will create request of UpdateFirewallAttribute action.
func (c *UNetClient) NewUpdateFirewallAttributeRequest() *UpdateFirewallAttributeRequest {
	req := &UpdateFirewallAttributeRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

// UpdateFirewallAttribute - 更新防火墙规则
func (c *UNetClient) UpdateFirewallAttribute(req *UpdateFirewallAttributeRequest) (*UpdateFirewallAttributeResponse, error) {
	var err error
	var res UpdateFirewallAttributeResponse

	err = c.Client.InvokeAction("UpdateFirewallAttribute", req, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}
