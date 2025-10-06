package product

import (
	"ecommerce/utility"
	"fmt"
	"net/http"
	"strconv"
)

func (h *Handler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	//utility.SendData(w, database.List(), 200)

	productID := r.PathValue("id")
	pId, err := strconv.Atoi(productID)

	if err != nil {
		fmt.Println(err)
		utility.SendError(w, http.StatusBadRequest, " Invalid pID")
		return
	}

	err = h.productRepo.Delete(pId)
	if err != nil {
		fmt.Println(err)
		utility.SendError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	utility.SendData(w, http.StatusOK, "deleted")

}
