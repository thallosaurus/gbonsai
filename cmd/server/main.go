package main

import (
	"bytes"
	"io"
	"log"
	"net/http"

	"github.com/aymanbagabas/go-pty"
	"github.com/gorilla/mux"
	"github.com/leaanthony/go-ansi-parser"
)

func bonsai() []byte {
	pty, err := pty.New()

	if err != nil {
		log.Fatalf("failed to open pty: %s", err)
	}

	pty.Resize(100, 100)

	defer pty.Close()

	c := pty.Command("cbonsai", "-p", "-s", "10")

	if err := c.Start(); err != nil {
		log.Fatalf("failed to start child: %s", err)
	}

	var buf bytes.Buffer
	io.Copy(&buf, pty)

	//log.Print(buf.String())
	cleaner, err := ansi.Cleanse(buf.String())
	if err != nil {
		log.Fatal(err)
	}
	log.Print(cleaner)

	//return ansihtml.ConvertToHTML(buf.Bytes())
	return buf.Bytes()
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/bonsai", func(w http.ResponseWriter, r *http.Request) {
		b := bonsai()
		//w.Write(b)
		w.Header().Add("Content-Type", "text/html")
		w.Write(b)
	})
	http.Handle("/", router)

	log.Fatal(http.ListenAndServe(":8080", router))
}
