package interfaces

type IRouter interface {
	SetRoute(pattern, contentType string, hosts []string, port string)
}
