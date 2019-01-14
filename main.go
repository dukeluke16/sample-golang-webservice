package main

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/dukeluke16/sample-golang-webservice/logger"
	"github.com/dukeluke16/sample-golang-webservice/web"
)

// start for mocking out web.Start
var start = web.Start

// fatal for mocking out log.Fatal
var fatal = log.Fatal

func main() {
	logger.Init(ioutil.Discard, os.Stdout, os.Stdout, os.Stderr)
	err := start()
	if err != nil {
		fatal(err)
	}
}
