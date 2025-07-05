package main

import (
	"flag"
	"log"
	"time"

	"github.com/riandyhasan/imperio/config"
	"github.com/riandyhasan/imperio/db"
	"github.com/riandyhasan/imperio/model"
	"github.com/riandyhasan/imperio/operation"
	"github.com/riandyhasan/imperio/runner"
)

func main() {
	// Parse CLI flags
	configPath := flag.String("config", "config.yaml", "Path to configuration file")
	flag.StringVar(configPath, "c", "config.yaml", "Alias for --config")
	flag.Parse()

	// 1. Load configuration
	cfg, err := config.Load(*configPath)
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// 2. Initialize database strategy
	database, err := db.NewDatabase(cfg.DBConfig, db.DBMS(cfg.Database))
	if err != nil {
		log.Fatalf("Unsupported database: %v", err)
	}
	defer database.Close()

	// 3. Load schema
	schemaLoader := model.NewFileSchemaLoader()
	schema, err := schemaLoader.LoadSchema(cfg.SchemaFile)
	if err != nil {
		log.Fatalf("Failed to load schema: %v", err)
	}

	// 4. Create operation executor
	exec, err := operation.NewExecutor(database, schema)
	if err != nil {
		log.Fatalf("Failed to initialize executor: %v", err)
	}

	// 5. Determine run duration (0 or < 0 means run forever)
	duration := time.Duration(0)
	if cfg.RunnerDuration > 0 {
		duration = cfg.RunnerDuration
	}

	// 6. Start runner
	simulator := runner.NewRunner(cfg, exec, duration)
	if err := simulator.Start(); err != nil {
		log.Fatalf("Execution failed: %v", err)
	}
}
