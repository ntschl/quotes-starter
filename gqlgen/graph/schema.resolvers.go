package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/ntschl/quotes-starter/gqlgen/graph/generated"
	"github.com/ntschl/quotes-starter/gqlgen/graph/model"
)

// RandomQuote is the resolver for the randomQuote field.
func (r *queryResolver) RandomQuote(ctx context.Context) (*model.Quote, error) {
	request, err := http.NewRequest("GET", "http://34.160.33.1:80/quotes", nil)
	request.Header.Set("x-api-key", "COCKTAILSAUCE")
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	response, _ := client.Do(request)
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var quote model.Quote
	json.Unmarshal(data, &quote)

	return &quote, nil
}

// QuoteByID is the resolver for the quoteByID field.
func (r *queryResolver) QuoteByID(ctx context.Context, id string) (*model.Quote, error) {
	panic(fmt.Errorf("not implemented: QuoteByID - quoteByID"))
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
