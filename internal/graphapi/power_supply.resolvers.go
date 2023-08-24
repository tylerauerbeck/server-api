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

// ServerPowerSupply is the resolver for the serverPowerSupply field.
func (r *mutationResolver) ServerPowerSupply(ctx context.Context, input generated.CreateServerPowerSupplyInput) (*ServerPowerSupplyCreatePayload, error) {
	panic(fmt.Errorf("not implemented: ServerPowerSupply - serverPowerSupply"))
}

// ServerPowerSupplyUpdate is the resolver for the serverPowerSupplyUpdate field.
func (r *mutationResolver) ServerPowerSupplyUpdate(ctx context.Context, id gidx.PrefixedID, input generated.UpdateServerPowerSupplyInput) (*ServerPowerSupplyUpdatePayload, error) {
	panic(fmt.Errorf("not implemented: ServerPowerSupplyUpdate - serverPowerSupplyUpdate"))
}

// ServerPowerSupplyDelete is the resolver for the serverPowerSupplyDelete field.
func (r *mutationResolver) ServerPowerSupplyDelete(ctx context.Context, id gidx.PrefixedID) (*ServerPowerSupplyTypeDeletePayload, error) {
	panic(fmt.Errorf("not implemented: ServerPowerSupplyDelete - serverPowerSupplyDelete"))
}

// ServerPowerSupply is the resolver for the serverPowerSupply field.
func (r *queryResolver) ServerPowerSupply(ctx context.Context, id gidx.PrefixedID) (*generated.ServerPowerSupply, error) {
	panic(fmt.Errorf("not implemented: ServerPowerSupply - serverPowerSupply"))
}
