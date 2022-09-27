package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"gql-go/graph/generated"
	"gql-go/graph/model"
)

// CreateBook is the resolver for the CreateBook field.
func (r *mutationResolver) CreateBook(ctx context.Context, input model.BookInput) (*model.Book, error) {
	panic(fmt.Errorf("not implemented: CreateBook - CreateBook"))
}

// DeleteBook is the resolver for the DeleteBook field.
func (r *mutationResolver) DeleteBook(ctx context.Context, id int) (string, error) {
	panic(fmt.Errorf("not implemented: DeleteBook - DeleteBook"))
}

// UpdateBook is the resolver for the UpdateBook field.
func (r *mutationResolver) UpdateBook(ctx context.Context, id int) (string, error) {
	panic(fmt.Errorf("not implemented: UpdateBook - UpdateBook"))
}

// GetAllBooks is the resolver for the GetAllBooks field.
func (r *queryResolver) GetAllBooks(ctx context.Context) ([]*model.Book, error) {
	panic(fmt.Errorf("not implemented: GetAllBooks - GetAllBooks"))
}

// GetOneBook is the resolver for the GetOneBook field.
func (r *queryResolver) GetOneBook(ctx context.Context, id int) (*model.Book, error) {
	panic(fmt.Errorf("not implemented: GetOneBook - GetOneBook"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
