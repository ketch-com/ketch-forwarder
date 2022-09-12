package client

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/ketch-com/ketch-forwarder/pkg/types"
	"go.ketch.com/lib/orlop/v2/errors"
	"net/http"
)

type ForwarderClientProvider struct {
	client *http.Client
}

func NewForwarderClientProvider(client *http.Client) *ForwarderClientProvider {
	return &ForwarderClientProvider{
		client: client,
	}
}

func (p *ForwarderClientProvider) Provide(_ context.Context, url string) *ForwarderClient {
	return NewForwarderClient(url, p.client)
}

type ForwarderClient struct {
	url    string
	client *http.Client
}

func NewForwarderClient(url string, client *http.Client) *ForwarderClient {
	return &ForwarderClient{
		url:    url,
		client: client,
	}
}

func (c *ForwarderClient) SendAccessRequest(ctx context.Context, request *types.AccessRequest) (*types.AccessResponse, error) {
	var out types.AccessResponse
	err := c.send(ctx, request, &out)
	return &out, err
}

func (c *ForwarderClient) SendDeleteRequest(ctx context.Context, request *types.DeleteRequest) (*types.DeleteResponse, error) {
	var out types.DeleteResponse
	err := c.send(ctx, request, &out)
	return &out, err
}

func (c *ForwarderClient) SendRestrictProcessingRequest(ctx context.Context, request *types.RestrictProcessingRequest) (*types.RestrictProcessingResponse, error) {
	var out types.RestrictProcessingResponse
	err := c.send(ctx, request, &out)
	return &out, err
}

func (c *ForwarderClient) send(ctx context.Context, request any, response any) error {
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
		var errResp types.Error
		err = json.NewDecoder(resp.Body).Decode(&errResp)
		if err == nil && errResp.Error != nil {
			return errors.WithStatusCode(errors.New(errResp.Error.Message), errResp.Error.Code)
		}

		return errors.WithStatusCode(nil, resp.StatusCode)
	}

	err = json.NewDecoder(resp.Body).Decode(response)
	if err != nil {
		return err
	}

	return nil
}
