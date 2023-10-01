package filter

import (
	"github.com/Masterminds/squirrel"
)

type Filter struct {
	Key   string
	Value interface{}
}

type Filters []Filter

func (f Filters) Eq() map[string]interface{} {
	var s = squirrel.Eq{}
	for _, filter := range f {
		s[filter.Key] = filter.Value
	}
	return s
}
