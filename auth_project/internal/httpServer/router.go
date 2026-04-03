package httpserver

import (
	"auth_project/internal/app"
	"time"

	// "auth_project/internal/auth"
	"auth_project/internal/middleware"
	"auth_project/internal/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter(a *app.App) *gin.Engine {
	// r := gin.Default()
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/health", healthCheckHandler)

	userRepo := user.NewRepo(a.DB)
	userService := user.NewService(userRepo, a.Config.JWTSecret)
	userHandler := user.NewHandler(userService)

	r.POST("/u/register", userHandler.Register)
	r.POST("/u/login", userHandler.Login)

	// grouping protected routes
	api := r.Group("/api")

	authMiddleware := middleware.AuthMiddleware(a.Config.JWTSecret)
	api.Use(authMiddleware)

	api.GET("/files", func(ctx *gin.Context) {
		userId, _ := middleware.GetUserID(ctx)
		role, _ := middleware.GetUserRole(ctx)

		ctx.JSON(http.StatusOK, gin.H{
			"files":     []any{},
			"userID":    userId,
			"role":      role,
			"success":   true,
			"timestamp": time.Now().UTC(),
		})
	})

	api.GET("/products", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"products":  []any{},
			"success":   true,
			"timestamp": ctx.MustGet("timestamp"),
			"userID":    ctx.MustGet("userID"),
			"role":      ctx.MustGet("role"),
		})
	})

	admin := api.Group("/admin")
	admin.Use(middleware.RolesMiddleware("admin"))
	admin.GET("/dashboard", func(ctx *gin.Context) {
		role, _ := middleware.GetUserRole(ctx)

		ctx.JSON(http.StatusOK, gin.H{
			"dashboard": []any{},
			"success":   true,
			"role":      role,
			"timestamp": time.Now().UTC(),
		})
	})

	return r
}
