package response

type Pagination struct {
	Total       int `json:"total"`
	CurrentPage int `json:"page"`
	PerPage     int `json:"page_size"`
}

type ResponseList struct {
	Data interface{} `json:"data"`
	Pagination
}

