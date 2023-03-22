package request_struct

type DemoRequest struct {
	Account string `json:"account" binding:"required"`
}
