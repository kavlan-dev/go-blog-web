package routers

import (
	"go-blog-web/internal/handlers"
	"go-blog-web/internal/middleware"
	"go-blog-web/internal/services"
	"net/http"
)

func SetupRoutes(handler *handlers.Handler, service *services.Service) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", handler.HealthCheck)
	mux.HandleFunc("GET /api/posts", handler.Posts)
	mux.HandleFunc("GET /api/posts/{id}", handler.PostById)
	mux.HandleFunc("GET /api/posts/title/{title}", handler.PostByTitle)
	mux.HandleFunc("POST /api/auth/register", handler.CreateUser)

	mux.HandleFunc("POST /api/posts", middleware.AuthMiddleware(service, handler.CreatePost))

	mux.HandleFunc("PUT /api/posts/{id}", middleware.AuthAdminMiddleware(service, handler.UpdatePost))
	mux.HandleFunc("DELETE /api/posts/{id}", middleware.AuthAdminMiddleware(service, handler.DeletePost))
	mux.HandleFunc("PUT /api/users/{id}", middleware.AuthAdminMiddleware(service, handler.UpdateUser))

	return mux
}
