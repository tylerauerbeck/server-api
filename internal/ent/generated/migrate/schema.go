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
		{Name: "resource_provider_id", Type: field.TypeString},
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
				Name:    "provider_resource_provider_id",
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
	// ServerCpUsColumns holds the columns for the "server_cp_us" table.
	ServerCpUsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "serial", Type: field.TypeString, Size: 2147483647},
		{Name: "server_id", Type: field.TypeString},
		{Name: "server_cpu_type_id", Type: field.TypeString},
	}
	// ServerCpUsTable holds the schema information for the "server_cp_us" table.
	ServerCpUsTable = &schema.Table{
		Name:       "server_cp_us",
		Columns:    ServerCpUsColumns,
		PrimaryKey: []*schema.Column{ServerCpUsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "server_cp_us_servers_server",
				Columns:    []*schema.Column{ServerCpUsColumns[4]},
				RefColumns: []*schema.Column{ServersColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "server_cp_us_server_cpu_types_server_cpu_type",
				Columns:    []*schema.Column{ServerCpUsColumns[5]},
				RefColumns: []*schema.Column{ServerCPUTypesColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "servercpu_created_at",
				Unique:  false,
				Columns: []*schema.Column{ServerCpUsColumns[1]},
			},
			{
				Name:    "servercpu_updated_at",
				Unique:  false,
				Columns: []*schema.Column{ServerCpUsColumns[2]},
			},
			{
				Name:    "servercpu_server_cpu_type_id",
				Unique:  false,
				Columns: []*schema.Column{ServerCpUsColumns[5]},
			},
			{
				Name:    "servercpu_server_id",
				Unique:  false,
				Columns: []*schema.Column{ServerCpUsColumns[4]},
			},
		},
	}
	// ServerCPUTypesColumns holds the columns for the "server_cpu_types" table.
	ServerCPUTypesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "vendor", Type: field.TypeString, Size: 2147483647},
		{Name: "model", Type: field.TypeString, Size: 2147483647},
		{Name: "clock_speed", Type: field.TypeString, Size: 2147483647},
		{Name: "core_count", Type: field.TypeInt},
	}
	// ServerCPUTypesTable holds the schema information for the "server_cpu_types" table.
	ServerCPUTypesTable = &schema.Table{
		Name:       "server_cpu_types",
		Columns:    ServerCPUTypesColumns,
		PrimaryKey: []*schema.Column{ServerCPUTypesColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "servercputype_created_at",
				Unique:  false,
				Columns: []*schema.Column{ServerCPUTypesColumns[1]},
			},
			{
				Name:    "servercputype_updated_at",
				Unique:  false,
				Columns: []*schema.Column{ServerCPUTypesColumns[2]},
			},
		},
	}
	// ServerChassesColumns holds the columns for the "server_chasses" table.
	ServerChassesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "parent_chassis_id", Type: field.TypeString, Size: 2147483647},
		{Name: "serial", Type: field.TypeString, Size: 2147483647},
		{Name: "server_id", Type: field.TypeString},
		{Name: "server_chassis_type_id", Type: field.TypeString},
	}
	// ServerChassesTable holds the schema information for the "server_chasses" table.
	ServerChassesTable = &schema.Table{
		Name:       "server_chasses",
		Columns:    ServerChassesColumns,
		PrimaryKey: []*schema.Column{ServerChassesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "server_chasses_servers_server",
				Columns:    []*schema.Column{ServerChassesColumns[5]},
				RefColumns: []*schema.Column{ServersColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "server_chasses_server_chassis_types_server_chassis_type",
				Columns:    []*schema.Column{ServerChassesColumns[6]},
				RefColumns: []*schema.Column{ServerChassisTypesColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "serverchassis_created_at",
				Unique:  false,
				Columns: []*schema.Column{ServerChassesColumns[1]},
			},
			{
				Name:    "serverchassis_updated_at",
				Unique:  false,
				Columns: []*schema.Column{ServerChassesColumns[2]},
			},
			{
				Name:    "serverchassis_server_chassis_type_id",
				Unique:  false,
				Columns: []*schema.Column{ServerChassesColumns[6]},
			},
			{
				Name:    "serverchassis_server_id",
				Unique:  false,
				Columns: []*schema.Column{ServerChassesColumns[5]},
			},
			{
				Name:    "serverchassis_parent_chassis_id",
				Unique:  false,
				Columns: []*schema.Column{ServerChassesColumns[3]},
			},
		},
	}
	// ServerChassisTypesColumns holds the columns for the "server_chassis_types" table.
	ServerChassisTypesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "vendor", Type: field.TypeString, Size: 2147483647},
		{Name: "model", Type: field.TypeString, Size: 2147483647},
		{Name: "height", Type: field.TypeString, Size: 2147483647},
		{Name: "is_full_depth", Type: field.TypeBool},
		{Name: "parent_chassis_type_id", Type: field.TypeString, Size: 2147483647},
	}
	// ServerChassisTypesTable holds the schema information for the "server_chassis_types" table.
	ServerChassisTypesTable = &schema.Table{
		Name:       "server_chassis_types",
		Columns:    ServerChassisTypesColumns,
		PrimaryKey: []*schema.Column{ServerChassisTypesColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "serverchassistype_created_at",
				Unique:  false,
				Columns: []*schema.Column{ServerChassisTypesColumns[1]},
			},
			{
				Name:    "serverchassistype_updated_at",
				Unique:  false,
				Columns: []*schema.Column{ServerChassisTypesColumns[2]},
			},
			{
				Name:    "serverchassistype_parent_chassis_type_id",
				Unique:  false,
				Columns: []*schema.Column{ServerChassisTypesColumns[7]},
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
	// ServerHardDrivesColumns holds the columns for the "server_hard_drives" table.
	ServerHardDrivesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "serial", Type: field.TypeString, Size: 2147483647},
		{Name: "server_id", Type: field.TypeString},
		{Name: "server_hard_drive_type_id", Type: field.TypeString},
	}
	// ServerHardDrivesTable holds the schema information for the "server_hard_drives" table.
	ServerHardDrivesTable = &schema.Table{
		Name:       "server_hard_drives",
		Columns:    ServerHardDrivesColumns,
		PrimaryKey: []*schema.Column{ServerHardDrivesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "server_hard_drives_servers_server",
				Columns:    []*schema.Column{ServerHardDrivesColumns[4]},
				RefColumns: []*schema.Column{ServersColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "server_hard_drives_server_hard_drive_types_hard_drive_type",
				Columns:    []*schema.Column{ServerHardDrivesColumns[5]},
				RefColumns: []*schema.Column{ServerHardDriveTypesColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "serverharddrive_created_at",
				Unique:  false,
				Columns: []*schema.Column{ServerHardDrivesColumns[1]},
			},
			{
				Name:    "serverharddrive_updated_at",
				Unique:  false,
				Columns: []*schema.Column{ServerHardDrivesColumns[2]},
			},
			{
				Name:    "serverharddrive_server_id",
				Unique:  false,
				Columns: []*schema.Column{ServerHardDrivesColumns[4]},
			},
			{
				Name:    "serverharddrive_server_hard_drive_type_id",
				Unique:  false,
				Columns: []*schema.Column{ServerHardDrivesColumns[5]},
			},
		},
	}
	// ServerHardDriveTypesColumns holds the columns for the "server_hard_drive_types" table.
	ServerHardDriveTypesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "vendor", Type: field.TypeString, Size: 2147483647},
		{Name: "model", Type: field.TypeString, Size: 2147483647},
		{Name: "speed", Type: field.TypeString, Size: 2147483647},
		{Name: "type", Type: field.TypeString, Size: 2147483647},
		{Name: "capacity", Type: field.TypeString, Size: 2147483647},
	}
	// ServerHardDriveTypesTable holds the schema information for the "server_hard_drive_types" table.
	ServerHardDriveTypesTable = &schema.Table{
		Name:       "server_hard_drive_types",
		Columns:    ServerHardDriveTypesColumns,
		PrimaryKey: []*schema.Column{ServerHardDriveTypesColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "serverharddrivetype_created_at",
				Unique:  false,
				Columns: []*schema.Column{ServerHardDriveTypesColumns[1]},
			},
			{
				Name:    "serverharddrivetype_updated_at",
				Unique:  false,
				Columns: []*schema.Column{ServerHardDriveTypesColumns[2]},
			},
		},
	}
	// ServerMemoriesColumns holds the columns for the "server_memories" table.
	ServerMemoriesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "serial", Type: field.TypeString, Size: 2147483647},
		{Name: "server_id", Type: field.TypeString},
		{Name: "server_memory_type_id", Type: field.TypeString},
	}
	// ServerMemoriesTable holds the schema information for the "server_memories" table.
	ServerMemoriesTable = &schema.Table{
		Name:       "server_memories",
		Columns:    ServerMemoriesColumns,
		PrimaryKey: []*schema.Column{ServerMemoriesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "server_memories_servers_server",
				Columns:    []*schema.Column{ServerMemoriesColumns[4]},
				RefColumns: []*schema.Column{ServersColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "server_memories_server_memory_types_server_memory_type",
				Columns:    []*schema.Column{ServerMemoriesColumns[5]},
				RefColumns: []*schema.Column{ServerMemoryTypesColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "servermemory_created_at",
				Unique:  false,
				Columns: []*schema.Column{ServerMemoriesColumns[1]},
			},
			{
				Name:    "servermemory_updated_at",
				Unique:  false,
				Columns: []*schema.Column{ServerMemoriesColumns[2]},
			},
			{
				Name:    "servermemory_server_id",
				Unique:  false,
				Columns: []*schema.Column{ServerMemoriesColumns[4]},
			},
			{
				Name:    "servermemory_server_memory_type_id",
				Unique:  false,
				Columns: []*schema.Column{ServerMemoriesColumns[5]},
			},
		},
	}
	// ServerMemoryTypesColumns holds the columns for the "server_memory_types" table.
	ServerMemoryTypesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "vendor", Type: field.TypeString, Size: 2147483647},
		{Name: "model", Type: field.TypeString, Size: 2147483647},
		{Name: "speed", Type: field.TypeString, Size: 2147483647},
		{Name: "size", Type: field.TypeString, Size: 2147483647},
	}
	// ServerMemoryTypesTable holds the schema information for the "server_memory_types" table.
	ServerMemoryTypesTable = &schema.Table{
		Name:       "server_memory_types",
		Columns:    ServerMemoryTypesColumns,
		PrimaryKey: []*schema.Column{ServerMemoryTypesColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "servermemorytype_created_at",
				Unique:  false,
				Columns: []*schema.Column{ServerMemoryTypesColumns[1]},
			},
			{
				Name:    "servermemorytype_updated_at",
				Unique:  false,
				Columns: []*schema.Column{ServerMemoryTypesColumns[2]},
			},
		},
	}
	// ServerMotherboardsColumns holds the columns for the "server_motherboards" table.
	ServerMotherboardsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "serial", Type: field.TypeString, Size: 2147483647},
		{Name: "server_id", Type: field.TypeString},
		{Name: "server_motherboard_type_id", Type: field.TypeString},
	}
	// ServerMotherboardsTable holds the schema information for the "server_motherboards" table.
	ServerMotherboardsTable = &schema.Table{
		Name:       "server_motherboards",
		Columns:    ServerMotherboardsColumns,
		PrimaryKey: []*schema.Column{ServerMotherboardsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "server_motherboards_servers_server",
				Columns:    []*schema.Column{ServerMotherboardsColumns[4]},
				RefColumns: []*schema.Column{ServersColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "server_motherboards_server_motherboard_types_server_motherboard_type",
				Columns:    []*schema.Column{ServerMotherboardsColumns[5]},
				RefColumns: []*schema.Column{ServerMotherboardTypesColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "servermotherboard_created_at",
				Unique:  false,
				Columns: []*schema.Column{ServerMotherboardsColumns[1]},
			},
			{
				Name:    "servermotherboard_updated_at",
				Unique:  false,
				Columns: []*schema.Column{ServerMotherboardsColumns[2]},
			},
		},
	}
	// ServerMotherboardTypesColumns holds the columns for the "server_motherboard_types" table.
	ServerMotherboardTypesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "vendor", Type: field.TypeString, Size: 2147483647},
		{Name: "model", Type: field.TypeString, Size: 2147483647},
	}
	// ServerMotherboardTypesTable holds the schema information for the "server_motherboard_types" table.
	ServerMotherboardTypesTable = &schema.Table{
		Name:       "server_motherboard_types",
		Columns:    ServerMotherboardTypesColumns,
		PrimaryKey: []*schema.Column{ServerMotherboardTypesColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "servermotherboardtype_created_at",
				Unique:  false,
				Columns: []*schema.Column{ServerMotherboardTypesColumns[1]},
			},
			{
				Name:    "servermotherboardtype_updated_at",
				Unique:  false,
				Columns: []*schema.Column{ServerMotherboardTypesColumns[2]},
			},
		},
	}
	// ServerPowerSuppliesColumns holds the columns for the "server_power_supplies" table.
	ServerPowerSuppliesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "serial", Type: field.TypeString, Size: 2147483647},
		{Name: "server_id", Type: field.TypeString},
		{Name: "server_power_supply_type_id", Type: field.TypeString},
	}
	// ServerPowerSuppliesTable holds the schema information for the "server_power_supplies" table.
	ServerPowerSuppliesTable = &schema.Table{
		Name:       "server_power_supplies",
		Columns:    ServerPowerSuppliesColumns,
		PrimaryKey: []*schema.Column{ServerPowerSuppliesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "server_power_supplies_servers_server",
				Columns:    []*schema.Column{ServerPowerSuppliesColumns[4]},
				RefColumns: []*schema.Column{ServersColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "server_power_supplies_server_power_supply_types_server_power_supply_type",
				Columns:    []*schema.Column{ServerPowerSuppliesColumns[5]},
				RefColumns: []*schema.Column{ServerPowerSupplyTypesColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "serverpowersupply_created_at",
				Unique:  false,
				Columns: []*schema.Column{ServerPowerSuppliesColumns[1]},
			},
			{
				Name:    "serverpowersupply_updated_at",
				Unique:  false,
				Columns: []*schema.Column{ServerPowerSuppliesColumns[2]},
			},
			{
				Name:    "serverpowersupply_server_power_supply_type_id",
				Unique:  false,
				Columns: []*schema.Column{ServerPowerSuppliesColumns[5]},
			},
			{
				Name:    "serverpowersupply_server_id",
				Unique:  false,
				Columns: []*schema.Column{ServerPowerSuppliesColumns[4]},
			},
		},
	}
	// ServerPowerSupplyTypesColumns holds the columns for the "server_power_supply_types" table.
	ServerPowerSupplyTypesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "vendor", Type: field.TypeString, Size: 2147483647},
		{Name: "model", Type: field.TypeString, Size: 2147483647},
		{Name: "watts", Type: field.TypeString, Size: 2147483647},
	}
	// ServerPowerSupplyTypesTable holds the schema information for the "server_power_supply_types" table.
	ServerPowerSupplyTypesTable = &schema.Table{
		Name:       "server_power_supply_types",
		Columns:    ServerPowerSupplyTypesColumns,
		PrimaryKey: []*schema.Column{ServerPowerSupplyTypesColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "serverpowersupplytype_created_at",
				Unique:  false,
				Columns: []*schema.Column{ServerPowerSupplyTypesColumns[1]},
			},
			{
				Name:    "serverpowersupplytype_updated_at",
				Unique:  false,
				Columns: []*schema.Column{ServerPowerSupplyTypesColumns[2]},
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
		ServerCpUsTable,
		ServerCPUTypesTable,
		ServerChassesTable,
		ServerChassisTypesTable,
		ServerComponentsTable,
		ServerComponentTypesTable,
		ServerHardDrivesTable,
		ServerHardDriveTypesTable,
		ServerMemoriesTable,
		ServerMemoryTypesTable,
		ServerMotherboardsTable,
		ServerMotherboardTypesTable,
		ServerPowerSuppliesTable,
		ServerPowerSupplyTypesTable,
		ServerTypesTable,
	}
)

