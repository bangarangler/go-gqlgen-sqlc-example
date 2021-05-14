package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/bangarangler/go-gqlgen-sqlc-example/dataloaders"
	"github.com/bangarangler/go-gqlgen-sqlc-example/gqlgen"
	"github.com/bangarangler/go-gqlgen-sqlc-example/pg"
)

func main() {
	db, err := pg.Open(pg.PgConnStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// initialize the repository
	repo := pg.NewRepository(db)

	// initialize the dataloaders
	dl := dataloaders.NewRetriever()

	// configure the server
	mux := http.NewServeMux()
	mux.Handle("/", gqlgen.NewPlaygroundHandler("/query"))
	dlMiddleware := dataloaders.Middleware(repo)     // <- initialize the middleware
	queryHandler := gqlgen.NewHandler(repo, dl)      // <- use dataloader.Retriever
	mux.Handle("/query", dlMiddleware(queryHandler)) // <- use dataloaders.Middleware

	// run the server
	port := ":8080"
	fmt.Fprintf(os.Stdout, "🚀 Server ready at http://localhost%s\n", port)
	fmt.Fprintln(os.Stderr, http.ListenAndServe(port, mux))
}
