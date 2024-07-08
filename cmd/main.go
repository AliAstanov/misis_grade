package main

import (
	"log"
	"misis_baholar/api"
	"misis_baholar/config"
	db "misis_baholar/pkg"
	"misis_baholar/pkg/utils"
	"misis_baholar/storage"
)

func main() {

	// to := []string{"aliastan1997@gmail.com","muhammadjanw1@gmail.com"}

	// smt.SendEmail(to,"hello quchqor")
	utils.Otp()

	cfg := config.Load()

	db, err := db.ConnectToDb(cfg.PgConfig)
	if err != nil {
		log.Println("error on connect to ConToDb:", err)
		return
	}
	storage := storage.NewStorage(db)

	api.Api(storage)

}
