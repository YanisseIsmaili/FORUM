package Forum

import (
	"html/template"
	"net/http"
)

var templates = template.Must(template.ParseFiles("./Frontend/index.html"))

func HomeHandler(w http.ResponseWriter, r *http.Request) {

	temp := ""
	err := templates.ExecuteTemplate(w, "index.html", temp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// func postHome(w http.ResponseWriter, r *http.Request) {
// 	r.ParseForm()                         // parse le formulaire
// 	EmailUser := r.FormValue("EmailUser") // prend la valeur du formulaire
// 	fmt.Println(EmailUser)
// 	PasswordUser := r.FormValue("passwordUser") // prend la valeur du formulaire
// 	fmt.Println(EmailUser)
// 	fmt.Println(PasswordUser)
// }

//dalda
