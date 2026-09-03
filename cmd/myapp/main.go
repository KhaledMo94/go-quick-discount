package main

import (
	"context"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/KhaledMo94/quick-discount/internal/config"
	"github.com/KhaledMo94/quick-discount/internal/db"
	"github.com/KhaledMo94/quick-discount/internal/redis"
	"github.com/KhaledMo94/quick-discount/internal/logger"
	"github.com/KhaledMo94/quick-discount/internal/server"
	"github.com/KhaledMo94/quick-discount/internal/handler"
	"github.com/KhaledMo94/quick-discount/internal/meliesearch"
	
)

func main(){

	logger.Setup()

	cnfg , err := config.Load()
	if err != nil {
		log.Fatalf("config loading failed : %v",err)
	}

	dbConn , err := db.New(cnfg.DSN(),cnfg.Driver)
	if err != nil {
		log.Fatalf("mysql connection failed: %v", err)
	}
	defer dbConn.Close()

	redisConn , err := redis.New(cnfg.RedisAddr(),cnfg.RedisPassword , cnfg.RedisDB)
	if err != nil {
		log.Fatalf("Redis Connection failed : %v",err)
	}
	defer redisConn.Close()

	msConn , err := meliesearch.New(cnfg.MEILISEARCH_HOST , cnfg.MEILISEARCH_KEY)
	if err != nil {
		log.Fatalf("Meilesearch Connection failed : %v",err)
	}
	defer msConn.Close() //it`s http request closed after data retrieving

	slog.Info("app started", "mysql", cnfg.DBHost, "redis", cnfg.RedisAddr() , "ms health",msConn.IsHealthy())

	h := handler.New(dbConn , redisConn , msConn)
	srv := server.New(":8000",h.Routes())
	srv.Start()

	ctx , stop := signal.NotifyContext(context.Background(),os.Interrupt,syscall.SIGTERM)
	defer stop()

	<- ctx.Done()
	slog.Info("shutdown signal received, cleaning up")

	shutDownContext , cancel := context.WithTimeout(context.Background(),5 * time.Second)
	defer cancel()
	_ = shutDownContext
}