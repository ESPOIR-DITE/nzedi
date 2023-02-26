package main

import (
	"github.com/ESPOIR-DITE/nzedi.git/pkg/config/nzedi/env"
	"github.com/ESPOIR-DITE/nzedi.git/pkg/logger"
	"github.com/ESPOIR-DITE/nzedi.git/pkg/wire"
	"os"
)

const version = "0.0.1"

func main() {
	logger.Log.Infof("starting Nzedi Api v: %s", version)
	os.Exit(wire.Run(env.NewNzediApiConfigManagerImpl()))
}
