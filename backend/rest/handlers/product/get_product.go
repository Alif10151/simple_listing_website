package product

import (
	"ecommerce/utility"
	"net/http"
	"strconv"
)

func (h *Handler) GetProduct(w http.ResponseWriter, r *http.Request) {
	productID := r.PathValue("id")
	pId, err := strconv.Atoi(productID)
	if err != nil {
		utility.SendError(w, http.StatusBadRequest, "Invalid req body")
		return
	}

	product, err := h.productRepo.Get(pId)

	if err != nil {
		utility.SendError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	if product == nil {
		utility.SendError(w, http.StatusNotFound, "Product not found")
		return
	}

	// for _, product := range database.ProductList {
	// 	if product.ID == pId {
	// 		utility.SendData(w, product, 200)
	// 		return
	// 	}
	// }

	utility.SendData(w, http.StatusOK, product)

}
