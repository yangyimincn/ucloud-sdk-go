// Code is generated by ucloud-model, DO NOT EDIT IT.

package vpc

import (
	"github.com/yangyimincn/ucloud-sdk-go/ucloud/request"
	"github.com/yangyimincn/ucloud-sdk-go/ucloud/response"
)

// ListSubnetForNATGWRequest is request schema for ListSubnetForNATGW action
type ListSubnetForNATGWRequest struct {
	request.CommonBase

	// [公共参数] 项目Id。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](../summary/get_project_list.html)
	// ProjectId *string `required:"false"`

	// [公共参数] 地域。 参见 [地域和可用区列表](../summary/regionlist.html)
	// Region *string `required:"true"`

	// NAT网关所属VPC Id。默认值为Default VPC Id
	VPCId *string `required:"false"`
}

// ListSubnetForNATGWResponse is response schema for ListSubnetForNATGW action
type ListSubnetForNATGWResponse struct {
	response.CommonBase

	// 具体参数请见NatgwSubnetDataSet
	DataSet []NatgwSubnetDataSet
}

// NewListSubnetForNATGWRequest will create request of ListSubnetForNATGW action.
func (c *VPCClient) NewListSubnetForNATGWRequest() *ListSubnetForNATGWRequest {
	req := &ListSubnetForNATGWRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

// ListSubnetForNATGW - 展示NAT网关可绑定的子网列表
func (c *VPCClient) ListSubnetForNATGW(req *ListSubnetForNATGWRequest) (*ListSubnetForNATGWResponse, error) {
	var err error
	var res ListSubnetForNATGWResponse

	reqCopier := *req

	err = c.Client.InvokeAction("ListSubnetForNATGW", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}
