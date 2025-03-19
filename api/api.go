package main

import (
	"context"
	"net/http"
	"strconv"

	errs "github.com/Alonso-Arias/test-falabella/errors"
	"github.com/Alonso-Arias/test-falabella/log"
	"github.com/Alonso-Arias/test-falabella/services/book"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

var loggerf = log.LoggerJSON().WithField("package", "main")

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:1323
// @BasePath /api/v1
func main() {
	e := echo.New()
	e.GET("/api/v1/books/findAll", findAllBooksGet)
	e.POST("/api/v1/book", bookPost)
	e.GET("/api/v1/book/:bookID", bookGet)
	e.GET("/api/v1/book/:bookID/boxPrice", bookBoxPriceGet)
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.Logger.Fatal(e.Start(":1323"))

}

// find all Books
// @Summary Find all Books
// @tags Book
// @Description obtiene todos los book
// @ID findAllBooksGet
// @Accept  json
// @Produce  json
// @Success 200  {object} book.FindAllBooksResponse
// @Failure 404 {object}  errors.CustomError
// @Failure 500 {object}  errors.CustomError
// @Router /book/findAll [get]
func findAllBooksGet(c echo.Context) error {

	res, err := book.BookService{}.FindAllBooks(context.TODO())
	if ce, ok := err.(errs.CustomError); ok {
		return c.JSON(ce.Code, err)
	} else if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, res)
}

// get book
// @Summary get book by id
// @tags book
// @Description obtiene book por id
// @ID bookGet
// @Accept  json
// @Produce  json
// @Param id path string true "Id"
// @Success 200  {object} book.GetBookResponse
// @Failure 404 {object}  errors.CustomError
// @Failure 500 {object}  errors.CustomError
// @Router /book/{id} [get]
func bookGet(c echo.Context) error {

	idStr := c.Param("bookID")
	idInt, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	req := book.GetBookRequest{
		BookId: int32(idInt),
	}

	res, err := book.BookService{}.GetBook(context.TODO(), req)
	if ce, ok := err.(errs.CustomError); ok {
		return c.JSON(ce.Code, err)
	} else if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, res)
}

// save book
// @Summary save book
// @tags book
// @Description guarda un book
// @ID bookPost
// @Accept  json
// @Produce  json
// @Param SaveBookRequest body book.SaveBookRequest true "book"
// @Success 200  {object} book.SaveBookResponse
// @Failure 404 {object}  errors.CustomError
// @Failure 500 {object}  errors.CustomError
// @Router /book [post]
func bookPost(c echo.Context) error {

	log := loggerf.WithField("func", "bookPost")

	req := book.SaveBookRequest{}

	if err := c.Bind(&req); err != nil {
		log.WithError(err).Error("Binding error")
		return c.JSON(http.StatusBadRequest, err)
	}

	res, err := book.BookService{}.SaveBook(context.TODO(), req)
	if ce, ok := err.(errs.CustomError); ok {
		return c.JSON(ce.Code, err)
	} else if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, res)
}

// get book
// @Summary get book by id
// @tags book
// @Description obtiene book por id
// @ID BookGet
// @Accept  json
// @Produce  json
// @Param id path string true "Id"
// @Success 200  {object} book.GetBookResponse
// @Failure 404 {object}  errors.CustomError
// @Failure 500 {object}  errors.CustomError
// @Router /book/{id} [get]
func bookBoxPriceGet(c echo.Context) error {

	idStr := c.Param("bookID")
	idInt, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	quantityIn := c.QueryParam("quantity")
	if quantityIn == "" {
		return c.JSON(http.StatusBadRequest, errs.BadRequest.SetMessage("Quantity is required"))
	}
	quantity, err := strconv.Atoi(quantityIn)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	req := book.GetBookBoxPriceRequest{
		BookId:   int32(idInt),
		Currency: c.QueryParam("currency"),
		Quantity: quantity,
	}

	res, err := book.BookService{}.GetBookBoxPrice(context.TODO(), req)
	if ce, ok := err.(errs.CustomError); ok {
		return c.JSON(ce.Code, err)
	} else if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, res)
}
