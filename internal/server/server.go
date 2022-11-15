package server

type Server struct {
	httpServer *HttpServer
}

func NewServer(httpServerPort string) *Server {
	httpServer := NewHttpServer(httpServerPort)

	return &Server{
		httpServer: httpServer,
	}
}

func (s *Server) Run() error {
	err := s.httpServer.Run()
	return err
}
