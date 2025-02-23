package handlers

import (
	"fmt"
	"form_basic/controllers"
	"net/http"
)

func SubscriptionHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "Error parsing form: %v", err)
			return
		}

		err := controllers.Create(r.PostForm.Get("name"), r.PostForm.Get("email"))
		if err != nil {
			fmt.Fprintf(w, "Error creating subscription: %v", err)
			return
		}
	}

	http.ServeFile(w, r, "handlers/templates/subscription_form.html")
}
