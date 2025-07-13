package main

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"x_bot/pkg/calc"
	"x_bot/pkg/twitter"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// Handler はLambda関数のメインハンドラー
func Handler(ctx context.Context, event events.CloudWatchEvent) error {
	log.Printf("🚀 X Bot Lambda function started at: %s", time.Now().Format(time.RFC3339))

	// Twitterクライアントを初期化
	twitterClient := twitter.NewTwitterClient()

	// 認証情報をチェック
	if err := twitterClient.ValidateCredentials(); err != nil {
		log.Printf("❌ Twitter credentials validation failed: %v", err)
		return fmt.Errorf("twitter credentials error: %w", err)
	}

	// 日本の現在時刻を取得
	jst := time.FixedZone("JST", 9*60*60)
	jstNow := time.Now().In(jst)

	// 日付をintで取得
	todayInt, err := strconv.Atoi(jstNow.Format("20060102"))
	if err != nil {
		return fmt.Errorf("time format error: %w", err)
	}

	// 日付を素因数分解した結果を文字列で取得
	strNumbers := []string{}
	for _, num := range calc.PrimeFactorize(todayInt) {
		strNumbers = append(strNumbers, strconv.Itoa(num))
	}
	primeFactorization := strings.Join(strNumbers, " * ")

	// テストツイートを投稿
	message := fmt.Sprintf("%d = %s", todayInt, primeFactorization)

	response, err := twitterClient.PostTweet(message)
	if err != nil {
		log.Printf("❌ Failed to post tweet: %v", err)
		return fmt.Errorf("failed to post tweet: %w", err)
	}

	log.Printf("✅ Tweet posted successfully! ID: %s", response.Data.ID)
	log.Printf("📝 Tweet content: %s", response.Data.Text)

	return nil
}

func main() {
	lambda.Start(Handler)
}
