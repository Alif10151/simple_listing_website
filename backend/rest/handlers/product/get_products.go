package product

import (
	"ecommerce/database"
	"ecommerce/utility"
	"net/http"
)

func (h *Handler) GetProducts(w http.ResponseWriter, r *http.Request) {
	utility.SendData(w, database.List(), 200)
}
