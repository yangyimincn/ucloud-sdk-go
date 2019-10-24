//Code is generated by ucloud code generator, don't modify it by hand, it will cause undefined behaviors.
//go:generate ucloud-gen-go-api UDPN DescribeUDPN

package udpn

import (
	"github.com/yangyimincn/ucloud-sdk-go/ucloud/request"
	"github.com/yangyimincn/ucloud-sdk-go/ucloud/response"
)

// DescribeUDPNRequest is request schema for DescribeUDPN action
type DescribeUDPNRequest struct {
	request.CommonBase

	// [公共参数] 地域。 参见 [地域和可用区列表](../summary/regionlist.html)
	// Region *string `required:"false"`

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](../summary/get_project_list.html)
	// ProjectId *string `required:"false"`

	// 申请到的 UDPN 资源 ID。若为空，则查询该用户在机房所有的专线信息。非默认项目资源，需填写ProjectId
	UDPNId *string `required:"false"`

	// 列表起始位置偏移量，默认为 0
	Offset *int `required:"false"`

	// 返回数据长度，默认为 20
	Limit *int `required:"false"`
}

// DescribeUDPNResponse is response schema for DescribeUDPN action
type DescribeUDPNResponse struct {
	response.CommonBase

	// 查询到的总数量
	TotalCount int

	// UDPN详情
	DataSet []UDPNData
}

// NewDescribeUDPNRequest will create request of DescribeUDPN action.
func (c *UDPNClient) NewDescribeUDPNRequest() *DescribeUDPNRequest {
	req := &DescribeUDPNRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

// DescribeUDPN - 描述 UDPN
func (c *UDPNClient) DescribeUDPN(req *DescribeUDPNRequest) (*DescribeUDPNResponse, error) {
	var err error
	var res DescribeUDPNResponse

	err = c.Client.InvokeAction("DescribeUDPN", req, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}
