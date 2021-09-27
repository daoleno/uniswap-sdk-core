package utils

import "fmt"

// SortedInsert givens an array of items sorted by `comparator`, insert an item into its sort index and constrain the size to
// `maxSize` by removing the last item
// TODO: verify this
func SortedInsert(items []interface{}, item interface{}, maxSize int, comparator func(a, b interface{}) bool) ([]interface{}, error) {
	if maxSize <= 0 {
		return nil, fmt.Errorf("maxSize must be greater than 0")
	}

	if len(items) > maxSize {
		return nil, fmt.Errorf("maxSize must be greater than the current size of items")

	}

	for i, v := range items {
		if comparator(v, item) {
			items = append(items[:i], append([]interface{}{item}, items[i:]...)...)
			return items, nil
		}
	}
	items = append(items, item)
	return items, nil
}
