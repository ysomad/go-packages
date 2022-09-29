package pagination

// List is generic object for pagination implementations.
// Usage: `func findRecords(p Params) List[YourModel]`.
type List[T any] struct {
	Result  []T  `json:"result"`
	HasNext bool `json:"has_next"`
}

// NewList creates new List with result and has next indicator which is
// indicates is there more records in db.
// It's assumed that List will be used when querying limit + 1 items from db,
// last item will be removed from `List.Result`.
// Whether there are more records in db will be determined by length of objList.
func NewList[T any](objList []T, limit uint16) List[T] {
	objLen := uint16(len(objList))
	hasNext := objLen == limit+1
	if hasNext {
		objList = objList[:objLen-1]
	}
	return List[T]{
		Result:  objList,
		HasNext: hasNext,
	}
}
