package monitor

import (
	"bufio"
	"log"
	"math/rand"
	"os"
	"time"
)

var randomGenerator = rand.New(rand.NewSource(time.Now().UnixNano()))

func GetRandomProxyFromArray(proxies []string) string {
	randomIndex := randomGenerator.Intn(len(proxies))

	return proxies[randomIndex]
}

func GetProxyFromFile(filename string) []string {
	file, err := os.Open(filename)
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

func GetRandomProxyByName(filename string) string {
	file, err := os.Open(filename)

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
