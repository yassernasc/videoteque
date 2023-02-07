package server

import (
	"github.com/labstack/echo/v4"
	"github.com/olahol/melody"
	"net/http"
)

func getWsHandler(m *melody.Melody) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		m.HandleRequest(w, r)
	})
}

func WsRoutes(e *echo.Echo, m *melody.Melody) {
	e.GET("/ws", echo.WrapHandler(getWsHandler(m)))

	m.HandleMessage(func(_ *melody.Session, msg []byte) {
		m.Broadcast([]byte(msg))
	})
}
