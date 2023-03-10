package http

type Option struct {
	Header map[string]string
	Cookie string
	Query map[string]string
	Body string
}

func (h *Option) SetHeader(key string, value string) {
	h.Header[key] = value
}

func (h *Option) SetCookie(cookieValue string) {
	h.Cookie = cookieValue
}

func (h *Option) SetQuery(query map[string]string) {
	h.Query = query
}

func (h *Option) SetBody(body string) {
	h.Body = body
}

func (h *Option) copyWith(option Option) {
	for key, value := range option.Header {
		h.SetHeader(key, value)
	}
	if option.Cookie != "" {
		h.SetCookie(option.Cookie)
	}
	if len(option.Query) != 0 {
		h.Query = option.Query
	}
	if option.Body != "" {
		h.Body = option.Body
	}
}

func defaultOption() Option {
	return Option{
		Header: map[string]string{
			"User-Agent": "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/106.0.0.0 Safari/537.36 Edg/106.0.1370.37",
		},
	}
}
