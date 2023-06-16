package main

import (
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/valexz/opsgbot/internal/config"
	"github.com/valexz/opsgbot/internal/updater"
	"github.com/voxelbrain/goptions"
)

type options struct {
	Verbose bool          `goptions:"-v, --verbose, description='Log verbosely'"`
	Pretty  bool          `goptions:"-p, --pretty, description='Log multi-line pretty-print json'"`
	Help    goptions.Help `goptions:"-h, --help, description='Show help'"`
	Config  string        `goptions:"-c, --config, description='Path to yaml config file'"`
	EnvFile string        `goptions:"-e, --env-file, description='Path to environment variable file'"`
}

func main() {
	parsedOptions := options{}

	parsedOptions.Config = "./config.yml"

	goptions.ParseAndFail(&parsedOptions)

	if parsedOptions.Verbose {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}

	log.SetFormatter(&log.JSONFormatter{
		PrettyPrint: parsedOptions.Pretty,
		FieldMap: log.FieldMap{
			log.FieldKeyLevel: "severity", // required for StackDriver GKE parsing
		},
		TimestampFormat: "2006-01-02T15:04:05.000000000Z", // nanoseconds
	})

	log.Debug("Logging verbosely!")

	err := config.Load(parsedOptions.Config, parsedOptions.EnvFile)
	if err == nil {
		err = config.Config.Validate()
	}

	if err != nil {
		log.WithFields(log.Fields{
			"configFile": parsedOptions.Config,
			"error":      err,
		}).Error("Could not load config file")
		os.Exit(1)
	}

	u, err := updater.New()
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("Could not start updater")
		os.Exit(1)
	}

	u.Start()
	u.Wg.Wait()
}
