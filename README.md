## Getting Started
1. このリポジトリをフォーク
2. `Settings > General > Template repository Loading`にチェックをつける
![Image](https://github.com/user-attachments/assets/617d2ae4-9248-4e1d-b37a-7d1f48e31ac2)
3. 新規リポジトリ作成時、`Repository template`のリストから当該リポジトリを選択
![Image](https://github.com/user-attachments/assets/b1937b9d-660b-4f1f-8f70-c7361a0766c3)
4. `.env.copy`を複製し、`.env`ファイルにリネーム
5. `.env`の環境変数に任意の値を入れる
6. `make build_up`コマンドでプロジェクトコンテナ立ち上げ（Docker Desktopダウンロード必須）

## Project Structure
```
project_template/
├── backend/                 # Go製のバックエンド（クリーンアーキテクチャ）
│   ├── cmd/                 # エントリーポイント（main.goのみ）
│   │   └── api/             # API起動用のmainパッケージ
│   ├── domain/              # ドメイン層：ビジネスエンティティとコアロジック
│   │   ├── entity/          # ビジネスエンティティの定義
│   │   └── repository/      # リポジトリのインターフェース定義
│   ├── usecase/             # ユースケース層：アプリケーションのビジネスロジック
│   │   ├── dto/             # ユースケース層のデータ転送オブジェクト
│   │   └── interactor/      # ユースケースの具体的な実装
│   ├── adapter/             # アダプター層：外部とのインターフェース
│   │   ├── handler/         # HTTPリクエストの処理（エンドポイントの実装）
│   │   ├── router/          # エンドポイントルーティングの設定
│   │   └── repository/      # リポジトリの具体的な実装
│   ├── infrastructure/      # インフラストラクチャ層：外部サービスとのやりとり
│   │   ├── bootstrap/       # 初期化処理（DB接続、環境変数の読み込みなど）
│   │   ├── config/          # 設定に関する処理（環境変数の読み取りなど）
│   │   └── db/              # データベースに関する処理
│   │       ├── migration/   # マイグレーションファイル群
│   │       └── seed/        # シードデータ群
│   ├── go.mod               # Goモジュールの設定ファイル（依存関係の管理など）
│   └── tmp/                 # 一時ファイル置き場（airによるビルド結果など）
├── frontend/                # Next.js製のフロントエンド（クリーンアーキテクチャ）
│   ├── __tests__/           # テストコードファイル群
│   ├── public/              # 画像などの静的アセット
│   └── src/
│       ├── app/             # App Router ディレクトリ（Next.jsのルーティング設定）
│       ├── domain/          # ドメイン層：ビジネスエンティティとロジック
│       │   ├── entity/      # ビジネスエンティティの定義
│       │   └── repository/  # リポジトリのインターフェース定義
│       ├── application/     # アプリケーション層：ユースケースとビジネスロジック
│       │   ├── usecase/     # ユースケースの定義
│       │   └── dto/         # データ転送オブジェクト
│       ├── adapter/         # アダプター層：外部とのインターフェース
│       │   ├── components/  # 再利用可能なReactコンポーネント群
│       │   ├── hooks/       # カスタムReactフック
│       │   ├── pages/       # ページコンポーネント（App Router使用時は省略可）
│       │   └── state/       # 状態管理（Recoilなど）
│       │       ├── atoms/   # 状態（Atom）の定義
│       │       ├── selectors/# 派生状態（Selector）の定義
│       │       └── effects/ # Atomに副作用を与えるEffectの定義
│       ├── infrastructure/  # インフラストラクチャ層：外部サービスとの連携
│       │   ├── api/         # APIクライアント
│       │   └── storage/     # ローカルストレージなどの永続化
│       ├── styles/          # CSSやスタイル関連ファイル
│       ├── types/           # 型定義ファイル（TypeScript用）
│       └── utils/           # 共通ユーティリティ関数群
├── .env                     # 環境変数設定ファイル
├── .gitignore               # Gitの追跡対象から除外するファイル設定
├── docker-compose.yml       # Dockerサービスの設定ファイル
└── Makefile                 # makeコマンドによるタスク定義ファイル
```
