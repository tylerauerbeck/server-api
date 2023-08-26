// Copyright 2023 The Infratographer Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by entc, DO NOT EDIT.

package generated

import (
	"go.infratographer.com/x/gidx"
)

// CreateServerProviderInput represents a mutation input for creating serverproviders.
type CreateServerProviderInput struct {
	Name               string
	ResourceProviderID gidx.PrefixedID
}

// Mutate applies the CreateServerProviderInput on the ProviderMutation builder.
func (i *CreateServerProviderInput) Mutate(m *ProviderMutation) {
	m.SetName(i.Name)
	m.SetResourceProviderID(i.ResourceProviderID)
}

// SetInput applies the change-set in the CreateServerProviderInput on the ProviderCreate builder.
func (c *ProviderCreate) SetInput(i CreateServerProviderInput) *ProviderCreate {
	i.Mutate(c.Mutation())
	return c
}

// UpdateServerProviderInput represents a mutation input for updating serverproviders.
type UpdateServerProviderInput struct {
	Name *string
}

// Mutate applies the UpdateServerProviderInput on the ProviderMutation builder.
func (i *UpdateServerProviderInput) Mutate(m *ProviderMutation) {
	if v := i.Name; v != nil {
		m.SetName(*v)
	}
}

// SetInput applies the change-set in the UpdateServerProviderInput on the ProviderUpdate builder.
func (c *ProviderUpdate) SetInput(i UpdateServerProviderInput) *ProviderUpdate {
	i.Mutate(c.Mutation())
	return c
}

// SetInput applies the change-set in the UpdateServerProviderInput on the ProviderUpdateOne builder.
func (c *ProviderUpdateOne) SetInput(i UpdateServerProviderInput) *ProviderUpdateOne {
	i.Mutate(c.Mutation())
	return c
}

// CreateServerInput represents a mutation input for creating servers.
type CreateServerInput struct {
	Name         string
	Description  *string
	OwnerID      gidx.PrefixedID
	LocationID   gidx.PrefixedID
	ProviderID   gidx.PrefixedID
	ServerTypeID gidx.PrefixedID
	ComponentIDs []gidx.PrefixedID
}

// Mutate applies the CreateServerInput on the ServerMutation builder.
func (i *CreateServerInput) Mutate(m *ServerMutation) {
	m.SetName(i.Name)
	if v := i.Description; v != nil {
		m.SetDescription(*v)
	}
	m.SetOwnerID(i.OwnerID)
	m.SetLocationID(i.LocationID)
	m.SetProviderID(i.ProviderID)
	m.SetServerTypeID(i.ServerTypeID)
	if v := i.ComponentIDs; len(v) > 0 {
		m.AddComponentIDs(v...)
	}
}

// SetInput applies the change-set in the CreateServerInput on the ServerCreate builder.
func (c *ServerCreate) SetInput(i CreateServerInput) *ServerCreate {
	i.Mutate(c.Mutation())
	return c
}

// UpdateServerInput represents a mutation input for updating servers.
type UpdateServerInput struct {
	Name               *string
	ClearDescription   bool
	Description        *string
	ClearComponents    bool
	AddComponentIDs    []gidx.PrefixedID
	RemoveComponentIDs []gidx.PrefixedID
}

// Mutate applies the UpdateServerInput on the ServerMutation builder.
func (i *UpdateServerInput) Mutate(m *ServerMutation) {
	if v := i.Name; v != nil {
		m.SetName(*v)
	}
	if i.ClearDescription {
		m.ClearDescription()
	}
	if v := i.Description; v != nil {
		m.SetDescription(*v)
	}
	if i.ClearComponents {
		m.ClearComponents()
	}
	if v := i.AddComponentIDs; len(v) > 0 {
		m.AddComponentIDs(v...)
	}
	if v := i.RemoveComponentIDs; len(v) > 0 {
		m.RemoveComponentIDs(v...)
	}
}

// SetInput applies the change-set in the UpdateServerInput on the ServerUpdate builder.
func (c *ServerUpdate) SetInput(i UpdateServerInput) *ServerUpdate {
	i.Mutate(c.Mutation())
	return c
}

// SetInput applies the change-set in the UpdateServerInput on the ServerUpdateOne builder.
func (c *ServerUpdateOne) SetInput(i UpdateServerInput) *ServerUpdateOne {
	i.Mutate(c.Mutation())
	return c
}

// CreateServerCPUInput represents a mutation input for creating servercpus.
type CreateServerCPUInput struct {
	Serial          string
	ServerID        gidx.PrefixedID
	ServerCPUTypeID gidx.PrefixedID
}

// Mutate applies the CreateServerCPUInput on the ServerCPUMutation builder.
func (i *CreateServerCPUInput) Mutate(m *ServerCPUMutation) {
	m.SetSerial(i.Serial)
	m.SetServerID(i.ServerID)
	m.SetServerCPUTypeID(i.ServerCPUTypeID)
}

// SetInput applies the change-set in the CreateServerCPUInput on the ServerCPUCreate builder.
func (c *ServerCPUCreate) SetInput(i CreateServerCPUInput) *ServerCPUCreate {
	i.Mutate(c.Mutation())
	return c
}

// UpdateServerCPUInput represents a mutation input for updating servercpus.
type UpdateServerCPUInput struct {
	Serial *string
}

// Mutate applies the UpdateServerCPUInput on the ServerCPUMutation builder.
func (i *UpdateServerCPUInput) Mutate(m *ServerCPUMutation) {
	if v := i.Serial; v != nil {
		m.SetSerial(*v)
	}
}

// SetInput applies the change-set in the UpdateServerCPUInput on the ServerCPUUpdate builder.
func (c *ServerCPUUpdate) SetInput(i UpdateServerCPUInput) *ServerCPUUpdate {
	i.Mutate(c.Mutation())
	return c
}

