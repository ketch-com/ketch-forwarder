package smoketest

import (
	"context"
	"github.com/google/uuid"
	"github.com/ketch-com/ketch-forwarder/pkg/types"
	"time"
)

func (suite *SmokeTestSuite) TestSendAccessRequest() {
	ctx := context.Background()
	client := suite.params.Forwarder.Provide(ctx, suite.params.Config.ReceiverURL)
	resp, err := client.SendAccessRequest(ctx, &types.AccessRequest{
		ApiVersion: types.ApiVersion,
		Kind:       types.AccessRequestKind,
		Metadata: &types.Metadata{
			UID:    uuid.New().String(),
			Tenant: "axonic",
		},
		Request: &types.AccessRequestBody{
			Controller:   "",
			Property:     "axonic.io",
			Environment:  "production",
			Regulation:   "gdpr",
			Jurisdiction: "gdpr",
			Identities: []*types.Identity{
				{
					IdentitySpace:  "email",
					IdentityFormat: "raw",
					IdentityValue:  "test@email.com",
				},
			},
			Callbacks: []*types.Callback{
				{
					URL: suite.params.Config.CallbackURL,
					Headers: map[string]string{
						"Authorization": "Bearer auth",
					},
				},
			},
			Subject: &types.DataSubject{
				Email:           "test@email.com",
				FirstName:       "Test",
				LastName:        "Subject",
				AddressLine1:    "123 Main St",
				AddressLine2:    "",
				City:            "Anytown",
				StateRegionCode: "MA",
				PostalCode:      "10123",
				CountryCode:     "US",
				Description:     "",
			},
			Claims:             nil,
			SubmittedTimestamp: time.Now().Unix(),
			DueTimestamp:       time.Now().Add(45 * 24 * time.Hour).Unix(),
		},
	})
	suite.Require().NoError(err)
	suite.Require().NotNil(resp)
}

func (suite *SmokeTestSuite) TestSendDeleteRequest() {
	ctx := context.Background()
	client := suite.params.Forwarder.Provide(ctx, suite.params.Config.ReceiverURL)
	resp, err := client.SendDeleteRequest(ctx, &types.DeleteRequest{
		ApiVersion: types.ApiVersion,
		Kind:       types.DeleteRequestKind,
		Metadata: &types.Metadata{
			UID:    uuid.New().String(),
			Tenant: "axonic",
		},
		Request: &types.DeleteRequestBody{
			Controller:   "",
			Property:     "axonic.io",
			Environment:  "production",
			Regulation:   "gdpr",
			Jurisdiction: "gdpr",
			Identities: []*types.Identity{
				{
					IdentitySpace:  "email",
					IdentityFormat: "raw",
					IdentityValue:  "test@email.com",
				},
			},
			Callbacks: []*types.Callback{
				{
					URL: suite.params.Config.CallbackURL,
					Headers: map[string]string{
						"Authorization": "Bearer auth",
					},
				},
			},
			Subject: &types.DataSubject{
				Email:           "test@email.com",
				FirstName:       "Test",
				LastName:        "Subject",
				AddressLine1:    "123 Main St",
				AddressLine2:    "",
				City:            "Anytown",
				StateRegionCode: "MA",
				PostalCode:      "10123",
				CountryCode:     "US",
				Description:     "",
			},
			Claims:             nil,
			SubmittedTimestamp: time.Now().Unix(),
			DueTimestamp:       time.Now().Add(45 * 24 * time.Hour).Unix(),
		},
	})
	suite.Require().NoError(err)
	suite.Require().NotNil(resp)
}
