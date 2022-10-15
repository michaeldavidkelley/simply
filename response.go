package simply

import "net/http"

type response struct {
	status   int
	contents string
}

func (r *response) GetContents() string {
	return r.contents
}

func (r *response) Raw(str string) Response {
	r.contents = str

	return r
}

func (r *response) GetStatus() int {
	return r.status
}

func (r *response) SetStatus(status int) Response {
	r.status = status

	return r
}

type Response interface {
	SetStatus(status int) Response
	GetStatus() int

	Raw(str string) Response
	GetContents() string
}

func NewResponse() Response {
	return &response{
		status: http.StatusOK,
	}
}
