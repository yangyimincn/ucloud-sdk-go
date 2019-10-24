//Code is generated by ucloud code generator, don't modify it by hand, it will cause undefined behaviors.
//go:generate ucloud-gen-go-api UDisk CreateUDisk

package udisk

import (
	"github.com/yangyimincn/ucloud-sdk-go/ucloud/request"
	"github.com/yangyimincn/ucloud-sdk-go/ucloud/response"
)

// CreateUDiskRequest is request schema for CreateUDisk action
type CreateUDiskRequest struct {
	request.CommonBase

	// [公共参数] 地域。 参见 [地域和可用区列表](../summary/regionlist.html)
	// Region *string `required:"true"`

	// [公共参数] 可用区。参见 [可用区列表](../summary/regionlist.html)
	// Zone *string `required:"true"`

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](../summary/get_project_list.html)
	// ProjectId *string `required:"false"`

	// 购买UDisk大小,单位:GB,普通盘: 范围[1~2000], 权限位控制可达8T,若需要请申请开通相关权限;SSD盘： 范围[1~4000]。
	Size *int `required:"true"`

	// 实例名称
	Name *string `required:"true"`

	// Year , Month, Dynamic, Trial 默认: Dynamic
	ChargeType *string `required:"false"`

	// 购买时长 默认: 1
	Quantity *int `required:"false"`

	// 是否开启数据方舟
	UDataArkMode *string `required:"false"`

	// 业务组 默认：Default
	Tag *string `required:"false"`

	// UDisk 类型: DataDisk（普通数据盘），SSDDataDisk（SSD数据盘），RSSDDataDisk（RSSD数据盘），默认值（DataDisk）
	DiskType *string `required:"false"`

	// 使用的代金券id
	CouponId *string `required:"false"`
}

// CreateUDiskResponse is response schema for CreateUDisk action
type CreateUDiskResponse struct {
	response.CommonBase

	// UDisk实例Id
	UDiskId []string
}

// NewCreateUDiskRequest will create request of CreateUDisk action.
func (c *UDiskClient) NewCreateUDiskRequest() *CreateUDiskRequest {
	req := &CreateUDiskRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(false)
	return req
}

// CreateUDisk - 创建UDisk磁盘
func (c *UDiskClient) CreateUDisk(req *CreateUDiskRequest) (*CreateUDiskResponse, error) {
	var err error
	var res CreateUDiskResponse

	err = c.Client.InvokeAction("CreateUDisk", req, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}
