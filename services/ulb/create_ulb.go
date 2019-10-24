//Code is generated by ucloud code generator, don't modify it by hand, it will cause undefined behaviors.
//go:generate ucloud-gen-go-api ULB CreateULB

package ulb

import (
	"github.com/yangyimincn/ucloud-sdk-go/ucloud/request"
	"github.com/yangyimincn/ucloud-sdk-go/ucloud/response"
)

// CreateULBRequest is request schema for CreateULB action
type CreateULBRequest struct {
	request.CommonBase

	// [公共参数] 地域。 参见 [地域和可用区列表](../summary/regionlist.html)
	// Region *string `required:"true"`

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](../summary/get_project_list.html)
	// ProjectId *string `required:"true"`

	// 负载均衡的名字，默认值为“ULB”
	ULBName *string `required:"false"`

	// 业务组
	Tag *string `required:"false"`

	// 备注
	Remark *string `required:"false"`

	// 创建的ULB是否为外网模式，默认即为外网模式
	OuterMode *string `required:"false"`

	// 创建的ULB是否为内网模式
	InnerMode *string `required:"false"`

	// 创建内网ULB时指定内网IP。若不传值，则随机分配当前子网下的IP（暂时不对外开放，创建外网ULB该字段会忽略）
	PrivateIp *string `required:"false"`

	// 付费方式
	ChargeType *string `required:"false"`

	// ULB所在的VPC的ID, 如果不传则使用默认的VPC
	VPCId *string `required:"false"`

	// 内网ULB 所属的子网ID，如果不传则使用默认的子网
	SubnetId *string `required:"false"`

	// ULB 所属的业务组ID，如果不传则使用默认的业务组
	BusinessId *string `required:"false"`
}

// CreateULBResponse is response schema for CreateULB action
type CreateULBResponse struct {
	response.CommonBase

	// 负载均衡实例的Id
	ULBId string
}

// NewCreateULBRequest will create request of CreateULB action.
func (c *ULBClient) NewCreateULBRequest() *CreateULBRequest {
	req := &CreateULBRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(false)
	return req
}

// CreateULB - 创建负载均衡实例，可以选择内网或者外网
func (c *ULBClient) CreateULB(req *CreateULBRequest) (*CreateULBResponse, error) {
	var err error
	var res CreateULBResponse

	err = c.Client.InvokeAction("CreateULB", req, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}
