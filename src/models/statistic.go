package models

type Statistic struct {
	TotalBlogs    int
	TotalUsers    int
	TotalComments int
	ValueBlogs    int
	ValueUsers    int
	ValueComments int
}

type StatisticRes struct {
	Id    string `json:"id"`
	Total int    `json:"total"`
	Value int    `json:"value"`
	Title string `json:"title"`
}
