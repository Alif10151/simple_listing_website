package product

import (
	"ecommerce/database"
	"ecommerce/utility"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func (h *Handler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	//utility.SendData(w, database.List(), 200)

	productID := r.PathValue("id")
	pId, err := strconv.Atoi(productID)
	if err != nil {
		http.Error(w, "invalid id", 400)
		return
	}

	var newProduct database.Product

	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&newProduct)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "give me valid json", 400)
		return
	}

	newProduct.ID = pId
	database.Update(newProduct)
	utility.SendData(w, "updated", 201)

}