// SetInput applies the change-set in the UpdateServerCPUInput on the ServerCPUUpdateOne builder.
func (c *ServerCPUUpdateOne) SetInput(i UpdateServerCPUInput) *ServerCPUUpdateOne {
	i.Mutate(c.Mutation())
	return c
}

// CreateServerCPUTypeInput represents a mutation input for creating servercputypes.
type CreateServerCPUTypeInput struct {
	Vendor     string
	Model      string
	ClockSpeed string
	CoreCount  int
	CPUIDs     []gidx.PrefixedID
}

// Mutate applies the CreateServerCPUTypeInput on the ServerCPUTypeMutation builder.
func (i *CreateServerCPUTypeInput) Mutate(m *ServerCPUTypeMutation) {
	m.SetVendor(i.Vendor)
	m.SetModel(i.Model)
	m.SetClockSpeed(i.ClockSpeed)
	m.SetCoreCount(i.CoreCount)
	if v := i.CPUIDs; len(v) > 0 {
		m.AddCPUIDs(v...)
	}
}

// SetInput applies the change-set in the CreateServerCPUTypeInput on the ServerCPUTypeCreate builder.
func (c *ServerCPUTypeCreate) SetInput(i CreateServerCPUTypeInput) *ServerCPUTypeCreate {
	i.Mutate(c.Mutation())
	return c
}

// UpdateServerCPUTypeInput represents a mutation input for updating servercputypes.
type UpdateServerCPUTypeInput struct {
	Vendor       *string
	Model        *string
	ClockSpeed   *string
	CoreCount    *int
	ClearCPU     bool
	AddCPUIDs    []gidx.PrefixedID
	RemoveCPUIDs []gidx.PrefixedID
}

// Mutate applies the UpdateServerCPUTypeInput on the ServerCPUTypeMutation builder.
func (i *UpdateServerCPUTypeInput) Mutate(m *ServerCPUTypeMutation) {
	if v := i.Vendor; v != nil {
		m.SetVendor(*v)
	}
	if v := i.Model; v != nil {
		m.SetModel(*v)
	}
	if v := i.ClockSpeed; v != nil {
		m.SetClockSpeed(*v)
	}
	if v := i.CoreCount; v != nil {
		m.SetCoreCount(*v)
	}
	if i.ClearCPU {
		m.ClearCPU()
	}
	if v := i.AddCPUIDs; len(v) > 0 {
		m.AddCPUIDs(v...)
	}
	if v := i.RemoveCPUIDs; len(v) > 0 {
		m.RemoveCPUIDs(v...)
	}
}

// SetInput applies the change-set in the UpdateServerCPUTypeInput on the ServerCPUTypeUpdate builder.
func (c *ServerCPUTypeUpdate) SetInput(i UpdateServerCPUTypeInput) *ServerCPUTypeUpdate {
	i.Mutate(c.Mutation())
	return c
}

// SetInput applies the change-set in the UpdateServerCPUTypeInput on the ServerCPUTypeUpdateOne builder.
func (c *ServerCPUTypeUpdateOne) SetInput(i UpdateServerCPUTypeInput) *ServerCPUTypeUpdateOne {
	i.Mutate(c.Mutation())
	return c
}

// CreateServerChassisInput represents a mutation input for creating serverchasses.
type CreateServerChassisInput struct {
	ParentChassisID     gidx.PrefixedID
	Serial              string
	ServerID            gidx.PrefixedID
	ServerChassisTypeID gidx.PrefixedID
}

// Mutate applies the CreateServerChassisInput on the ServerChassisMutation builder.
func (i *CreateServerChassisInput) Mutate(m *ServerChassisMutation) {
	m.SetParentChassisID(i.ParentChassisID)
	m.SetSerial(i.Serial)
	m.SetServerID(i.ServerID)
	m.SetServerChassisTypeID(i.ServerChassisTypeID)
}

// SetInput applies the change-set in the CreateServerChassisInput on the ServerChassisCreate builder.
func (c *ServerChassisCreate) SetInput(i CreateServerChassisInput) *ServerChassisCreate {
	i.Mutate(c.Mutation())
	return c
}

// UpdateServerChassisInput represents a mutation input for updating serverchasses.
type UpdateServerChassisInput struct {
	Serial *string
}

// Mutate applies the UpdateServerChassisInput on the ServerChassisMutation builder.
func (i *UpdateServerChassisInput) Mutate(m *ServerChassisMutation) {
	if v := i.Serial; v != nil {
		m.SetSerial(*v)
	}
}

// SetInput applies the change-set in the UpdateServerChassisInput on the ServerChassisUpdate builder.
func (c *ServerChassisUpdate) SetInput(i UpdateServerChassisInput) *ServerChassisUpdate {
	i.Mutate(c.Mutation())
	return c
}

// SetInput applies the change-set in the UpdateServerChassisInput on the ServerChassisUpdateOne builder.
func (c *ServerChassisUpdateOne) SetInput(i UpdateServerChassisInput) *ServerChassisUpdateOne {
	i.Mutate(c.Mutation())
	return c
}

// CreateServerChassisTypeInput represents a mutation input for creating serverchassistypes.
type CreateServerChassisTypeInput struct {
	Vendor              string
	Model               string
	Height              string
	IsFullDepth         bool
	ParentChassisTypeID gidx.PrefixedID
	ChassiIDs           []gidx.PrefixedID
}

