// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package graphapi

import (
	"go.infratographer.com/server-api/internal/ent/generated"
	"go.infratographer.com/x/gidx"
)

// Return response from serverChassisTypeCreate
type ServerChassisTypeCreatePayload struct {
	// The created server chassis type.
	ServerChassisType *generated.ServerChassisType `json:"serverChassisType"`
}

// Return response from serverChassisTypeDelete
type ServerChassisTypeDeletePayload struct {
	// The ID of the deleted server chassis type.
	DeletedID gidx.PrefixedID `json:"deletedID"`
}

// Return response from serverChassisTypeUpdate
type ServerChassisTypeUpdatePayload struct {
	// The updated server chassis type.
	ServerChassisType *generated.ServerChassisType `json:"serverChassisType"`
}

// Return response from serverComponentCreate
type ServerComponentCreatePayload struct {
	// The created server component.
	ServerComponent *generated.ServerComponent `json:"serverComponent"`
}

// Return response from serverComponentDelete
type ServerComponentDeletePayload struct {
	// The ID of the deleted server component.
	DeletedID gidx.PrefixedID `json:"deletedID"`
}

// Return response from serverComponentTypeCreate
type ServerComponentTypeCreatePayload struct {
	// The created server component type.
	ServerComponentType *generated.ServerComponentType `json:"serverComponentType"`
}

// Return response from serverComponentTypeDelete
type ServerComponentTypeDeletePayload struct {
	// The ID of the deleted server component type.
	DeletedID gidx.PrefixedID `json:"deletedID"`
}

// Return response from serverComponentTypeUpdate
type ServerComponentTypeUpdatePayload struct {
	// The updated server component type.
	ServerComponentType *generated.ServerComponentType `json:"serverComponentType"`
}

// Return response from serverComponentUpdate
type ServerComponentUpdatePayload struct {
	// The updated server component.
	ServerComponent *generated.ServerComponent `json:"serverComponent"`
}

// Return response from serverCreate
type ServerCreatePayload struct {
	// The created server.
	Server *generated.Server `json:"server"`
}

// Return response from serverDelete
type ServerDeletePayload struct {
	// The ID of the deleted server.
	DeletedID gidx.PrefixedID `json:"deletedID"`
}

// Return response from serverProviderCreate
type ServerProviderCreatePayload struct {
	// The created server provider.
	ServerProvider *generated.Provider `json:"serverProvider"`
}

// Return response from serverProviderDelete
type ServerProviderDeletePayload struct {
	// The ID of the deleted server provider.
	DeletedID gidx.PrefixedID `json:"deletedID"`
}

// Return response from serverProviderUpdate
type ServerProviderUpdatePayload struct {
	// The updated server provider.
	ServerProvider *generated.Provider `json:"serverProvider"`
}

// Return response from serverTypeCreate
type ServerTypeCreatePayload struct {
	// The created server type.
	ServerType *generated.ServerType `json:"serverType"`
}

// Return response from serverTypeDelete
type ServerTypeDeletePayload struct {
	// The ID of the deleted server type.
	DeletedID gidx.PrefixedID `json:"deletedID"`
}

// Return response from serverTypeUpdate
type ServerTypeUpdatePayload struct {
	// The updated server type.
	ServerType *generated.ServerType `json:"serverType"`
}

// Return response from serverUpdate
type ServerUpdatePayload struct {
	// The updated server.
	Server *generated.Server `json:"server"`
}
