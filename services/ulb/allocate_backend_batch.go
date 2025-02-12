//Code is generated by ucloud code generator, don't modify it by hand, it will cause undefined behaviors.
//go:generate ucloud-gen-go-api ULB AllocateBackendBatch

package ulb

import (
	"github.com/yangyimincn/ucloud-sdk-go/ucloud/request"
	"github.com/yangyimincn/ucloud-sdk-go/ucloud/response"
)

// AllocateBackendBatchRequest is request schema for AllocateBackendBatch action
type AllocateBackendBatchRequest struct {
	request.CommonBase

	// [公共参数] 地域。 参见 [地域和可用区列表](../summary/regionlist.html)
	// Region *string `required:"true"`

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](../summary/get_project_list.html)
	// ProjectId *string `required:"true"`

	// 负载均衡实例的ID
	ULBId *string `required:"true"`

	// VServer实例的ID
	VServerId *string `required:"true"`

	// 用| 分割字段，格式：ResourceId| ResourceType| Port| Enabled|IP| Weight。ResourceId:所添加的后端资源的资源ID；ResourceType:所添加的后端资源的类型，枚举值：UHost -> 云主机；UPM -> 物理云主机； UDHost -> 私有专区主机；UDocker -> 容器，默认值为“UHost”；Port:所添加的后端资源服务端口，取值范围[1-65535]；Enabled:后端实例状态开关，枚举值： 1：启用； 0：禁用；IP:后端资源内网ip；Weight：所添加的后端RS权重（在加权轮询算法下有效），取值范围[0-100]，默认为1
	Backends []string `required:"true"`

	// 已弃用，指定 Api 版本
	ApiVersion *int `required:"false"`
}

// AllocateBackendBatchResponse is response schema for AllocateBackendBatch action
type AllocateBackendBatchResponse struct {
	response.CommonBase

	// 所添加的后端资源ID，（为ULB系统中使用，与资源自身ID无关），可用于 UpdateBackendAttribute/UpdateBackendAttributeBatch/ReleaseBackend
	BackendSet []BackendSet
}

// NewAllocateBackendBatchRequest will create request of AllocateBackendBatch action.
func (c *ULBClient) NewAllocateBackendBatchRequest() *AllocateBackendBatchRequest {
	req := &AllocateBackendBatchRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

// AllocateBackendBatch - 批量添加VServer后端节点
func (c *ULBClient) AllocateBackendBatch(req *AllocateBackendBatchRequest) (*AllocateBackendBatchResponse, error) {
	var err error
	var res AllocateBackendBatchResponse

	err = c.Client.InvokeAction("AllocateBackendBatch", req, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}
