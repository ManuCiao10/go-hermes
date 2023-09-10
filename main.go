package main

import "github.com/ManuCiao10/go-hermes/monitor"

//fare package per usare in altri funzioni

//init a certificate for a software (learn it and test)

func main() {
	logger := monitor.NewCustomLogger("9373DK", "DE")
	logger.Info("staring task", "200 OK")
}
