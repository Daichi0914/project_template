-- 拡張機能（UUID生成用）
CREATE EXTENSION IF NOT EXISTS "pgcrypto";

-- タスクテーブル
CREATE TABLE IF NOT EXISTS tasks (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name TEXT NOT NULL,
    estimate_minutes INTEGER NOT NULL,
    started_at TIMESTAMPTZ NOT NULL,
    stopped_at TIMESTAMPTZ
);
