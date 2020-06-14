package main

import (
	"flag"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jimareed/drawing"
	"github.com/jimareed/slides"
)

var mainDeck = slides.SlideDeck{}

func getHandler(w http.ResponseWriter, r *http.Request) {

	svg := ""

	if len(mainDeck.Slides) > 0 {
		svg, err := drawing.ToSvg(mainDeck.Slides[0].Drawing)
		if err != nil {
			log.Fatal(err)
		}
	}

	io.WriteString(w, "<html><body>"+svg+"</body></html>\n")
}

func main() {

	server := flag.Bool("server", false, "run in server mode")
	input := flag.String("input", "", "path to source")
	output := flag.String("output", "", "path to source")
	help := flag.Bool("help", false, "help")

	flag.Parse()

	if *help {
		log.Fatal("usage: slide-generator [-input <path>][-output <path>][-server][-help]")
	}

	if *input == "" {
		*input = "./slides"
	}

	log.Print("reading deck from ", *input)

	mainDeck, err := slides.Read(*input)
	if err == nil {
		log.Print(deck.Title, " read successful.")
	} else {
		log.Fatal(err)
	}

	if *output != "" {
		log.Print("writing ", deck.Title, " to ", *output)
		err = slides.Write(deck, *output)
		if err == nil {
			log.Print(deck.Title, " write successful.")
		} else {
			log.Fatal(err)
		}
	}

	if *server {
		r := mux.NewRouter()
		r.HandleFunc("/", getHandler).Methods("GET")

		log.Print("Server started on localhost:8080")
		log.Fatal(http.ListenAndServe(":8080", r))
	}
}
