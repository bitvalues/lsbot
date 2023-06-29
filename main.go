package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"

	"github.com/bitvalues/lsbot/pkg/bot"
	"github.com/bitvalues/lsbot/pkg/config"
	"github.com/bitvalues/lsbot/pkg/plugins/auctions"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

const LogLevelError = 0
const LogLevelWarn = 1
const LogLevelInfo = 2
const LogLevelDebug = 3

var signalChannel chan os.Signal
var log *logrus.Logger

func init() {
	// create our channel for signal interrupts
	signalChannel = make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt, syscall.SIGINT)

	// setup the logging level & controller list (csv)
	logLevel := flag.Int("v", 0, "the level of verbosity for logging")
	flag.Parse()

	// create a new logger
	log = logrus.New()
	log.SetOutput(os.Stdout)
	// set the log level
	if *logLevel == LogLevelError {
		log.SetLevel(logrus.ErrorLevel)
	} else if *logLevel == LogLevelWarn {
		log.SetLevel(logrus.WarnLevel)
	} else if *logLevel == LogLevelInfo {
		log.SetLevel(logrus.InfoLevel)
	} else if *logLevel >= LogLevelDebug {
		log.SetLevel(logrus.DebugLevel)
	}

	// figure out where to load the .env file from
	envFileLocation := ".env"
	if location, exists := os.LookupEnv("ENV_FILE"); exists {
		envFileLocation = location
	}

	// attempt to load our .env file
	log.WithField("location", envFileLocation).Debug("Loading environment file from location")

	err := godotenv.Load(envFileLocation)
	if err != nil {
		log.WithError(err).Warn("Could not load environment file from location")
	}
}

func main() {
	// get the config first
	cfg := config.GetConfig()

	// setup the new discord bot
	bot, err := bot.NewBot(cfg, log)
	if err != nil {
		log.WithError(err).Fatal("could not create discord session")
		return
	}

	// register all of our plugins
	bot.LoadPlugin(auctions.NewAuctionsPlugin(bot))

	// make sure to shutdown the bot before we close out
	defer bot.Shutdown()

	// startup the bot
	go bot.Startup()

	// wait for an interrupt signal
	<-signalChannel
}
