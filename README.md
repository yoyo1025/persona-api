# docker-go-postgresql-sampleapp

- docker x go x postgresql 環境のサンプルアプリケーション

# コマンド

1. このプロジェクトをpull
2. ルート直下に移動し以下のコマンドでコンテナをビルド起動する

```bash
docker-compose up -d --build
```

3. app コンテナにアクセス

```bash
docker-compose exec app sh
```

4. db コンテナにアクセス

```bash
docker-compose exec db bash
```

5. サンプルプログラムで動作の確認

```bash
$ docker-compose exec app sh
$ go run main.go
```

参考リンク
- https://www.wantedly.com/id/yosei_tadano/items/35b83014-e819-4203-990c-1eadd22a2417
