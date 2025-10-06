package product

import (
	"ecommerce/utility"
	"net/http"
)

func (h *Handler) GetProducts(w http.ResponseWriter, r *http.Request) {
	productList, err := h.productRepo.List()
	if err != nil {
		utility.SendError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}
	utility.SendData(w, http.StatusOK, productList)
}
