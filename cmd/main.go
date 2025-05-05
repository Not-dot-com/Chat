package main

import (
	services "Chatmod/internal/services"
	"log"
	"net/http"
)


func main() {
	// Регистрируем два новых обработчика и соответствующие URL-шаблоны в
	// маршрутизаторе servemux
	mux := http.NewServeMux()
	mux.HandleFunc("/", services.HomePageHandler)

	// Adding css styles handler from file assets
	fileServer := http.FileServer(http.Dir("../ui/templates/assets"))
	mux.Handle("/assets/", http.StripPrefix("/assets", fileServer))

	log.Println("Запуск веб-сервера на http://127.0.0.1:4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}

// Прикольно было бы добавить автоматическое открытие сайта при запуске программы
