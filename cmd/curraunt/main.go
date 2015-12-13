// Copyright (C) 2015 Max Riveiro <kavu13@gmail.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
// The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package main

import (
	"flag"
	"github.com/davecheney/profile"
	"github.com/julienschmidt/httprouter"
	"github.com/kavu/curraunt"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {
	portFlag := flag.Int("p", 80, "a port to run server on")
	flag.Parse()
	port := strconv.Itoa(*portFlag)

	if len(os.Getenv("PPROF")) > 0 {
		profile.Start(&profile.Config{
			CPUProfile:     true,
			MemProfile:     true,
			BlockProfile:   false,
			ProfilePath:    ".",
			NoShutdownHook: false,
		})
	}

	// set logging parameters
	log.SetFlags(log.LstdFlags | log.LUTC | log.Lmicroseconds)

	// init and start internal db
	curraunt.InitDB()

	// build http routes
	router := httprouter.New()
	router.HandlerFunc("GET", "/latest", curraunt.LatestHandler)

	// start http server
	log.Println("curraunt on port " + port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
