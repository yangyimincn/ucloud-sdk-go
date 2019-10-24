// yangyimincn created
//go:generate ucloud-gen-go-api UBill DescribeOrderDetailInfo

package ubill

import (
	"github.com/yangyimincn/ucloud-sdk-go/ucloud/request"
	"github.com/yangyimincn/ucloud-sdk-go/ucloud/response"
)

// DescribeTransactionListRequest is request schema for DescribeTransactionList action
type DescribeTransactionListRequest struct {
	request.CommonBase

	// 	开始时间
	BeginTime *int `required:"true"`

	// 结束时间(时间跨度不超过3个月)
	EndTime *int `required:"true"`

	// 返回数据长度(默认为25)[1~100]
	Limit *int `required:"false"`

	// 数据偏移量(默认为0)
	Offset *int `required:"false"`
}

// DescribeTransactionListResponse is response schema for DescribeTransactionList action
type  DescribeTransactionListResponse struct {
	response.CommonBase

	// JSON格式的订单信息
	TransactionInfos []TransactionInfo
}


func (c *UBillClient) NewDescribeTransactionListRequest() *DescribeTransactionListRequest {
	req := &DescribeTransactionListRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

func (c *UBillClient) DescribeTransactionList(req *DescribeTransactionListRequest) (*DescribeTransactionListResponse, error) {
	var err error
	var res DescribeTransactionListResponse

	err = c.Client.InvokeAction("DescribeTransactionList", req, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}