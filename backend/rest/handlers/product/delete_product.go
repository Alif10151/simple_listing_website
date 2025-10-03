package product

import (
	"ecommerce/database"
	"ecommerce/utility"
	"net/http"
	"strconv"
)

func (h *Handler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	//utility.SendData(w, database.List(), 200)

	productID := r.PathValue("id")
	pId, err := strconv.Atoi(productID)
	if err != nil {
		http.Error(w, "invalid id", 400)
		return
	}
	database.Delete(pId)
	utility.SendData(w, "deleted", 201)

}
