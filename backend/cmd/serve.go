package cmd

import (
	"ecommerce/config"
	"ecommerce/infrastructure/db"
	"ecommerce/repo"
	"ecommerce/rest"
	prdctHandler "ecommerce/rest/handlers/product"
	usrHandler "ecommerce/rest/handlers/user"
	middleware "ecommerce/rest/middlewares"
	"ecommerce/user"
	"fmt"
	"os"
)

func Serve() {

	cnf := config.GetConfig()
	fmt.Printf("%+v", cnf.DB)

	dbCon, err := db.NewConnection(cnf.DB)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	err = db.MigrateDB(dbCon, "./migrations")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	productRepo := repo.NewProductRepo(dbCon)
	userRepo := repo.NewUserRepo(dbCon)

	usrSvc := user.NewService(userRepo)

	middlewares := middleware.NewMiddleswares(cnf)

	productHandler := prdctHandler.NewHandler(middlewares, productRepo)
	userHandler := usrHandler.NewHandler(cnf, usrSvc)

	server := rest.NewServer(
		cnf,
		productHandler,
		userHandler,
	)
	server.Start()
}
