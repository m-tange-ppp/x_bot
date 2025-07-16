package twitter

import (
	"os"
	"testing"
)

func TestNewTwitterClient(t *testing.T) {
	// 環境変数をテスト用に設定
	os.Setenv("TWITTER_API_KEY", "test_api_key")
	os.Setenv("TWITTER_API_SECRET", "test_api_secret")
	os.Setenv("TWITTER_ACCESS_TOKEN", "test_access_token")
	os.Setenv("TWITTER_ACCESS_TOKEN_SECRET", "test_access_token_secret")

	client := NewTwitterClient()

	if client.APIKey != "test_api_key" {
		t.Errorf("Expected APIKey to be 'test_api_key', got '%s'", client.APIKey)
	}

	if client.APISecret != "test_api_secret" {
		t.Errorf("Expected APISecret to be 'test_api_secret', got '%s'", client.APISecret)
	}

	if client.AccessToken != "test_access_token" {
		t.Errorf("Expected AccessToken to be 'test_access_token', got '%s'", client.AccessToken)
	}

	if client.AccessSecret != "test_access_token_secret" {
		t.Errorf("Expected AccessSecret to be 'test_access_token_secret', got '%s'", client.AccessSecret)
	}

	if client.baseURL != "https://api.twitter.com/2" {
		t.Errorf("Expected baseURL to be 'https://api.twitter.com/2', got '%s'", client.baseURL)
	}

	// テスト後にクリーンアップ
	os.Unsetenv("TWITTER_API_KEY")
	os.Unsetenv("TWITTER_API_SECRET")
	os.Unsetenv("TWITTER_ACCESS_TOKEN")
	os.Unsetenv("TWITTER_ACCESS_TOKEN_SECRET")
}

func TestValidateCredentials(t *testing.T) {
	tests := []struct {
		name         string
		apiKey       string
		apiSecret    string
		accessToken  string
		accessSecret string
		expectError  bool
	}{
		{
			name:         "Valid credentials",
			apiKey:       "valid_key",
			apiSecret:    "valid_secret",
			accessToken:  "valid_token",
			accessSecret: "valid_access_secret",
			expectError:  false,
		},
		{
			name:         "Empty API key",
			apiKey:       "",
			apiSecret:    "valid_secret",
			accessToken:  "valid_token",
			accessSecret: "valid_access_secret",
			expectError:  true,
		},
		{
			name:         "Empty API secret",
			apiKey:       "valid_key",
			apiSecret:    "",
			accessToken:  "valid_token",
			accessSecret: "valid_access_secret",
			expectError:  true,
		},
		{
			name:         "Empty access token",
			apiKey:       "valid_key",
			apiSecret:    "valid_secret",
			accessToken:  "",
			accessSecret: "valid_access_secret",
			expectError:  true,
		},
		{
			name:         "Empty access secret",
			apiKey:       "valid_key",
			apiSecret:    "valid_secret",
			accessToken:  "valid_token",
			accessSecret: "",
			expectError:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &TwitterClient{
				APIKey:       tt.apiKey,
				APISecret:    tt.apiSecret,
				AccessToken:  tt.accessToken,
				AccessSecret: tt.accessSecret,
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
