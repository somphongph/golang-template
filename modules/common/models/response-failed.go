package models

func CannotBindData() (res Response) {
	res = Response{}
	res.Code = "cannot_bind_data"
	res.Message = "Cannot bind data"
	res.Data = nil

	return res
}

func DataNotFound() (res Response) {
	res = Response{}
	res.Code = "data_not_found"
	res.Message = "Data not found."
	res.Data = nil

	return res
}

func OperationFailed() (res Response) {
	res = Response{}
	res.Code = "operation_failed"
	res.Message = "The operation failed."
	res.Data = nil

	return res
}
