# 概要

# 主な使用技術

# 使用方法
## 1．Docker イメージのビルド&コンテナの起動

```
$ docker-compose up -d --build
```

## 2．データベースの作成
① DB コンテナ内へ移動

```
$ docker exec -it react-echo-db bash
```

② DB 接続

```
root@ec19d85976f4:/# mysql -u root -h db -p
Enter password: root
```

③ DB 作成

```
mysql> CREATE DATABASE reactechosample;
```

## 3．マイグレーションファイルの実行
① アプリケーションコンテナ内へ移動

```
$ docker exec -it react-echo-backend bash
```

② マイグレーションファイルの実行

```
root@fe385569a625:/go/src/app# goose up
```

## 4．アプリケーションの起動

① main.goがあるsrc配下に移動

```
root@fe385569a625:/go/app# cd src
```

② アプリケーションの起動

```
root@fe385569a625:/go/app/src# go run main.go
```

# 注意点