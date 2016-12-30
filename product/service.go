package product

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"net/http"
)

type Service struct {
	Port   string
	Config ServiceConfig
	Db     *gorm.DB
}

func NewProductService(env string) (*Service, error) {
	s := &Service{}
	config, err := NewProductServiceConfig(fmt.Sprintf("./config/%s.yml", env))
	s.Config = config

	if err != nil {
		return s, err
	}

	s.Db, err = DbOpen(s.Config.Db)

	if err != nil {
		return s, err
	}

	return s, nil
}

func (s *Service) Run() (err error) {
	resource := &Resource{s.Db}

	// Register Handlers
	http.HandleFunc("/api/products", resource.FindAll)
	http.HandleFunc("/api/products/", resource.Find)

	return http.ListenAndServe(":8080", nil)
}

func (s *Service) MigrateDb() error {
	// Migrate the schema
	s.Db.AutoMigrate(&Product{})
	s.Db.AutoMigrate(&ProductReview{})

	return nil
}

func (s *Service) LoadFixtures() error {
	dummyProduct := Product{
		Name:             "Product 1",
		ShortDescription: "Short Description",
		FullDescription:  "Full Product 1 Description",
		ImageSrc:         "http://img.com/test.png",
	}
	s.Db.Create(&dummyProduct)

	dummyProductReview := ProductReview{
		Content:   "Dummy Review",
		ProductID: dummyProduct.ID,
	}
	s.Db.Create(&dummyProductReview)

	return nil
}

func DbOpen(dbStr string) (*gorm.DB, error) {
	db, err := gorm.Open("mysql", dbStr)

	if err != nil {
		return db, err
	}

	db.SingularTable(true)

	return db, nil
}
