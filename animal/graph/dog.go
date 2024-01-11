package graph

import (
	"main/graph/model"
)

func dog(id, name string) *model.Dog {
	return &model.Dog{
		ID:   id,
		Name: name,
	}
}
