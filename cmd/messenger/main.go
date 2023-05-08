package main

import (
	"log"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/staurran/messengerKR.git/internal/app/dsn"
	"github.com/staurran/messengerKR.git/internal/app/servicedefault"
	"github.com/staurran/messengerKR.git/internal/pkg/app"
	"github.com/staurran/messengerKR.git/internal/pkg/app/server"
	"github.com/staurran/messengerKR.git/service/proto/authProto"
	"github.com/staurran/messengerKR.git/utils/logger"
)

func main() {
	logger.Init(servicedefault.NameService)
	log.Println("Application is starting")

	a := app.New()

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("failed to connect env" + err.Error())
	}
	db, err := gorm.Open(postgres.Open(dsn.FromEnv()), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect db" + err.Error())
	}

	connAuth, err := grpc.Dial(":8082", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer connAuth.Close()
	authClient := authProto.NewAuthClient(connAuth)

	serv := new(server.Server)
	opts := server.GetServerOptions()
	a.InitRoutes(db, authClient)
	err = serv.Run(opts, a.Router)
	if err != nil {
		log.Fatalf("error occured while server starting: %v", err)
	}
	log.Println("Application terminate")
}
