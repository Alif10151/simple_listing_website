package product

import (
	"ecommerce/repo"
	"ecommerce/utility"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type ReqUpdateProduct struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImageURL    string  `json:"image_url"`
}

func (h *Handler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	//utility.SendData(w, database.List(), 200)

	productID := r.PathValue("id")
	pId, err := strconv.Atoi(productID)
	if err != nil {
		utility.SendError(w, http.StatusBadRequest, "invalid id")
		return
	}

	var req ReqUpdateProduct

	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&req)
	if err != nil {
		fmt.Println(err)
		utility.SendError(w, http.StatusBadRequest, "invalid re method")
		return
	}

	_, err = h.productRepo.Update(repo.Product{
		ID:          pId,
		Title:       req.Title,
		Description: req.Description,
		Price:       req.Price,
		ImageURL:    req.ImageURL,
	})
	if err != nil {
		utility.SendError(w, http.StatusInternalServerError, "Internal server error")
		return
	}
	utility.SendData(w, http.StatusOK, "updated")

}
