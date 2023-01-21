package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/roviery/catetin-api/domain"
	"github.com/roviery/catetin-api/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TodoHandler struct {
	todoUsecase domain.TodoUsecase
}

func NewTodoHandler(todoUsecase domain.TodoUsecase) *TodoHandler {
	return &TodoHandler{
		todoUsecase: todoUsecase,
	}
}

func (to *TodoHandler) Create(c echo.Context) error {
	var req entity.TodoRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, entity.BaseResponse{
			Error: err.Error(),
		})
	}

	if len(req.Todo) == 0 || len(req.Status) == 0 {
		return c.JSON(http.StatusBadRequest, entity.BaseResponse{
			Error: "empty request data exist",
		})
	}

	if req.Status != "todo" && req.Status != "in progress" && req.Status != "done" {
		return c.JSON(http.StatusBadRequest, entity.BaseResponse{
			Error: "status request invalid",
		})
	}

	ownerId, err := primitive.ObjectIDFromHex(c.Get("user_id").(string))
	if err != nil {
		return c.JSON(http.StatusBadRequest, entity.BaseResponse{
			Error: err.Error(),
		})
	}

	res, err := to.todoUsecase.CreateTodo(c.Request().Context(), &ownerId, &req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, entity.BaseResponse{
			Error: err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, entity.BaseResponse{
		Message: "success",
		Data:    res,
	})
}

func (to *TodoHandler) GetTodo(c echo.Context) error {
	todoId, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, entity.BaseResponse{
			Error: err.Error(),
		})
	}

	res, err := to.todoUsecase.GetTodo(c.Request().Context(), &todoId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, entity.BaseResponse{
			Error: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, entity.BaseResponse{
		Message: "success",
		Data:    res,
	})
}

func (to *TodoHandler) GetTodos(c echo.Context) error {
	ownerId, err := primitive.ObjectIDFromHex(c.Get("user_id").(string))
	if err != nil {
		return c.JSON(http.StatusBadRequest, entity.BaseResponse{
			Error: err.Error(),
		})
	}

	res, err := to.todoUsecase.GetTodos(c.Request().Context(), &ownerId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, entity.BaseResponse{
			Error: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, entity.BaseResponse{
		Message: "success",
		Data:    res,
	})
}

func (to *TodoHandler) Update(c echo.Context) error {
	var req entity.TodoRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, entity.BaseResponse{
			Error: err.Error(),
		})
	}

	if len(req.Todo) == 0 || len(req.Status) == 0 {
		return c.JSON(http.StatusBadRequest, entity.BaseResponse{
			Error: "empty request data exist",
		})
	}

	if req.Status != "todo" && req.Status != "in progress" && req.Status != "done" {
		return c.JSON(http.StatusBadRequest, entity.BaseResponse{
			Error: "status request invalid",
		})
	}

	todoId, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, entity.BaseResponse{
			Error: err.Error(),
		})
	}

	res, err := to.todoUsecase.UpdateTodoByID(c.Request().Context(), &todoId, &req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, entity.BaseResponse{
			Error: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, entity.BaseResponse{
		Message: "success",
		Data:    res,
	})
}

func (to *TodoHandler) Delete(c echo.Context) error {
	todoId, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, entity.BaseResponse{
			Error: err.Error(),
		})
	}

	err = to.todoUsecase.DeleteTodoByID(c.Request().Context(), &todoId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, entity.BaseResponse{
			Error: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, entity.BaseResponse{
		Message: "success",
	})
}
