## Project Structure
```
project_template/
├── backend/                 # Go製のバックエンド（クリーンアーキテクチャ）
│   ├── cmd/                 # エントリーポイント（main.goなど）
│   │   └── api/             # API起動用のmainパッケージ
│   ├── config/              # 設定に関する処理（環境変数の読み取りなど）
│   ├── handler/             # HTTPリクエストの処理（エンドポイントの実装）
│   ├── infrastructure/      # 外部サービスとのやりとりの実装
│   │   └── db/              # データベースに関する処理
│   │       ├── generated/   # sqlcなどで自動生成されたコード
│   │       ├── query/       # SQLクエリファイル
│   │       ├── init.sql     # テーブル定義
│   │       ├── migration.go # テーブルマイグレーション処理
│   │       └── sqlc.yaml    # sqlcの設定ファイル
│   ├── repository/          # ドメインレイヤーのインターフェース定義
│   ├── repository_impl/     # リポジトリインターフェースの具体実装
│   ├── usecase/             # ビジネスロジック（ユースケース層）
│   ├── go.mod               # Goモジュールの設定ファイル（依存関係の管理など）
│   └── tmp/                 # 一時ファイル置き場（airによるビルド結果など）
├── frontend/                # Next.js製のフロントエンド
│   ├── public/              # 画像などの静的アセット
│   └── src/
│       ├── app/             # App Router ディレクトリ（Next.jsのルーティング設定）
│       ├── components/      # 再利用可能なReactコンポーネント群
│       ├── hooks/           # カスタムReactフック
│       ├── page/            # 画面単位のページコンポーネント
│       ├── provider/        # グローバル状態などのContextプロバイダー
│       ├── recoil/          # Recoilの状態定義
│       │   ├── atoms/       # 状態（Atom）の定義
│       │   ├── selectors/   # 派生状態（Selector）の定義
│       │   └── effects/     # Atomに副作用を与えるEffectの定義
│       ├── styles/          # CSSやスタイル関連ファイル
│       ├── types/           # 型定義ファイル（TypeScript用）
│       └── utils/           # 共通ユーティリティ関数群
├── .env                     # 環境変数設定ファイル
├── .gitignore               # Gitの追跡対象から除外するファイル設定
├── docker-compose.yml       # Dockerサービスの設定ファイル
└── Makefile                 # makeコマンドによるタスク定義ファイル
```