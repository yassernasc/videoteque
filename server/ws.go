package server

import (
	"github.com/olahol/melody"
	"net/http"
)

var m *melody.Melody

func init() {
	m = melody.New()

	m.HandleMessage(func(s *melody.Session, msg []byte) {
		m.Broadcast(msg)
	})
}

func handleWs(w http.ResponseWriter, r *http.Request) {
	m.HandleRequest(w, r)
}
