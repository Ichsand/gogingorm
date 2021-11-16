package dtos

type Multi struct {
	Publishers []string `json:"publishers" binding:"required"`
}