// Mutate applies the CreateServerChassisTypeInput on the ServerChassisTypeMutation builder.
func (i *CreateServerChassisTypeInput) Mutate(m *ServerChassisTypeMutation) {
	m.SetVendor(i.Vendor)
	m.SetModel(i.Model)
	m.SetHeight(i.Height)
	m.SetIsFullDepth(i.IsFullDepth)
	m.SetParentChassisTypeID(i.ParentChassisTypeID)
	if v := i.ChassiIDs; len(v) > 0 {
		m.AddChassiIDs(v...)
	}
}

// SetInput applies the change-set in the CreateServerChassisTypeInput on the ServerChassisTypeCreate builder.
func (c *ServerChassisTypeCreate) SetInput(i CreateServerChassisTypeInput) *ServerChassisTypeCreate {
	i.Mutate(c.Mutation())
	return c
}

// UpdateServerChassisTypeInput represents a mutation input for updating serverchassistypes.
type UpdateServerChassisTypeInput struct {
	Vendor          *string
	Model           *string
	Height          *string
	IsFullDepth     *bool
	ClearChassis    bool
	AddChassiIDs    []gidx.PrefixedID
	RemoveChassiIDs []gidx.PrefixedID
}

// Mutate applies the UpdateServerChassisTypeInput on the ServerChassisTypeMutation builder.
func (i *UpdateServerChassisTypeInput) Mutate(m *ServerChassisTypeMutation) {
	if v := i.Vendor; v != nil {
		m.SetVendor(*v)
	}
	if v := i.Model; v != nil {
		m.SetModel(*v)
	}
	if v := i.Height; v != nil {
		m.SetHeight(*v)
	}
	if v := i.IsFullDepth; v != nil {
		m.SetIsFullDepth(*v)
	}
	if i.ClearChassis {
		m.ClearChassis()
	}
	if v := i.AddChassiIDs; len(v) > 0 {
		m.AddChassiIDs(v...)
	}
	if v := i.RemoveChassiIDs; len(v) > 0 {
		m.RemoveChassiIDs(v...)
	}
}

// SetInput applies the change-set in the UpdateServerChassisTypeInput on the ServerChassisTypeUpdate builder.
func (c *ServerChassisTypeUpdate) SetInput(i UpdateServerChassisTypeInput) *ServerChassisTypeUpdate {
	i.Mutate(c.Mutation())
	return c
}

// SetInput applies the change-set in the UpdateServerChassisTypeInput on the ServerChassisTypeUpdateOne builder.
func (c *ServerChassisTypeUpdateOne) SetInput(i UpdateServerChassisTypeInput) *ServerChassisTypeUpdateOne {
	i.Mutate(c.Mutation())
	return c
}

// CreateServerComponentInput represents a mutation input for creating servercomponents.
type CreateServerComponentInput struct {
	Name            string
	Vendor          string
	Model           string
	Serial          string
	ComponentTypeID gidx.PrefixedID
	ServerID        gidx.PrefixedID
}

// Mutate applies the CreateServerComponentInput on the ServerComponentMutation builder.
func (i *CreateServerComponentInput) Mutate(m *ServerComponentMutation) {
	m.SetName(i.Name)
	m.SetVendor(i.Vendor)
	m.SetModel(i.Model)
	m.SetSerial(i.Serial)
	m.SetComponentTypeID(i.ComponentTypeID)
	m.SetServerID(i.ServerID)
}

// SetInput applies the change-set in the CreateServerComponentInput on the ServerComponentCreate builder.
func (c *ServerComponentCreate) SetInput(i CreateServerComponentInput) *ServerComponentCreate {
	i.Mutate(c.Mutation())
	return c
}

// UpdateServerComponentInput represents a mutation input for updating servercomponents.
type UpdateServerComponentInput struct {
	Name   *string
	Vendor *string
	Model  *string
	Serial *string
}

// Mutate applies the UpdateServerComponentInput on the ServerComponentMutation builder.
func (i *UpdateServerComponentInput) Mutate(m *ServerComponentMutation) {
	if v := i.Name; v != nil {
		m.SetName(*v)
	}
	if v := i.Vendor; v != nil {
		m.SetVendor(*v)
	}
	if v := i.Model; v != nil {
		m.SetModel(*v)
	}
	if v := i.Serial; v != nil {
		m.SetSerial(*v)
	}
}

// SetInput applies the change-set in the UpdateServerComponentInput on the ServerComponentUpdate builder.
func (c *ServerComponentUpdate) SetInput(i UpdateServerComponentInput) *ServerComponentUpdate {
	i.Mutate(c.Mutation())
	return c
}

// SetInput applies the change-set in the UpdateServerComponentInput on the ServerComponentUpdateOne builder.
func (c *ServerComponentUpdateOne) SetInput(i UpdateServerComponentInput) *ServerComponentUpdateOne {
	i.Mutate(c.Mutation())
	return c
}

// CreateServerComponentTypeInput represents a mutation input for creating servercomponenttypes.
type CreateServerComponentTypeInput struct {
	Name string
}

// Mutate applies the CreateServerComponentTypeInput on the ServerComponentTypeMutation builder.
func (i *CreateServerComponentTypeInput) Mutate(m *ServerComponentTypeMutation) {
	m.SetName(i.Name)
}

// SetInput applies the change-set in the CreateServerComponentTypeInput on the ServerComponentTypeCreate builder.
func (c *ServerComponentTypeCreate) SetInput(i CreateServerComponentTypeInput) *ServerComponentTypeCreate {
	i.Mutate(c.Mutation())
	return c
}

// UpdateServerComponentTypeInput represents a mutation input for updating servercomponenttypes.
type UpdateServerComponentTypeInput struct {
	Name *string
}

// Mutate applies the UpdateServerComponentTypeInput on the ServerComponentTypeMutation builder.
func (i *UpdateServerComponentTypeInput) Mutate(m *ServerComponentTypeMutation) {
	if v := i.Name; v != nil {
		m.SetName(*v)
	}
}

