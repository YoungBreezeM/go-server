package models

type R[T any] struct {
	Status  int16  `json:"status"`
	Data    T      `json:"data"`
	Message string `json:"message"`
}
