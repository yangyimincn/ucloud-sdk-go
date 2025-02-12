// Code is generated by ucloud-model, DO NOT EDIT IT.

package ipsecvpn

import (
	"github.com/yangyimincn/ucloud-sdk-go/ucloud/request"
	"github.com/yangyimincn/ucloud-sdk-go/ucloud/response"
)

// UpdateVPNTunnelAttributeRequest is request schema for UpdateVPNTunnelAttribute action
type UpdateVPNTunnelAttributeRequest struct {
	request.CommonBase

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](../summary/get_project_list.html)
	// ProjectId *string `required:"true"`

	// [公共参数] 地域。 参见 [地域和可用区列表](../summary/regionlist.html)
	// Region *string `required:"true"`

	// IKE协商过程中使用的认证算法
	IKEAuthenticationAlgorithm *string `required:"false"`

	// IKE协商过程中使用的DH组
	IKEDhGroup *string `required:"false"`

	// IKE协商过程中使用的加密算法
	IKEEncryptionAlgorithm *string `required:"false"`

	// IKE协商过程中使用的模式，可选“主动模式”与“野蛮模式”。IKEV2不使用该参数。
	IKEExchangeMode *string `required:"false"`

	// 本端标识。不填时默认使用之前的参数，结合IKEversion进行校验，IKEV2时不能为auto。
	IKELocalId *string `required:"false"`

	// 预共享密钥
	IKEPreSharedKey *string `required:"false"`

	// 客户端标识。不填时默认使用之前的参数，结合IKEversion进行校验，IKEV2时不能为auto。
	IKERemoteId *string `required:"false"`

	// IKE中SA的生存时间
	IKESALifetime *string `required:"false"`

	// 枚举值："IKE V1","IKE V2"
	IKEVersion *string `required:"false"`

	// IPSec隧道中使用的认证算法
	IPSecAuthenticationAlgorithm *string `required:"false"`

	// IPSec隧道中使用的加密算法
	IPSecEncryptionAlgorithm *string `required:"false"`

	// 指定VPN连接的本地子网的id，用逗号分隔
	IPSecLocalSubnetIds []string `required:"false"`

	// IPSec中的PFS是否开启
	IPSecPFSDhGroup *string `required:"false"`

	// 使用的安全协议，ESP或AH
	IPSecProtocol *string `required:"false"`

	// 指定VPN连接的客户网段，用逗号分隔
	IPSecRemoteSubnets []string `required:"false"`

	// IPSec中SA的生存时间
	IPSecSALifetime *string `required:"false"`

	// IPSec中SA的生存时间（以字节计）
	IPSecSALifetimeBytes *string `required:"false"`

	// VPN隧道的资源ID
	VPNTunnelId *string `required:"true"`
}

// UpdateVPNTunnelAttributeResponse is response schema for UpdateVPNTunnelAttribute action
type UpdateVPNTunnelAttributeResponse struct {
	response.CommonBase
}

// NewUpdateVPNTunnelAttributeRequest will create request of UpdateVPNTunnelAttribute action.
func (c *IPSecVPNClient) NewUpdateVPNTunnelAttributeRequest() *UpdateVPNTunnelAttributeRequest {
	req := &UpdateVPNTunnelAttributeRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

// UpdateVPNTunnelAttribute - 更新VPN隧道属性
func (c *IPSecVPNClient) UpdateVPNTunnelAttribute(req *UpdateVPNTunnelAttributeRequest) (*UpdateVPNTunnelAttributeResponse, error) {
	var err error
	var res UpdateVPNTunnelAttributeResponse

	reqCopier := *req

	err = c.Client.InvokeAction("UpdateVPNTunnelAttribute", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}
