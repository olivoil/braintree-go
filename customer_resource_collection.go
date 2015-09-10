package braintree

import "math"

// CustomerResourceCollection wraps a customer collection where the actual
// fetching of the customers is paged. Use `Each` to iterate over all items
// in the collection.
type CustomerResourceCollection struct {
	*Braintree
	ids         []string
	pageSize    int64
	pagingBlock func(ids []string) ([]*Customer, error)
}

// Each iterates over all items in the collection, calling `f` for each
// customer.
func (r *CustomerResourceCollection) Each(f func(*Customer)) error {
	pageCount := int64(math.Ceil(float64(len(r.ids)) / float64(r.pageSize)))
	var page int64
	for page = 0; page < pageCount; page++ {
		pageStart := (page * r.pageSize)
		pageEnd := int64(math.Min(float64((page+1)*r.pageSize), float64(len(r.ids))))
		idBlock := r.ids[pageStart:pageEnd]
		customers, err := r.pagingBlock(idBlock)
		if err != nil {
			return err
		}
		for _, c := range customers {
			f(c)
		}
	}
	return nil
}

// IsEmpty true if the collection has no items, false otherwise.
func (r *CustomerResourceCollection) IsEmpty() bool {
	return len(r.ids) == 0
}

// First returns the first item in the collection or nil if the collection is
// empty
func (r *CustomerResourceCollection) First() (*Customer, error) {
	if len(r.ids) == 0 {
		return nil, nil
	}
	c, err := r.pagingBlock([]string{r.ids[0]})
	if err != nil {
		return nil, err
	}
	if len(c) == 0 {
		return nil, nil
	}
	return c[0], nil
}

// MaximumSize returns the maximum size of the resource collection.
// Only the maximum size of a resource collection can be determined since the
// data on the server can change while fetching blocks of results for iteration.
// For example, customers can be deleted while iterating, so the number of
// results iterated over may be less than the maximum_size. In general, this
// method should be avoided.
func (r *CustomerResourceCollection) MaximumSize() int {
	return len(r.ids)
}
