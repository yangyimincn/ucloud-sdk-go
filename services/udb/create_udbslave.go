//Code is generated by ucloud code generator, don't modify it by hand, it will cause undefined behaviors.
//go:generate ucloud-gen-go-api UDB CreateUDBSlave

package udb

import (
	"github.com/yangyimincn/ucloud-sdk-go/ucloud/request"
	"github.com/yangyimincn/ucloud-sdk-go/ucloud/response"
)

// CreateUDBSlaveRequest is request schema for CreateUDBSlave action
type CreateUDBSlaveRequest struct {
	request.CommonBase

	// [公共参数] 地域。 参见 [地域和可用区列表](../summary/regionlist.html)
	// Region *string `required:"true"`

	// [公共参数] 可用区。参见 [可用区列表](../summary/regionlist.html)
	// Zone *string `required:"false"`

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](../summary/get_project_list.html)
	// ProjectId *string `required:"false"`

	// master实例的DBId,该值可以通过DescribeUDBInstance获取
	SrcId *string `required:"true"`

	// 实例名称，至少6位
	Name *string `required:"true"`

	// 端口号，mysql默认3306
	Port *int `required:"false"`

	// 是否使用SSD，默认为false
	UseSSD *bool `required:"false"`

	// SSD类型，可选值为"SATA"、"PCI-E"，如果UseSSD为true ，则必选
	SSDType *string `required:"false"`

	// 内存限制(MB)，目前支持以下几档 1000M/2000M/4000M/ 6000M/8000M/12000M/16000M/ 24000M/32000M/48000M/ 64000M/96000M
	MemoryLimit *int `required:"false"`

	// 磁盘空间(GB), 暂时支持20G - 3000G（API支持，前端暂时只开放内存定制）
	DiskSpace *int `required:"false"`

	// 是否锁主库，默认为true
	IsLock *bool `required:"false"`

	// UDB实例部署模式，可选值如下：Normal: 普通单点实例HA: 高可用部署实例
	InstanceMode *string `required:"false"`

	// UDB实例类型：Normal和SATA_SSD
	InstanceType *string `required:"false"`

	// 使用的代金券id
	CouponId *string `required:"false"`

	// 配置参数组 ID
	ParamGroupId *string `required:"false"`
}

// CreateUDBSlaveResponse is response schema for CreateUDBSlave action
type CreateUDBSlaveResponse struct {
	response.CommonBase

	// 创建slave的DBId
	DBId string
}

// NewCreateUDBSlaveRequest will create request of CreateUDBSlave action.
func (c *UDBClient) NewCreateUDBSlaveRequest() *CreateUDBSlaveRequest {
	req := &CreateUDBSlaveRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(false)
	return req
}

// CreateUDBSlave - 创建UDB实例的slave
func (c *UDBClient) CreateUDBSlave(req *CreateUDBSlaveRequest) (*CreateUDBSlaveResponse, error) {
	var err error
	var res CreateUDBSlaveResponse

	err = c.Client.InvokeAction("CreateUDBSlave", req, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}
