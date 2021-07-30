package service

type Server struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	IP   string `json:"ip"`
	Port string `json:"port"`
	Desc string `json:"desc"`
}

func (s Server) RegisterServer() {

}

// 心跳检测服务
func Heartbeat() {

}

// 服务详情
func ServerDetail() {

}
