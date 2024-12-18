package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.55

import (
	"context"
	"fmt"

	graph1 "github.com/israelalvesmelo/CleanArch/internal/infra/graph/generated"
	"github.com/israelalvesmelo/CleanArch/internal/infra/graph/model"
)

// CreateOrder is the resolver for the createOrder field.
func (r *mutationResolver) CreateOrder(ctx context.Context, input *model.OrderInput) (*model.Order, error) {
	panic(fmt.Errorf("not implemented: CreateOrder - createOrder"))
}

// Orders is the resolver for the orders field.
func (r *queryResolver) Orders(ctx context.Context) ([]*model.Order, error) {
	panic(fmt.Errorf("not implemented: Orders - orders"))
}

// Mutation returns graph1.MutationResolver implementation.
func (r *Resolver) Mutation() graph1.MutationResolver { return &mutationResolver{r} }

// Query returns graph1.QueryResolver implementation.
func (r *Resolver) Query() graph1.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
