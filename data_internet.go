package gofakeit

// DataInternet consists of various internet information
var DataInternet = map[string][]string{
	"browser":       {"firefox", "chrome", "internetExplorer", "opera", "safari"},
	"domain_suffix": {"com", "biz", "info", "name", "net", "org"},
	"http_method":   {"HEAD", "GET", "POST", "PUT", "PATCH", "DELETE"},
}
