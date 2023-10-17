package queryParameter

import (
	"github.com/Orendev/gokeeper/pkg/type/filter"
	"github.com/Orendev/gokeeper/pkg/type/pagination"
	"github.com/Orendev/gokeeper/pkg/type/sort"
)

type QueryParameter struct {
	Sorts      sort.Sorts
	Pagination pagination.Pagination
	Filters    filter.Filters
}
