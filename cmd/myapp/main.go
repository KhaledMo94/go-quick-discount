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


	slog.Info("app started", "mysql", cnfg.DBHost, "redis", cnfg.RedisAddr())

	ctx , stop := signal.NotifyContext(context.Background(),os.Interrupt,syscall.SIGTERM)
	defer stop()

	<- ctx.Done()
	slog.Info("shutdown signal received, cleaning up")

	shutDownContext , cancel := context.WithTimeout(context.Background(),5 * time.Second)
	defer cancel()
	_ = shutDownContext
}