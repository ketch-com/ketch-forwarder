//go:build !unit && !integration && smoke

package smoketest

import (
	"context"
	"github.com/google/uuid"
	"github.com/ketch-com/ketch-forwarder/pkg/types"
	"time"
)

func (suite *SmokeTestSuite) TestAccessCallback() {
	ctx := context.Background()
	client := suite.params.Callbacks.Provide(ctx, suite.params.Config.CallbackURL)
	err := client.SendAccessStatusEvent(ctx, &types.AccessStatusEvent{
		ApiVersion: types.ApiVersion,
		Kind:       types.AccessStatusEventKind,
		Metadata: &types.Metadata{
			UID:    uuid.New().String(),
			Tenant: "axonic",
		},
		Event: &types.AccessResponseBody{
			Status:                      types.CompletedRequestStatus,
			Reason:                      types.OtherRequestStatusReason,
			ExpectedCompletionTimestamp: time.Now().Unix(),
			Results:                     nil, // TODO
		},
	})
	suite.Require().NoError(err)
}

func (suite *SmokeTestSuite) TestDeleteCallback() {
	ctx := context.Background()
	client := suite.params.Callbacks.Provide(ctx, suite.params.Config.CallbackURL)
	err := client.SendDeleteStatusEvent(ctx, &types.DeleteStatusEvent{
		ApiVersion: types.ApiVersion,
		Kind:       types.DeleteStatusEventKind,
		Metadata: &types.Metadata{
			UID:    uuid.New().String(),
			Tenant: "axonic",
		},
		Event: &types.DeleteResponseBody{
			Status:                      types.CompletedRequestStatus,
			Reason:                      types.OtherRequestStatusReason,
			ExpectedCompletionTimestamp: time.Now().Unix(),
		},
	})
	suite.Require().NoError(err)
}
