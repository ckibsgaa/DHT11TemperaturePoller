package main

import (
    "os"
	"log"
	"strings"
	"net/http"
	"fmt"
	"github.com/jacobsa/go-serial/serial"
)

var portnavn = os.Args[1]

func woot(w http.ResponseWriter, r *http.Request) {
	options := serial.OpenOptions{
		PortName: portnavn,
      BaudRate: 9600,
      DataBits: 8,
      StopBits: 1,
      MinimumReadSize: 4,
	}

	// Open the port.
    port, err := serial.Open(options)
    if err != nil {
      log.Fatalf("serial.Open: %v", err)
    }

    // Make sure to close it later.
    defer port.Close()

    buf := make([]byte, 128)
    br, err := port.Read(buf)
    if err != nil {
		log.Fatal(err)
	}

	rawstring := string(buf[:br])

	stringarr := strings.Split(rawstring, ";")

	fmt.Fprintf(w, "temperature " + stringarr[0] + "\n")
	fmt.Fprintf(w, "humidity " + stringarr[1] + "\n")
}

func main() {
	// The Handler function provides a default handler to expose metrics
	// via an HTTP server. "/metrics" is the usual endpoint for that.
	

	http.HandleFunc("/woot", woot)
	log.Fatal(http.ListenAndServe(":8100", nil))
}
