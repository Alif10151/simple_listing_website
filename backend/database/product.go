package database

var productList []Product

type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImageURL    string  `json:"image_url"`
}

func Store(p Product) Product {
	p.ID = len(productList) + 1
	productList = append(productList, p)
	return p
}

func List() []Product {
	return productList
}

func GET(productID int) *Product {
	for _, product := range productList {
		if product.ID == productID {
			return &product
		}
	}

	return nil
}

func Update(product Product) {
	for idx, p := range productList {
		if p.ID == product.ID {
			productList[idx] = product
		}
	}
}

func Delete(productID int) {

	var tempList []Product

	for _, p := range productList {
		if p.ID != productID {
			tempList = append(tempList, p)
		}
	}

	productList = tempList
}

func init() {
	prd1 := Product{
		ID:          1,
		Title:       "Orange",
		Description: "i love it",
		Price:       20,
		ImageURL:    "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcT-6GoLpPXoOkT0lAuFcxXwJSQ7nxtRQqVJLg&s",
	}

	prd2 := Product{
		ID:          2,
		Title:       "Jackfruit",
		Description: "i love it",
		Price:       50,
		ImageURL:    "https://static1.squarespace.com/static/578753d7d482e9c3a909de40/578756a859cc6802a0ce9dfd/608824420809513ad892ca0e/1693742657517/ripe-jackfruit.jpeg?format=1500w",
	}

	prd3 := Product{
		ID:          3,
		Title:       "mango",
		Description: "i love it",
		Price:       5,
		ImageURL:    "https://img.freepik.com/free-photo/mango-still-life_23-2151542114.jpg?w=360",
	}

	prd4 := Product{
		ID:          4,
		Title:       "banana",
		Description: "i  don't love it",
		Price:       15,
		ImageURL:    "https://www.lovefoodhatewaste.com/sites/default/files/styles/16_9_two_column/public/2022-07/Bananas.jpg.webp?itok=bqo03oZ1",
	}

	prd5 := Product{
		ID:          5,
		Title:       "apple",
		Description: "well",
		Price:       20,
		ImageURL:    "https://static.vecteezy.com/system/resources/previews/012/086/172/non_2x/green-apple-with-green-leaf-isolated-on-white-background-vector.jpg",
	}

	productList = append(productList, prd1)
	productList = append(productList, prd2)
	productList = append(productList, prd3)
	productList = append(productList, prd4)
	productList = append(productList, prd5)

}
