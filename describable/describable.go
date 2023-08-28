package describable

import (
	item "app/Item"
	"fmt"
)

type Describable interface {
	Description() string
}

func ConvertItemSliceToDescribableSlice(slice []item.Item) ([]Describable, error) {
	describableSlice := make([]Describable, len(slice))
	if slice == nil {
		return nil, fmt.Errorf("empty slice")
	}
	for i, val := range slice {
		describableSlice[i] = val
	}
	return describableSlice, nil
}
