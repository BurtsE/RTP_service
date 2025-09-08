package main

import (
	"context"
	"flag"
	"log"
	"multiplicator/internal/config"
	"multiplicator/internal/server"
	"multiplicator/internal/service"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)
		<-c
		cancel()
	}()

	rtp := flag.Float64("rtp", 1, "set desired rtp as a float value")
	flag.Parse()

	if *rtp <= 0 || *rtp > 1 {
		log.Fatal("invalid rtp value")
	}

	log.Printf("rtp values set to: %f\n", *rtp)

	calibration := config.NewCalibration("./config/calibration.json")
	servise := service.NewService(*rtp, calibration)

	config := config.NewConfig("./config/config.json")

	log.Printf("starting server with port: %d\n", config.Port)

	server := server.NewHTTPServer(config, servise)
	go server.Start()

	<-ctx.Done()
	server.Stop(context.Background())
	log.Println("app shutdown")

}
