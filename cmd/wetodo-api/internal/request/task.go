package request

type TaskListRequest struct {
	Completed bool   `json:"completed,omitempty"`
	Limit     int    `json:"limit,omitempty"`
	Page      int    `json:"page,omitempty"`
	SortField string `json:"sortField,omitempty"`
	SortOrder string `json:"sortOrder,omitempty"`
}
