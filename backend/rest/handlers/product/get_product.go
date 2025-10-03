package product

import (
	"ecommerce/database"
	"ecommerce/utility"
	"net/http"
	"strconv"
)

func (h *Handler) GetProduct(w http.ResponseWriter, r *http.Request) {
	productID := r.PathValue("id")
	pId, err := strconv.Atoi(productID)
	if err != nil {
		http.Error(w, "invalid id", 400)
		return
	}

	product := database.GET(pId)
	if product == nil {
		utility.SendError(w, 404, "Product not found")
		return
	}

	// for _, product := range database.ProductList {
	// 	if product.ID == pId {
	// 		utility.SendData(w, product, 200)
	// 		return
	// 	}
	// }

	utility.SendData(w, product, 200)

}
