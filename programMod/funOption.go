/*
函数式选项模式
实例化serve通过传入不同的函数来配置不同的参数
*/

type Server struct {
	Addr     string        //必填
	Port     int           //必填
	Protocol string        //选填
	Timeout  time.Duration //选填
	MaxConns int           //选填
	TLS      *tls.Config   //选填
}

type Option func(*Server)

//每个配置项都要有一个函数
func Protocol(p string) Option {
	return func(s *Server) {
		s.Protocol = p
	}
}
func Timeout(timeout time.Duration) Option {
	return func(s *Server) {
		s.Timeout = timeout
	}
}
func MaxConns(maxconns int) Option {
	return func(s *Server) {
		s.MaxConns = maxconns
	}
}
func TLS(tls *tls.Config) Option {
	return func(s *Server) {
		s.TLS = tls
	}
}

func NewServer(addr string, port int, options ...func(*Server)) (*Server, error) {
	srv := Server{
		Addr:     addr,
		Port:     port,
		Protocol: "tcp",
		Timeout:  30 * time.Second,
		MaxConns: 1000,
		TLS:      nil,
	}
	for _, option := range options {
		option(&srv)
	}
	return &srv, nil
}

//通过传入不同的函数来实现配置
// s1, _ := NewServer("localhost", 1024)
// s2, _ := NewServer("localhost", 2048, Protocol("udp"))
// s3, _ := NewServer("0.0.0.0", 8080, Timeout(300*time.Second), MaxConns(1000))