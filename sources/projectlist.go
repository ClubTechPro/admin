package sources

import (
	"fmt"
	"html/template"
	"net/http"
)

func ProjectList(w http.ResponseWriter, r *http.Request){


	// Generate the template
	tmpl, err := template.New("").ParseFiles("templates/base.html", "templates/projectlist.html")
	if err != nil {
		fmt.Println(err.Error())
	}else {
		tmpl.ExecuteTemplate(w, "base", nil) 
	}
}