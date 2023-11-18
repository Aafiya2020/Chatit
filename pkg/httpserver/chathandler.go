package httpserver

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"gochatapp/pkg/redisrepo"
)
// userReq represents the structure of the JSON request for user-related operations.
type userReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Client   string `json:"client"`
}
// response represents the structure of the JSON response for API calls.
type response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Total   int         `json:"total,omitempty"`
}

// registerHandler handles user registration requests.
func registerHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Decode JSON request into userReq struct
	u := &userReq{}
	if err := json.NewDecoder(r.Body).Decode(u); err != nil {
		http.Error(w, "error decoidng request object", http.StatusBadRequest)
		return
	}

	// Process registration and send JSON response
	res := register(u)
	json.NewEncoder(w).Encode(res)
}

// loginHandler handles user login requests.
func loginHandler(w http.ResponseWriter, r *http.Request) {
	// Set response header
	w.Header().Set("Content-Type", "application/json")

	// Decode JSON request into userReq struct
	u := &userReq{}
	if err := json.NewDecoder(r.Body).Decode(u); err != nil {
		http.Error(w, "error decoding request object", http.StatusBadRequest)
		return
	}

	// Process login and send JSON response
	res := login(u)
	json.NewEncoder(w).Encode(res)
}

// verifyContactHandler handles contact verification requests.
func verifyContactHandler(w http.ResponseWriter, r *http.Request) {
	// Set response header
	w.Header().Set("Content-Type", "application/json")

	// Decode JSON request into userReq struct
	u := &userReq{}
	if err := json.NewDecoder(r.Body).Decode(u); err != nil {
		http.Error(w, "error decoding request object", http.StatusBadRequest)
		return
	}

	// Process contact verification and send JSON response
	res := verifyContact(u.Username)
	json.NewEncoder(w).Encode(res)
}

// chatHistoryHandler handles requests for retrieving chat history between two users.
func chatHistoryHandler(w http.ResponseWriter, r *http.Request) {
	// Set response header
	w.Header().Set("Content-Type", "application/json")

	// Extract query parameters
	u1 := r.URL.Query().Get("u1")
	u2 := r.URL.Query().Get("u2")
	fromTS, toTS := "0", "+inf"

	// Extract additional parameters for chat history
	if r.URL.Query().Get("from-ts") != "" && r.URL.Query().Get("to-ts") != "" {
		fromTS = r.URL.Query().Get("from-ts")
		toTS = r.URL.Query().Get("to-ts")
	}

	// Process chat history request and send JSON response
	res := chatHistory(u1, u2, fromTS, toTS)
	json.NewEncoder(w).Encode(res)
}

// contactListHandler handles requests for retrieving the contact list of a user.
func contactListHandler(w http.ResponseWriter, r *http.Request) {
	// Set response header
	w.Header().Set("Content-Type", "application/json")

	// Extract username from query parameters
	u := r.URL.Query().Get("username")

	// Process contact list request and send JSON response
	res := contactList(u)
	json.NewEncoder(w).Encode(res)
}

// register processes user registration.
func register(u *userReq) *response {
	// Create response structure
	res := &response{Status: true}

	// Check if the username already exists
	status := redisrepo.IsUserExist(u.Username)
	if status {
		res.Status = false
		res.Message = "username already taken. try something else."
		return res
	}

	// Register a new user and handle errors
	err := redisrepo.RegisterNewUser(u.Username, u.Password)
	if err != nil {
		res.Status = false
		res.Message = "something went wrong while registering the user. please try again after sometime."
		return res
	}

	return res
}

// login processes user login.
func login(u *userReq) *response {
	// Create response structure
	res := &response{Status: true}

	// Check if the provided username and password are valid
	err := redisrepo.IsUserAuthentic(u.Username, u.Password)
	if err != nil {
		res.Status = false
		res.Message = err.Error()
		return res
	}

	return res
}

// verifyContact processes contact verification.
func verifyContact(username string) *response {
	// Create response structure
	res := &response{Status: true}

	// Check if the provided username is valid
	status := redisrepo.IsUserExist(username)
	if !status {
		res.Status = false
		res.Message = "invalid username"
	}

	return res
}

// chatHistory processes requests for retrieving chat history between two users.
func chatHistory(username1, username2, fromTS, toTS string) *response {
	// Create response structure
	res := &response{}

	// Check if both usernames are valid
	if !redisrepo.IsUserExist(username1) || !redisrepo.IsUserExist(username2) {
		res.Message = "incorrect username"
		return res
	}

	// Fetch chat history and handle errors
	chats, err := redisrepo.FetchChatBetween(username1, username2, fromTS, toTS)
	if err != nil {
		log.Println("error in fetch chat between", err)
		res.Message = "unable to fetch chat history. please try again later."
		return res
	}

	res.Status = true
	res.Data = chats
	res.Total = len(chats)
	return res
}

// contactList processes requests for retrieving the contact list of a user.
func contactList(username string) *response {
	// if invalid username return error
	// if valid users fetch chats
	res := &response{}

	// check if user exists
	if !redisrepo.IsUserExist(username) {
		res.Message = "incorrect username"
		return res
	}

	contactList, err := redisrepo.FetchContactList(username)
	if err != nil {
		log.Println("error in fetch contact list of username: ", username, err)
		res.Message = "unable to fetch contact list. please try again later."
		return res
	}

	res.Status = true
	res.Data = contactList
	res.Total = len(contactList)
	return res
}
