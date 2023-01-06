package requests

type ContentSearch struct {
	Tag       string `form:"tag" json:"tag"`
	Category  string `form:"category" json:"category"`
	Status    string `form:"status" json:"status"`
	Title     string `form:"title" json:"title"`
	Content   string `form:"content" json:"content"`
	Type      string `form:"type" json:"type"`
	StartTime string `form:"startTime" json:"startTime"`
	EndTime   string `form:"endTime" json:"endTime"`
}
