# metrica-api

A clean, idiomatic Go client for the ChunkByte Metrica API.
## Installation

go get github.com/GrandTheBest/metrica-api

## Usage

Here's a basic example of fetching a user profile by username.

```go
package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/GrandTheBest/metrica-api"
)

func main() {
	// 1. Initialize the client with desired options.
	// An authentication token is required.
	client := metrica.NewClient(
		metrica.WithAuthToken("YOUR_SECRET_TOKEN"),
		metrica.WithTimeout(10*time.Second),
	)

	// 2. Call an API method.
	// Always use a context for cancellation and deadlines.
	user, err := client.GetByUsername(context.Background(), "@derxS4C")
	if err != nil {
		log.Fatalf("API call failed: %v", err)
	}

	// 3. Process the response.
	// Always check that pointer fields are not nil before accessing them.
	if user.Profile != nil {
		fmt.Printf("Successfully fetched user: %s (ID: %d)\n", user.Profile.FirstName, user.Profile.ID)
	} else {
		fmt.Println("User data was returned, but the profile object is empty.")
	}
}
```

## Configuration

The client is configured using functional options passed to the NewClient constructor.

    WithAuthToken(string): (Required) Sets the authentication token for all requests.

    WithTimeout(time.Duration): Sets a timeout for all HTTP requests made by the client.

    WithHTTPClient(*http.Client): Overrides the default HTTP client with a custom implementation. Useful for advanced cases like custom transport or middleware.

API

The client provides the following methods to retrieve user data:

    GetByID(ctx context.Context, id string) (*User, error)

    GetByUsername(ctx context.Context, username string) (*User, error)

    GetByPhoneNumber(ctx context.Context, phone string) (*User, error)

All methods return a *User struct which contains the complete, nested user profile information.
License

This project is licensed under the GPL 3.0 License. See the LICENSE file for details.