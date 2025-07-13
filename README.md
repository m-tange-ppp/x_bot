# 🤖 X Bot (Twitter Bot with AWS Lambda & Go)

AWS LambdaとTwitter API v2を使ったGoのXボットです！

## 🚀 特徴

- ⚡ AWS Lambdaで軽量・高速実行
- 🐦 Twitter API v2対応
- 🛡️ セキュアな認証情報管理
- 📦 簡単デプロイ
- 🧪 テスト対応

## 📋 必要な準備

### 1. Twitter Developer Account
1. [Twitter Developer Platform](https://developer.twitter.com/) でアカウント作成
2. 新しいアプリを作成
3. 以下の情報を取得：
   - Bearer Token
   - API Key & Secret
   - Access Token & Secret

### 2. AWS環境
1. AWS CLIのインストール・設定
2. Lambda実行用のIAMロール作成
3. Lambda関数の作成

### 3. 開発環境
- Go 1.21以上
- Make (ビルド用)
- Git

## ⚙️ セットアップ

### 1. リポジトリのクローン
```bash
git clone <this-repository>
cd x_bot
```

### 2. 依存関係のインストール
```bash
make deps
```

### 3. 環境変数の設定
```bash
# .env.exampleをコピーして設定
cp .env.example .env
# .envファイルを編集してTwitter認証情報を設定
```

### 4. ローカルテスト用ビルド
```bash
make build-local
```

## 🛠️ 使い方

### ローカルでのテスト
```bash
# 環境変数を設定してから実行
export TWITTER_BEARER_TOKEN="your_token_here"
./main.exe
```

### AWS Lambdaへのデプロイ
```bash
# 1. Lambda用ビルド
make build

# 2. デプロイパッケージ作成
make package

# 3. AWS Lambdaにデプロイ（AWS CLI設定済みの場合）
make deploy
```

## 📁 プロジェクト構造

```
x_bot/
├── main.go                 # Lambda関数のエントリーポイント
├── pkg/
│   └── twitter/
│       └── client.go       # Twitter API v2クライアント
├── Makefile               # ビルド・デプロイスクリプト
├── .env.example           # 環境変数設定例
├── .gitignore            # Git除外設定
└── README.md             # このファイル
```

## 🧪 テスト

```bash
make test
```

## 🔧 開発コマンド

```bash
# 依存関係インストール
make deps

# ローカル用ビルド
make build-local

# Lambda用ビルド
make build

# デプロイパッケージ作成
make package

# クリーンアップ
make clean

# テスト実行
make test

# リント実行
make lint

# ヘルプ表示
make help
```

## 🛡️ セキュリティ

- 認証情報は環境変数で管理
- .envファイルはGitに含めない
- AWS Lambda環境変数で本番運用

## 📈 今後の拡張予定

- [ ] 定期ツイート機能
- [ ] メンション監視・自動返信
- [ ] 画像付きツイート
- [ ] DM機能
- [ ] Analytics連携

## 🤝 コントリビューション

1. このリポジトリをフォーク
2. 新しいブランチを作成
3. 変更をコミット
4. プルリクエストを作成

## 📝 ライセンス

MIT License

## 💬 サポート

何か問題があったらIssueを作成してください！

---

🎉 Happy Botting! 🤖✨
