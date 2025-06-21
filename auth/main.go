package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/globallstudent/go-auth/docs"
	"github.com/globallstudent/go-auth/internal"
	ginSwaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Gopher Notes Auth API
// @version         1.0
// @description     Advanced authentication and authorization API with JWT and role-based access control.
// @host      localhost:8080
// @BasePath  /

func SwaggerHomeHandler(c *gin.Context) {
	// @Summary      Swagger Home Redirect
	// @Description  Redirects to Swagger UI
	// @Tags         docs
	// @Produce      plain
	// @Success      307 {string} string "Redirect"
	c.Redirect(http.StatusTemporaryRedirect, "/swagger/index.html")
}

// @Summary      Swagger UI
// @Description  Swagger documentation endpoint
// @Tags         docs
// @Produce      json
// @Success      200 {object} string
func SwaggerUIHandler(c *gin.Context) {
	ginSwagger.WrapHandler(ginSwaggerFiles.Handler)(c)
}

// @Summary      Login
// @Description  Login with username and password
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        credentials body struct{Username string `json:"username"`; Password string `json:"password"`} true "Login credentials"
// @Success      200 {object} map[string]string
// @Failure      400 {object} map[string]string
// @Failure      401 {object} map[string]string
func LoginHandler(c *gin.Context) {
	var req struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	user, ok := internal.GetUserByEmail(req.Username)
	if !ok || req.Password != "password" { // Demo: password is always 'password'
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}
	token, err := internal.GenerateJWT(user.Email, user.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "JWT generation failed"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}

// @Summary      Get Profile
// @Description  Get user profile (JWT required)
// @Tags         user
// @Produce      json
// @Success      200 {object} map[string]string
// @Failure      401 {object} map[string]string
func ProfileHandler(c *gin.Context) {
	email := c.GetString("email")
	role := c.GetString("role")
	c.JSON(http.StatusOK, gin.H{"email": email, "role": role})
}

// @Summary      Admin Only
// @Description  Admin endpoint (JWT + admin role required)
// @Tags         admin
// @Produce      json
// @Success      200 {object} map[string]string
// @Failure      401 {object} map[string]string
// @Failure      403 {object} map[string]string
func AdminHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Welcome, admin!"})
}

// @Summary      User Only
// @Description  User endpoint (JWT + user role required)
// @Tags         user
// @Produce      json
// @Success      200 {object} map[string]string
// @Failure      401 {object} map[string]string
// @Failure      403 {object} map[string]string
func UserHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Welcome, user!"})
}

func main() {
	r := gin.Default()

	r.GET("/", SwaggerHomeHandler)
	r.GET("/swagger/*any", SwaggerUIHandler)
	r.POST("/auth/login", LoginHandler)
	r.GET("/profile", internal.JWTAuthMiddleware(), ProfileHandler)
	r.GET("/admin", internal.JWTAuthMiddleware(), internal.RoleAuthMiddleware("admin"), AdminHandler)
	r.GET("/user", internal.JWTAuthMiddleware(), internal.RoleAuthMiddleware("user"), UserHandler)

	r.Run(":8080")
}
