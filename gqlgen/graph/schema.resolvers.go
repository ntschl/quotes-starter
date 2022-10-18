package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/ntschl/quotes-starter/gqlgen/graph/generated"
	"github.com/ntschl/quotes-starter/gqlgen/graph/model"
)

// CreateQuote is the resolver for the createQuote field.
func (r *mutationResolver) CreateQuote(ctx context.Context, input model.NewQuote) (*model.Quote, error) {
	quote := &model.Quote{
		Quote:  input.Quote,
		Author: input.Author,
	}

	byteArray, _ := json.Marshal(quote)
	buffer := bytes.NewBuffer(byteArray)

	request, _ := http.NewRequest("POST", "http://34.160.33.1:80/quotes", buffer)
	request.Header.Set("x-api-key", fmt.Sprint(ctx.Value("myKey")))

	client := &http.Client{}
	response, _ := client.Do(request)

	data, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	json.Unmarshal(data, quote)

	return quote, nil
}

// DeleteQuote is the resolver for the deleteQuote field.
func (r *mutationResolver) DeleteQuote(ctx context.Context, id string) (*string, error) {
	url := "http://34.160.33.1:80/quotes/" + id
	request, err := http.NewRequest("DELETE", url, nil)
	request.Header.Set("x-api-key", fmt.Sprint(ctx.Value("myKey")))
	if err != nil {
		return nil, err
	}
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	return &response.Status, nil
}

// RandomQuote is the resolver for the randomQuote field.
func (r *queryResolver) RandomQuote(ctx context.Context) (*model.Quote, error) {
	request, err := http.NewRequest("GET", "http://34.160.33.1:80/quotes", nil)
	request.Header.Set("x-api-key", fmt.Sprint(ctx.Value("myKey")))
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	response, _ := client.Do(request)
	data, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var quote model.Quote
	json.Unmarshal(data, &quote)

	return &quote, nil
}

// QuoteByID is the resolver for the quoteByID field.
func (r *queryResolver) QuoteByID(ctx context.Context, id string) (*model.Quote, error) {
	url := "http://34.160.33.1:80/quotes/" + id
	request, err := http.NewRequest("GET", url, nil)
	request.Header.Set("x-api-key", fmt.Sprint(ctx.Value("myKey")))
	if err != nil {
		return nil, err
	}
	client := &http.Client{}
	response, _ := client.Do(request)
	data, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	var quote model.Quote
	json.Unmarshal(data, &quote)
	return &quote, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
