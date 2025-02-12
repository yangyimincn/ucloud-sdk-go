// Code is generated by ucloud-model, DO NOT EDIT IT.

package ulb

import (
	"github.com/yangyimincn/ucloud-sdk-go/ucloud"
	"github.com/yangyimincn/ucloud-sdk-go/ucloud/auth"
)

// ULBClient is the client of ULB
type ULBClient struct {
	*ucloud.Client
}

// NewClient will return a instance of ULBClient
func NewClient(config *ucloud.Config, credential *auth.Credential) *ULBClient {
	meta := ucloud.ClientMeta{Product: "ULB"}
	client := ucloud.NewClientWithMeta(config, credential, meta)
	return &ULBClient{
		client,
	}
}
