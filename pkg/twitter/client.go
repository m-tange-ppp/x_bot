package twitter

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

// TwitterClient はTwitter API v2のクライアント
type TwitterClient struct {
	BearerToken  string
	APIKey       string
	APISecret    string
	AccessToken  string
	AccessSecret string
	baseURL      string
	httpClient   *http.Client
}

// TweetRequest はツイート投稿のリクエスト構造体
type TweetRequest struct {
	Text string `json:"text"`
}

// TweetResponse はツイート投稿のレスポンス構造体
type TweetResponse struct {
	Data struct {
		ID   string `json:"id"`
		Text string `json:"text"`
	} `json:"data"`
}

// NewTwitterClient は新しいTwitterクライアントを作成
func NewTwitterClient() *TwitterClient {
	return &TwitterClient{
		BearerToken:  os.Getenv("TWITTER_BEARER_TOKEN"),
		APIKey:       os.Getenv("TWITTER_API_KEY"),
		APISecret:    os.Getenv("TWITTER_API_SECRET"),
		AccessToken:  os.Getenv("TWITTER_ACCESS_TOKEN"),
		AccessSecret: os.Getenv("TWITTER_ACCESS_TOKEN_SECRET"),
		baseURL:      "https://api.twitter.com/2",
		httpClient:   &http.Client{Timeout: 30 * time.Second},
	}
}

// PostTweet はツイートを投稿する（OAuth 1.0a認証）
func (c *TwitterClient) PostTweet(text string) (*TweetResponse, error) {
	if c.APIKey == "" || c.APISecret == "" || c.AccessToken == "" || c.AccessSecret == "" {
		return nil, fmt.Errorf("OAuth 1.0a credentials are required: API_KEY, API_SECRET, ACCESS_TOKEN, ACCESS_TOKEN_SECRET")
	}

	url := fmt.Sprintf("%s/tweets", c.baseURL)

	tweetReq := TweetRequest{Text: text}
	jsonData, err := json.Marshal(tweetReq)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal tweet request: %w", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// OAuth 1.0a認証ヘッダーを生成
	authHeader, err := c.generateOAuth1Header("POST", url, "")
	if err != nil {
		return nil, fmt.Errorf("failed to generate OAuth header: %w", err)
	}

	req.Header.Set("Authorization", authHeader)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return nil, c.handleAPIError(resp)
	}

	var tweetResp TweetResponse
	if err := json.NewDecoder(resp.Body).Decode(&tweetResp); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &tweetResp, nil
}

// generateOAuth1Header はOAuth 1.0a認証ヘッダーを生成
func (c *TwitterClient) generateOAuth1Header(method, requestURL, queryString string) (string, error) {
	// OAuth 1.0aパラメータ
	nonce := c.generateNonce()
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)

	params := map[string]string{
		"oauth_consumer_key":     c.APIKey,
		"oauth_nonce":            nonce,
		"oauth_signature_method": "HMAC-SHA1",
		"oauth_timestamp":        timestamp,
		"oauth_token":            c.AccessToken,
		"oauth_version":          "1.0",
	}

	// 署名ベース文字列を作成
	signature := c.generateSignature(method, requestURL, params, queryString)
	params["oauth_signature"] = signature

	// 認証ヘッダーを構築
	var headerParts []string
	for key, value := range params {
		headerParts = append(headerParts, fmt.Sprintf(`%s="%s"`, key, url.QueryEscape(value)))
	}
	sort.Strings(headerParts)

	return "OAuth " + strings.Join(headerParts, ", "), nil
}

// generateSignature はOAuth 1.0a署名を生成
func (c *TwitterClient) generateSignature(method, requestURL string, params map[string]string, queryString string) string {
	// パラメータを収集
	var paramPairs []string
	for key, value := range params {
		paramPairs = append(paramPairs, fmt.Sprintf("%s=%s", url.QueryEscape(key), url.QueryEscape(value)))
	}

	// クエリ文字列も追加
	if queryString != "" {
		paramPairs = append(paramPairs, queryString)
	}

	sort.Strings(paramPairs)
	paramString := strings.Join(paramPairs, "&")

	// 署名ベース文字列を作成
	signatureBaseString := fmt.Sprintf("%s&%s&%s",
		url.QueryEscape(method),
		url.QueryEscape(requestURL),
		url.QueryEscape(paramString))

	// 署名キーを作成
	signingKey := fmt.Sprintf("%s&%s", url.QueryEscape(c.APISecret), url.QueryEscape(c.AccessSecret))

	// HMAC-SHA1署名を生成
	h := hmac.New(sha1.New, []byte(signingKey))
	h.Write([]byte(signatureBaseString))
	signature := base64.StdEncoding.EncodeToString(h.Sum(nil))

	return signature
}

// generateNonce はランダムなnonceを生成
func (c *TwitterClient) generateNonce() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, 32)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

// handleAPIError はAPIエラーを処理する
func (c *TwitterClient) handleAPIError(resp *http.Response) error {
	switch resp.StatusCode {
	case 401:
		return fmt.Errorf("unauthorized: check your Twitter API credentials")
	case 403:
		return fmt.Errorf("forbidden: insufficient permissions")
	case 429:
		return fmt.Errorf("rate limit exceeded: please wait before making another request")
	default:
		return fmt.Errorf("API request failed with status: %s", resp.Status)
	}
}

// ValidateCredentials は認証情報が設定されているかチェック
func (c *TwitterClient) ValidateCredentials() error {
	if c.APIKey == "" {
		return fmt.Errorf("TWITTER_API_KEY is required")
	}
	if c.APISecret == "" {
		return fmt.Errorf("TWITTER_API_SECRET is required")
	}
	if c.AccessToken == "" {
		return fmt.Errorf("TWITTER_ACCESS_TOKEN is required")
	}
	if c.AccessSecret == "" {
		return fmt.Errorf("TWITTER_ACCESS_TOKEN_SECRET is required")
	}
	return nil
}
