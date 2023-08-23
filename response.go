package molit

import (
	"encoding/xml"
	"io"
)

type Header struct {
	ResultCode int    `xml:"resultCode"`
	ResultMsg  string `xml:"resultMsg"`
}

type Body interface {
	[]Apt | []Office | []RH | []SH | []Land
}

type Response[T Body] struct {
	Header Header `xml:"header"`
	Body   T      `xml:"body>items>item"`
}

func NewResponse[T Body](body T) *Response[T] {
	return &Response[T]{
		Body: body,
	}
}

func ResponseFromXML[T Body](body T, data io.Reader) *Response[T] {
	res := NewResponse(body)

	xml.NewDecoder(data).Decode(&res)

	return res
}
