//Code is generated by ucloud code generator, don't modify it by hand, it will cause undefined behaviors.
//go:generate ucloud-gen-go-api UDisk CloneUDisk

package udisk

import (
	"github.com/yangyimincn/ucloud-sdk-go/ucloud/request"
	"github.com/yangyimincn/ucloud-sdk-go/ucloud/response"
)

// CloneUDiskRequest is request schema for CloneUDisk action
type CloneUDiskRequest struct {
	request.CommonBase

	// [公共参数] 地域。 参见 [地域和可用区列表](../summary/regionlist.html)
	// Region *string `required:"true"`

	// [公共参数] 可用区。参见 [可用区列表](../summary/regionlist.html)
	// Zone *string `required:"true"`

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](../summary/get_project_list.html)
	// ProjectId *string `required:"false"`

	// 实例名称
	Name *string `required:"true"`

	// 克隆父Disk的Id
	SourceId *string `required:"true"`

	// 方舟是否开启，"Yes":开启，"No":关闭；默认为"No"
	UDataArkMode *string `required:"false"`

	// 购买时长 默认: 1
	Quantity *int `required:"false"`

	// Disk注释
	Comment *string `required:"false"`

	// Year , Month, Dynamic 默认: Dynamic
	ChargeType *string `required:"false"`

	// 使用的代金券id
	CouponId *string `required:"false"`
}

// CloneUDiskResponse is response schema for CloneUDisk action
type CloneUDiskResponse struct {
	response.CommonBase

	// 创建UDisk Id
	UDiskId []string
}

// NewCloneUDiskRequest will create request of CloneUDisk action.
func (c *UDiskClient) NewCloneUDiskRequest() *CloneUDiskRequest {
	req := &CloneUDiskRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(false)
	return req
}

// CloneUDisk - 从UDisk创建UDisk克隆
func (c *UDiskClient) CloneUDisk(req *CloneUDiskRequest) (*CloneUDiskResponse, error) {
	var err error
	var res CloneUDiskResponse

	err = c.Client.InvokeAction("CloneUDisk", req, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}
