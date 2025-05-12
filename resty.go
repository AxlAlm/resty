package resty

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"reflect"
)

// Generic Resource and Controller
type Controller[InBody any, OutBody any, Params any] interface {
	Create(context.Context, InBody) (OutBody, error)
	GetOne(context.Context, Params) (OutBody, error)
}

func Resource[InBody any, OutBody any, Params any](
	mux *http.ServeMux,
	name string,
	controller Controller[InBody, OutBody, Params],
	omit []string,
) {
	mux.Handle(fmt.Sprintf("POST /%s", name), handleMutate(controller.Create))
	mux.Handle(fmt.Sprintf("GET /%s/{id}", name), handleQuery(controller.GetOne))

}

// Generic Handlers
type TargetFunc[In any, Out any] func(context.Context, In) (Out, error)

func handleMutate[In any, Out any](f TargetFunc[In, Out]) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var in In

		err := json.NewDecoder(r.Body).Decode(&in)
		if err != nil {
			http.Error(w, "invalid json", http.StatusBadRequest)
			return
		}

		out, err := f(r.Context(), in)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Format and write response
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		err = json.NewEncoder(w).Encode(out)
		if err != nil {
			log.Printf("failed to encode created note: %v", err)
			return
		}
	})
}

func handleQuery[In any, Out any](f TargetFunc[In, Out]) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var in In

		err := r.ParseForm()
		if err != nil {
			http.Error(w, "invalid json", http.StatusBadRequest)
			return
		}

		fmt.Println(r.PathValue("id"))

		fmt.Println(reflect.TypeOf(in))

		var zero [0]In
		fmt.Println("zero", zero)
		fmt.Println(reflect.TypeOf(zero))

		fmt.Println(reflect.TypeFor[TargetFunc[In, Out]]().Name())

		values := r.Form
		for k, v := range values {
			fmt.Println(k, v)
		}

		out, err := f(r.Context(), in)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Format and write response
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		err = json.NewEncoder(w).Encode(out)
		if err != nil {
			log.Printf("failed to encode created note: %v", err)
			return
		}
	})
}
