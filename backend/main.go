package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humachi"
	"github.com/danielgtaylor/huma/v2/humacli"
	"github.com/go-chi/chi/v5"
)

// GreetingInput represents the input for our greeting operation.
type GreetingInput struct {
	Name string `path:"name" maxLength:"30" example:"world" doc:"Name to greet"`
}

// GreetingOutput represents the output for our greeting operation.
type GreetingOutput struct {
	Body struct {
		Message string `json:"message" example:"Hello, world!" doc:"Greeting message"`
	}
}

// Options for the CLI.
type Options struct {
	Port int    `help:"Port to listen on" short:"p" default:"8888"`
	Host string `help:"Host to listen on" default:"localhost"`
}

func main() {
	// Create a CLI app which takes a function that is called with the
	// CLI hooks and options.
	cli := humacli.New(func(hooks humacli.Hooks, options *Options) {
		// 1. Create a new router & API instance.
		router := chi.NewMux()
		api := humachi.New(router, huma.DefaultConfig("My API", "1.0.0"))

		// 2. Register a new operation.
		huma.Register(api, huma.Operation{
			OperationID: "get-greeting",
			Method:      http.MethodGet,
			Path:        "/greeting/{name}",
			Summary:     "Get a greeting",
			Description: "Get a friendly greeting for a user.",
			Tags:        []string{"Greetings"},
		}, func(ctx context.Context, input *GreetingInput) (*GreetingOutput, error) {
			resp := &GreetingOutput{}
			resp.Body.Message = fmt.Sprintf("Hello, %s!", input.Name)
			return resp, nil
		})

		// 3. Start the server!
		hooks.OnStart(func() {
			fmt.Printf("Server running at http://%s:%d\n", options.Host, options.Port)
			http.ListenAndServe(fmt.Sprintf("%s:%d", options.Host, options.Port), router)
		})
	})

	// Run the CLI.
	cli.Run()
}
