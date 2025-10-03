package product

import (
	"ecommerce/database"
	"ecommerce/utility"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
)

func (h *Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {

	var newProduct database.Product
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "give me valid json", 400)
		return
	}

	createdProduct := database.Store(newProduct)

	utility.SendData(w, createdProduct, 201)
}

func base64UrlEncode(data []byte) string {
	return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(data)
}
