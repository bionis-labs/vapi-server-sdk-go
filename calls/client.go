// This file was auto-generated by Fern from our API Definition.

package calls

import (
	context "context"
	serversdkgo "github.com/bionis-labs/vapi-server-sdk-go"
	core "github.com/bionis-labs/vapi-server-sdk-go/core"
	internal "github.com/bionis-labs/vapi-server-sdk-go/internal"
	option "github.com/bionis-labs/vapi-server-sdk-go/option"
	http "net/http"
)

type Client struct {
	baseURL string
	caller  *internal.Caller
	header  http.Header
}

func NewClient(opts ...option.RequestOption) *Client {
	options := core.NewRequestOptions(opts...)
	return &Client{
		baseURL: options.BaseURL,
		caller: internal.NewCaller(
			&internal.CallerParams{
				Client:      options.HTTPClient,
				MaxAttempts: options.MaxAttempts,
			},
		),
		header: options.ToHeader(),
	}
}

func (c *Client) List(
	ctx context.Context,
	request *serversdkgo.CallsListRequest,
	opts ...option.RequestOption,
) ([]*serversdkgo.Call, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.vapi.ai",
	)
	endpointURL := baseURL + "/call"
	queryParams, err := internal.QueryValues(request)
	if err != nil {
		return nil, err
	}
	if len(queryParams) > 0 {
		endpointURL += "?" + queryParams.Encode()
	}
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)

	var response []*serversdkgo.Call
	if err := c.caller.Call(
		ctx,
		&internal.CallParams{
			URL:             endpointURL,
			Method:          http.MethodGet,
			Headers:         headers,
			MaxAttempts:     options.MaxAttempts,
			BodyProperties:  options.BodyProperties,
			QueryParameters: options.QueryParameters,
			Client:          options.HTTPClient,
			Response:        &response,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

func (c *Client) Create(
	ctx context.Context,
	request *serversdkgo.CreateCallDto,
	opts ...option.RequestOption,
) (*serversdkgo.Call, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.vapi.ai",
	)
	endpointURL := baseURL + "/call"
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)
	headers.Set("Content-Type", "application/json")

	var response *serversdkgo.Call
	if err := c.caller.Call(
		ctx,
		&internal.CallParams{
			URL:             endpointURL,
			Method:          http.MethodPost,
			Headers:         headers,
			MaxAttempts:     options.MaxAttempts,
			BodyProperties:  options.BodyProperties,
			QueryParameters: options.QueryParameters,
			Client:          options.HTTPClient,
			Request:         request,
			Response:        &response,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

func (c *Client) Get(
	ctx context.Context,
	id string,
	opts ...option.RequestOption,
) (*serversdkgo.Call, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.vapi.ai",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/call/%v",
		id,
	)
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)

	var response *serversdkgo.Call
	if err := c.caller.Call(
		ctx,
		&internal.CallParams{
			URL:             endpointURL,
			Method:          http.MethodGet,
			Headers:         headers,
			MaxAttempts:     options.MaxAttempts,
			BodyProperties:  options.BodyProperties,
			QueryParameters: options.QueryParameters,
			Client:          options.HTTPClient,
			Response:        &response,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

func (c *Client) Delete(
	ctx context.Context,
	id string,
	opts ...option.RequestOption,
) (*serversdkgo.Call, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.vapi.ai",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/call/%v",
		id,
	)
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)

	var response *serversdkgo.Call
	if err := c.caller.Call(
		ctx,
		&internal.CallParams{
			URL:             endpointURL,
			Method:          http.MethodDelete,
			Headers:         headers,
			MaxAttempts:     options.MaxAttempts,
			BodyProperties:  options.BodyProperties,
			QueryParameters: options.QueryParameters,
			Client:          options.HTTPClient,
			Response:        &response,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

func (c *Client) Update(
	ctx context.Context,
	id string,
	request *serversdkgo.UpdateCallDto,
	opts ...option.RequestOption,
) (*serversdkgo.Call, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.vapi.ai",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/call/%v",
		id,
	)
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)
	headers.Set("Content-Type", "application/json")

	var response *serversdkgo.Call
	if err := c.caller.Call(
		ctx,
		&internal.CallParams{
			URL:             endpointURL,
			Method:          http.MethodPatch,
			Headers:         headers,
			MaxAttempts:     options.MaxAttempts,
			BodyProperties:  options.BodyProperties,
			QueryParameters: options.QueryParameters,
			Client:          options.HTTPClient,
			Request:         request,
			Response:        &response,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}
