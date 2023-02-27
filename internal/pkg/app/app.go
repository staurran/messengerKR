package app

import (
	"context"

	"github.com/gin-gonic/gin"
	"messengerKR/internal/app/dsn"
	"messengerKR/internal/app/repository"
)

type Application struct {
	Conf   *config.Config
	repo   *repository.Repository
	Router gin.IRouter
}

func New(ctx context.Context) (*Application, error) {
	cnf, err := config.NewConfig(ctx)
	if err != nil {
		return nil, err
	}

	dsnStr := dsn.FromEnv()
	rep, err := repository.New(dsnStr)
	if err != nil {
		return nil, err
	}
	a := &Application{Conf: cnf, repo: rep}

	return a, nil
}
