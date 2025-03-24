package controllers

import (
    "net/http"
    "time"
    "goAuthApi/services"
    "github.com/gin-gonic/gin"
    "github.com/golang-jwt/jwt/v5"
)

type AuthController struct {
    userService *services.UserService
    jwtSecret   string
}


func NewAuthController(userService *services.UserService, jwtSecret string) *AuthController {
    return &AuthController{userService: userService}
}

func (c *AuthController) Login(ctx *gin.Context) {
    var input struct {
        Username string `json:"username"`
        Password string `json:"password"`
    }
    if err := ctx.ShouldBindJSON(&input); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    user, err := c.userService.Authenticate(input.Username, input.Password)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    if user == nil {
        ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
        return
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "sub": user.ID,
        "username": user.Username,
        "exp": time.Now().Add(time.Hour * 24).Unix(),
    })

    tokenString, err := token.SignedString([]byte(c.jwtSecret))
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{
        "message": "Login successful",
        "token":   tokenString,
    })
}