// SetInput applies the change-set in the UpdateServerComponentTypeInput on the ServerComponentTypeUpdate builder.
func (c *ServerComponentTypeUpdate) SetInput(i UpdateServerComponentTypeInput) *ServerComponentTypeUpdate {
	i.Mutate(c.Mutation())
	return c
}

// SetInput applies the change-set in the UpdateServerComponentTypeInput on the ServerComponentTypeUpdateOne builder.
func (c *ServerComponentTypeUpdateOne) SetInput(i UpdateServerComponentTypeInput) *ServerComponentTypeUpdateOne {
	i.Mutate(c.Mutation())
	return c
}

// CreateServerHardDriveInput represents a mutation input for creating serverharddrives.
type CreateServerHardDriveInput struct {
	Serial          string
	ServerID        gidx.PrefixedID
	HardDriveTypeID gidx.PrefixedID
}

// Mutate applies the CreateServerHardDriveInput on the ServerHardDriveMutation builder.
func (i *CreateServerHardDriveInput) Mutate(m *ServerHardDriveMutation) {
	m.SetSerial(i.Serial)
	m.SetServerID(i.ServerID)
	m.SetHardDriveTypeID(i.HardDriveTypeID)
}

// SetInput applies the change-set in the CreateServerHardDriveInput on the ServerHardDriveCreate builder.
func (c *ServerHardDriveCreate) SetInput(i CreateServerHardDriveInput) *ServerHardDriveCreate {
	i.Mutate(c.Mutation())
	return c
}

// UpdateServerHardDriveInput represents a mutation input for updating serverharddrives.
type UpdateServerHardDriveInput struct {
	Serial *string
}

// Mutate applies the UpdateServerHardDriveInput on the ServerHardDriveMutation builder.
func (i *UpdateServerHardDriveInput) Mutate(m *ServerHardDriveMutation) {
	if v := i.Serial; v != nil {
		m.SetSerial(*v)
	}
}

// SetInput applies the change-set in the UpdateServerHardDriveInput on the ServerHardDriveUpdate builder.
func (c *ServerHardDriveUpdate) SetInput(i UpdateServerHardDriveInput) *ServerHardDriveUpdate {
	i.Mutate(c.Mutation())
	return c
}

// SetInput applies the change-set in the UpdateServerHardDriveInput on the ServerHardDriveUpdateOne builder.
func (c *ServerHardDriveUpdateOne) SetInput(i UpdateServerHardDriveInput) *ServerHardDriveUpdateOne {
	i.Mutate(c.Mutation())
	return c
}

// CreateServerHardDriveTypeInput represents a mutation input for creating serverharddrivetypes.
type CreateServerHardDriveTypeInput struct {
	Vendor       string
	Model        string
	Speed        string
	Type         string
	Capacity     string
	HardDriveIDs []gidx.PrefixedID
}

// Mutate applies the CreateServerHardDriveTypeInput on the ServerHardDriveTypeMutation builder.
func (i *CreateServerHardDriveTypeInput) Mutate(m *ServerHardDriveTypeMutation) {
	m.SetVendor(i.Vendor)
	m.SetModel(i.Model)
	m.SetSpeed(i.Speed)
	m.SetType(i.Type)
	m.SetCapacity(i.Capacity)
	if v := i.HardDriveIDs; len(v) > 0 {
		m.AddHardDriveIDs(v...)
	}
}

// SetInput applies the change-set in the CreateServerHardDriveTypeInput on the ServerHardDriveTypeCreate builder.
func (c *ServerHardDriveTypeCreate) SetInput(i CreateServerHardDriveTypeInput) *ServerHardDriveTypeCreate {
	i.Mutate(c.Mutation())
	return c
}

// UpdateServerHardDriveTypeInput represents a mutation input for updating serverharddrivetypes.
type UpdateServerHardDriveTypeInput struct {
	Vendor             *string
	Model              *string
	Speed              *string
	Type               *string
	Capacity           *string
	ClearHardDrive     bool
	AddHardDriveIDs    []gidx.PrefixedID
	RemoveHardDriveIDs []gidx.PrefixedID
}

// Mutate applies the UpdateServerHardDriveTypeInput on the ServerHardDriveTypeMutation builder.
func (i *UpdateServerHardDriveTypeInput) Mutate(m *ServerHardDriveTypeMutation) {
	if v := i.Vendor; v != nil {
		m.SetVendor(*v)
	}
	if v := i.Model; v != nil {
		m.SetModel(*v)
	}
	if v := i.Speed; v != nil {
		m.SetSpeed(*v)
	}
	if v := i.Type; v != nil {
		m.SetType(*v)
	}
	if v := i.Capacity; v != nil {
		m.SetCapacity(*v)
	}
	if i.ClearHardDrive {
		m.ClearHardDrive()
	}
	if v := i.AddHardDriveIDs; len(v) > 0 {
		m.AddHardDriveIDs(v...)
	}
	if v := i.RemoveHardDriveIDs; len(v) > 0 {
		m.RemoveHardDriveIDs(v...)
	}
}

// SetInput applies the change-set in the UpdateServerHardDriveTypeInput on the ServerHardDriveTypeUpdate builder.
func (c *ServerHardDriveTypeUpdate) SetInput(i UpdateServerHardDriveTypeInput) *ServerHardDriveTypeUpdate {
	i.Mutate(c.Mutation())
	return c
}

// SetInput applies the change-set in the UpdateServerHardDriveTypeInput on the ServerHardDriveTypeUpdateOne builder.
func (c *ServerHardDriveTypeUpdateOne) SetInput(i UpdateServerHardDriveTypeInput) *ServerHardDriveTypeUpdateOne {
	i.Mutate(c.Mutation())
	return c
}

