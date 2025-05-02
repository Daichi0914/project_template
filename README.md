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
│   ├── cmd/                 # エントリーポイント（main.goのみ、他処理はinternalに分離）
│   │   └── api/             # API起動用のmainパッケージ
│   ├── config/              # 設定に関する処理（環境変数の読み取りなど）
│   ├── handler/             # HTTPリクエストの処理（エンドポイントの実装）
│   ├── infrastructure/      # 外部サービスとのやりとりの実装
│   │   └── db/              # データベースに関する処理
│   │       ├── migration/   # マイグレーションファイル群
│   │       └── seed/        # シードデータ群
│   ├── repository/          # ドメインレイヤーのインターフェース定義
│   ├── usecase/             # ビジネスロジック（ユースケース層）
│   ├── go.mod               # Goモジュールの設定ファイル（依存関係の管理など）
│   ├── tmp/                 # 一時ファイル置き場（airによるビルド結果など）
│   └── internal/            # アプリケーション内部ロジック（クリーンアーキテクチャ補助）
│       ├── bootstrap/       # 初期化処理（DB接続、環境変数の読み込みなど）
│       └── router/          # エンドポイントルーティングの設定
├── frontend/                # Next.js製のフロントエンド
│   ├── __tests__/           # テストコードファイル群
│   ├── public/              # 画像などの静的アセット
│   └── src/
│       ├── app/             # App Router ディレクトリ（Next.jsのルーティング設定）
│       ├── components/      # 再利用可能なReactコンポーネント群
│       ├── hooks/           # カスタムReactフック
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