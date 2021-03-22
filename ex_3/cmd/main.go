package main

import (
	"ex_3/api/controller"
	"github.com/rs/zerolog/log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	ctrl := controller.NewController()
	ctrl.Start(":1323")

	{
		<-stop

		log.Info().Msg("Closing service..")
		ctrl.Stop()

		log.Info().Msg("Service closed! Have a nice day :)")
	}
}
