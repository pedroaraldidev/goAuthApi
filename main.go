package main

import (
    "github.com/gin-gonic/gin"
    "goAuthApi/routes"
)

func main() {
    router := gin.Default()
    routes.SetupAuthRoutes(router)
    router.Run(":8080")
}