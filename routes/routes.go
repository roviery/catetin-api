package routes

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	userController "github.com/roviery/catetin-api/app/users/controller"
	userRepo "github.com/roviery/catetin-api/app/users/repository"
	userUsecase "github.com/roviery/catetin-api/app/users/usecase"
	"github.com/roviery/catetin-api/constant"
	"github.com/roviery/catetin-api/db"
	"github.com/roviery/catetin-api/entity"
)

func Init() *echo.Echo {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome to CatetinAPI")
	})

	sqlDB := db.DB()

	userRepo := userRepo.NewUserRepo(sqlDB)
	userUsecase := userUsecase.NewUserUsecase(userRepo)
	userHandler := userController.NewUserHandler(userUsecase)

	r := e.Group("/api/v1")
	r.POST("/user/register", userHandler.Register)
	r.POST("/user/login", userHandler.Login)
	r.GET("/user/logout", userHandler.Logout)

	return e
}

func middlewareJWTAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if !strings.Contains(authHeader, "Bearer") {
			return c.JSON(http.StatusBadRequest, entity.BaseResponse{
				Error: "token invalid",
			})
		}

		tokenString := strings.Split(authHeader, " ")[1]
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("signing method invalid")
			} else if method != constant.JWTSigningMethod {
				return nil, fmt.Errorf("signing method invalid")
			}
			return constant.JWTSecretKey, nil
		})
		if err != nil {
			return c.JSON(http.StatusBadRequest, entity.BaseResponse{
				Error: err.Error(),
			})
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			return c.JSON(http.StatusBadRequest, entity.BaseResponse{
				Error: "failed to parse token claims",
			})
		}

		c.Set("user_id", claims["id"])
		return next(c)
	}
}
