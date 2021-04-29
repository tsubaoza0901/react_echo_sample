package main

import (
	"log"
	"react-echo-sample/adapter/http/router"
	"react-echo-sample/infrastructure/middleware"
	"react-echo-sample/infrastructure/rdb/connection"
	"react-echo-sample/interactor"
	"time"

	"github.com/labstack/echo/v4"
)

const location = "Asia/Tokyo"

func setTimeZone() {
	loc, err := time.LoadLocation(location)
	if err != nil {
		loc = time.FixedZone(location, 9*60*60)
	}
	time.Local = loc
}

func main() {
	// ローカル時間設定
	setTimeZone()

	// DBセットアップ
	conn := connection.InitRDB()
	sqlDB, err := conn.DB()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if sqlDB != nil {
			if err := sqlDB.Close(); err != nil {
				log.Fatal(err)
			}
		}
	}()

	// Interactorのインスタンス生成
	i := interactor.NewInteractor(conn)

	// Handlerのインスタンス生成
	h := i.NewAppHandler()

	// Echoのインスタンス生成
	e := echo.New()

	// 全てのリクエストで差し込みたいミドルウェアをセット
	middleware.InitMiddleware(e)

	// ルーティングをセット
	router.InitRouting(e, h)

	err = e.Start(":9090")
	if err != nil {
		log.Fatal(err)
		// zap.S().Fatalw("HTTP Server 起動エラー", zap.Error(err))
	}
}
