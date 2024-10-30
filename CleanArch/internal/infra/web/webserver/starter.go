package webserver

type WebServerStarter struct {
	WebServer WebServer
}

// TODO: Validar se alguém está usando
func NewWebServerStarter(webServer WebServer) *WebServerStarter {
	return &WebServerStarter{
		WebServer: webServer,
	}
}
