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

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// ProvidersColumns holds the columns for the "providers" table.
	ProvidersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "name", Type: field.TypeString, Size: 2147483647},
		{Name: "owner_id", Type: field.TypeString},
	}
	// ProvidersTable holds the schema information for the "providers" table.
	ProvidersTable = &schema.Table{
		Name:       "providers",
		Columns:    ProvidersColumns,
		PrimaryKey: []*schema.Column{ProvidersColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "provider_created_at",
				Unique:  false,
				Columns: []*schema.Column{ProvidersColumns[1]},
			},
			{
				Name:    "provider_updated_at",
				Unique:  false,
				Columns: []*schema.Column{ProvidersColumns[2]},
			},
			{
				Name:    "provider_owner_id",
				Unique:  false,
				Columns: []*schema.Column{ProvidersColumns[4]},
			},
		},
	}
	// ServersColumns holds the columns for the "servers" table.
	ServersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "name", Type: field.TypeString, Size: 2147483647},
		{Name: "description", Type: field.TypeString, Nullable: true, Size: 2147483647},
		{Name: "owner_id", Type: field.TypeString},
		{Name: "location_id", Type: field.TypeString},
		{Name: "provider_id", Type: field.TypeString},
		{Name: "server_type_id", Type: field.TypeString},
	}
	// ServersTable holds the schema information for the "servers" table.
	ServersTable = &schema.Table{
		Name:       "servers",
		Columns:    ServersColumns,
		PrimaryKey: []*schema.Column{ServersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "servers_providers_provider",
				Columns:    []*schema.Column{ServersColumns[7]},
				RefColumns: []*schema.Column{ProvidersColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "servers_server_types_server_type",
				Columns:    []*schema.Column{ServersColumns[8]},
				RefColumns: []*schema.Column{ServerTypesColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "server_created_at",
				Unique:  false,
				Columns: []*schema.Column{ServersColumns[1]},
			},
			{
				Name:    "server_updated_at",
				Unique:  false,
				Columns: []*schema.Column{ServersColumns[2]},
			},
			{
				Name:    "server_owner_id",
				Unique:  false,
				Columns: []*schema.Column{ServersColumns[5]},
			},
			{
				Name:    "server_location_id",
				Unique:  false,
				Columns: []*schema.Column{ServersColumns[6]},
			},
			{
				Name:    "server_provider_id",
				Unique:  false,
				Columns: []*schema.Column{ServersColumns[7]},
			},
			{
				Name:    "server_server_type_id",
				Unique:  false,
				Columns: []*schema.Column{ServersColumns[8]},
			},
		},
	}
	// ServerAttributesColumns holds the columns for the "server_attributes" table.
	ServerAttributesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "name", Type: field.TypeString, Size: 2147483647},
		{Name: "server_id", Type: field.TypeString},
	}
	// ServerAttributesTable holds the schema information for the "server_attributes" table.
	ServerAttributesTable = &schema.Table{
		Name:       "server_attributes",
		Columns:    ServerAttributesColumns,
		PrimaryKey: []*schema.Column{ServerAttributesColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "serverattribute_created_at",
				Unique:  false,
				Columns: []*schema.Column{ServerAttributesColumns[1]},
			},
			{
				Name:    "serverattribute_updated_at",
				Unique:  false,
				Columns: []*schema.Column{ServerAttributesColumns[2]},
			},
		},
	}
	// ServerComponentsColumns holds the columns for the "server_components" table.
	ServerComponentsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "name", Type: field.TypeString, Size: 2147483647},
		{Name: "vendor", Type: field.TypeString, Size: 2147483647},
		{Name: "model", Type: field.TypeString, Size: 2147483647},
		{Name: "serial", Type: field.TypeString, Size: 2147483647},
		{Name: "component_type_id", Type: field.TypeString},
		{Name: "server_id", Type: field.TypeString},
	}
	// ServerComponentsTable holds the schema information for the "server_components" table.
	ServerComponentsTable = &schema.Table{
		Name:       "server_components",
		Columns:    ServerComponentsColumns,
		PrimaryKey: []*schema.Column{ServerComponentsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "server_components_server_component_types_component_type",
				Columns:    []*schema.Column{ServerComponentsColumns[7]},
				RefColumns: []*schema.Column{ServerComponentTypesColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "server_components_servers_server",
				Columns:    []*schema.Column{ServerComponentsColumns[8]},
				RefColumns: []*schema.Column{ServersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "servercomponent_created_at",
				Unique:  false,
				Columns: []*schema.Column{ServerComponentsColumns[1]},
			},
			{
				Name:    "servercomponent_updated_at",
				Unique:  false,
				Columns: []*schema.Column{ServerComponentsColumns[2]},
			},
			{
				Name:    "servercomponent_server_id",
				Unique:  false,
				Columns: []*schema.Column{ServerComponentsColumns[8]},
			},
			{
				Name:    "servercomponent_component_type_id",
				Unique:  false,
				Columns: []*schema.Column{ServerComponentsColumns[7]},
			},
		},
	}
	// ServerComponentTypesColumns holds the columns for the "server_component_types" table.
	ServerComponentTypesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "name", Type: field.TypeString, Size: 2147483647},
	}
	// ServerComponentTypesTable holds the schema information for the "server_component_types" table.
	ServerComponentTypesTable = &schema.Table{
		Name:       "server_component_types",
		Columns:    ServerComponentTypesColumns,
		PrimaryKey: []*schema.Column{ServerComponentTypesColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "servercomponenttype_created_at",
				Unique:  false,
				Columns: []*schema.Column{ServerComponentTypesColumns[1]},
			},
			{
				Name:    "servercomponenttype_updated_at",
				Unique:  false,
				Columns: []*schema.Column{ServerComponentTypesColumns[2]},
			},
		},
	}
	// ServerTypesColumns holds the columns for the "server_types" table.
	ServerTypesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "name", Type: field.TypeString, Size: 2147483647},
		{Name: "owner_id", Type: field.TypeString},
	}
	// ServerTypesTable holds the schema information for the "server_types" table.
	ServerTypesTable = &schema.Table{
		Name:       "server_types",
		Columns:    ServerTypesColumns,
		PrimaryKey: []*schema.Column{ServerTypesColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "servertype_created_at",
				Unique:  false,
				Columns: []*schema.Column{ServerTypesColumns[1]},
			},
			{
				Name:    "servertype_updated_at",
				Unique:  false,
				Columns: []*schema.Column{ServerTypesColumns[2]},
			},
			{
				Name:    "servertype_owner_id",
				Unique:  false,
				Columns: []*schema.Column{ServerTypesColumns[4]},
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		ProvidersTable,
		ServersTable,
		ServerAttributesTable,
		ServerComponentsTable,
		ServerComponentTypesTable,
		ServerTypesTable,
	}
)

func init() {
	ServersTable.ForeignKeys[0].RefTable = ProvidersTable
	ServersTable.ForeignKeys[1].RefTable = ServerTypesTable
	ServerComponentsTable.ForeignKeys[0].RefTable = ServerComponentTypesTable
	ServerComponentsTable.ForeignKeys[1].RefTable = ServersTable
}
