package book

import (
	"context"

	"github.com/Alonso-Arias/test-falabella/db/dao"
	md "github.com/Alonso-Arias/test-falabella/db/model"
	errs "github.com/Alonso-Arias/test-falabella/errors"
	"github.com/Alonso-Arias/test-falabella/log"
	"github.com/Alonso-Arias/test-falabella/services/model"
	"gopkg.in/dealancer/validate.v2"
	"gorm.io/gorm"
)

var loggerf = log.LoggerJSON().WithField("package", "services")

// BookService contiene los métodos relacionados con las tareas.
type BookService struct{}

// FindAllBooksResponse es la respuesta para FindAllBooks.
type FindAllBooksResponse struct {
	Books []model.Book `json:"Books"`
}

// FindAllBooks recupera todas las tareas.
func (bk BookService) FindAllBooks(ctx context.Context) (FindAllBooksResponse, error) {
	log := loggerf.WithField("service", "BookService").WithField("func", "FindAllBooks")

	BookDAO := dao.NewBookDAO()

	Books, err := BookDAO.FindAll(ctx)
	if err != nil && err != gorm.ErrRecordNotFound {
		log.WithError(err).Error("problems with getting Books")
		return FindAllBooksResponse{}, err
	} else if err == gorm.ErrRecordNotFound {
		return FindAllBooksResponse{}, errs.BooksNotFound
	}

	results := []model.Book{}

	for _, v := range Books {
		Book := model.Book{
			ID:        v.ID,
			Title:     v.Title,
			Author:    v.Author,
			Publisher: v.Publisher,
			Country:   v.Country,
			Price:     v.Price,
			Currency:  v.Currency,
		}
		results = append(results, Book)
	}

	return FindAllBooksResponse{Books: results}, nil
}

// GetBookRequest es la solicitud para GetBook.
type GetBookRequest struct {
	BookId int32 `json:"bookID"`
}

// GetBookResponse es la respuesta para GetBook.
type GetBookResponse struct {
	Book model.Book `json:"Book"`
}

// GetBook obtiene una tarea por su ID.
func (bk BookService) GetBook(ctx context.Context, in GetBookRequest) (GetBookResponse, error) {
	log := loggerf.WithField("service", "BookService").WithField("func", "GetBook")

	if in.BookId == 0 {
		return GetBookResponse{}, errs.BadRequest
	}

	BookDAO := dao.NewBookDAO()

	v, err := BookDAO.Get(ctx, in.BookId)
	if err != nil && err != gorm.ErrRecordNotFound {
		log.WithError(err).Error("problems with getting Book")
		return GetBookResponse{}, err
	} else if err == gorm.ErrRecordNotFound {
		return GetBookResponse{}, errs.BooksNotFound
	}

	Book := model.Book{
		ID:        v.ID,
		Title:     v.Title,
		Author:    v.Author,
		Publisher: v.Publisher,
		Country:   v.Country,
		Price:     v.Price,
		Currency:  v.Currency,
	}

	return GetBookResponse{Book: Book}, nil
}

// SaveBookRequest es la solicitud para SaveBook.
type SaveBookRequest struct {
	Book model.Book `json:"Book"`
}

// SaveBookResponse es la respuesta para SaveBook.
type SaveBookResponse struct{}

// SaveBook guarda una nueva tarea.
func (bk BookService) SaveBook(ctx context.Context, in SaveBookRequest) (SaveBookResponse, error) {
	log := loggerf.WithField("service", "BookService").WithField("func", "SaveBook")

	// Valida la solicitud de entrada
	if err := validate.Validate(in); err != nil {
		log.WithError(err).Error("validation problems")
		return SaveBookResponse{}, errs.BadRequest
	}

	BookDAO := dao.NewBookDAO()

	err := BookDAO.Save(ctx, md.Book(md.Book{
		ID:        in.Book.ID,
		Title:     in.Book.Title,
		Author:    in.Book.Author,
		Publisher: in.Book.Publisher,
		Country:   in.Book.Country,
		Price:     in.Book.Price,
		Currency:  in.Book.Currency,
	}))
	if err != nil {
		return SaveBookResponse{}, err
	}

	return SaveBookResponse{}, nil
}

// GetBookBoxPriceRequest es la solicitud para GetBook.
type GetBookBoxPriceRequest struct {
	BookId   int32  `json:"bookID"`
	Currency string `json:"currency"`
	Quantity int    `json:"quantity"`
}

// GetBookBoxPriceResponse es la respuesta para GetBook.
type GetBookBoxPriceResponse struct {
	TotalPrice float64 `json:"totalPrice"`
}

// GetBook obtiene una tarea por su ID.
func (bk BookService) GetBookBoxPrice(ctx context.Context, in GetBookBoxPriceRequest) (GetBookBoxPriceResponse, error) {
	log := loggerf.WithField("service", "BookService").WithField("func", "GetBook")

	if in.BookId == 0 {
		return GetBookBoxPriceResponse{}, errs.BadRequest
	}

	BookDAO := dao.NewBookDAO()

	v, err := BookDAO.Get(ctx, in.BookId)
	if err != nil && err != gorm.ErrRecordNotFound {
		log.WithError(err).Error("problems with getting Book")
		return GetBookBoxPriceResponse{}, err
	} else if err == gorm.ErrRecordNotFound {
		return GetBookBoxPriceResponse{}, errs.BooksNotFound
	}

	totalPrice := v.Price * float64(in.Quantity)

	return GetBookBoxPriceResponse{TotalPrice: totalPrice}, nil
}
