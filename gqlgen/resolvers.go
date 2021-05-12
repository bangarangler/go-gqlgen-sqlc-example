package gqlgen

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

import (
	"context"

	"github.com/bangarangler/go-gqlgen-sqlc-example/pg"
)

type Resolver struct {
	Repository pg.Repository
}

func (r *agentResolver) Authors(ctx context.Context, obj *pg.Agent) ([]pg.Author, error) {
	panic("not implemented")
}

func (r *authorResolver) Website(ctx context.Context, obj *pg.Author) (*string, error) {
	panic("not implemented")
}

func (r *authorResolver) Agent(ctx context.Context, obj *pg.Author) (*pg.Agent, error) {
	panic("not implemented")
}

func (r *authorResolver) Books(ctx context.Context, obj *pg.Author) ([]pg.Book, error) {
	panic("not implemented")
}

func (r *bookResolver) Authors(ctx context.Context, obj *pg.Book) ([]pg.Author, error) {
	panic("not implemented")
}

func (r *mutationResolver) CreateAgent(ctx context.Context, data AgentInput) (*pg.Agent, error) {
	panic("not implemented")
}

func (r *mutationResolver) UpdateAgent(ctx context.Context, id string, data AgentInput) (*pg.Agent, error) {
	panic("not implemented")
}

func (r *mutationResolver) DeleteAgent(ctx context.Context, id string) (*pg.Agent, error) {
	panic("not implemented")
}

func (r *mutationResolver) CreateAuthor(ctx context.Context, data AuthorInput) (*pg.Author, error) {
	panic("not implemented")
}

func (r *mutationResolver) UpdateAuthor(ctx context.Context, id string, data AuthorInput) (*pg.Author, error) {
	panic("not implemented")
}

func (r *mutationResolver) DeleteAuthor(ctx context.Context, id string) (*pg.Author, error) {
	panic("not implemented")
}

func (r *mutationResolver) CreateBook(ctx context.Context, data BookInput) (*pg.Book, error) {
	panic("not implemented")
}

func (r *mutationResolver) UpdateBook(ctx context.Context, id string, data BookInput) (*pg.Book, error) {
	panic("not implemented")
}

func (r *mutationResolver) DeleteBook(ctx context.Context, id string) (*pg.Book, error) {
	panic("not implemented")
}

func (r *queryResolver) Agent(ctx context.Context, id string) (*pg.Agent, error) {
	panic("not implemented")
}

func (r *queryResolver) Agents(ctx context.Context) ([]pg.Agent, error) {
	panic("not implemented")
}

func (r *queryResolver) Author(ctx context.Context, id string) (*pg.Author, error) {
	panic("not implemented")
}

func (r *queryResolver) Authors(ctx context.Context) ([]pg.Author, error) {
	panic("not implemented")
}

func (r *queryResolver) Book(ctx context.Context, id string) (*pg.Book, error) {
	panic("not implemented")
}

func (r *queryResolver) Books(ctx context.Context) ([]pg.Book, error) {
	panic("not implemented")
}

// Agent returns AgentResolver implementation.
func (r *Resolver) Agent() AgentResolver { return &agentResolver{r} }

// Author returns AuthorResolver implementation.
func (r *Resolver) Author() AuthorResolver { return &authorResolver{r} }

// Book returns BookResolver implementation.
func (r *Resolver) Book() BookResolver { return &bookResolver{r} }

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type agentResolver struct{ *Resolver }

func (r *agentResolver) Authors(ctx context.Context, obj *pg.Agent) ([]pg.Author, error) {
	panic("not implemented")
}

type authorResolver struct{ *Resolver }
type bookResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
