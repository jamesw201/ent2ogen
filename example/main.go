package main

import (
	"context"
	"flag"
	"log"
	"os"
  "net/http"

	"entgo.io/ent/dialect"
	_ "github.com/jackc/pgx/v4/stdlib"

	entsql "entgo.io/ent/dialect/sql"

	"github.com/go-faster/errors"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"

  "github.com/XSAM/otelsql"
  "github.com/joho/godotenv"

  "github.com/jamesw201/go-starter/example/driver"
	"github.com/jamesw201/go-starter/example/ent"
	"github.com/jamesw201/go-starter/example/api"
	"github.com/jamesw201/go-starter/internal/app"
	handler "github.com/jamesw201/go-starter/example/handler"
)

func main() {
  handler.Hello()

  err := godotenv.Load(".env")
  log.Println(os.Getenv("COCKROACH_DSN"))

  // db, err := sql.Open("pgx", os.Getenv("COCKROACH_DSN"))
  db, err := otelsql.Open("pgx", os.Getenv("COCKROACH_DSN"))
	if err != nil {
		log.Fatal(err)
	}
	drv := entsql.OpenDB(dialect.Postgres, db)
	cd := driver.New(drv)
	client := ent.NewClient(ent.Driver(cd))
	defer client.Close()
	ctx := context.Background()
	// run the auto migration tool.
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	app.Run(func(ctx context.Context, lg *zap.Logger) error {
		var arg struct {
			Addr        string
			MetricsAddr string
		}
		flag.StringVar(&arg.Addr, "addr", "127.0.0.1:8082", "listen address")
		flag.StringVar(&arg.MetricsAddr, "metrics.addr", "127.0.0.1:9090", "metrics listen address")
		flag.Parse()

		lg.Info("Initializing",
			zap.String("http.addr", arg.Addr),
			zap.String("metrics.addr", arg.MetricsAddr),
		)

		m, err := app.NewMetrics(lg, app.Config{
			Addr: arg.MetricsAddr,
			Name: "api",
		})
		if err != nil {
			return errors.Wrap(err, "metrics")
		}

		oasServer, err := api.NewServer(
      handler.NewOgentHandler(client), 
			api.WithTracerProvider(m.TracerProvider()),
			api.WithMeterProvider(m.MeterProvider()),
		)

		if err != nil {
			return errors.Wrap(err, "server init")
		}
		httpServer := http.Server{
			Addr:    arg.Addr,
			Handler: oasServer,
		}

		g, ctx := errgroup.WithContext(ctx)
		g.Go(func() error {
			return m.Run(ctx)
		})
		g.Go(func() error {
			<-ctx.Done()
			return httpServer.Shutdown(ctx)
		})
		g.Go(func() error {
			defer lg.Info("Server stopped")
			if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				return errors.Wrap(err, "http")
			}
			return nil
		})

		return g.Wait()
	})
}
