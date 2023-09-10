package monitor

import (
	"bufio"
	"log"
	"math/rand"
	"os"
	"time"
)

var randomGenerator = rand.New(rand.NewSource(time.Now().UnixNano()))

func GetProxyFromFile(filename string) []string {
	pfolder := "proxies/"
	file, err := os.Open(pfolder + filename)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	var proxies []string

	for scanner.Scan() {
		proxies = append(proxies, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return proxies
}

func GetRandomPorxyFromArray(proxies []string) string {
	randomIndex := randomGenerator.Intn(len(proxies))

	return proxies[randomIndex]
}

func GetRandomProxyByName(filename string) string {
	pfolder := "proxies/"
	file, err := os.Open(pfolder + filename)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	var proxies []string

	for scanner.Scan() {
		proxies = append(proxies, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	randomIndex := randomGenerator.Intn(len(proxies))

	return proxies[randomIndex]
}
