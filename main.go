
package main

import (
	"net/http"

	"techpro.club.admin/sources"
	"techpro.club.admin/sources/common"
)

func main(){

	fs := http.FileServer(http.Dir("templates/assets"))
	http.Handle("/assets/", http.StripPrefix("/assets", fs))

	// Login
	http.HandleFunc("/", sources.Login)


	// Projects
	http.HandleFunc("/projectlist", sources.ProjectList)

	// Users
	http.HandleFunc("/userlist", sources.UserList)


	// Start the web server
    http.ListenAndServe(common.CONST_APP_PORT, nil)
}
