//Code is generated by ucloud code generator, don't modify it by hand, it will cause undefined behaviors.
//go:generate ucloud-gen-go-api UMem DescribeUMemcacheUpgradePrice

package umem

import (
	"github.com/yangyimincn/ucloud-sdk-go/ucloud/request"
	"github.com/yangyimincn/ucloud-sdk-go/ucloud/response"
)

// DescribeUMemcacheUpgradePriceRequest is request schema for DescribeUMemcacheUpgradePrice action
type DescribeUMemcacheUpgradePriceRequest struct {
	request.CommonBase

	// [公共参数] 地域。 参见 [地域和可用区列表](../summary/regionlist.html)
	// Region *string `required:"true"`

	// [公共参数] 可用区。参见 [可用区列表](../summary/regionlist.html)
	// Zone *string `required:"false"`

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](../summary/get_project_list.html)
	// ProjectId *string `required:"false"`

	// 购买umemcache大小,单位:GB
	Size *int `required:"true"`

	// 需要升级的空间的GroupId,请参考DescribeUMemcacheGroup接口
	GroupId *string `required:"true"`
}

// DescribeUMemcacheUpgradePriceResponse is response schema for DescribeUMemcacheUpgradePrice action
type DescribeUMemcacheUpgradePriceResponse struct {
	response.CommonBase

	// 价格，单位：元
	Price float64
}

// NewDescribeUMemcacheUpgradePriceRequest will create request of DescribeUMemcacheUpgradePrice action.
func (c *UMemClient) NewDescribeUMemcacheUpgradePriceRequest() *DescribeUMemcacheUpgradePriceRequest {
	req := &DescribeUMemcacheUpgradePriceRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

// DescribeUMemcacheUpgradePrice - 获取umemcache升级价格信息
func (c *UMemClient) DescribeUMemcacheUpgradePrice(req *DescribeUMemcacheUpgradePriceRequest) (*DescribeUMemcacheUpgradePriceResponse, error) {
	var err error
	var res DescribeUMemcacheUpgradePriceResponse

	err = c.Client.InvokeAction("DescribeUMemcacheUpgradePrice", req, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}
