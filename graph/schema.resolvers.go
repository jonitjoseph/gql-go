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
	fmt.Println("CreateBook")
	book, err := r.BookRepository.CreateBook(&input)
	bookCreated := &model.Book{
		Author:    book.Author,
		Publisher: book.Publisher,
		Title:     book.Title,
		ID:        book.ID,
	}
	if err != nil {
		return nil, err
	}
	return bookCreated, nil
}

// DeleteBook is the resolver for the DeleteBook field.
func (r *mutationResolver) DeleteBook(ctx context.Context, id int) (string, error) {
	fmt.Println("DeleteBook")
	err := r.BookRepository.DeleteBook(id)
	if err != nil {
		return "", err
	}
	successMessage := "successfully deleted"
	return successMessage, nil
}

// UpdateBook is the resolver for the UpdateBook field.
func (r *mutationResolver) UpdateBook(ctx context.Context, id int, input model.BookInput) (*model.Book, error) {
	fmt.Println("UpdateBook")
	book, err := r.BookRepository.UpdateBook(&input, id)
	bookUpdated := &model.Book{
		Author:    book.Author,
		Publisher: book.Publisher,
		Title:     book.Title,
		ID:        book.ID,
	}
	if err != nil {
		return nil, err
	}
	// successMessage := "successfully updated"

	return bookUpdated, nil
}

// GetAllBooks is the resolver for the GetAllBooks field.
func (r *queryResolver) GetAllBooks(ctx context.Context) ([]*model.Book, error) {
	fmt.Println("GetAllBooks")
	books, err := r.BookRepository.GetAllBooks()
	if err != nil {
		return nil, err
	}
	return books, nil
}

// GetOneBook is the resolver for the GetOneBook field.
func (r *queryResolver) GetOneBook(ctx context.Context, id int) (*model.Book, error) {
	fmt.Println("GetOneBook")
	book, err := r.BookRepository.GetOneBook(id)
	selectedBook := &model.Book{
		ID:        book.ID,
		Author:    book.Author,
		Publisher: book.Publisher,
		Title:     book.Title,
	}
	if err != nil {
		return nil, err
	}
	return selectedBook, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
