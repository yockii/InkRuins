package domain

type ListProjectReq struct {
	PaginateReq
	Title string `query:"title"`
}

type CreateProjectReq struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Genre       string `json:"genre"`
}
