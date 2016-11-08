package model

type JsonResult struct {
	Success bool
	Error   []byte
	Data    []byte
}
