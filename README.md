# X Bot (Twitter Bot with AWS Lambda & Go)

AWS LambdaとTwitter API v2を使ったGoのXボットです。

## 特徴

- AWS Lambdaで軽量・高速実行
- Twitter API v2対応
- セキュアな認証情報管理
- 簡単デプロイ
- テスト対応

## 必要な準備

### 1. Twitter Developer Account
1. Twitter Developer Platform でアカウント作成
2. 新しいアプリを作成
3. アプリの権限を「Read and Write」に設定
4. 以下の情報を取得：
   - API Key & Secret
   - Access Token & Secret

### 2. AWS環境
1. AWS CLIのインストール・設定
2. Lambda実行用のIAMロール作成
3. Lambda関数の作成

### 3. 開発環境
- Go 1.21以上
- Make (ビルド用)
- AWS CLI (デプロイ用)
- build-lambda-zip.exe (パッケージ作成用)
- Git

## セットアップ

### 1. リポジトリのクローン
```bash
git clone <this-repository>
cd x_bot
```

### 2. 環境変数の設定
```bash
# .envファイルを編集してTwitter認証情報を設定
```

### 3. ローカルでのビルドテスト
```bash
make build
```

## 使い方

### AWS Lambdaへのデプロイ
```bash
# 1. Lambda用ビルド
make build

# 2. デプロイパッケージ作成
make package

# 3. AWS Lambdaにデプロイ
make deploy

# 4. Lambda関数をテスト
make test-lambda

# 5. ログを確認
make logs
```

## プロジェクト構造

```
x_bot/
├── main.go                 # Lambda関数のエントリーポイント
├── pkg/
│   ├── twitter/
│   │   ├── client.go       # Twitter API v2クライアント
│   └── calc/
│       └── prime_factorization.go  # 素因数分解機能
├── Makefile               # ビルド・デプロイスクリプト
├── .env.example           # 環境変数設定例
├── .gitignore            # Git除外設定
└── README.md             # このファイル
```

## テスト

```bash
# Lambda関数をテスト
make test-lambda

# ログを確認
make logs
```

## 開発コマンド

```bash
# Lambda用ビルド
make build

# デプロイパッケージ作成
make package

# AWS Lambdaにデプロイ
make deploy

# Lambda関数をテスト
make test-lambda

# ログを確認
make logs

# クリーンアップ
make clean

# ヘルプ表示
make help
```

## セキュリティ

- 認証情報は環境変数で管理
- .envファイルはGitに含めない
- AWS Lambda環境変数で本番運用

## コントリビューション

1. このリポジトリをフォーク
2. 新しいブランチを作成
3. 変更をコミット
4. プルリクエストを作成

## ライセンス

このプロジェクトのライセンスは未定です。
