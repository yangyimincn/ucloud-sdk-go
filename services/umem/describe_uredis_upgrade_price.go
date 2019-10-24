//Code is generated by ucloud code generator, don't modify it by hand, it will cause undefined behaviors.
//go:generate ucloud-gen-go-api UMem DescribeURedisUpgradePrice

package umem

import (
	"github.com/yangyimincn/ucloud-sdk-go/ucloud/request"
	"github.com/yangyimincn/ucloud-sdk-go/ucloud/response"
)

// DescribeURedisUpgradePriceRequest is request schema for DescribeURedisUpgradePrice action
type DescribeURedisUpgradePriceRequest struct {
	request.CommonBase

	// [公共参数] 地域。 参见 [地域和可用区列表](../summary/regionlist.html)
	// Region *string `required:"true"`

	// [公共参数] 可用区。参见 [可用区列表](../summary/regionlist.html)
	// Zone *string `required:"false"`

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](../summary/get_project_list.html)
	// ProjectId *string `required:"false"`

	// 购买uredis大小,单位:GB,范围是[1-32]
	Size *int `required:"true"`

	// 要升级的空间的GroupId,请参考DescribeURedisGroup接口
	GroupId *string `required:"true"`

	//
	Type *string `required:"false"`
}

// DescribeURedisUpgradePriceResponse is response schema for DescribeURedisUpgradePrice action
type DescribeURedisUpgradePriceResponse struct {
	response.CommonBase

	// 扩容差价，单位: 元，保留小数点后两位有效数字
	Price float64
}

// NewDescribeURedisUpgradePriceRequest will create request of DescribeURedisUpgradePrice action.
func (c *UMemClient) NewDescribeURedisUpgradePriceRequest() *DescribeURedisUpgradePriceRequest {
	req := &DescribeURedisUpgradePriceRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

// DescribeURedisUpgradePrice - 获取uredis升级价格信息
func (c *UMemClient) DescribeURedisUpgradePrice(req *DescribeURedisUpgradePriceRequest) (*DescribeURedisUpgradePriceResponse, error) {
	var err error
	var res DescribeURedisUpgradePriceResponse

	err = c.Client.InvokeAction("DescribeURedisUpgradePrice", req, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}
