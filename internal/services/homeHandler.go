package services

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// Handler for home page
func HomePageHandler(w http.ResponseWriter, r *http.Request) {
	// Проверяется, если текущий путь URL запроса точно совпадает с шаблоном "/". Если нет, вызывается
	// функция http.NotFound() для возвращения клиенту ошибки 404.
	// Важно, чтобы мы завершили работу обработчика через return. Если мы забудем про "return", то обработчик
	// продолжит работу и выведет сообщение "Привет из SnippetBox" как ни в чем не бывало.
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		log.Printf("HomeHandler error with finding path")
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
		log.Printf("homeHandler error with file parsing: %v", err)
		fmt.Printf("homeHandler error with file parsing: %v", err)
		return
	}

	// Send data to page
	err = tmpl.Execute(w, data)
	if err != nil {
		log.Printf("homeHandler error with file executing: %v", err)
		return
	}

}
