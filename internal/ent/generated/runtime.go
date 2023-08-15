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
	"time"

	"go.infratographer.com/server-api/internal/ent/generated/provider"
	"go.infratographer.com/server-api/internal/ent/generated/server"
	"go.infratographer.com/server-api/internal/ent/generated/serverattribute"
	"go.infratographer.com/server-api/internal/ent/generated/servercomponent"
	"go.infratographer.com/server-api/internal/ent/generated/servercomponenttype"
	"go.infratographer.com/server-api/internal/ent/generated/servertype"
	"go.infratographer.com/server-api/internal/ent/schema"
	"go.infratographer.com/x/gidx"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	providerMixin := schema.Provider{}.Mixin()
	providerMixinFields0 := providerMixin[0].Fields()
	_ = providerMixinFields0
	providerFields := schema.Provider{}.Fields()
	_ = providerFields
	// providerDescCreatedAt is the schema descriptor for created_at field.
	providerDescCreatedAt := providerMixinFields0[0].Descriptor()
	// provider.DefaultCreatedAt holds the default value on creation for the created_at field.
	provider.DefaultCreatedAt = providerDescCreatedAt.Default.(func() time.Time)
	// providerDescUpdatedAt is the schema descriptor for updated_at field.
	providerDescUpdatedAt := providerMixinFields0[1].Descriptor()
	// provider.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	provider.DefaultUpdatedAt = providerDescUpdatedAt.Default.(func() time.Time)
	// provider.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	provider.UpdateDefaultUpdatedAt = providerDescUpdatedAt.UpdateDefault.(func() time.Time)
	// providerDescName is the schema descriptor for name field.
	providerDescName := providerFields[1].Descriptor()
	// provider.NameValidator is a validator for the "name" field. It is called by the builders before save.
	provider.NameValidator = providerDescName.Validators[0].(func(string) error)
	// providerDescID is the schema descriptor for id field.
	providerDescID := providerFields[0].Descriptor()
	// provider.DefaultID holds the default value on creation for the id field.
	provider.DefaultID = providerDescID.Default.(func() gidx.PrefixedID)
	serverMixin := schema.Server{}.Mixin()
	serverMixinFields0 := serverMixin[0].Fields()
	_ = serverMixinFields0
	serverFields := schema.Server{}.Fields()
	_ = serverFields
	// serverDescCreatedAt is the schema descriptor for created_at field.
	serverDescCreatedAt := serverMixinFields0[0].Descriptor()
	// server.DefaultCreatedAt holds the default value on creation for the created_at field.
	server.DefaultCreatedAt = serverDescCreatedAt.Default.(func() time.Time)
	// serverDescUpdatedAt is the schema descriptor for updated_at field.
	serverDescUpdatedAt := serverMixinFields0[1].Descriptor()
	// server.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	server.DefaultUpdatedAt = serverDescUpdatedAt.Default.(func() time.Time)
	// server.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	server.UpdateDefaultUpdatedAt = serverDescUpdatedAt.UpdateDefault.(func() time.Time)
	// serverDescName is the schema descriptor for name field.
	serverDescName := serverFields[1].Descriptor()
	// server.NameValidator is a validator for the "name" field. It is called by the builders before save.
	server.NameValidator = serverDescName.Validators[0].(func(string) error)
	// serverDescLocationID is the schema descriptor for location_id field.
	serverDescLocationID := serverFields[4].Descriptor()
	// server.LocationIDValidator is a validator for the "location_id" field. It is called by the builders before save.
	server.LocationIDValidator = serverDescLocationID.Validators[0].(func(string) error)
	// serverDescProviderID is the schema descriptor for provider_id field.
	serverDescProviderID := serverFields[5].Descriptor()
	// server.ProviderIDValidator is a validator for the "provider_id" field. It is called by the builders before save.
	server.ProviderIDValidator = serverDescProviderID.Validators[0].(func(string) error)
	// serverDescID is the schema descriptor for id field.
	serverDescID := serverFields[0].Descriptor()
	// server.DefaultID holds the default value on creation for the id field.
	server.DefaultID = serverDescID.Default.(func() gidx.PrefixedID)
	serverattributeMixin := schema.ServerAttribute{}.Mixin()
	serverattributeMixinFields0 := serverattributeMixin[0].Fields()
	_ = serverattributeMixinFields0
	serverattributeFields := schema.ServerAttribute{}.Fields()
	_ = serverattributeFields
	// serverattributeDescCreatedAt is the schema descriptor for created_at field.
	serverattributeDescCreatedAt := serverattributeMixinFields0[0].Descriptor()
	// serverattribute.DefaultCreatedAt holds the default value on creation for the created_at field.
	serverattribute.DefaultCreatedAt = serverattributeDescCreatedAt.Default.(func() time.Time)
	// serverattributeDescUpdatedAt is the schema descriptor for updated_at field.
	serverattributeDescUpdatedAt := serverattributeMixinFields0[1].Descriptor()
	// serverattribute.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	serverattribute.DefaultUpdatedAt = serverattributeDescUpdatedAt.Default.(func() time.Time)
	// serverattribute.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	serverattribute.UpdateDefaultUpdatedAt = serverattributeDescUpdatedAt.UpdateDefault.(func() time.Time)
	// serverattributeDescName is the schema descriptor for name field.
	serverattributeDescName := serverattributeFields[1].Descriptor()
	// serverattribute.NameValidator is a validator for the "name" field. It is called by the builders before save.
	serverattribute.NameValidator = serverattributeDescName.Validators[0].(func(string) error)
	// serverattributeDescServerID is the schema descriptor for server_id field.
	serverattributeDescServerID := serverattributeFields[2].Descriptor()
	// serverattribute.ServerIDValidator is a validator for the "server_id" field. It is called by the builders before save.
	serverattribute.ServerIDValidator = serverattributeDescServerID.Validators[0].(func(string) error)
	// serverattributeDescID is the schema descriptor for id field.
	serverattributeDescID := serverattributeFields[0].Descriptor()
	// serverattribute.DefaultID holds the default value on creation for the id field.
	serverattribute.DefaultID = serverattributeDescID.Default.(func() gidx.PrefixedID)
	servercomponentMixin := schema.ServerComponent{}.Mixin()
	servercomponentMixinFields0 := servercomponentMixin[0].Fields()
	_ = servercomponentMixinFields0
	servercomponentFields := schema.ServerComponent{}.Fields()
	_ = servercomponentFields
	// servercomponentDescCreatedAt is the schema descriptor for created_at field.
	servercomponentDescCreatedAt := servercomponentMixinFields0[0].Descriptor()
	// servercomponent.DefaultCreatedAt holds the default value on creation for the created_at field.
	servercomponent.DefaultCreatedAt = servercomponentDescCreatedAt.Default.(func() time.Time)
	// servercomponentDescUpdatedAt is the schema descriptor for updated_at field.
	servercomponentDescUpdatedAt := servercomponentMixinFields0[1].Descriptor()
	// servercomponent.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	servercomponent.DefaultUpdatedAt = servercomponentDescUpdatedAt.Default.(func() time.Time)
	// servercomponent.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	servercomponent.UpdateDefaultUpdatedAt = servercomponentDescUpdatedAt.UpdateDefault.(func() time.Time)
	// servercomponentDescName is the schema descriptor for name field.
	servercomponentDescName := servercomponentFields[1].Descriptor()
	// servercomponent.NameValidator is a validator for the "name" field. It is called by the builders before save.
	servercomponent.NameValidator = servercomponentDescName.Validators[0].(func(string) error)
	// servercomponentDescVendor is the schema descriptor for vendor field.
	servercomponentDescVendor := servercomponentFields[2].Descriptor()
	// servercomponent.VendorValidator is a validator for the "vendor" field. It is called by the builders before save.
	servercomponent.VendorValidator = servercomponentDescVendor.Validators[0].(func(string) error)
	// servercomponentDescModel is the schema descriptor for model field.
	servercomponentDescModel := servercomponentFields[3].Descriptor()
	// servercomponent.ModelValidator is a validator for the "model" field. It is called by the builders before save.
	servercomponent.ModelValidator = servercomponentDescModel.Validators[0].(func(string) error)
	// servercomponentDescSerial is the schema descriptor for serial field.
	servercomponentDescSerial := servercomponentFields[4].Descriptor()
	// servercomponent.SerialValidator is a validator for the "serial" field. It is called by the builders before save.
	servercomponent.SerialValidator = servercomponentDescSerial.Validators[0].(func(string) error)
	// servercomponentDescServerID is the schema descriptor for server_id field.
	servercomponentDescServerID := servercomponentFields[5].Descriptor()
	// servercomponent.ServerIDValidator is a validator for the "server_id" field. It is called by the builders before save.
	servercomponent.ServerIDValidator = servercomponentDescServerID.Validators[0].(func(string) error)
	// servercomponentDescComponentTypeID is the schema descriptor for component_type_id field.
	servercomponentDescComponentTypeID := servercomponentFields[6].Descriptor()
	// servercomponent.ComponentTypeIDValidator is a validator for the "component_type_id" field. It is called by the builders before save.
	servercomponent.ComponentTypeIDValidator = servercomponentDescComponentTypeID.Validators[0].(func(string) error)
	// servercomponentDescID is the schema descriptor for id field.
	servercomponentDescID := servercomponentFields[0].Descriptor()
	// servercomponent.DefaultID holds the default value on creation for the id field.
	servercomponent.DefaultID = servercomponentDescID.Default.(func() gidx.PrefixedID)
	servercomponenttypeMixin := schema.ServerComponentType{}.Mixin()
	servercomponenttypeMixinFields0 := servercomponenttypeMixin[0].Fields()
	_ = servercomponenttypeMixinFields0
	servercomponenttypeFields := schema.ServerComponentType{}.Fields()
	_ = servercomponenttypeFields
	// servercomponenttypeDescCreatedAt is the schema descriptor for created_at field.
	servercomponenttypeDescCreatedAt := servercomponenttypeMixinFields0[0].Descriptor()
	// servercomponenttype.DefaultCreatedAt holds the default value on creation for the created_at field.
	servercomponenttype.DefaultCreatedAt = servercomponenttypeDescCreatedAt.Default.(func() time.Time)
	// servercomponenttypeDescUpdatedAt is the schema descriptor for updated_at field.
	servercomponenttypeDescUpdatedAt := servercomponenttypeMixinFields0[1].Descriptor()
	// servercomponenttype.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	servercomponenttype.DefaultUpdatedAt = servercomponenttypeDescUpdatedAt.Default.(func() time.Time)
	// servercomponenttype.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	servercomponenttype.UpdateDefaultUpdatedAt = servercomponenttypeDescUpdatedAt.UpdateDefault.(func() time.Time)
	// servercomponenttypeDescName is the schema descriptor for name field.
	servercomponenttypeDescName := servercomponenttypeFields[1].Descriptor()
	// servercomponenttype.NameValidator is a validator for the "name" field. It is called by the builders before save.
	servercomponenttype.NameValidator = servercomponenttypeDescName.Validators[0].(func(string) error)
	// servercomponenttypeDescID is the schema descriptor for id field.
	servercomponenttypeDescID := servercomponenttypeFields[0].Descriptor()
	// servercomponenttype.DefaultID holds the default value on creation for the id field.
	servercomponenttype.DefaultID = servercomponenttypeDescID.Default.(func() gidx.PrefixedID)
	servertypeMixin := schema.ServerType{}.Mixin()
	servertypeMixinFields0 := servertypeMixin[0].Fields()
	_ = servertypeMixinFields0
	servertypeFields := schema.ServerType{}.Fields()
	_ = servertypeFields
	// servertypeDescCreatedAt is the schema descriptor for created_at field.
	servertypeDescCreatedAt := servertypeMixinFields0[0].Descriptor()
	// servertype.DefaultCreatedAt holds the default value on creation for the created_at field.
	servertype.DefaultCreatedAt = servertypeDescCreatedAt.Default.(func() time.Time)
	// servertypeDescUpdatedAt is the schema descriptor for updated_at field.
	servertypeDescUpdatedAt := servertypeMixinFields0[1].Descriptor()
	// servertype.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	servertype.DefaultUpdatedAt = servertypeDescUpdatedAt.Default.(func() time.Time)
	// servertype.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	servertype.UpdateDefaultUpdatedAt = servertypeDescUpdatedAt.UpdateDefault.(func() time.Time)
	// servertypeDescName is the schema descriptor for name field.
	servertypeDescName := servertypeFields[1].Descriptor()
	// servertype.NameValidator is a validator for the "name" field. It is called by the builders before save.
	servertype.NameValidator = servertypeDescName.Validators[0].(func(string) error)
	// servertypeDescID is the schema descriptor for id field.
	servertypeDescID := servertypeFields[0].Descriptor()
	// servertype.DefaultID holds the default value on creation for the id field.
	servertype.DefaultID = servertypeDescID.Default.(func() gidx.PrefixedID)
}
