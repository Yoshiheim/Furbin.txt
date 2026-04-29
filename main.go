package main

import (
	"flag"
	"fmt"
	"hoxt/data"
	"hoxt/internal/db"
	"hoxt/internal/helpers"
	"hoxt/internal/router"
	"log"
	"net/http"
)

var logoflag *string

func main() {
	db.InitDataBase()

	// Get config.json
	data.InitConfig("./data/config.json")
	fmt.Printf("GETS config.json - OK\n")

	// Handle static directory.
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	if data.Configs.FaviconPath != "" {
		// /HOXT/data/config.json
		http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
			http.ServeFile(w, r, data.Configs.FaviconPath)
		})

		fmt.Printf("GET THE favicon.ico - OK\n")
	}

	// Init Route from /HOAX/internal/router/*
	router.InitRoute()
	fmt.Println("INIT ROUTE")

	// Init of timer for clear pastes
	helpers.Timer()
	fmt.Println("INIT TIMER CLEARER")

	// flags of app.
	hostflag := flag.String("host", data.Configs.Host, "Host Of Website")
	portflag := flag.String("port", data.Configs.Port, "Port Of Website")
	//logoflag := flag.String("logo_path", "./data/cute_furry_raptor.png", "Path to Image for Convert to ASCII")

	flag.Parse()

	/*
		if logoflag != nil {

			// for convert image to ASCII
			flags := aic_package.DefaultFlags()

			flags.Dimensions = []int{450, 300}
			flags.CustomMap = " .:-~+=*%&)@"
			flags.Full = false
			flags.Dither = true

			asciiArt, err := aic_package.Convert(*logoflag, flags)
			if err != nil {
				fmt.Printf("-- %s --\n", err)
				data.Logo = []byte("")
			} else {
				fmt.Printf("CONVERT IMAGE TO ASCIII - OK\n")
				data.Logo = []byte(asciiArt)
			}
		}
	*/
	c, err := data.LoadDynamicConfig("./data/textconf.json")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("\n%v\n\n%v\n", data.Configs, c)

	if *hostflag != data.Configs.Host && *portflag == data.Configs.Port {

		fmt.Printf("[NOT EQUAL]\n")

		// without any flags, website will use host nd port from ./data/conifg.json
		// for avoiding hardcoding whats i did before
		// because hosting use 0.0.0.0:10000, but for test i'll use 127.0.0.1:8080.
		// you can change it for your facilities.

		// im lazy so just use "./run.sh"

		log.Printf("Server ran on http://%s:%s\n", data.Configs.Host, data.Configs.Port)

		http.ListenAndServe(data.Configs.Host+":"+data.Configs.Port, nil)
	} else {

		// if command to run website use flag like "go run main.go -host=127.0.0.1 -port=8080"
		// but im lazy so just use "./run.sh local"

		log.Printf("Server ran on http://%s:%s\n", *hostflag, *portflag)

		http.ListenAndServe(*hostflag+":"+*portflag, nil)

	}
}
