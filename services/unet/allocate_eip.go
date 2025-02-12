//Code is generated by ucloud code generator, don't modify it by hand, it will cause undefined behaviors.
//go:generate ucloud-gen-go-api UNet AllocateEIP

package unet

import (
	"github.com/yangyimincn/ucloud-sdk-go/ucloud/request"
	"github.com/yangyimincn/ucloud-sdk-go/ucloud/response"
)

// AllocateEIPRequest is request schema for AllocateEIP action
type AllocateEIPRequest struct {
	request.CommonBase

	// [公共参数] 地域。
	// Region *string `required:"true"`

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。
	// ProjectId *string `required:"false"`

	// 弹性IP的线路如下: 国际: International BGP: Bgp  各地域允许的线路参数如下:  cn-sh1: Bgp cn-sh2: Bgp cn-gd: Bgp cn-bj1: Bgp cn-bj2: Bgp hk: International us-ca: International th-bkk: International  kr-seoul:International  us-ws:International  ge-fra:International  sg:International  tw-kh:International.其他海外线路均为 International
	OperatorName *string `required:"true"`

	// 弹性IP的外网带宽, 单位为Mbps. 共享带宽模式必须指定0M带宽, 非共享带宽模式必须指定非0Mbps带宽. 各地域非共享带宽的带宽范围如下： 流量计费[1-200]，带宽计费[1-800]
	Bandwidth *int `required:"true"`

	// 业务组名称, 默认为 "Default"
	Tag *string `required:"false"`

	// 付费方式, 枚举值为: Year, 按年付费; Month, 按月付费; Dynamic, 按需付费(需开启权限); Trial, 试用(需开启权限) 默认为按月付费
	ChargeType *string `required:"false"`

	// 购买时长, 默认: 1
	Quantity *int `required:"false"`

	// 弹性IP的计费模式. 枚举值: "Traffic", 流量计费; "Bandwidth", 带宽计费; "ShareBandwidth",共享带宽模式. 默认为 "Bandwidth".
	PayMode *string `required:"false"`

	// 绑定的共享带宽Id，仅当PayMode为ShareBandwidth时有效
	ShareBandwidthId *string `required:"false"`

	// 弹性IP的名称, 默认为 "EIP"
	Name *string `required:"false"`

	// 弹性IP的备注, 默认为空
	Remark *string `required:"false"`

	// 代金券ID, 默认不使用
	CouponId *string `required:"false"`
}

// AllocateEIPResponse is response schema for AllocateEIP action
type AllocateEIPResponse struct {
	response.CommonBase

	// 申请到的EIP资源详情 参见 UnetAllocateEIPSet
	EIPSet []UnetAllocateEIPSet
}

// NewAllocateEIPRequest will create request of AllocateEIP action.
func (c *UNetClient) NewAllocateEIPRequest() *AllocateEIPRequest {
	req := &AllocateEIPRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(false)
	return req
}

// AllocateEIP - 根据提供信息, 申请弹性IP
func (c *UNetClient) AllocateEIP(req *AllocateEIPRequest) (*AllocateEIPResponse, error) {
	var err error
	var res AllocateEIPResponse

	err = c.Client.InvokeAction("AllocateEIP", req, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}
