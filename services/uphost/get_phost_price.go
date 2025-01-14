//Code is generated by ucloud code generator, don't modify it by hand, it will cause undefined behaviors.
//go:generate ucloud-gen-go-api UPHost GetPHostPrice

package uphost

import (
	"github.com/yangyimincn/ucloud-sdk-go/ucloud/request"
	"github.com/yangyimincn/ucloud-sdk-go/ucloud/response"
)

// GetPHostPriceRequest is request schema for GetPHostPrice action
type GetPHostPriceRequest struct {
	request.CommonBase

	// 可用区。参见 [可用区列表](../summary/regionlist.html)
	// Zone *string `required:"false"`

	// 购买数量，范围[1-5]
	Count *int `required:"true"`

	// 计费模式，枚举值为： Year/Month/Trial/Dynamic
	ChargeType *string `required:"true"`

	// 购买时长，1-10个月或1-10年
	Quantity *int `required:"true"`

	// 默认为：DB(数据库型)
	Type *string `required:"false"`

	// 网络环境，可选千兆：1G ，万兆：10G
	Cluster *string `required:"false"`
}

// GetPHostPriceResponse is response schema for GetPHostPrice action
type GetPHostPriceResponse struct {
	response.CommonBase

	// 价格列表 见 PHostPriceSet
	PriceSet []PHostPriceSet
}

// NewGetPHostPriceRequest will create request of GetPHostPrice action.
func (c *UPHostClient) NewGetPHostPriceRequest() *GetPHostPriceRequest {
	req := &GetPHostPriceRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

// GetPHostPrice - 获取物理机价格列表
func (c *UPHostClient) GetPHostPrice(req *GetPHostPriceRequest) (*GetPHostPriceResponse, error) {
	var err error
	var res GetPHostPriceResponse

	err = c.Client.InvokeAction("GetPHostPrice", req, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}
