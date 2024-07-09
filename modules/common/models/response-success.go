package models

func Success(i interface{}) (res Response) {
	res = Response{}
	res.Code = "success"
	res.Message = "Success"
	res.Data = i

	return res
}

func ItemSuccess(i interface{}, c bool) (res ResponseItem) {
	res = ResponseItem{}
	res.Code = "success"
	res.Message = "Success"
	res.IsCached = c
	res.Data = i

	return res
}

func ListSuccess(i interface{}) (res ResponseList) {
	res = ResponseList{}
	res.Code = "success"
	res.Message = "Success"
	res.Data = i

	return res
}

func PagingSuccess(i interface{}, p Paging) (res ResponsePaging) {
	res = ResponsePaging{}
	res.Code = "success"
	res.Message = "Success"
	res.Data = i
	res.Page = p.Page
	res.Limit = p.Limit
	res.Total = p.Total

	return res
}
