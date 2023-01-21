package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/roviery/catetin-api/domain"
	"github.com/roviery/catetin-api/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type QuicknoteHandler struct {
	quicknoteUseCase domain.QuicknoteUsecase
}

func NewQuicknoteHandler(quicknoteUsecase domain.QuicknoteUsecase) *QuicknoteHandler {
	return &QuicknoteHandler{
		quicknoteUseCase: quicknoteUsecase,
	}
}

func (q *QuicknoteHandler) Create(c echo.Context) error {
	var req entity.QuicknoteRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, entity.BaseResponse{
			Error: err.Error(),
		})
	}

	if len(req.Note) == 0 {
		return c.JSON(http.StatusBadRequest, entity.BaseResponse{
			Error: "empty request data exist",
		})
	}

	ownerId, err := primitive.ObjectIDFromHex(c.Get("user_id").(string))
	if err != nil {
		return c.JSON(http.StatusBadRequest, entity.BaseResponse{
			Error: err.Error(),
		})
	}

	res, err := q.quicknoteUseCase.CreateQuicknote(c.Request().Context(), &ownerId, &req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, entity.BaseResponse{
			Error: err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, entity.BaseResponse{
		Message: "success",
		Data:    res,
	})
}

func (q *QuicknoteHandler) GetQuicknote(c echo.Context) error {
	quicknoteId, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, entity.BaseResponse{
			Error: err.Error(),
		})
	}

	res, err := q.quicknoteUseCase.GetQuicknote(c.Request().Context(), &quicknoteId)
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

func (q *QuicknoteHandler) GetQuicknotes(c echo.Context) error {
	ownerId, err := primitive.ObjectIDFromHex(c.Get("user_id").(string))
	if err != nil {
		return c.JSON(http.StatusBadRequest, entity.BaseResponse{
			Error: err.Error(),
		})
	}

	res, err := q.quicknoteUseCase.GetQuicknotes(c.Request().Context(), &ownerId)
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

func (q *QuicknoteHandler) Update(c echo.Context) error {
	var req entity.QuicknoteRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, entity.BaseResponse{
			Error: err.Error(),
		})
	}

	if len(req.Note) == 0 {
		return c.JSON(http.StatusBadRequest, entity.BaseResponse{
			Error: "empty request data exist",
		})
	}

	quicknoteId, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, entity.BaseResponse{
			Error: err.Error(),
		})
	}

	res, err := q.quicknoteUseCase.UpdateQuicknoteByID(c.Request().Context(), &quicknoteId, &req)
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

func (q *QuicknoteHandler) Delete(c echo.Context) error {
	quicknoteId, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, entity.BaseResponse{
			Error: err.Error(),
		})
	}

	err = q.quicknoteUseCase.DeleteQuicknoteByID(c.Request().Context(), &quicknoteId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, entity.BaseResponse{
			Error: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, entity.BaseResponse{
		Message: "success",
	})
}
