package examples

import (
	"log"
	"os"

	"github.com/robertogyn19/gmusic"
)

func Login() *gmusic.GMusic {
	user := os.Getenv("GOOGLE_USER")
	pass := os.Getenv("GOOGLE_PASS")

	if user == "" || pass == "" {
		log.Fatal("Invalid credentials")
	}

	gm, err := gmusic.Login(user, pass)

	if err != nil {
		log.Fatal("Login error ", err)
	}

	return gm
}
