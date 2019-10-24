//Code is generated by ucloud code generator, don't modify it by hand, it will cause undefined behaviors.
//go:generate ucloud-gen-go-api UHost DescribeUHostInstance

package uhost

import (
	"github.com/yangyimincn/ucloud-sdk-go/ucloud/request"
	"github.com/yangyimincn/ucloud-sdk-go/ucloud/response"
)

// DescribeUHostInstanceRequest is request schema for DescribeUHostInstance action
type DescribeUHostInstanceRequest struct {
	request.CommonBase

	// [公共参数] 地域。 参见 [地域和可用区列表](../summary/regionlist.html)
	// Region *string `required:"true"`

	// [公共参数] 可用区。参见 [可用区列表](../summary/regionlist.html)
	// Zone *string `required:"false"`

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](../summary/get_project_list.html)
	// ProjectId *string `required:"false"`

	// 【数组】UHost主机的资源ID，例如UHostIds.0代表希望获取信息 的主机1，UHostIds.1代表主机2。 如果不传入，则返回当前Region 所有符合条件的UHost实例。
	UHostIds []string `required:"false"`

	// 要查询的业务组名称
	Tag *string `required:"false"`

	// 1：普通云主机；2：抢占型云主机；如不传此参数，默认全部获取
	LifeCycle *int `required:"false"`

	// 列表起始位置偏移量，默认为0
	Offset *int `required:"false"`

	// 返回数据长度，默认为20，最大100
	Limit *int `required:"false"`

	// 硬件隔离组id。通过硬件隔离组筛选主机。
	IsolationGroup *string `required:"false"`

	// vpc id。通过VPC筛选主机。
	VPCId *string `required:"false"`

	// 子网id。通过子网筛选主机。
	SubnetId *string `required:"false"`
}

// DescribeUHostInstanceResponse is response schema for DescribeUHostInstance action
type DescribeUHostInstanceResponse struct {
	response.CommonBase

	// UHostInstance总数
	TotalCount int

	// 云主机实例列表，每项参数可见下面 UHostInstanceSet
	UHostSet []UHostInstanceSet
}

// NewDescribeUHostInstanceRequest will create request of DescribeUHostInstance action.
func (c *UHostClient) NewDescribeUHostInstanceRequest() *DescribeUHostInstanceRequest {
	req := &DescribeUHostInstanceRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

// DescribeUHostInstance - 获取主机或主机列表信息，并可根据数据中心，主机ID等参数进行过滤。
func (c *UHostClient) DescribeUHostInstance(req *DescribeUHostInstanceRequest) (*DescribeUHostInstanceResponse, error) {
	var err error
	var res DescribeUHostInstanceResponse

	err = c.Client.InvokeAction("DescribeUHostInstance", req, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}
