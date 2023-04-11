package models

type PaginationData struct {
	StartBlock int64
	EndBlock   int64
	Page       int
	PerPage    int
	Sort       string
}
