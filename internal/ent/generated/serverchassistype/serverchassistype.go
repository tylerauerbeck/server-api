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

package serverchassistype

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"go.infratographer.com/x/gidx"
)

const (
	// Label holds the string label denoting the serverchassistype type in the database.
	Label = "server_chassis_type"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldVendor holds the string denoting the vendor field in the database.
	FieldVendor = "vendor"
	// FieldModel holds the string denoting the model field in the database.
	FieldModel = "model"
	// FieldHeight holds the string denoting the height field in the database.
	FieldHeight = "height"
	// FieldIsFullDepth holds the string denoting the is_full_depth field in the database.
	FieldIsFullDepth = "is_full_depth"
	// FieldParentChassisTypeID holds the string denoting the parent_chassis_type_id field in the database.
	FieldParentChassisTypeID = "parent_chassis_type_id"
	// Table holds the table name of the serverchassistype in the database.
	Table = "server_chassis_types"
)

// Columns holds all SQL columns for serverchassistype fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldVendor,
	FieldModel,
	FieldHeight,
	FieldIsFullDepth,
	FieldParentChassisTypeID,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// VendorValidator is a validator for the "vendor" field. It is called by the builders before save.
	VendorValidator func(string) error
	// ModelValidator is a validator for the "model" field. It is called by the builders before save.
	ModelValidator func(string) error
	// HeightValidator is a validator for the "height" field. It is called by the builders before save.
	HeightValidator func(string) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() gidx.PrefixedID
)

// OrderOption defines the ordering options for the ServerChassisType queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByVendor orders the results by the vendor field.
func ByVendor(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldVendor, opts...).ToFunc()
}

// ByModel orders the results by the model field.
func ByModel(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldModel, opts...).ToFunc()
}

// ByHeight orders the results by the height field.
func ByHeight(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldHeight, opts...).ToFunc()
}

// ByIsFullDepth orders the results by the is_full_depth field.
func ByIsFullDepth(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsFullDepth, opts...).ToFunc()
}

// ByParentChassisTypeID orders the results by the parent_chassis_type_id field.
func ByParentChassisTypeID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldParentChassisTypeID, opts...).ToFunc()
}
