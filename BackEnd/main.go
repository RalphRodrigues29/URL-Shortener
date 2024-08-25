package main

import (
	"math/rand"
	"net/http"
	"encoding/json"

	"github.com/glebarez/sqlite"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/middleware"
	"gorm.io/gorm"
	"github.com/go-chi/cors"
)

type URL struct {
	ID       uint   `gorm:"primaryKey"`
	LongURL  string `json:"longUrl"`
	ShortURL string `json:"shortUrl"`
}

var db *gorm.DB

func main() {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Use(cors.New(cors.Options{
        AllowedOrigins: []string{"http://localhost:3000", "https://localhost:8080"},
        AllowedMethods: []string{"GET", "POST", "PUT", "OPTIONS"},
        AllowedHeaders: []string{
            "Accept", "Authorization", "Content-Type", "X-CSRF-Token", "X-Requested-With",
            "AuthorizationExpire", "AuthorizationSignature", "ImageId", "X-File-MIME-Type", "X-Login-Type",
        },
        AllowCredentials: true,
        MaxAge:           300,
        Debug:            true,
    }).Handler)
	
	var err error
	db, err = gorm.Open(sqlite.Open("urls.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}
	db.AutoMigrate(&URL{})

	router.Post("/shorten", shortenURL)
	router.Get("/{shortUrl}", redirectURL)
	router.Options("/shorten", shortenURL)

	err = http.ListenAndServe(":8080", router)
	if err != nil {
		panic(err)
	}
}

func shortenURL(w http.ResponseWriter, r *http.Request) {
	var request URL

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	request.ShortURL = generateShortURL()
	db.Create(&request)
	response := "http://localhost:8080/" + request.ShortURL
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(response))
}

func redirectURL(w http.ResponseWriter, r *http.Request) {
	shortUrl := chi.URLParam(r, "shortUrl")
	var url URL
	if result := db.Where("short_url = ?", shortUrl).First(&url); result.Error != nil {
		http.Error(w, "URL not found", http.StatusNotFound)
		return
	}
	http.Redirect(w, r, url.LongURL, http.StatusMovedPermanently)
}

func generateShortURL() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	shortURL := make([]byte, 6)
	for i := range shortURL {
		shortURL[i] = charset[rand.Intn(len(charset))]
	}
	return string(shortURL)
}