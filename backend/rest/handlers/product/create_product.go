package product

import (
	"ecommerce/repo"
	"ecommerce/utility"
	"encoding/json"
	"fmt"
	"net/http"
)

type ReqCreateProduct struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImageURL    string  `json:"image_url"`
}

func (h *Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {

	var req ReqCreateProduct
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		fmt.Println(err)
		utility.SendError(w, http.StatusBadRequest, " Invalid req body")
		return
	}

	// createdProduct := database.Store(newProduct)

	createdProduct, err := h.productRepo.Create(repo.Product{
		Title:       req.Title,
		Description: req.Description,
		Price:       req.Price,
		ImageURL:    req.ImageURL,
	})
	if err != nil {
		utility.SendError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	utility.SendData(w, http.StatusCreated, createdProduct)
}

// func base64UrlEncode(data []byte) string {
// 	return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(data)
// }
