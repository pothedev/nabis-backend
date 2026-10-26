package main

import (
	"fmt"
)

/* endpoints:
GET /api/sessions
	no params are sent, backend automatically grabs a cookie
	secretly passed by the frontend and queries it in redis,
	if it exists -> 200

GET /api/auth/spotify/login
	returns a redirect link to spotify login api

GET /api/auth/spotify/callback
	hz will figure out later


GET /api/concerts
	pulls users following list,
	checks whenever that list matches the ones saved in redis
	scrapes concerts info for new entires
	returns []Concert type
*/

// responses
type SessionsResponse struct {
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type Concert struct {
	Artist      string `json:"artist"`
	City        string `json:"city"`
	Start_time  string `json:"start_time"`
	Title       string `json:"title"`
	Spotify_uri string `json:"spotify_uri"`
}

type ConcertsResponse struct {
	Concerts []Concert `json:"concerts"`
}

// funcs
func scrapeConcerts(artist_id string) {
	//concerts []Concert
	//return concerts
}

func main() {
	fmt.Println("Hello, World!")
}
