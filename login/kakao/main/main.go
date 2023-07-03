package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"golang.org/x/oauth2"
)

var (
	state = "Osj6jGfD8zsr7ujxVw1t9b6tFJ9TR4EL"

	conf = &oauth2.Config{
		ClientID: "83aac3bc88df4a5b153d4349fed39ff8",

		// 카카오 app 설정에서 ClientSecret을 설정했다면 적어줘야 한다.
		ClientSecret: "Osj6jGfD8zsr7ujxVw1t9b6tFJ9TR4EL",
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://kauth.kakao.com/oauth/authorize",
			TokenURL: "https://kauth.kakao.com/oauth/token",
		},
		RedirectURL: "http://localhost/callback",
	}

	// conf = &oauth2.Config{
	// 	ClientID:     "YOUR_CLIENT_ID",
	// 	ClientSecret: "YOUR_CLIENT_SECRET",
	// 	Endpoint: oauth2.Endpoint{
	// 		AuthURL:  "https://nid.naver.com/oauth2.0/authorize",
	// 		TokenURL: "https://nid.naver.com/oauth2.0/token",
	// 	},
	// 	RedirectURL: "YOUR_REDIRECT_URL",
	// }
)

const htmlIndex = `<html><body>
Logged in with <a href="/login">KAKAO</a>
</body></html>
`

// /
func handleMain(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(htmlIndex))
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	u := conf.AuthCodeURL(state, oauth2.AccessTypeOffline)

	http.Redirect(w, r, u, http.StatusTemporaryRedirect)
}

func handleCallback(w http.ResponseWriter, r *http.Request) {

	s := r.FormValue("state")
	if s != state {
		fmt.Printf("invalid oauth state, expected '%s', got '%s'\n", state, s)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	fmt.Println("state", s)

	c := r.FormValue("code")

	fmt.Println("code", c)

	// Use the custom HTTP client when requesting a token.
	httpClient := &http.Client{Timeout: 2 * time.Second}
	ctx := context.WithValue(oauth2.NoContext, oauth2.HTTPClient, httpClient)

	token, err := conf.Exchange(ctx, c)
	if err != nil {
		fmt.Printf("conf.Exchange() failed with '%s'\n", err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	client := conf.Client(ctx, token)
	_ = client

	fmt.Println(client)
	fmt.Println(token)

	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}

func main() {
	http.HandleFunc("/", handleMain)
	http.HandleFunc("/login", handleLogin)
	http.HandleFunc("/callback", handleCallback)

	fmt.Print("Started running on Server\n")
	fmt.Println(http.ListenAndServe(":80", nil))
}
