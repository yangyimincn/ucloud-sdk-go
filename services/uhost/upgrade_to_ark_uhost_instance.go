//Code is generated by ucloud code generator, don't modify it by hand, it will cause undefined behaviors.
//go:generate ucloud-gen-go-api UHost UpgradeToArkUHostInstance

package uhost

import (
	"github.com/yangyimincn/ucloud-sdk-go/ucloud/request"
	"github.com/yangyimincn/ucloud-sdk-go/ucloud/response"
)

// UpgradeToArkUHostInstanceRequest is request schema for UpgradeToArkUHostInstance action
type UpgradeToArkUHostInstanceRequest struct {
	request.CommonBase

	// [公共参数] 地域。 参见 [地域和可用区列表](../summary/regionlist.html)
	// Region *string `required:"true"`

	// [公共参数] 可用区。参见 [可用区列表](../summary/regionlist.html)
	// Zone *string `required:"true"`

	// UHost主机的资源ID，例如UHostIds.0代表希望升级的主机1，UHostIds.1代表主机2。
	UHostIds []string `required:"true"`

	// 代金券ID 请参考DescribeCoupon接口
	CouponId *string `required:"false"`
}

// UpgradeToArkUHostInstanceResponse is response schema for UpgradeToArkUHostInstance action
type UpgradeToArkUHostInstanceResponse struct {
	response.CommonBase

	// UHost主机的资源ID数组
	UHostSet []string
}

// NewUpgradeToArkUHostInstanceRequest will create request of UpgradeToArkUHostInstance action.
func (c *UHostClient) NewUpgradeToArkUHostInstanceRequest() *UpgradeToArkUHostInstanceRequest {
	req := &UpgradeToArkUHostInstanceRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

// UpgradeToArkUHostInstance - 普通升级为方舟机型
func (c *UHostClient) UpgradeToArkUHostInstance(req *UpgradeToArkUHostInstanceRequest) (*UpgradeToArkUHostInstanceResponse, error) {
	var err error
	var res UpgradeToArkUHostInstanceResponse

	err = c.Client.InvokeAction("UpgradeToArkUHostInstance", req, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}
