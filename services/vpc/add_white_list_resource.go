// Code is generated by ucloud-model, DO NOT EDIT IT.

package vpc

import (
	"github.com/yangyimincn/ucloud-sdk-go/ucloud/request"
	"github.com/yangyimincn/ucloud-sdk-go/ucloud/response"
)

// AddWhiteListResourceRequest is request schema for AddWhiteListResource action
type AddWhiteListResourceRequest struct {
	request.CommonBase

	// [公共参数] 项目Id。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](../summary/get_project_list.html)
	// ProjectId *string `required:"false"`

	// [公共参数] 地域。 参见 [地域和可用区列表](../summary/regionlist.html)
	// Region *string `required:"true"`

	// NAT网关Id
	NATGWId *string `required:"true"`

	// 可添加白名单的资源Id
	ResourceIds []string `required:"true"`
}

// AddWhiteListResourceResponse is response schema for AddWhiteListResource action
type AddWhiteListResourceResponse struct {
	response.CommonBase
}

// NewAddWhiteListResourceRequest will create request of AddWhiteListResource action.
func (c *VPCClient) NewAddWhiteListResourceRequest() *AddWhiteListResourceRequest {
	req := &AddWhiteListResourceRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(false)
	return req
}

// AddWhiteListResource - 添加NAT网关白名单
func (c *VPCClient) AddWhiteListResource(req *AddWhiteListResourceRequest) (*AddWhiteListResourceResponse, error) {
	var err error
	var res AddWhiteListResourceResponse

	reqCopier := *req

	err = c.Client.InvokeAction("AddWhiteListResource", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}
