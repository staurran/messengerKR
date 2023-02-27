package messenger

import (
	"context"
	"log"
	"os"

	"github.com/staurran/messengerKR.git/internal/pkg/app"
)

func main() {
	log.Println("Application start")
	ctx := context.Background()
	a, err := app.New(ctx)
	if err != nil {
		log.Println("Application failed")
		os.Exit(2)
	}
	a.StartServer()
	log.Println("Application terminate")
}
