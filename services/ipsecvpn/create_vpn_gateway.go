// Code is generated by ucloud-model, DO NOT EDIT IT.

package ipsecvpn

import (
	"github.com/yangyimincn/ucloud-sdk-go/ucloud/request"
	"github.com/yangyimincn/ucloud-sdk-go/ucloud/response"
)

// CreateVPNGatewayRequest is request schema for CreateVPNGateway action
type CreateVPNGatewayRequest struct {
	request.CommonBase

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](../summary/get_project_list.html)
	// ProjectId *string `required:"true"`

	// [公共参数] 地域。 参见 [地域和可用区列表](../summary/regionlist.html)
	// Region *string `required:"true"`

	// 业务组ID
	BusinessId *string `required:"false"`

	// 付费方式, 枚举值为: Year, 按年付费; Month, 按月付费；Dynamic, 按需付费(需开启权限)；Trial, 试用(需开启权限)；默认为按月付费
	ChargeType *string `required:"false"`

	// 代金券ID, 默认不使用
	CouponId *string `required:"false"`

	// 若要绑定EIP，在此填上EIP的资源ID
	EIPId *string `required:"false"`

	// 购买的VPN网关规格，枚举值为: Standard, 标准型; Enhanced, 增强型
	Grade *string `required:"true"`

	// 购买时长, 默认: 1
	Quantity *int `required:"false"`

	// 备注，默认为空
	Remark *string `required:"false"`

	// 业务组名称，默认为 "Default"
	Tag *string `required:"false"`

	// 新建VPN网关所属VPC的资源ID
	VPCId *string `required:"true"`

	// 新建VPN网关名称
	VPNGatewayName *string `required:"true"`
}

// CreateVPNGatewayResponse is response schema for CreateVPNGateway action
type CreateVPNGatewayResponse struct {
	response.CommonBase

	// 新建VPN网关的资源ID
	VPNGatewayId string
}

// NewCreateVPNGatewayRequest will create request of CreateVPNGateway action.
func (c *IPSecVPNClient) NewCreateVPNGatewayRequest() *CreateVPNGatewayRequest {
	req := &CreateVPNGatewayRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(false)
	return req
}

// CreateVPNGateway - 创建VPN网关
func (c *IPSecVPNClient) CreateVPNGateway(req *CreateVPNGatewayRequest) (*CreateVPNGatewayResponse, error) {
	var err error
	var res CreateVPNGatewayResponse

	reqCopier := *req

	err = c.Client.InvokeAction("CreateVPNGateway", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}