// CreateServerMemoryInput represents a mutation input for creating servermemories.
type CreateServerMemoryInput struct {
	Serial             string
	ServerID           gidx.PrefixedID
	ServerMemoryTypeID gidx.PrefixedID
}

// Mutate applies the CreateServerMemoryInput on the ServerMemoryMutation builder.
func (i *CreateServerMemoryInput) Mutate(m *ServerMemoryMutation) {
	m.SetSerial(i.Serial)
	m.SetServerID(i.ServerID)
	m.SetServerMemoryTypeID(i.ServerMemoryTypeID)
}

// SetInput applies the change-set in the CreateServerMemoryInput on the ServerMemoryCreate builder.
func (c *ServerMemoryCreate) SetInput(i CreateServerMemoryInput) *ServerMemoryCreate {
	i.Mutate(c.Mutation())
	return c
}

// UpdateServerMemoryInput represents a mutation input for updating servermemories.
type UpdateServerMemoryInput struct {
	Serial *string
}

// Mutate applies the UpdateServerMemoryInput on the ServerMemoryMutation builder.
func (i *UpdateServerMemoryInput) Mutate(m *ServerMemoryMutation) {
	if v := i.Serial; v != nil {
		m.SetSerial(*v)
	}
}

// SetInput applies the change-set in the UpdateServerMemoryInput on the ServerMemoryUpdate builder.
func (c *ServerMemoryUpdate) SetInput(i UpdateServerMemoryInput) *ServerMemoryUpdate {
	i.Mutate(c.Mutation())
	return c
}

// SetInput applies the change-set in the UpdateServerMemoryInput on the ServerMemoryUpdateOne builder.
func (c *ServerMemoryUpdateOne) SetInput(i UpdateServerMemoryInput) *ServerMemoryUpdateOne {
	i.Mutate(c.Mutation())
	return c
}

// CreateServerMemoryTypeInput represents a mutation input for creating servermemorytypes.
type CreateServerMemoryTypeInput struct {
	Vendor    string
	Model     string
	Speed     string
	Size      string
	MemoryIDs []gidx.PrefixedID
}

// Mutate applies the CreateServerMemoryTypeInput on the ServerMemoryTypeMutation builder.
func (i *CreateServerMemoryTypeInput) Mutate(m *ServerMemoryTypeMutation) {
	m.SetVendor(i.Vendor)
	m.SetModel(i.Model)
	m.SetSpeed(i.Speed)
	m.SetSize(i.Size)
	if v := i.MemoryIDs; len(v) > 0 {
		m.AddMemoryIDs(v...)
	}
}

// SetInput applies the change-set in the CreateServerMemoryTypeInput on the ServerMemoryTypeCreate builder.
func (c *ServerMemoryTypeCreate) SetInput(i CreateServerMemoryTypeInput) *ServerMemoryTypeCreate {
	i.Mutate(c.Mutation())
	return c
}

// UpdateServerMemoryTypeInput represents a mutation input for updating servermemorytypes.
type UpdateServerMemoryTypeInput struct {
	Vendor          *string
	Model           *string
	Speed           *string
	Size            *string
	ClearMemory     bool
	AddMemoryIDs    []gidx.PrefixedID
	RemoveMemoryIDs []gidx.PrefixedID
}

// Mutate applies the UpdateServerMemoryTypeInput on the ServerMemoryTypeMutation builder.
func (i *UpdateServerMemoryTypeInput) Mutate(m *ServerMemoryTypeMutation) {
	if v := i.Vendor; v != nil {
		m.SetVendor(*v)
	}
	if v := i.Model; v != nil {
		m.SetModel(*v)
	}
	if v := i.Speed; v != nil {
		m.SetSpeed(*v)
	}
	if v := i.Size; v != nil {
		m.SetSize(*v)
	}
	if i.ClearMemory {
		m.ClearMemory()
	}
	if v := i.AddMemoryIDs; len(v) > 0 {
		m.AddMemoryIDs(v...)
	}
	if v := i.RemoveMemoryIDs; len(v) > 0 {
		m.RemoveMemoryIDs(v...)
	}
}

// SetInput applies the change-set in the UpdateServerMemoryTypeInput on the ServerMemoryTypeUpdate builder.
func (c *ServerMemoryTypeUpdate) SetInput(i UpdateServerMemoryTypeInput) *ServerMemoryTypeUpdate {
	i.Mutate(c.Mutation())
	return c
}

// SetInput applies the change-set in the UpdateServerMemoryTypeInput on the ServerMemoryTypeUpdateOne builder.
func (c *ServerMemoryTypeUpdateOne) SetInput(i UpdateServerMemoryTypeInput) *ServerMemoryTypeUpdateOne {
	i.Mutate(c.Mutation())
	return c
}

// CreateServerMotherboardInput represents a mutation input for creating servermotherboards.
type CreateServerMotherboardInput struct {
	Serial                  string
	ServerID                gidx.PrefixedID
	ServerMotherboardTypeID gidx.PrefixedID
}

// Mutate applies the CreateServerMotherboardInput on the ServerMotherboardMutation builder.
func (i *CreateServerMotherboardInput) Mutate(m *ServerMotherboardMutation) {
	m.SetSerial(i.Serial)
	m.SetServerID(i.ServerID)
	m.SetServerMotherboardTypeID(i.ServerMotherboardTypeID)
}

// SetInput applies the change-set in the CreateServerMotherboardInput on the ServerMotherboardCreate builder.
func (c *ServerMotherboardCreate) SetInput(i CreateServerMotherboardInput) *ServerMotherboardCreate {
	i.Mutate(c.Mutation())
	return c
}

// UpdateServerMotherboardInput represents a mutation input for updating servermotherboards.
type UpdateServerMotherboardInput struct {
	Serial *string
}

