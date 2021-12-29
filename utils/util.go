package utils

func ComputePaginationFromRequest(rowcountpage uint32, currentpage uint32) (limit uint32, offset uint32) {

	limit = uint32(rowcountpage)
	offset = uint32(currentpage) * uint32(rowcountpage)
	return
}
