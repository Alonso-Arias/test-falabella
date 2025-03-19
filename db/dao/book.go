package dao

import (
	"context"

	"github.com/Alonso-Arias/test-falabella/db/base"
	"github.com/Alonso-Arias/test-falabella/db/model"
	"github.com/Alonso-Arias/test-falabella/log"
	"gorm.io/gorm"
)

var loggerf = log.LoggerJSON().WithField("package", "dao")

// BookDAO - Book dao interface
type BookDAO interface {
	FindAll(ctx context.Context) ([]model.Book, error)
	Save(ctx context.Context, Book model.Book) error
	Get(ctx context.Context, bookID int) (model.Book, error)
}

// BookDAOImpl - Book dao implementation
type BookDAOImpl struct {
}

// NewBookDAO - gets an BookDAOImpl instance
func NewBookDAO() *BookDAOImpl {
	return &BookDAOImpl{}
}

// FindAll -
func (pd *BookDAOImpl) FindAll(ctx context.Context) ([]model.Book, error) {

	log := loggerf.WithField("struct", "BookDAOImpl").WithField("function", "FindAll")

	db := base.GetDB()

	Books := []model.Book{}
	err := db.Find(&Books).Error

	if err != nil {
		log.WithError(err).Error("get Books fails")
		return []model.Book{}, err
	} else if Books == nil {
		return []model.Book{}, gorm.ErrRecordNotFound
	}

	log.Debugf("%v", Books)

	return Books, nil

}

// Get -
func (pd *BookDAOImpl) Get(ctx context.Context, id int32) (model.Book, error) {

	log := loggerf.WithField("struct", "BookDAOImpl").WithField("function", "Get")

	db := base.GetDB()

	Book := model.Book{}
	err := db.Where("ID = ?", id).FirstOrInit(&Book).Error

	if err != nil {
		log.WithError(err).Error("get Books fails")
		return model.Book{}, err
	} else if Book.ID == 0 {
		return model.Book{}, gorm.ErrRecordNotFound
	}

	log.Debugf("%v", Book)

	return Book, nil

}

// Save -
func (pd *BookDAOImpl) Save(ctx context.Context, Book model.Book) error {

	log := loggerf.WithField("struct", "BookDAOImpl").WithField("function", "Save")

	db := base.GetDB()

	err := db.Create(&Book)

	if err.Error != nil {
		log.Debugf("%v", err.Error)
		return err.Error
	}

	log.Infof("Save Book Sucessfull\n")

	return nil

}
