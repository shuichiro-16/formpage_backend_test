# Go バックエンドサーバ (フォームデータ保存)

## 概要

React フロントエンドから送信されたフォームデータを API 経由で受け取り、MySQL に保存する Go バックエンドサーバです。

## 構成

- `/form` (POST): フォームデータ受け取り・保存 API
- MySQL: データ保存用

## セットアップ

1. `.env`ファイルの MySQL 情報を編集
2. MySQL で以下のテーブルを作成

```sql
CREATE DATABASE formdb;
USE formdb;
CREATE TABLE forms (
	id INT AUTO_INCREMENT PRIMARY KEY,
	name VARCHAR(255),
	email VARCHAR(255),
	message TEXT
);
```

3. 依存パッケージインストール

```bash
go mod tidy
```

4. サーバ起動

```bash
cd cmd
go run main.go
```

## API 例

POST `/form`

```json
{
  "name": "山田太郎",
  "email": "taro@example.com",
  "message": "こんにちは"
}
```

# formpage_backend_test
