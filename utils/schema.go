package utils

import "github.com/hashicorp/go-memdb"

// Person keeps details about single user
type Person struct {
	Id		string
	Email string
	Name  string
	Age   int
	Group *Group
}

// Group keeps group details
type Group struct {
	Id string
	Name string
}

// CreateSchema creates schema for in-memory db
func CreateSchema() *memdb.MemDB {

	// Create the DB schema
	schema := &memdb.DBSchema{
		Tables: map[string]*memdb.TableSchema{
			"person": &memdb.TableSchema{
				Name: "person",
				Indexes: map[string]*memdb.IndexSchema{
					"id": &memdb.IndexSchema{
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "Email"},
					},
				},
			},
			"group": &memdb.TableSchema{
				Name: "group",
				Indexes: map[string]*memdb.IndexSchema{
					"id": 
				}
			}
		},
	}

	// Create a new data base
	db, err := memdb.NewMemDB(schema)
	if err != nil {
		panic(err)
	}

	return db
}