// Mutate applies the UpdateServerMotherboardInput on the ServerMotherboardMutation builder.
func (i *UpdateServerMotherboardInput) Mutate(m *ServerMotherboardMutation) {
	if v := i.Serial; v != nil {
		m.SetSerial(*v)
	}
}

// SetInput applies the change-set in the UpdateServerMotherboardInput on the ServerMotherboardUpdate builder.
func (c *ServerMotherboardUpdate) SetInput(i UpdateServerMotherboardInput) *ServerMotherboardUpdate {
	i.Mutate(c.Mutation())
	return c
}

// SetInput applies the change-set in the UpdateServerMotherboardInput on the ServerMotherboardUpdateOne builder.
func (c *ServerMotherboardUpdateOne) SetInput(i UpdateServerMotherboardInput) *ServerMotherboardUpdateOne {
	i.Mutate(c.Mutation())
	return c
}

// CreateServerMotherboardTypeInput represents a mutation input for creating servermotherboardtypes.
type CreateServerMotherboardTypeInput struct {
	Vendor         string
	Model          string
	MotherboardIDs []gidx.PrefixedID
}

// Mutate applies the CreateServerMotherboardTypeInput on the ServerMotherboardTypeMutation builder.
func (i *CreateServerMotherboardTypeInput) Mutate(m *ServerMotherboardTypeMutation) {
	m.SetVendor(i.Vendor)
	m.SetModel(i.Model)
	if v := i.MotherboardIDs; len(v) > 0 {
		m.AddMotherboardIDs(v...)
	}
}

// SetInput applies the change-set in the CreateServerMotherboardTypeInput on the ServerMotherboardTypeCreate builder.
func (c *ServerMotherboardTypeCreate) SetInput(i CreateServerMotherboardTypeInput) *ServerMotherboardTypeCreate {
	i.Mutate(c.Mutation())
	return c
}

// UpdateServerMotherboardTypeInput represents a mutation input for updating servermotherboardtypes.
type UpdateServerMotherboardTypeInput struct {
	Vendor               *string
	Model                *string
	ClearMotherboard     bool
	AddMotherboardIDs    []gidx.PrefixedID
	RemoveMotherboardIDs []gidx.PrefixedID
}

// Mutate applies the UpdateServerMotherboardTypeInput on the ServerMotherboardTypeMutation builder.
func (i *UpdateServerMotherboardTypeInput) Mutate(m *ServerMotherboardTypeMutation) {
	if v := i.Vendor; v != nil {
		m.SetVendor(*v)
	}
	if v := i.Model; v != nil {
		m.SetModel(*v)
	}
	if i.ClearMotherboard {
		m.ClearMotherboard()
	}
	if v := i.AddMotherboardIDs; len(v) > 0 {
		m.AddMotherboardIDs(v...)
	}
	if v := i.RemoveMotherboardIDs; len(v) > 0 {
		m.RemoveMotherboardIDs(v...)
	}
}

// SetInput applies the change-set in the UpdateServerMotherboardTypeInput on the ServerMotherboardTypeUpdate builder.
func (c *ServerMotherboardTypeUpdate) SetInput(i UpdateServerMotherboardTypeInput) *ServerMotherboardTypeUpdate {
	i.Mutate(c.Mutation())
	return c
}

// SetInput applies the change-set in the UpdateServerMotherboardTypeInput on the ServerMotherboardTypeUpdateOne builder.
func (c *ServerMotherboardTypeUpdateOne) SetInput(i UpdateServerMotherboardTypeInput) *ServerMotherboardTypeUpdateOne {
	i.Mutate(c.Mutation())
	return c
}

// CreateServerNetworkCardInput represents a mutation input for creating servernetworkcards.
type CreateServerNetworkCardInput struct {
	Serial            string
	NetworkCardTypeID gidx.PrefixedID
	ServerID          gidx.PrefixedID
	NetworkPortIDs    []gidx.PrefixedID
}

// Mutate applies the CreateServerNetworkCardInput on the ServerNetworkCardMutation builder.
func (i *CreateServerNetworkCardInput) Mutate(m *ServerNetworkCardMutation) {
	m.SetSerial(i.Serial)
	m.SetNetworkCardTypeID(i.NetworkCardTypeID)
	m.SetServerID(i.ServerID)
	if v := i.NetworkPortIDs; len(v) > 0 {
		m.AddNetworkPortIDs(v...)
	}
}

// SetInput applies the change-set in the CreateServerNetworkCardInput on the ServerNetworkCardCreate builder.
func (c *ServerNetworkCardCreate) SetInput(i CreateServerNetworkCardInput) *ServerNetworkCardCreate {
	i.Mutate(c.Mutation())
	return c
}

// UpdateServerNetworkCardInput represents a mutation input for updating servernetworkcards.
type UpdateServerNetworkCardInput struct {
	Serial               *string
	ClearNetworkPort     bool
	AddNetworkPortIDs    []gidx.PrefixedID
	RemoveNetworkPortIDs []gidx.PrefixedID
}

// Mutate applies the UpdateServerNetworkCardInput on the ServerNetworkCardMutation builder.
func (i *UpdateServerNetworkCardInput) Mutate(m *ServerNetworkCardMutation) {
	if v := i.Serial; v != nil {
		m.SetSerial(*v)
	}
	if i.ClearNetworkPort {
		m.ClearNetworkPort()
	}
	if v := i.AddNetworkPortIDs; len(v) > 0 {
		m.AddNetworkPortIDs(v...)
	}
	if v := i.RemoveNetworkPortIDs; len(v) > 0 {
		m.RemoveNetworkPortIDs(v...)
	}
}

