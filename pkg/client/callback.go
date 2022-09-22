package client

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/ketch-com/ketch-forwarder/pkg/types"
	"go.ketch.com/lib/orlop/v2/errors"
	"net/http"
)

type CallbackClientProvider struct {
	client *http.Client
}

func NewCallbackClientProvider(client *http.Client) *CallbackClientProvider {
	return &CallbackClientProvider{
		client: client,
	}
}

func (p *CallbackClientProvider) Provide(_ context.Context, url string) *CallbackClient {
	return NewCallbackClient(url, p.client)
}

type CallbackClient struct {
	url    string
	client *http.Client
}

func NewCallbackClient(url string, client *http.Client) *CallbackClient {
	return &CallbackClient{
		url:    url,
		client: client,
	}
}

func (c *CallbackClient) SendAccessStatusEvent(ctx context.Context, request *types.AccessStatusEvent) error {
	return c.send(ctx, request)
}

func (c *CallbackClient) SendDeleteStatusEvent(ctx context.Context, request *types.DeleteStatusEvent) error {
	return c.send(ctx, request)
}

func (c *CallbackClient) SendCorrectionStatusEvent(ctx context.Context, request *types.CorrectionStatusEvent) error {
	return c.send(ctx, request)
}

func (c *CallbackClient) SendRestrictProcessingStatusEvent(ctx context.Context, request *types.RestrictProcessingStatusEvent) error {
	return c.send(ctx, request)
}

func (c *CallbackClient) send(ctx context.Context, request any) error {
	buf := new(bytes.Buffer)

	if err := json.NewEncoder(buf).Encode(request); err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.url, buf)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		errResp := &types.Error{}

		if resp.Header.Get("Content-Type") == "application/json" {
			err = json.NewDecoder(resp.Body).Decode(&errResp)
			if err != nil {
				return err
			}

			return errors.WithStatusCode(errors.New(errResp.Error.Message), errResp.Error.Code)
		}

		return errors.WithStatusCode(nil, resp.StatusCode)
	}

	return nil
}
