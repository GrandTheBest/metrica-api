package request

import (
	"context"
	"fmt"
	"io"
	"net/http"
)

func Request(ctx context.Context, client *http.Client, req *http.Request) ([]byte, error) {
	req = req.WithContext(ctx)

	resp, err := client.Do(req)

	if err != nil {
		return nil, fmt.Errorf("Failed to execute http request: %w", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		// Try to read the body to include it in the error message.
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			// If we can't even read the body, just return the status code.
			return nil, fmt.Errorf("API request failed with status code: %d (and failed to read body)", resp.StatusCode)
		}
		return nil, fmt.Errorf("API request failed with status code: %d, body: %s", resp.StatusCode, string(bodyBytes))
	}
	bytes, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, fmt.Errorf("Failed to read response body: %w", err)
	}

	return bytes, nil
}
