package main

import "github.com/ManuCiao10/go-hermes/monitor"

func main() {
	log := monitor.NewCustomLogger("test")
	log.Info("test", 4)
}

//init a certificate for a software (learn it and test)
