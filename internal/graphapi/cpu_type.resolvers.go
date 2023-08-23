package graphapi

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.34

import (
	"context"
	"fmt"

	"go.infratographer.com/server-api/internal/ent/generated"
	"go.infratographer.com/x/gidx"
)

// ServerCPUType is the resolver for the serverCPUType field.
func (r *mutationResolver) ServerCPUType(ctx context.Context, input generated.CreateServerCPUTypeInput) (*ServerCPUTypeCreatePayload, error) {
	panic(fmt.Errorf("not implemented: ServerCPUType - serverCPUType"))
}

// ServerCPUTypeUpdate is the resolver for the serverCPUTypeUpdate field.
func (r *mutationResolver) ServerCPUTypeUpdate(ctx context.Context, id gidx.PrefixedID, input generated.UpdateServerCPUTypeInput) (*ServerCPUTypeUpdatePayload, error) {
	panic(fmt.Errorf("not implemented: ServerCPUTypeUpdate - serverCPUTypeUpdate"))
}

// ServerCPUTypeDelete is the resolver for the serverCPUTypeDelete field.
func (r *mutationResolver) ServerCPUTypeDelete(ctx context.Context, id gidx.PrefixedID) (*ServerCPUTypeDeletePayload, error) {
	panic(fmt.Errorf("not implemented: ServerCPUTypeDelete - serverCPUTypeDelete"))
}

// ServerCPUType is the resolver for the serverCPUType field.
func (r *queryResolver) ServerCPUType(ctx context.Context, id gidx.PrefixedID) (*generated.ServerCPUType, error) {
	panic(fmt.Errorf("not implemented: ServerCPUType - serverCPUType"))
}
