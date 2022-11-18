// Resolver type.
//
// ref: Eli Bendersky [https://eli.thegreenplace.net]
package graph

import (
	"rests.com/internal/taskstore"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Store *taskstore.TaskStore
}
