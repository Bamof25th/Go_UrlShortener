package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type URL struct {
	Id          string    `json:"id"`
	OriginalUrl string    `json:"original_url"`
	ShortUrl    string    `json:"short_url"`
	CreatedAt   time.Time `json:"createdAt"`
}

var urlDb = make(map[string]URL)

func generateShortUrl(OriginalUrl string) string {
	hasher := md5.New()

	hasher.Write([]byte(OriginalUrl))
	fmt.Println("hasher generated", hasher)
	data := hasher.Sum(nil)
	fmt.Println("data generated", data)
	hash := hex.EncodeToString(data)
	fmt.Println("hash generated", hash[:8])
	return hash[:8] // only 8 chars to show so [:8]
}

func createUrl(originalUrl string) string {
	shortUrl := generateShortUrl(originalUrl)
	id := shortUrl
	urlDb[id] = URL{
		Id:          id,
		OriginalUrl: originalUrl,
		ShortUrl:    shortUrl,
		CreatedAt:   time.Now(),
	}
	return shortUrl
}

func getUrl(shortUrl string) (URL, error) {
	url, ok := urlDb[shortUrl]
	if !ok {
		return URL{}, fmt.Errorf("url not found")
	}
	return url, nil
}

func Homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world")
}

func ShortUrlHandler(w http.ResponseWriter, r *http.Request) {
	var data struct {
		URL string `json:"url"`
	}

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error()+"Invalid request Body", http.StatusBadRequest)
		return
	}
	shortUrl := createUrl(data.URL)

	// fmt.Fprintf(w, shortUrl)

	response := struct {
		ShortUrl string `json:"shortUrl"`
	}{ShortUrl: shortUrl}

	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func RedirectUrlHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/redirect/"):]
	url, err := getUrl(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	http.Redirect(w, r, url.OriginalUrl, http.StatusFound)
}

func main() {
	fmt.Println("making url shortener")
	// var OriginalUrl = "https://github.com/Bamof25th"

	// generateShortUrl(OriginalUrl)

	// handel function  rootUrl('/')
	http.HandleFunc("/", Homepage)
	http.HandleFunc("/getUrl", ShortUrlHandler)
	http.HandleFunc("/redirect/", RedirectUrlHandler)

	// start the server
	port := "3000"
	fmt.Println("server started on PORT:", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Println("error:", err)
	}

}
