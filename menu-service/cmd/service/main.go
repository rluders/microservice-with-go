package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/jmoiron/sqlx"

	"github.com/gorilla/mux"
	"golang.org/x/sync/errgroup"
	"menu-service/internal/config"
	database "menu-service/internal/database/postgres"
	"menu-service/internal/domain"
	repository "menu-service/internal/repository/postgres"
	"menu-service/internal/transport/rest"
)

func main() {
	log.Println("service starting")

	var configPath string
	flag.StringVar(&configPath, "config", "", "Path to config file")
	flag.Parse()

	// Initialize the basics
	cfg := loadConfig(configPath)
	db := loadDatabase(cfg.Database)

	// Initialize the repositories
	itemRepository := repository.NewItemRepository(db)
	categoryRepository := repository.NewCategoryRepository(db)

	// Initialize the services
	itemService := domain.NewItemService(itemRepository, categoryRepository)
	categoryService := domain.NewCategoryService(categoryRepository)

	// Gracefully shutdown
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
	defer signal.Stop(interrupt)

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	g, ctx := errgroup.WithContext(ctx)

	// REST Server
	var restServer *rest.Server
	g.Go(func() (err error) {
		router := mux.NewRouter().StrictSlash(true)

		// Register the handlers
		rest.NewItemHandler(itemService).Register(router)
		rest.NewCategoryHandler(categoryService).Register(router)

		// Create the server
		restServer, err = rest.NewServer(cfg.Server.HTTP, router)
		if err != nil {
			return err
		}

		return restServer.Start()
	})

	log.Println("service started")

	select {
	case <-interrupt:
		break
	case <-ctx.Done():
		break
	}

	log.Println("shutdown signal received")

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdownCancel()

	if restServer != nil {
		restServer.Stop(shutdownCtx)
	}

	err := g.Wait()
	if err != nil {
		log.Printf("server shutdown returned an error: %v", err)
		os.Exit(2)
	}

	log.Println("service shutdown")
}

func loadConfig(configPath string) *config.Config {
	if configPath == "" {
		configPath = os.Getenv("APP_CONFIG_PATH")
		if configPath == "" {
			configPath = "./config.yaml"
		}
	}

	cfg, err := config.NewConfig(configPath)
	if err != nil {
		log.Printf("configuration error: %v", err)
		os.Exit(-1)
	}

	return cfg
}

func loadDatabase(cfg *config.Database) *sqlx.DB {
	db, err := database.Connect(cfg)
	if err != nil {
		log.Printf("database error: %v", err)
		os.Exit(-1)
	}

	return db
}
