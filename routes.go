package main

import "goAuthApi/routes"

func SetupRoutes(router *gin.Engine) {
    routes.SetupAuthRoutes(router)
}