package sources

import (
	"fmt"
	"html/template"
	"net/http"
)

func UserList(w http.ResponseWriter, r *http.Request){


	// Generate the template
	tmpl, err := template.New("").ParseFiles("templates/base.html", "templates/userlist.html")
	if err != nil {
		fmt.Println(err.Error())
	}else {
		tmpl.ExecuteTemplate(w, "base", nil) 
	}
}