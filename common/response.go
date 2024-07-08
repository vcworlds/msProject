package common

type Response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

func (r *Response) Success(msg string, data any) *Response {
	r.Code = 200
	r.Msg = msg
	r.Data = data
	return r
}

func (r *Response) Fail(msg string, code int) *Response {
	r.Code = code
	r.Msg = msg
	return r
}
