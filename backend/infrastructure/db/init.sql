-- 拡張機能（UUID生成用）
CREATE EXTENSION IF NOT EXISTS "pgcrypto";

-- ユーザーテーブル
CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    username TEXT NOT NULL UNIQUE,
    email TEXT NOT NULL UNIQUE,  -- 追加: emailカラムをユニークに設定
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- サンプルユーザーデータ
INSERT INTO users (username, email)
VALUES
    ('demo_user', 'demo@example.com')
ON CONFLICT (username) DO NOTHING;
