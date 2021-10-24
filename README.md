# Auth0でCookieセッション張るサンプル

## サーバサイド
- go + ginで実装
- ほぼ　https://github.com/auth0-samples/auth0-golang-web-app　の内容
- session storageとしてredisを利用
- `.env.sample` を `.env` にリネームし、Auth0の各種設定を記入すること

## クライアントサイド
- node + redis (hooks) で実装

## How to run
```bash
docker-compose build
docker-compose up -d
```
http://localhost:3000 にブラウザからアクセスする