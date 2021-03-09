package main

type Server struct {
	Addr string
	Port int
	Protocol string
	MaxConns int
}
type Options func(*Server)
func MaxConns(maxconns int) Options { return func(s *Server) { s.MaxConns = maxconns }}
func Protocols (p string) Options { return func(s *Server) { s.Protocol = p }}
func NewServer(addr string, port int, options ...func(*Server)) (*Server, error) {

	srv := Server{
		Addr:     addr,
		Port:     port,
		Protocol: "tcp",
		MaxConns: 1000,
	}
	for _, option := range options {
		option(&srv)
	}
	//...
	return &srv, nil
}

func main()  {
	s2, _ := NewServer("localhost", 2048, Protocols("udp"))
	println(s2)
}