package main

import (
	"bufio"
	"log"
	"net/http"
	"os"
	"regexp"

	app "github.com/marcos-dev88/checkpass/application"
	"github.com/marcos-dev88/checkpass/handler"
)

var regexEnvs = regexp.MustCompile(`(\S+)=(\S+)`)

func init() {
	if err := defineEnvs(".env"); err != nil {
		log.Printf("error to define envs: %v", err)
	}
}

func main() {

	defaultPort := os.Getenv("API_PORT")

	a := app.NewApp(nil)
	h := handler.NewHandler(a)

	router := http.NewServeMux()
	middleware := handler.NewMiddleware()

	router.HandleFunc("/verify", middleware.CallMiddleware(h.HandlePassword))

	if len(defaultPort) == 0 {
		defaultPort = "8000"
	}

	log.Println("server is running on port: " + defaultPort)
	log.Fatal(http.ListenAndServe(":"+defaultPort, router))
}

// defineEnvs: get all environment variables from .env file
// curious about this code?
// check: https://gist.github.com/marcos-dev88/8df0c09d4dfcfc1cb0e0aa30311c37b5
func defineEnvs(filename string) error {
	file, err := os.Open(filename)

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	if err != nil {
		return err
	}

	sc := bufio.NewScanner(file)

	for sc.Scan() {
		envMatch := regexEnvs.FindStringSubmatch(sc.Text())
		if envMatch != nil {
			err := os.Setenv(envMatch[1], envMatch[2])
			if err != nil {
				return err
			}
		}
	}

	return nil
}
