//Code is generated by ucloud code generator, don't modify it by hand, it will cause undefined behaviors.
//go:generate ucloud-gen-go-api UDB CreateUDBRouteInstance

package udb

import (
	"github.com/yangyimincn/ucloud-sdk-go/ucloud/request"
	"github.com/yangyimincn/ucloud-sdk-go/ucloud/response"
)

// CreateUDBRouteInstanceRequest is request schema for CreateUDBRouteInstance action
type CreateUDBRouteInstanceRequest struct {
	request.CommonBase

	// [公共参数] 地域。 参见 [地域和可用区列表](../summary/regionlist.html)
	// Region *string `required:"true"`

	// [公共参数] 可用区。参见 [可用区列表](../summary/regionlist.html)
	// Zone *string `required:"false"`

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](../summary/get_project_list.html)
	// ProjectId *string `required:"false"`

	// DB类型id，mongodb按版本细分有1：mongodb-2.4，2：mongodb-2.6,3：mongodb-3.0，4：mongodb-3.2
	DBTypeId *string `required:"true"`

	// 实例名称，至少6位
	Name *string `required:"true"`

	// 端口号，mongodb默认27017
	Port *int `required:"true"`

	// DB实例使用的配置参数组id
	ParamGroupId *int `required:"true"`

	// 内存限制(MB)，目前支持以下几档 600M/1500M/3000M /6000M/15000M/30000M
	MemoryLimit *int `required:"true"`

	// 磁盘空间(GB), 暂时支持20G - 500G
	DiskSpace *int `required:"true"`

	// 配置服务器的dbid，允许一个或者三个
	ConfigsvrId []string `required:"true"`

	// Year， Month， Dynamic，Trial，默认: Month
	ChargeType *string `required:"false"`

	// 购买时长，默认值1
	Quantity *int `required:"false"`

	// 是否使用SSD，默认为false
	UseSSD *bool `required:"false"`

	// 使用的代金券id
	CouponId *string `required:"false"`
}

// CreateUDBRouteInstanceResponse is response schema for CreateUDBRouteInstance action
type CreateUDBRouteInstanceResponse struct {
	response.CommonBase

	// db实例id
	DBId string
}

// NewCreateUDBRouteInstanceRequest will create request of CreateUDBRouteInstance action.
func (c *UDBClient) NewCreateUDBRouteInstanceRequest() *CreateUDBRouteInstanceRequest {
	req := &CreateUDBRouteInstanceRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(false)
	return req
}

// CreateUDBRouteInstance - 创建mongos实例
func (c *UDBClient) CreateUDBRouteInstance(req *CreateUDBRouteInstanceRequest) (*CreateUDBRouteInstanceResponse, error) {
	var err error
	var res CreateUDBRouteInstanceResponse

	err = c.Client.InvokeAction("CreateUDBRouteInstance", req, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}
