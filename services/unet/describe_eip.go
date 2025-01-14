//Code is generated by ucloud code generator, don't modify it by hand, it will cause undefined behaviors.
//go:generate ucloud-gen-go-api UNet DescribeEIP

package unet

import (
	"github.com/yangyimincn/ucloud-sdk-go/ucloud/request"
	"github.com/yangyimincn/ucloud-sdk-go/ucloud/response"
)

// DescribeEIPRequest is request schema for DescribeEIP action
type DescribeEIPRequest struct {
	request.CommonBase

	// [公共参数] 地域
	// Region *string `required:"true"`

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写
	// ProjectId *string `required:"false"`

	// 弹性IP的资源ID如果为空, 则返回当前 Region中符合条件的的所有EIP
	EIPIds []string `required:"false"`

	// 数据偏移量, 默认为0
	Offset *int `required:"false"`

	// 数据分页值, 默认为20
	Limit *int `required:"false"`
}

// DescribeEIPResponse is response schema for DescribeEIP action
type DescribeEIPResponse struct {
	response.CommonBase

	// 满足条件的弹性IP总数
	TotalCount int

	// 满足条件的弹性IP带宽总和, 单位Mbps
	TotalBandwidth int

	// 弹性IP列表, 每项参数详见 UnetEIPSet
	EIPSet []UnetEIPSet
}

// NewDescribeEIPRequest will create request of DescribeEIP action.
func (c *UNetClient) NewDescribeEIPRequest() *DescribeEIPRequest {
	req := &DescribeEIPRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

// DescribeEIP - 获取弹性IP信息
func (c *UNetClient) DescribeEIP(req *DescribeEIPRequest) (*DescribeEIPResponse, error) {
	var err error
	var res DescribeEIPResponse

	err = c.Client.InvokeAction("DescribeEIP", req, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}
