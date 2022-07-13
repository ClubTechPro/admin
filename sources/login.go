package sources

import (
	"fmt"
	"html/template"
	"net/http"
)

// Function to handle login of admin
func Login(w http.ResponseWriter, r *http.Request) {
	
	// Generate the template
	tmpl, err := template.ParseFiles("templates/login.html")
	if err != nil {
		fmt.Println(err.Error())
	}else {
		tmpl.Execute(w, nil) 
	}
}