// SetInput applies the change-set in the UpdateServerNetworkCardInput on the ServerNetworkCardUpdate builder.
func (c *ServerNetworkCardUpdate) SetInput(i UpdateServerNetworkCardInput) *ServerNetworkCardUpdate {
	i.Mutate(c.Mutation())
	return c
}

// SetInput applies the change-set in the UpdateServerNetworkCardInput on the ServerNetworkCardUpdateOne builder.
func (c *ServerNetworkCardUpdateOne) SetInput(i UpdateServerNetworkCardInput) *ServerNetworkCardUpdateOne {
	i.Mutate(c.Mutation())
	return c
}

// CreateServerNetworkCardTypeInput represents a mutation input for creating servernetworkcardtypes.
type CreateServerNetworkCardTypeInput struct {
	Vendor         string
	Model          string
	PortCount      int
	NetworkCardIDs []gidx.PrefixedID
}

// Mutate applies the CreateServerNetworkCardTypeInput on the ServerNetworkCardTypeMutation builder.
func (i *CreateServerNetworkCardTypeInput) Mutate(m *ServerNetworkCardTypeMutation) {
	m.SetVendor(i.Vendor)
	m.SetModel(i.Model)
	m.SetPortCount(i.PortCount)
	if v := i.NetworkCardIDs; len(v) > 0 {
		m.AddNetworkCardIDs(v...)
	}
}

// SetInput applies the change-set in the CreateServerNetworkCardTypeInput on the ServerNetworkCardTypeCreate builder.
func (c *ServerNetworkCardTypeCreate) SetInput(i CreateServerNetworkCardTypeInput) *ServerNetworkCardTypeCreate {
	i.Mutate(c.Mutation())
	return c
}

// UpdateServerNetworkCardTypeInput represents a mutation input for updating servernetworkcardtypes.
type UpdateServerNetworkCardTypeInput struct {
	Vendor               *string
	Model                *string
	PortCount            *int
	ClearNetworkCard     bool
	AddNetworkCardIDs    []gidx.PrefixedID
	RemoveNetworkCardIDs []gidx.PrefixedID
}

// Mutate applies the UpdateServerNetworkCardTypeInput on the ServerNetworkCardTypeMutation builder.
func (i *UpdateServerNetworkCardTypeInput) Mutate(m *ServerNetworkCardTypeMutation) {
	if v := i.Vendor; v != nil {
		m.SetVendor(*v)
	}
	if v := i.Model; v != nil {
		m.SetModel(*v)
	}
	if v := i.PortCount; v != nil {
		m.SetPortCount(*v)
	}
	if i.ClearNetworkCard {
		m.ClearNetworkCard()
	}
	if v := i.AddNetworkCardIDs; len(v) > 0 {
		m.AddNetworkCardIDs(v...)
	}
	if v := i.RemoveNetworkCardIDs; len(v) > 0 {
		m.RemoveNetworkCardIDs(v...)
	}
}

// SetInput applies the change-set in the UpdateServerNetworkCardTypeInput on the ServerNetworkCardTypeUpdate builder.
func (c *ServerNetworkCardTypeUpdate) SetInput(i UpdateServerNetworkCardTypeInput) *ServerNetworkCardTypeUpdate {
	i.Mutate(c.Mutation())
	return c
}

// SetInput applies the change-set in the UpdateServerNetworkCardTypeInput on the ServerNetworkCardTypeUpdateOne builder.
func (c *ServerNetworkCardTypeUpdateOne) SetInput(i UpdateServerNetworkCardTypeInput) *ServerNetworkCardTypeUpdateOne {
	i.Mutate(c.Mutation())
	return c
}

// CreateServerNetworkPortInput represents a mutation input for creating servernetworkports.
type CreateServerNetworkPortInput struct {
	MACAddress    string
	NetworkCardID gidx.PrefixedID
}

// Mutate applies the CreateServerNetworkPortInput on the ServerNetworkPortMutation builder.
func (i *CreateServerNetworkPortInput) Mutate(m *ServerNetworkPortMutation) {
	m.SetMACAddress(i.MACAddress)
	m.SetNetworkCardID(i.NetworkCardID)
}

// SetInput applies the change-set in the CreateServerNetworkPortInput on the ServerNetworkPortCreate builder.
func (c *ServerNetworkPortCreate) SetInput(i CreateServerNetworkPortInput) *ServerNetworkPortCreate {
	i.Mutate(c.Mutation())
	return c
}

// UpdateServerNetworkPortInput represents a mutation input for updating servernetworkports.
type UpdateServerNetworkPortInput struct {
	MACAddress *string
}

// Mutate applies the UpdateServerNetworkPortInput on the ServerNetworkPortMutation builder.
func (i *UpdateServerNetworkPortInput) Mutate(m *ServerNetworkPortMutation) {
	if v := i.MACAddress; v != nil {
		m.SetMACAddress(*v)
	}
}

// SetInput applies the change-set in the UpdateServerNetworkPortInput on the ServerNetworkPortUpdate builder.
func (c *ServerNetworkPortUpdate) SetInput(i UpdateServerNetworkPortInput) *ServerNetworkPortUpdate {
	i.Mutate(c.Mutation())
	return c
}

// SetInput applies the change-set in the UpdateServerNetworkPortInput on the ServerNetworkPortUpdateOne builder.
func (c *ServerNetworkPortUpdateOne) SetInput(i UpdateServerNetworkPortInput) *ServerNetworkPortUpdateOne {
	i.Mutate(c.Mutation())
	return c
}

// CreateServerPowerSupplyInput represents a mutation input for creating serverpowersupplies.
type CreateServerPowerSupplyInput struct {
	Serial                  string
	ServerID                gidx.PrefixedID
	ServerPowerSupplyTypeID gidx.PrefixedID
}

