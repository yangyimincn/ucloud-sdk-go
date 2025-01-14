// Code is generated by ucloud-model, DO NOT EDIT IT.

package vpc

import (
	"github.com/yangyimincn/ucloud-sdk-go/ucloud/request"
	"github.com/yangyimincn/ucloud-sdk-go/ucloud/response"
)

// GetAvailableResourceForPolicyRequest is request schema for GetAvailableResourceForPolicy action
type GetAvailableResourceForPolicyRequest struct {
	request.CommonBase

	// [公共参数] 项目Id。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](../summary/get_project_list.html)
	// ProjectId *string `required:"false"`

	// [公共参数] 地域。 参见 [地域和可用区列表](../summary/regionlist.html)
	// Region *string `required:"true"`

	// 返回数据长度，默认为10000
	Limit *int `required:"false"`

	// NAT网关Id
	NATGWId *string `required:"true"`

	// 列表起始位置偏移量，默认为0
	Offset *int `required:"false"`
}

// GetAvailableResourceForPolicyResponse is response schema for GetAvailableResourceForPolicy action
type GetAvailableResourceForPolicyResponse struct {
	response.CommonBase

	// 支持资源类型的信息
	DataSet []GetAvailableResourceForPolicyDataSet
}

// NewGetAvailableResourceForPolicyRequest will create request of GetAvailableResourceForPolicy action.
func (c *VPCClient) NewGetAvailableResourceForPolicyRequest() *GetAvailableResourceForPolicyRequest {
	req := &GetAvailableResourceForPolicyRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

// GetAvailableResourceForPolicy - 获取NAT网关可配置端口转发规则的资源信息
func (c *VPCClient) GetAvailableResourceForPolicy(req *GetAvailableResourceForPolicyRequest) (*GetAvailableResourceForPolicyResponse, error) {
	var err error
	var res GetAvailableResourceForPolicyResponse

	reqCopier := *req

	err = c.Client.InvokeAction("GetAvailableResourceForPolicy", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}
