package models

type Request struct {
	URL    string                 `json:"url"`
	Method string                 `json:"method"`
	Params map[string]string      `json:"params,omitempty"`
	Form   map[string]string      `json:"form,omitempty"`
	Body   map[string]interface{} `json:"body,omitempty"`
}
