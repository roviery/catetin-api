package controller

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/roviery/catetin-api/domain"
	"github.com/roviery/catetin-api/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PDFHandler struct {
	pdfUsecase domain.PDFUsecase
}

func NewPDFHandler(pdfUsecase domain.PDFUsecase) *PDFHandler {
	return &PDFHandler{
		pdfUsecase: pdfUsecase,
	}
}

func (p *PDFHandler) Upload(c echo.Context) error {
	file, err := c.FormFile("file")
	if err != nil {
		return err
	}

	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	data, err := ioutil.ReadAll(src)
	if err != nil {
		return err
	}

	pdf := &entity.PDFRequest{
		Name: file.Filename,
		Data: data,
	}

	res, err := p.pdfUsecase.UploadPDF(c.Request().Context(), pdf)
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

func (p *PDFHandler) Download(c echo.Context) error {
	fileId, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, entity.BaseResponse{
			Error: err.Error(),
		})
	}

	pdf, err := p.pdfUsecase.DownloadPDF(c.Request().Context(), &fileId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, entity.BaseResponse{
			Error: err.Error(),
		})
	}

	c.Response().Header().Set("Content-Type", "application/pdf")
	c.Response().Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", pdf.Name))

	if _, err := io.Copy(c.Response(), bytes.NewReader(pdf.Data)); err != nil {
		log.Println("Error: ", err)
		return c.JSON(http.StatusInternalServerError, "Error writing file to response")
	}

	return nil
}
