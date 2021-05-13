package gqlgen

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

import (
	"context"
	"fmt"
	// "strconv"

	"github.com/bangarangler/go-gqlgen-sqlc-example/pg"
)

type Resolver struct {
	Repository pg.Repository
}

func (r *authorResolver) Website(ctx context.Context, obj *pg.Author) (*string, error) {
	var w string
	if obj.Website.Valid {
		w = obj.Website.String
		return &w, nil
	}
	return nil, nil
}

//TODO: possible N+1 query issue to be resolved with dataloader
// many authors along with agents
func (r *authorResolver) Agent(ctx context.Context, obj *pg.Author) (*pg.Agent, error) {
	agent, err := r.Repository.GetAgent(ctx, obj.AgentID)
	if err != nil {
		return nil, err
	}
	return &agent, nil
}

// TODO: n+1 query issue to be fixed with dataloader
func (r *authorResolver) Books(ctx context.Context, obj *pg.Author) ([]pg.Book, error) {
	return r.Repository.ListBooksByAuthorID(ctx, obj.ID)
}

func (r *bookResolver) Authors(ctx context.Context, obj *pg.Book) ([]pg.Author, error) {
	return r.Repository.ListAuthorsByBookID(ctx, obj.ID)
}

func (r *mutationResolver) CreateAgent(ctx context.Context, data AgentInput) (*pg.Agent, error) {
	agent, err := r.Repository.CreateAgent(ctx, pg.CreateAgentParams{
		Name:  data.Name,
		Email: data.Email,
	})
	if err != nil {
		return nil, err
	}
	return &agent, nil
}

func (r *mutationResolver) UpdateAgent(ctx context.Context, id string, data AgentInput) (*pg.Agent, error) {
	panic("not implemented")
}

func (r *mutationResolver) DeleteAgent(ctx context.Context, id string) (*pg.Agent, error) {
	panic("not implemented")
}

func (r *mutationResolver) CreateAuthor(ctx context.Context, data AuthorInput) (*pg.Author, error) {
	// var convert data.AgentID
	// i, _ := strconv.ParseInt(convert, 10, 64)
	fmt.Printf("val: %v; type %T\n", data.AgentID)
	author, err := r.Repository.CreateAuthor(ctx, pg.CreateAuthorParams{
		Name:    data.Name,
		Website: pg.StringPtrToNullString(data.Website),
		// AgentID: data.AgentID,
		// AgentID: i,
	})
	if err != nil {
		return nil, err
	}
	return &author, nil
}

func (r *mutationResolver) UpdateAuthor(ctx context.Context, id string, data AuthorInput) (*pg.Author, error) {
	panic("not implemented")
}

func (r *mutationResolver) DeleteAuthor(ctx context.Context, id string) (*pg.Author, error) {
	panic("not implemented")
}

func (r *mutationResolver) CreateBook(ctx context.Context, data BookInput) (*pg.Book, error) {
	return r.Repository.CreateBook(ctx, pg.CreateBookParams{
		Title:       data.Title,
		Description: data.Description,
		Cover:       data.Cover,
		// }, data.AuthorIDs)
	}, []int64{1})
}

func (r *mutationResolver) UpdateBook(ctx context.Context, id string, data BookInput) (*pg.Book, error) {
	panic("not implemented")
}

func (r *mutationResolver) DeleteBook(ctx context.Context, id string) (*pg.Book, error) {
	panic("not implemented")
}

// func (r *queryResolver) Agent(ctx context.Context, id string) (*pg.Agent, error) {
func (r *queryResolver) Agent(ctx context.Context, id int64) (*pg.Agent, error) {
	agent, err := r.Repository.GetAgent(ctx, id)
	if err != nil {
		return nil, err
	}
	return &agent, nil
}

func (r *queryResolver) Agents(ctx context.Context) ([]pg.Agent, error) {
	return r.Repository.ListAgents(ctx)
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

// TODO add dataloader this will suffer from n+1 query issue
func (r *agentResolver) Authors(ctx context.Context, obj *pg.Agent) ([]pg.Author, error) {
	return r.Repository.ListAuthorsByAgentID(ctx, obj.ID)
}

type authorResolver struct{ *Resolver }
type bookResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
