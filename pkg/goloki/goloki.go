package goloki

import (
	"time"

	"github.com/sirupsen/logrus"
	"github.com/yukitsune/lokirus"
)

func Generate(interval string) {

	// Configure the Loki hook
	opts := lokirus.NewLokiHookOptions().
		// Grafana doesn't have a "panic" level, but it does have a "critical" level
		// https://grafana.com/docs/grafana/latest/explore/logs-integration/
		WithLevelMap(lokirus.LevelMap{logrus.PanicLevel: "critical"}).
		WithFormatter(&logrus.JSONFormatter{}).
		WithStaticLabels(lokirus.Labels{
			"app":         "goloki",
			"environment": "development",
		}).
		WithBasicAuth("admin", "secretpassword") // Optional

	hook := lokirus.NewLokiHookWithOpts(
		"http://localhost:3100",
		opts,
		logrus.InfoLevel,
		logrus.WarnLevel,
		logrus.ErrorLevel,
		logrus.FatalLevel)

	// Configure the logger
	logger := logrus.New()
	logger.AddHook(hook)

	// Log all the things!
	duration, _ := time.ParseDuration(interval + "s")

	for {
		time.Sleep(duration)
		logger.WithField("fizz", "buzz").Warnln("warning")
	}
}
