package twitter

import (
	"os"
	"testing"
)

func TestNewTwitterClient(t *testing.T) {
	// 環境変数をテスト用に設定
	os.Setenv("TWITTER_BEARER_TOKEN", "test_bearer_token")
	os.Setenv("TWITTER_API_KEY", "test_api_key")
	os.Setenv("TWITTER_API_SECRET", "test_api_secret")

	client := NewTwitterClient()

	if client.BearerToken != "test_bearer_token" {
		t.Errorf("Expected BearerToken to be 'test_bearer_token', got '%s'", client.BearerToken)
	}

	if client.APIKey != "test_api_key" {
		t.Errorf("Expected APIKey to be 'test_api_key', got '%s'", client.APIKey)
	}

	if client.APISecret != "test_api_secret" {
		t.Errorf("Expected APISecret to be 'test_api_secret', got '%s'", client.APISecret)
	}

	if client.baseURL != "https://api.twitter.com/2" {
		t.Errorf("Expected baseURL to be 'https://api.twitter.com/2', got '%s'", client.baseURL)
	}

	// テスト後にクリーンアップ
	os.Unsetenv("TWITTER_BEARER_TOKEN")
	os.Unsetenv("TWITTER_API_KEY")
	os.Unsetenv("TWITTER_API_SECRET")
}

func TestValidateCredentials(t *testing.T) {
	tests := []struct {
		name        string
		bearerToken string
		expectError bool
	}{
		{
			name:        "Valid bearer token",
			bearerToken: "valid_token",
			expectError: false,
		},
		{
			name:        "Empty bearer token",
			bearerToken: "",
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &TwitterClient{
				BearerToken: tt.bearerToken,
			}

			err := client.ValidateCredentials()

			if tt.expectError && err == nil {
				t.Error("Expected error but got none")
			}

			if !tt.expectError && err != nil {
				t.Errorf("Expected no error but got: %v", err)
			}
		})
	}
}
