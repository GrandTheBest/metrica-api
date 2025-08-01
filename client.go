package metrica

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/GrandTheBest/metrica-api/internal/request"
)

type Client struct {
	baseURL    string
	httpClient *http.Client
	authToken  string
}

func NewClient(opts ...Option) *Client {
	client := &Client{
		baseURL: "https://api.chunkbyte.space",
	}

	client.httpClient = &http.Client{Timeout: 30}

	for _, opt := range opts {
		opt(client)
	}

	return client
}

// Return the pointer to instance of User by Telegram ID.
// See https://dev.chunkbyte.space/api/metrica/GetByID
func (C *Client) GetByID(ctx context.Context, id string) (*User, error) {
	u, err := url.Parse(C.baseURL)

	u.Path = "/metrica/about"

	query := u.Query()
	query.Set("token", C.authToken)
	query.Set("id", id)

	u.RawQuery = query.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)

	if err != nil {
		return nil, fmt.Errorf("Failed to create request: %w", err)
	}

	resp, err := request.Request(ctx, C.httpClient, req)

	if err != nil {
		return nil, fmt.Errorf("Failed to execute request: %w", err)
	}

	var user User

	if err := json.Unmarshal(resp, &user); err != nil {
		return nil, fmt.Errorf("Failed to decode api response: %w", err)
	}

	return &user, nil
}

// Return the pointer to instance of User by @Username.
// See https://dev.chunkbyte.space/api/metrica/GetByUsername
func (C *Client) GetByUsername(ctx context.Context, username string) (*User, error) {
	u, err := url.Parse(C.baseURL)

	u.Path = "/metrica/about"

	query := u.Query()
	query.Set("token", C.authToken)
	query.Set("id", username)

	u.RawQuery = query.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)

	if err != nil {
		return nil, fmt.Errorf("Failed to create request: %w", err)
	}

	resp, err := request.Request(ctx, C.httpClient, req)

	if err != nil {
		return nil, fmt.Errorf("Failed to execute request: %w", err)
	}

	var user User

	if err := json.Unmarshal(resp, &user); err != nil {
		return nil, fmt.Errorf("Failed to decode api response: %w", err)
	}

	return &user, nil
}

// Return the pointer to instance of User by Phone Number.
// See https://dev.chunkbyte.space/api/metrica/GetByPhoneNumber
func (C *Client) GetByPhoneNumber(ctx context.Context, phone string) (*User, error) {
	u, err := url.Parse(C.baseURL)

	u.Path = "/metrica/about"

	query := u.Query()
	query.Set("token", C.authToken)
	query.Set("id", phone)

	u.RawQuery = query.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)

	if err != nil {
		return nil, fmt.Errorf("Failed to create request: %w", err)
	}

	resp, err := request.Request(ctx, C.httpClient, req)

	if err != nil {
		return nil, fmt.Errorf("Failed to execute request: %w", err)
	}

	var user User

	if err := json.Unmarshal(resp, &user); err != nil {
		return nil, fmt.Errorf("Failed to decode api response: %w", err)
	}

	return &user, nil
}
