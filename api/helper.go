package api

func (e *ApiError) Error() string {
	return e.Message
}

func CalculatePageCount(total uint, pageSize uint) (pages int) {
	pages = int(total / pageSize)
	if total%pageSize > 0 {
		pages++
	}
	return
}
