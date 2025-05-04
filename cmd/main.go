package main

import (
	"log"
	"net/http"
	"text/template"
)

// Creating new page with url /home. Create connection with html template
func homePageHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	// Create simple info about page
	data := struct {
		PageName string
		Content  string
	}{
		PageName: "Home",
		Content:  "Some Information about ass",
	}

	// Create adresses to html files. Home files and layout
	tmpl, err := template.ParseFiles("../ui/templates/layout.html",
		"../ui/templates/homePage.html") // error system can't fine path home.html
	if err != nil {
		log.Printf("homePageHandler error with file parsing: %v", err)
		return
	}

	// Send data to page
	err = tmpl.Execute(w, data)
	if err != nil {
		log.Printf("homePageHandler error with file executing: %v", err)
	}

}

func showSnippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Отображение заметки..."))
}

func createSnippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Форма для создания новой заметки.."))
}

// Need new functions for page registrationPage(ui/templates/registrationPage)

// Need the same thing as with homePageHandler
func registrationPageHandler(w http.ResponseWriter, r *http.Request) {

}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", homePageHandler)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	// Adding css styles handler from file assets
	fileServer := http.FileServer(http.Dir("../ui/templates/assets"))
	mux.Handle("/assets/", http.StripPrefix("/assets", fileServer))

	log.Println("Запуск веб-сервера на http://127.0.0.1:4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}

// Прикольно было бы добавить автоматическое открытие сайта при запуске программы
