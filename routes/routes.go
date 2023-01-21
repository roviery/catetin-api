package routes

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	deadlineController "github.com/roviery/catetin-api/app/deadlines/controller"
	deadlineRepo "github.com/roviery/catetin-api/app/deadlines/repository"
	deadlineUsecase "github.com/roviery/catetin-api/app/deadlines/usecase"
	financeController "github.com/roviery/catetin-api/app/finances/controller"
	financeRepo "github.com/roviery/catetin-api/app/finances/repository"
	financeUsecase "github.com/roviery/catetin-api/app/finances/usecase"
	quicknoteController "github.com/roviery/catetin-api/app/quicknotes/controller"
	quicknoteRepo "github.com/roviery/catetin-api/app/quicknotes/repository"
	quicknoteUsecase "github.com/roviery/catetin-api/app/quicknotes/usecase"
	todoController "github.com/roviery/catetin-api/app/todos/controller"
	todoRepo "github.com/roviery/catetin-api/app/todos/repository"
	todoUsecase "github.com/roviery/catetin-api/app/todos/usecase"
	transactionController "github.com/roviery/catetin-api/app/transactions/controller"
	transactionRepository "github.com/roviery/catetin-api/app/transactions/repository"
	transactionUsecase "github.com/roviery/catetin-api/app/transactions/usecase"
	userController "github.com/roviery/catetin-api/app/users/controller"
	userRepo "github.com/roviery/catetin-api/app/users/repository"
	userUsecase "github.com/roviery/catetin-api/app/users/usecase"
	"github.com/roviery/catetin-api/config"
	"github.com/roviery/catetin-api/constant"
	"github.com/roviery/catetin-api/entity"
	"github.com/spf13/viper"
)

func Init() *echo.Echo {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome to CatetinAPI")
	})

	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	mongoUri := viper.GetString("mongo.uri")
	c, err := config.ConnectMongo(mongoUri)
	if err != nil {
		panic(err)
	}

	userRepo := userRepo.NewUserRepo(c)
	userUsecase := userUsecase.NewUserUsecase(userRepo)
	userHandler := userController.NewUserHandler(userUsecase)

	deadlineRepo := deadlineRepo.NewDeadlineRepo(c)
	deadlineUsecase := deadlineUsecase.NewDeadlineUsecase(deadlineRepo)
	deadlineHandler := deadlineController.NewDeadlineHandler(deadlineUsecase)

	quicknoteRepo := quicknoteRepo.NewQuicknoteRepo(c)
	quicknoteUsecase := quicknoteUsecase.NewQuicknoteUsecase(quicknoteRepo)
	quicknoteHandler := quicknoteController.NewQuicknoteHandler(quicknoteUsecase)

	financeRepo := financeRepo.NewFinanceRepo(c)
	financeUsecase := financeUsecase.NewFinanceUsecase(financeRepo)
	financeHandler := financeController.NewFinanceHandler(financeUsecase)

	transactionRepo := transactionRepository.NewTransactionRepo(c)
	transactionUsecase := transactionUsecase.NewTransactionUsecase(transactionRepo, financeRepo)
	transactionHandler := transactionController.NewTransactionHandler(transactionUsecase)

	todoRepo := todoRepo.NewTodoRepository(c)
	todoUsecase := todoUsecase.NewTodoUsecase(todoRepo)
	todoHandler := todoController.NewTodoHandler(todoUsecase)

	r := e.Group("/api/v1")
	r.POST("/user/register", userHandler.Register)
	r.POST("/user/login", userHandler.Login)
	r.GET("/user/logout", userHandler.Logout)

	r.POST("/deadline", deadlineHandler.Create, middlewareJWTAuth)
	r.GET("/deadline/:id", deadlineHandler.GetDeadline, middlewareJWTAuth)
	r.GET("/deadlines", deadlineHandler.GetDeadlines, middlewareJWTAuth)
	r.PUT("/deadline/:id", deadlineHandler.Update, middlewareJWTAuth)
	r.DELETE("/deadline/:id", deadlineHandler.Delete, middlewareJWTAuth)

	r.POST("/quicknote", quicknoteHandler.Create, middlewareJWTAuth)
	r.GET("/quicknote/:id", quicknoteHandler.GetQuicknote, middlewareJWTAuth)
	r.GET("/quicknotes", quicknoteHandler.GetQuicknotes, middlewareJWTAuth)
	r.PUT("/quicknote/:id", quicknoteHandler.Update, middlewareJWTAuth)
	r.DELETE("/quicknote/:id", quicknoteHandler.Delete, middlewareJWTAuth)

	r.POST("/finance", financeHandler.Create, middlewareJWTAuth)
	r.GET("/finance/:id", financeHandler.GetFinance, middlewareJWTAuth)
	r.GET("/finances", financeHandler.GetFinances, middlewareJWTAuth)
	r.PUT("/finance/:id", financeHandler.Update, middlewareJWTAuth)
	r.DELETE("/finance/:id", financeHandler.Delete, middlewareJWTAuth)

	r.POST("/transaction", transactionHandler.Create, middlewareJWTAuth)
	r.GET("/transaction/:id", transactionHandler.GetTransaction, middlewareJWTAuth)
	r.GET("/transactions", transactionHandler.GetTransactions, middlewareJWTAuth)
	r.PUT("/transaction/:id", transactionHandler.Update, middlewareJWTAuth)
	r.DELETE("/transaction/:id", transactionHandler.Delete, middlewareJWTAuth)

	r.POST("/todo", todoHandler.Create, middlewareJWTAuth)
	r.GET("/todo/:id", todoHandler.GetTodo, middlewareJWTAuth)
	r.GET("/todos", todoHandler.GetTodos, middlewareJWTAuth)
	r.PUT("/todo/:id", todoHandler.Update, middlewareJWTAuth)
	r.DELETE("/todo/:id", todoHandler.Delete, middlewareJWTAuth)

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
