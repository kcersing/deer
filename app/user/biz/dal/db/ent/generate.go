package ent

//go:generate go run -mod=readonly entgo.io/ent/cmd/ent generate --feature intercept,schema/snapshot,sql/versioned-migration ./schema
