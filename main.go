package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/ServiceWeaver/weaver"
	"github.com/rynowak/serviceweaver-calculator/pkg/calculator"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

func main() {
	// Initialize the Service Weaver application.
	root := weaver.Init(context.Background())

	adder, err := weaver.Get[calculator.Adder](root)
	if err != nil {
		log.Fatal(err)
	}

	subtracter, err := weaver.Get[calculator.Subtracter](root)
	if err != nil {
		log.Fatal(err)
	}

	multiplier, err := weaver.Get[calculator.Multiplier](root)
	if err != nil {
		log.Fatal(err)
	}

	divider, err := weaver.Get[calculator.Divider](root)
	if err != nil {
		log.Fatal(err)
	}

	// Get a network listener on address "localhost:12345".
	opts := weaver.ListenerOptions{LocalAddress: "localhost:12345"}
	listener, err := root.Listener("calculator", opts)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("calculator listener available on %v\n", listener)

	http.HandleFunc("/addapp/method/add", handler(adder.Add))
	http.HandleFunc("/subtractapp/method/subtract", handler(subtracter.Subtract))
	http.HandleFunc("/multiplyapp/method/multiply", handler(multiplier.Multiply))
	http.HandleFunc("/divideapp/method/divide", handler(divider.Divide))

	http.Serve(listener, otelhttp.NewHandler(http.DefaultServeMux, "http"))
}

func handler(calc func(ctx context.Context, x float64, y float64) (float64, error)) func(w http.ResponseWriter, r *http.Request) {
	// Adapts the JSON API used by React Calculator to the internal serviceweaver
	// calls.
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// The react app uses quoted numbers :-| ... why??
		type payload struct {
			X string `json:"operandOne"`
			Y string `json:"operandTwo"`
		}

		p := payload{}
		err = json.Unmarshal(body, &p)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		x, err := strconv.ParseFloat(p.X, 64)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		y, err := strconv.ParseFloat(p.Y, 64)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		result, err := calc(r.Context(), x, y)
		if errors.Is(err, &calculator.MathError{}) {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		} else if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "\"%v\"", result)
	}
}
