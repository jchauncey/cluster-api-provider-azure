package wrappers

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2018-02-01/resources"
	"context"
)

type DeploymentsDeleteFutureWrapper struct {
	mock bool
	future resources.DeploymentsDeleteFuture
}

func (wrapper *DeploymentsDeleteFutureWrapper) WaitForCompletion(ctx context.Context, client autorest.Client) error {
	if !wrapper.mock {
		return wrapper.future.Future.WaitForCompletionRef(ctx, client)
	}
	return nil
}

