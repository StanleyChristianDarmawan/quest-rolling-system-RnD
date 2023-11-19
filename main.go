package main

import (
	"log"
	"net/http"
	"quest-rolling-system-RnD-backup/config"
	"quest-rolling-system-RnD-backup/controller/cartcontroller"
	"quest-rolling-system-RnD-backup/controller/categorycontroller"
	"quest-rolling-system-RnD-backup/controller/homecontroller"
	"quest-rolling-system-RnD-backup/controller/productcontroller"
)

func main() {
	config.ConnectDB()

	// home page
	http.HandleFunc("/", homecontroller.Welcome)

	// Categories
	http.HandleFunc("/categories", categorycontroller.Index)
	http.HandleFunc("/categories/add", categorycontroller.Add)
	http.HandleFunc("/categories/edit", categorycontroller.Edit)
	http.HandleFunc("/categories/delete", categorycontroller.Delete)

	// Products
	http.HandleFunc("/products", productcontroller.Index)
	http.HandleFunc("/products/add", productcontroller.Add)
	http.HandleFunc("/products/detail", productcontroller.Detail)
	http.HandleFunc("/products/edit", productcontroller.Edit)
	http.HandleFunc("/products/delete", productcontroller.Delete)

	// Cart
	http.HandleFunc("/cart", cartcontroller.Index)

	log.Println("Server running...")
	http.ListenAndServe(":8000", nil)
}