// Mutate applies the CreateServerPowerSupplyInput on the ServerPowerSupplyMutation builder.
func (i *CreateServerPowerSupplyInput) Mutate(m *ServerPowerSupplyMutation) {
	m.SetSerial(i.Serial)
	m.SetServerID(i.ServerID)
	m.SetServerPowerSupplyTypeID(i.ServerPowerSupplyTypeID)
}

// SetInput applies the change-set in the CreateServerPowerSupplyInput on the ServerPowerSupplyCreate builder.
func (c *ServerPowerSupplyCreate) SetInput(i CreateServerPowerSupplyInput) *ServerPowerSupplyCreate {
	i.Mutate(c.Mutation())
	return c
}

// UpdateServerPowerSupplyInput represents a mutation input for updating serverpowersupplies.
type UpdateServerPowerSupplyInput struct {
	Serial *string
}

// Mutate applies the UpdateServerPowerSupplyInput on the ServerPowerSupplyMutation builder.
func (i *UpdateServerPowerSupplyInput) Mutate(m *ServerPowerSupplyMutation) {
	if v := i.Serial; v != nil {
		m.SetSerial(*v)
	}
}

// SetInput applies the change-set in the UpdateServerPowerSupplyInput on the ServerPowerSupplyUpdate builder.
func (c *ServerPowerSupplyUpdate) SetInput(i UpdateServerPowerSupplyInput) *ServerPowerSupplyUpdate {
	i.Mutate(c.Mutation())
	return c
}

// SetInput applies the change-set in the UpdateServerPowerSupplyInput on the ServerPowerSupplyUpdateOne builder.
func (c *ServerPowerSupplyUpdateOne) SetInput(i UpdateServerPowerSupplyInput) *ServerPowerSupplyUpdateOne {
	i.Mutate(c.Mutation())
	return c
}

// CreateServerPowerSupplyTypeInput represents a mutation input for creating serverpowersupplytypes.
type CreateServerPowerSupplyTypeInput struct {
	Vendor string
	Model  string
	Watts  string
}

// Mutate applies the CreateServerPowerSupplyTypeInput on the ServerPowerSupplyTypeMutation builder.
func (i *CreateServerPowerSupplyTypeInput) Mutate(m *ServerPowerSupplyTypeMutation) {
	m.SetVendor(i.Vendor)
	m.SetModel(i.Model)
	m.SetWatts(i.Watts)
}

// SetInput applies the change-set in the CreateServerPowerSupplyTypeInput on the ServerPowerSupplyTypeCreate builder.
func (c *ServerPowerSupplyTypeCreate) SetInput(i CreateServerPowerSupplyTypeInput) *ServerPowerSupplyTypeCreate {
	i.Mutate(c.Mutation())
	return c
}

// UpdateServerPowerSupplyTypeInput represents a mutation input for updating serverpowersupplytypes.
type UpdateServerPowerSupplyTypeInput struct {
	Vendor *string
	Model  *string
	Watts  *string
}

// Mutate applies the UpdateServerPowerSupplyTypeInput on the ServerPowerSupplyTypeMutation builder.
func (i *UpdateServerPowerSupplyTypeInput) Mutate(m *ServerPowerSupplyTypeMutation) {
	if v := i.Vendor; v != nil {
		m.SetVendor(*v)
	}
	if v := i.Model; v != nil {
		m.SetModel(*v)
	}
	if v := i.Watts; v != nil {
		m.SetWatts(*v)
	}
}

// SetInput applies the change-set in the UpdateServerPowerSupplyTypeInput on the ServerPowerSupplyTypeUpdate builder.
func (c *ServerPowerSupplyTypeUpdate) SetInput(i UpdateServerPowerSupplyTypeInput) *ServerPowerSupplyTypeUpdate {
	i.Mutate(c.Mutation())
	return c
}

// SetInput applies the change-set in the UpdateServerPowerSupplyTypeInput on the ServerPowerSupplyTypeUpdateOne builder.
func (c *ServerPowerSupplyTypeUpdateOne) SetInput(i UpdateServerPowerSupplyTypeInput) *ServerPowerSupplyTypeUpdateOne {
	i.Mutate(c.Mutation())
	return c
}

// CreateServerTypeInput represents a mutation input for creating servertypes.
type CreateServerTypeInput struct {
	Name    string
	OwnerID gidx.PrefixedID
}

// Mutate applies the CreateServerTypeInput on the ServerTypeMutation builder.
func (i *CreateServerTypeInput) Mutate(m *ServerTypeMutation) {
	m.SetName(i.Name)
	m.SetOwnerID(i.OwnerID)
}

// SetInput applies the change-set in the CreateServerTypeInput on the ServerTypeCreate builder.
func (c *ServerTypeCreate) SetInput(i CreateServerTypeInput) *ServerTypeCreate {
	i.Mutate(c.Mutation())
	return c
}

// UpdateServerTypeInput represents a mutation input for updating servertypes.
type UpdateServerTypeInput struct {
	Name *string
}

// Mutate applies the UpdateServerTypeInput on the ServerTypeMutation builder.
func (i *UpdateServerTypeInput) Mutate(m *ServerTypeMutation) {
	if v := i.Name; v != nil {
		m.SetName(*v)
	}
}

// SetInput applies the change-set in the UpdateServerTypeInput on the ServerTypeUpdate builder.
func (c *ServerTypeUpdate) SetInput(i UpdateServerTypeInput) *ServerTypeUpdate {
	i.Mutate(c.Mutation())
	return c
}

// SetInput applies the change-set in the UpdateServerTypeInput on the ServerTypeUpdateOne builder.
func (c *ServerTypeUpdateOne) SetInput(i UpdateServerTypeInput) *ServerTypeUpdateOne {
	i.Mutate(c.Mutation())
	return c
}
