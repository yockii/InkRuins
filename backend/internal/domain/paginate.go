package domain

type PaginateReq struct {
	Page     int `query:"page"`
	PageSize int `query:"size"`
}

type PaginateResponse struct {
	Response
	Total    int64 `json:"total"`
	Page     int   `json:"page"`
	PageSize int   `json:"size"`
}

func NewPaginateResponse[T any](code int, msg string, total int64, page int, pageSize int, data T) PaginateResponse {
	return PaginateResponse{
		Response: Response{
			Code:    code,
			Message: msg,
			Data:    data,
		},
		Total:    total,
		Page:     page,
		PageSize: pageSize,
	}
}

func NewSuccessPaginateResponse[T any](total int64, page int, pageSize int, data T) PaginateResponse {
	return NewPaginateResponse(0, "success", total, page, pageSize, data)
}
