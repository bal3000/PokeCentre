package main

import (
	"flag"
	"strings"
	"sync"

	"go.uber.org/zap"
)

type config struct {
	port int
	env  string
	db   struct {
		dsn          string
		maxOpenConns int
		maxIdleConns int
		maxIdleTime  string
	}
	limiter struct {
		rps     float64
		burst   int
		enabled bool
	}
	cors struct {
		trustedOrigins []string
	}
}

type application struct {
	config config
	logger *zap.Logger
	wg     sync.WaitGroup
}

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 8080, "port to listen on")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")

	flag.StringVar(&cfg.db.dsn, "db-dsn", "", "Database DSN")
	flag.IntVar(&cfg.db.maxOpenConns, "db-max-open-conns", 25, "Maximum number of open connections to the database")
	flag.IntVar(&cfg.db.maxIdleConns, "db-max-idle-conns", 25, "Maximum number of idle connections to the database")
	flag.StringVar(&cfg.db.maxIdleTime, "db-max-idle-time", "15m", "Maximum amount of time a connection may be reused")

	flag.Float64Var(&cfg.limiter.rps, "limiter-rps", 5, "Maximum number of requests per second")
	flag.IntVar(&cfg.limiter.burst, "limiter-burst", 10, "Maximum number of requests allowed in a burst")
	flag.BoolVar(&cfg.limiter.enabled, "limiter-enabled", true, "Enable request rate limiter")

	flag.Func("cors-trusted-origins", "Trusted CORS origins (comma separated)", func(s string) error {
		cfg.cors.trustedOrigins = strings.Split(s, ",")
		return nil
	})

	flag.Parse()

	if err := run(cfg); err != nil {
		panic(err)
	}
}

func run(cfg config) error {
	logger, err := getLogger(cfg.env)
	if err != nil {
		return err
	}
	defer logger.Sync()

	return nil
}

func getLogger(env string) (*zap.Logger, error) {
	switch {
	case env == "development":
		return zap.NewDevelopment()
	case env == "staging":
		return zap.NewProduction()
	case env == "production":
		return zap.NewProduction()
	default:
		return zap.NewDevelopment()
	}
}
