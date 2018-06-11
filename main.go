package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/sessions"
	"golang.org/x/crypto/bcrypt"
)

type GPSdata struct {
	Latitude  string
	Longitude string
}

var s GPSdata
var encryptionKey = "something-very-secret"
var loggedUserSession = sessions.NewCookieStore([]byte(encryptionKey))

func init() {

	loggedUserSession.Options = &sessions.Options{
		// change domain to match your machine. Can be localhost
		// IF the Domain name doesn't match, your session will be EMPTY!
		Domain:   "localhost",
		Path:     "/",
		MaxAge:   3600 * 3, // 3 hours
		HttpOnly: true,
	}
}

var dashboardTemplate = template.Must(template.ParseFiles("dashBoardPage.gtpl"))
var logUserTemplate = template.Must(template.ParseFiles("logUserPage.gtpl"))

func DashBoardPageHandler(w http.ResponseWriter, r *http.Request) {
	conditionsMap := map[string]interface{}{}
	//read from session
	session, err := loggedUserSession.Get(r, "authenticated-user-session")

	if err != nil {
		log.Println("Unable to retrieve session data!", err)
	}

	log.Println("Session name : ", session.Name())

	log.Println("Username : ", session.Values["username"])

	conditionsMap["Username"] = session.Values["username"]

	if err := dashboardTemplate.Execute(w, conditionsMap); err != nil {
		log.Println(err)
	}
	result, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal([]byte(result), &s)
	fmt.Fprintf(w, "Latitude: %s\n", s.Latitude)
	fmt.Fprintf(w, "Longitude: %s\n", s.Longitude)
}

func LoginPageHandler(w http.ResponseWriter, r *http.Request) {

	conditionsMap := map[string]interface{}{}

	// check if session is active
	session, _ := loggedUserSession.Get(r, "authenticated-user-session")

	if session != nil {
		conditionsMap["Username"] = session.Values["username"]
	}

	// verify username and password
	if r.FormValue("Login") != "" && r.FormValue("Username") != "" {
		username := r.FormValue("Username")
		password := r.FormValue("Password")

		// NOTE: here is where you want to query your database to retrieve the hashed password
		// for username.
		// For this tutorial and simplicity sake, we will simulate the retrieved hashed password
		// as $2a$10$4Yhs5bfGgp4vz7j6ScujKuhpRTA4l4OWg7oSukRbyRN7dc.C1pamu
		// the plain password is 'mynakedpassword'
		// see https://www.socketloop.com/tutorials/golang-bcrypting-password for more details
		// on how to generate bcrypted password

		hashedPasswordFromDatabase := []byte("$2a$10$4Yhs5bfGgp4vz7j6ScujKuhpRTA4l4OWg7oSukRbyRN7dc.C1pamu")

		if err := bcrypt.CompareHashAndPassword(hashedPasswordFromDatabase, []byte(password)); err != nil {
			log.Println("Either username or password is wrong")
			conditionsMap["LoginError"] = true
		} else {
			log.Println("Logged in :", username)
			conditionsMap["Username"] = username
			conditionsMap["LoginError"] = false

			// create a new session and redirect to dashboard
			session, _ := loggedUserSession.New(r, "authenticated-user-session")

			session.Values["username"] = username
			err := session.Save(r, w)

			if err != nil {
				log.Println(err)
			}

			http.Redirect(w, r, "/dashboard", http.StatusFound)
		}

	}

	if err := logUserTemplate.Execute(w, conditionsMap); err != nil {
		log.Println(err)
	}
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	//read from session
	session, _ := loggedUserSession.Get(r, "authenticated-user-session")

	// remove the username
	session.Values["username"] = ""
	err := session.Save(r, w)

	if err != nil {
		log.Println(err)
	}

	w.Write([]byte("Logged out!"))
}

func main() {
	fmt.Println("Server starting, point your browser to localhost:8080/login to start")
	http.HandleFunc("/login", LoginPageHandler)
	http.HandleFunc("/dashboard", DashBoardPageHandler)
	http.HandleFunc("/logout", LogoutHandler)
	http.ListenAndServe(":8080", nil)
}
