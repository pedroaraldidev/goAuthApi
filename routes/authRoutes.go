package routes

import (
    "github.com/gin-gonic/gin"
    "github.com/joho/godotenv"
    "goAuthApi/controllers"
    "goAuthApi/services"
    "goAuthApi/repositories"
    "log"
    "os"
)

func SetupAuthRoutes(router *gin.Engine) {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Erro ao carregar o arquivo .env")
    }
    jwtSecret := os.Getenv("JWT_SECRET")
    if jwtSecret == "" {
        log.Fatal("A chave secreta JWT não foi configurada no .env")
    }

    userRepo := repositories.NewInMemoryUserRepository()
    userService := services.NewUserService(userRepo)

    authController := controllers.NewAuthController(userService, jwtSecret)

    router.POST("/login", authController.Login)
}