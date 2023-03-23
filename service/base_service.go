package service

const SUCCESS = 0
const FAIL = 100

type ResultService struct {
	Code int
	Data interface{}
}
