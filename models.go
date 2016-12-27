package main

type ProductReview struct {
	Id int `json:"id"`
	Content string `json:"content"`
}

type Product struct {
	Id int `json:"id"`
	Name string `json:"name"`
	ShortDescription string `json:"shortDescription"`
	FullDescription string `json:"fullDescription"`
	ImageSrc string `json:"imageSrc"`
	Reviews []ProductReview `json:"productReviews"`
}
