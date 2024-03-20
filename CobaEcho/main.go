package main

import (
	"CobaEcho/db"
	"CobaEcho/routes"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db.Init()
	e := routes.Init()

	e.Logger.Fatal(e.Start(":1234"))
}
	// db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/db_latihan_pbp?parseTime=true&loc=Asia%2FJakarta")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer db.Close()

	// controllers.InitializeDB(db)
	// router := mux.NewRouter()
	// router.HandleFunc("/login", controllers.Login).Methods("POST")
	// router.HandleFunc("/login", controllers.CheckUserLogin).Methods("POST")
	// router.HandleFunc("/logout", controllers.Logout).Methods("POST")
	// router.HandleFunc("/users", controllers.Authenticate(controllers.GetAllUsers, 1)).Methods("GET")
	// router.HandleFunc("/users", controllers.InsertUser).Methods("POST")
	// router.HandleFunc("/users/{id}", controllers.UpdateUser).Methods("PUT")

	// router.HandleFunc("/products", controllers.GetAllProducts).Methods("GET")
	// router.HandleFunc("/products", controllers.InsertNewProduct).Methods("POST")
	// router.HandleFunc("/products/{id}", controllers.UpdateProduct).Methods("PUT")
	// router.HandleFunc("/products/{id}", controllers.DeleteProduct).Methods("DELETE")

	// // Transactions
	// router.HandleFunc("/transactions", controllers.GetAllTransactions).Methods("GET")
	// router.HandleFunc("/transactions", controllers.InsertNewTransaction).Methods("POST")
	// router.HandleFunc("/transactions/{id}", controllers.UpdateTransaction).Methods("PUT")
	// router.HandleFunc("/transactions/{id}", controllers.DeleteTransaction).Methods("DELETE")

	// router.HandleFunc("/users/{user_id}", controllers.Authenticate(controllers.DeleteUser, 2)).Methods("DELETE")
	// router.HandleFunc("/users/{user_id}/transactions/{transaction_id}/...", controllers.Authenticate(controllers.DeleteUser, 1)).Methods("GET")

	// USER
	// router.HandleFunc("/users", controllers.GetAllUsers).Methods("GET")
	// router.HandleFunc("/users", controllers.InsertNewUser).Methods("POST")
	// router.HandleFunc("/users", controllers.UpdateUser).Methods("PUT")
	// router.HandleFunc("/users", controllers.DeleteUser).Methods("DELETE")
	// router.HandleFunc("/users", controllers.Login).Methods("POST")

	// // PRODUCT
	// router.HandleFunc("/products", controllers.GetAllProducts).Methods("GET")
	// router.HandleFunc("/products", controllers.InsertNewProduct).Methods("POST")
	// router.HandleFunc("/products", controllers.UpdateProduct).Methods("PUT")
	// router.HandleFunc("/products", controllers.DeleteProduct).Methods("DELETE")

	// // TRANSACTION
	// router.HandleFunc("/transactions", controllers.GetAllTransactions).Methods("GET")
	// router.HandleFunc("/transactions", controllers.InsertNewTransaction).Methods("POST")
	// router.HandleFunc("/transactions", controllers.UpdateTransaction).Methods("PUT")
	// router.HandleFunc("/transactions", controllers.DeleteTransaction).Methods("DELETE")
	// router.HandleFunc("/transactions", controllers.GetDetailUserTransaction).Methods("GET")

	// fmt.Println("Connected to port 8888")
	// log.Println("Connected to port 8888")
	// log.Fatal(http.ListenAndServe(":8888", router))

