package entities

import "strconv"

type Metadata struct {
	Page int `json:"page"`
	Limit int `json:"limit"`
}

func (m Metadata) GetOffset() int {
	return (m.Page - 1) * m.Limit
}

func GetMetadata(page, limit string) Metadata {
	pageInt, err := strconv.Atoi(page)
	if err != nil || pageInt <= 0 {
		pageInt = 1
	}
	limitInt, err := strconv.Atoi(limit)
	if err != nil || limitInt <= 0 {
		limitInt = 10
	}
	return Metadata{Page: pageInt, Limit: limitInt}
}