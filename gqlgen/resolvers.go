package gqlgen

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

import (
	"context"
	"strconv"
	// "strconv"

	"github.com/bangarangler/go-gqlgen-sqlc-example/pg"
)

// Resolver connects individual resolvers with the datalayer.
type Resolver struct {
	Repository pg.Repository
}

// Agent returns an implementation of the AgentResolver interface.
func (r *Resolver) Agent() AgentResolver {
	return &agentResolver{r}
}

// Author returns an implementation of the AuthorResolver interface.
func (r *Resolver) Author() AuthorResolver {
	return &authorResolver{r}
}

// Book returns an implementation of the BookResolver interface.
func (r *Resolver) Book() BookResolver {
	return &bookResolver{r}
}

// Mutation returns an implementation of the MutationResolver interface.
func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}

// Query returns an implementation of the QueryResolver interface.
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type agentResolver struct{ *Resolver }

func (r *agentResolver) Authors(ctx context.Context, obj *pg.Agent) ([]pg.Author, error) {
	// convert := strconv.FormatInt(obj.ID, 10)
	return r.Repository.ListAuthorsByAgentID(ctx, obj.ID)
}

type authorResolver struct{ *Resolver }

func (r *authorResolver) Website(ctx context.Context, obj *pg.Author) (*string, error) {
	var w string
	if obj.Website.Valid {
		w = obj.Website.String
		return &w, nil
	}
	return nil, nil
}

func (r *authorResolver) Agent(ctx context.Context, obj *pg.Author) (*pg.Agent, error) {
	// convert := strconv.FormatInt(obj.AgentID, 10)
	agent, err := r.Repository.GetAgent(ctx, obj.AgentID)
	if err != nil {
		return nil, err
	}
	return &agent, nil
}

func (r *authorResolver) Books(ctx context.Context, obj *pg.Author) ([]pg.Book, error) {
	// convert := strconv.FormatInt(obj.AgentID, 10)
	return r.Repository.ListBooksByAuthorID(ctx, obj.AgentID)
}

type bookResolver struct{ *Resolver }

func (r *bookResolver) Authors(ctx context.Context, obj *pg.Book) ([]pg.Author, error) {
	// convert := strconv.FormatInt(obj.ID, 10)
	return r.Repository.ListAuthorsByBookID(ctx, obj.ID)
}

type mutationResolver struct{ *Resolver }

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

func (r *mutationResolver) UpdateAgent(ctx context.Context, id int64, data AgentInput) (*pg.Agent, error) {
	agent, err := r.Repository.UpdateAgent(ctx, pg.UpdateAgentParams{
		ID:    id,
		Name:  data.Name,
		Email: data.Email,
	})
	if err != nil {
		return nil, err
	}
	return &agent, nil
}

func (r *mutationResolver) DeleteAgent(ctx context.Context, id int64) (*pg.Agent, error) {
	agent, err := r.Repository.DeleteAgent(ctx, id)
	if err != nil {
		return nil, err
	}
	return &agent, nil
}

func (r *mutationResolver) CreateAuthor(ctx context.Context, data AuthorInput) (*pg.Author, error) {
	var convert = data.AgentID
	i, _ := strconv.ParseInt(convert, 10, 64)
	author, err := r.Repository.CreateAuthor(ctx, pg.CreateAuthorParams{
		Name:    data.Name,
		Website: pg.StringPtrToNullString(data.Website),
		// AgentID: data.AgentID,
		AgentID: i,
	})
	if err != nil {
		return nil, err
	}
	return &author, nil
}

func (r *mutationResolver) UpdateAuthor(ctx context.Context, id int64, data AuthorInput) (*pg.Author, error) {
	var convert = data.AgentID
	i, _ := strconv.ParseInt(convert, 10, 64)
	author, err := r.Repository.UpdateAuthor(ctx, pg.UpdateAuthorParams{
		ID:      id,
		Name:    data.Name,
		Website: pg.StringPtrToNullString(data.Website),
		// AgentID: data.AgentID,
		AgentID: i,
	})
	if err != nil {
		return nil, err
	}
	return &author, nil
}

func (r *mutationResolver) DeleteAuthor(ctx context.Context, id int64) (*pg.Author, error) {
	// convert := strconv.FormatInt(id, 10)
	author, err := r.Repository.DeleteAuthor(ctx, id)
	if err != nil {
		return nil, err
	}
	return &author, nil
}

func (r *mutationResolver) CreateBook(ctx context.Context, data BookInput) (*pg.Book, error) {
	var convert = data.AuthorIDs
	var c = []int64{}
	for _, i := range convert {
		println("i", i)
		j, err := strconv.ParseInt(i, 10, 64)
		if err != nil {
			panic(err)
		}
		c = append(c, j)
	}
	// i,_ := strconv.ParseInt(convert, 10, 64)
	return r.Repository.CreateBook(ctx, pg.CreateBookParams{
		Title:       data.Title,
		Description: data.Description,
		Cover:       data.Cover,
		// }, data.AuthorIDs)
	}, c)
}

func (r *mutationResolver) UpdateBook(ctx context.Context, id int64, data BookInput) (*pg.Book, error) {
	var convert = data.AuthorIDs
	var c = []int64{}
	for _, i := range convert {
		println("i", i)
		j, err := strconv.ParseInt(i, 10, 64)
		if err != nil {
			panic(err)
		}
		c = append(c, j)
	}
	return r.Repository.UpdateBook(ctx, pg.UpdateBookParams{
		ID:          id,
		Title:       data.Title,
		Description: data.Description,
		Cover:       data.Cover,
		// }, data.AuthorIDs)
	}, c)
}

func (r *mutationResolver) DeleteBook(ctx context.Context, id int64) (*pg.Book, error) {
	// BookAuthors associations will cascade automatically.
	// convert := strconv.FormatInt(id, 10)
	book, err := r.Repository.DeleteBook(ctx, id)
	if err != nil {
		return nil, err
	}
	return &book, nil
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Agent(ctx context.Context, id int64) (*pg.Agent, error) {
	// convert := strconv.FormatInt(id, 10)
	agent, err := r.Repository.GetAgent(ctx, id)
	if err != nil {
		return nil, err
	}
	return &agent, nil
}

func (r *queryResolver) Agents(ctx context.Context) ([]pg.Agent, error) {
	return r.Repository.ListAgents(ctx)
}

func (r *queryResolver) Author(ctx context.Context, id int64) (*pg.Author, error) {
	// convert := strconv.FormatInt(id, 10)
	author, err := r.Repository.GetAuthor(ctx, id)
	if err != nil {
		return nil, err
	}
	return &author, nil
}

func (r *queryResolver) Authors(ctx context.Context) ([]pg.Author, error) {
	return r.Repository.ListAuthors(ctx)
}

func (r *queryResolver) Book(ctx context.Context, id int64) (*pg.Book, error) {
	// convert := strconv.FormatInt(id, 10)
	book, err := r.Repository.GetBook(ctx, id)
	if err != nil {
		return nil, err
	}
	return &book, nil
}

func (r *queryResolver) Books(ctx context.Context) ([]pg.Book, error) {
	return r.Repository.ListBooks(ctx)
}
