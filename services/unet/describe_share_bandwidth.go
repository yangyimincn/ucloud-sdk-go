//Code is generated by ucloud code generator, don't modify it by hand, it will cause undefined behaviors.
//go:generate ucloud-gen-go-api UNet DescribeShareBandwidth

package unet

import (
	"github.com/yangyimincn/ucloud-sdk-go/ucloud/request"
	"github.com/yangyimincn/ucloud-sdk-go/ucloud/response"
)

// DescribeShareBandwidthRequest is request schema for DescribeShareBandwidth action
type DescribeShareBandwidthRequest struct {
	request.CommonBase

	// [公共参数] 地域。 参见 [地域和可用区列表](../summary/regionlist.html)
	// Region *string `required:"true"`

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](../summary/get_project_list.html)
	// ProjectId *string `required:"false"`

	// 需要返回的共享带宽Id
	ShareBandwidthIds []string `required:"false"`
}

// DescribeShareBandwidthResponse is response schema for DescribeShareBandwidth action
type DescribeShareBandwidthResponse struct {
	response.CommonBase

	// 共享带宽信息组 参见 UnetShareBandwidthSet
	DataSet []UnetShareBandwidthSet

	// 符合条件的共享带宽总数，大于等于返回DataSet长度
	TotalCount int
}

// NewDescribeShareBandwidthRequest will create request of DescribeShareBandwidth action.
func (c *UNetClient) NewDescribeShareBandwidthRequest() *DescribeShareBandwidthRequest {
	req := &DescribeShareBandwidthRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

// DescribeShareBandwidth - 获取共享带宽信息
func (c *UNetClient) DescribeShareBandwidth(req *DescribeShareBandwidthRequest) (*DescribeShareBandwidthResponse, error) {
	var err error
	var res DescribeShareBandwidthResponse

	err = c.Client.InvokeAction("DescribeShareBandwidth", req, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}
