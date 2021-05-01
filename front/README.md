# docker 起動

```
$ docker-compose up -d --build
```

# アプリケーション起動

① docker 内へ移動

```
$ docker exec -it react-echo-front sh
```

② アプリケーション起動

```
/var/www # cd app/
/var/www/app # yarn start
```

以下にアクセス
http://localhost:3100/
