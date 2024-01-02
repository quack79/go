package main

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	"github.com/quack79/golinks/backend"
	"github.com/quack79/golinks/backend/firestore"
	"github.com/quack79/golinks/backend/leveldb"
	"github.com/quack79/golinks/web"
)

func main() {
	pflag.String("addr", ":80", "Default port binding")
	pflag.Bool("adm", false, "Allow admin-level requests")
	pflag.String("version", "1.1", "Version string")
	pflag.String("backend", "leveldb", "Backend store to use. 'leveldb' and 'firestore' currently supported.")
	pflag.String("data", "data", "The location of the leveldb data directory")
	pflag.String("project", "", "The GCP project to use for the firestore backend. Will attempt to use application default creds if not defined.")
	pflag.String("host", "", "The host field to use when generating the source URL of a link. Defaults to the Host header of the generate request")
	pflag.Parse()

	if err := viper.BindPFlags(pflag.CommandLine); err != nil {
		log.Panic(err)
	}

	// allow env vars to set pflags
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))

	var backend backend.Backend

	switch viper.GetString("backend") {
	case "leveldb":
		var err error
		backend, err = leveldb.New(viper.GetString("data"))
		if err != nil {
			log.Panic(err)
		}
	case "firestore":
		var err error

		backend, err = firestore.New(context.Background(), viper.GetString("project"))
		if err != nil {
			log.Panic(err)
		}
	default:
		log.Panic(fmt.Sprintf("Unknown backend %s", viper.GetString("backend")))
	}

	defer backend.Close()

	log.Panic(web.ListenAndServe(backend))
}
