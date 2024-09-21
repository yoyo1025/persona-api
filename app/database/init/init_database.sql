-- CREATE DATABASE IF NOT EXISTS sample_app;
-- 1. usersテーブルの作成
CREATE TABLE IF NOT EXISTS users (
  id SERIAL PRIMARY KEY,
  name VARCHAR(100),
  email VARCHAR(100) UNIQUE NOT NULL,
  password VARCHAR(255)
);

-- id=1のユーザーを挿入
INSERT INTO users (id, name, email, password)
VALUES (1, 'test1', 'test1@example.com', 'password')
ON CONFLICT DO NOTHING;

-- 2. personaテーブルの作成
CREATE TABLE IF NOT EXISTS persona (
  id SERIAL PRIMARY KEY,
  name VARCHAR(100),
  user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
  sex VARCHAR(10) CHECK (sex IN ('male', 'female', 'other')),
  age INTEGER,
  profession VARCHAR(100),
  problems TEXT,
  behavior TEXT
);

-- 3. conversationテーブルの作成
CREATE TABLE IF NOT EXISTS conversation (
  id SERIAL PRIMARY KEY,
  user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
  persona_id INTEGER REFERENCES persona(id) ON DELETE CASCADE
);

-- 4. commentテーブルの作成
CREATE TABLE IF NOT EXISTS comment (
  id SERIAL PRIMARY KEY,
  user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
  persona_id INTEGER REFERENCES persona(id) ON DELETE CASCADE,
  comment TEXT NOT NULL,
  is_user_comment BOOLEAN,
  good BOOLEAN DEFAULT FALSE
);
