package product

import "time"

type ProductReview struct {
	Id int `json:"id"`
	Content string `json:"content"`
	ProductID uint `json:"productId"`
}

type Product struct {
	ID uint `gorm:"primary_key" json:"id"`
	Name string `json:"name"`
	ShortDescription string `json:"shortDescription"`
	FullDescription string `json:"fullDescription"`
	ImageSrc string `json:"imageSrc"`
	Reviews []ProductReview `json:"productReviews"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	DeletedAt *time.Time `json:"-"`
}