func init() {
	ServersTable.ForeignKeys[0].RefTable = ProvidersTable
	ServersTable.ForeignKeys[1].RefTable = ServerTypesTable
	ServerCpUsTable.ForeignKeys[0].RefTable = ServersTable
	ServerCpUsTable.ForeignKeys[1].RefTable = ServerCPUTypesTable
	ServerChassesTable.ForeignKeys[0].RefTable = ServersTable
	ServerChassesTable.ForeignKeys[1].RefTable = ServerChassisTypesTable
	ServerComponentsTable.ForeignKeys[0].RefTable = ServerComponentTypesTable
	ServerComponentsTable.ForeignKeys[1].RefTable = ServersTable
	ServerHardDrivesTable.ForeignKeys[0].RefTable = ServersTable
	ServerHardDrivesTable.ForeignKeys[1].RefTable = ServerHardDriveTypesTable
	ServerMemoriesTable.ForeignKeys[0].RefTable = ServersTable
	ServerMemoriesTable.ForeignKeys[1].RefTable = ServerMemoryTypesTable
	ServerMotherboardsTable.ForeignKeys[0].RefTable = ServersTable
	ServerMotherboardsTable.ForeignKeys[1].RefTable = ServerMotherboardTypesTable
	ServerPowerSuppliesTable.ForeignKeys[0].RefTable = ServersTable
	ServerPowerSuppliesTable.ForeignKeys[1].RefTable = ServerPowerSupplyTypesTable
}
