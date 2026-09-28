package database

import _ "embed"

//go:embed queries/schema.sql
var schema string
