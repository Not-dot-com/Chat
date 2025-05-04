package main

import (
	"log"
	"net/http"
	"text/template"
)

// New function home. Need to replace it to top.
// Handler for home page
func homePageHandler(w http.ResponseWriter, r *http.Request) {
	// Проверяется, если текущий путь URL запроса точно совпадает с шаблоном "/". Если нет, вызывается
	// функция http.NotFound() для возвращения клиенту ошибки 404.
	// Важно, чтобы мы завершили работу обработчика через return. Если мы забудем про "return", то обработчик
	// продолжит работу и выведет сообщение "Привет из SnippetBox" как ни в чем не бывало.
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

// Обработчик для отображения содержимого заметки.
func showSnippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Отображение заметки..."))
}

// Обработчик для создания новой заметки.
func createSnippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Форма для создания новой замтки.."))
}

func main() {
	// Регистрируем два новых обработчика и соответствующие URL-шаблоны в
	// маршрутизаторе servemux
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
