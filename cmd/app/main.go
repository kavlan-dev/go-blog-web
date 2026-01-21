package main

import (
	"context"
	"go-blog-web/internal/config"
	"go-blog-web/internal/handlers"
	"go-blog-web/internal/middleware"
	"go-blog-web/internal/routers"
	"go-blog-web/internal/services"
	"go-blog-web/internal/storage/memory"
	"go-blog-web/internal/utils"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalln("Ошибка загрузки настроек", err)
		return
	}
	log := utils.New(cfg.Env)

	storage := memory.New()
	service := services.New(storage)
	handler := handlers.New(service, log)

	if err := service.CreateFirstAdmin(cfg); err != nil {
		log.Error("Не удалось создать администратора", utils.Err(err))
		return
	}

	server := &http.Server{
		Addr:    cfg.ServerAddress(),
		Handler: middleware.CORSMiddleware(cfg, routers.SetupRoutes(handler, service)),
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		log.Info("Запуск сервера", slog.String("address", cfg.ServerAddress()))
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Error("Ошибка запуска сервера", utils.Err(err))
		}
	}()

	<-sigChan
	log.Info("Получен сигнал завершения, начинаем плавное завершение...")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Error("Ошибка при плавном завершении сервера", utils.Err(err))
		return
	}

	log.Info("Сервер успешно остановлен")
}
