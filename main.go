package main

import (
	"github.com/u00io/gomisc/logger"
	"github.com/u00io/u00node/localstorage"
	"github.com/u00io/u00node/mainform"
)

func main() {
	localstorage.Init("u00node")
	logger.Init(localstorage.Path() + "/logs")
	mainform.Run()
}
