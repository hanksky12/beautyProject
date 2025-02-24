package websocket

import (
	"github.com/olahol/melody"
	log "github.com/sirupsen/logrus"
	"sync"
)

var WebSocket *melody.Melody

// 用於存儲每個使用者的 WebSocket 連線，管理連線唯一性
var connections = struct {
	sync.Mutex
	data map[string]map[string]map[string]*melody.Session
}{
	data: make(map[string]map[string]map[string]*melody.Session),
}

//userID -> route -> connectionsManageParam -> session

func AddRouter() {
	WebSocket = melody.New()
	handler := Handler{}
	WebSocket.HandleMessage(func(s *melody.Session, msg []byte) {
		if h, ok := s.Get("httpInfo"); ok {
			httpInfo := h.(*HttpInfo)
			route := httpInfo.GetRoute()
			log.Info("Ws Path :", route)
			switch route {
			case "/api/user/mouse/tracking":
				handler.RecordAction(httpInfo, msg)
			case "/api/...":

			default:
				log.Printf("[Unknown Route] Received: %s", string(msg))

			}
		}
	})

	WebSocket.HandleConnect(func(s *melody.Session) {
		log.Println("Ws Client connected")
		if httpInfo, ok := s.Request.Context().Value("httpInfo").(*HttpInfo); ok {
			userID := httpInfo.GetUserID()
			route := httpInfo.GetRoute()
			paramName := httpInfo.GetConnectionsManageParam()
			connections.Lock()
			defer connections.Unlock()
			// 確保使用者在這個路由中只有一條連線
			if connections.data[userID] == nil {
				connections.data[userID] = make(map[string]map[string]*melody.Session)
			}
			if connections.data[userID][route] == nil {
				connections.data[userID][route] = make(map[string]*melody.Session)
			}
			// 如果該路由已經存在連線，關閉舊的連線
			if existingSession, exists := connections.data[userID][route][paramName]; exists {
				if existingSession != s {
					existingSession.Close()
				}
			}
			// 保存新連線
			connections.data[userID][route][paramName] = s
			//log.Info(connections)
			s.Set("httpInfo", httpInfo)
		}
	})

	WebSocket.HandleDisconnect(func(s *melody.Session) {
		if httpInfo, ok := s.Request.Context().Value("httpInfo").(*HttpInfo); ok {
			userID := httpInfo.GetUserID()
			route := httpInfo.GetRoute()
			paramName := httpInfo.GetConnectionsManageParam()

			connections.Lock()
			defer connections.Unlock()

			// 刪除斷開的連線
			if connections.data[userID] != nil && connections.data[userID][route] != nil {
				if currentSession, exists := connections.data[userID][route][paramName]; exists && currentSession == s {
					delete(connections.data[userID][route], paramName)
					// 如果該 route 下沒有連線了，刪除 route
					if len(connections.data[userID][route]) == 0 {
						delete(connections.data[userID], route)
					}
					// 如果該 userID 下沒有連線了，刪除 userID
					if len(connections.data[userID]) == 0 {
						delete(connections.data, userID)
					}
				}
			}
		}
		log.Println("Ws Client closed")
	})

	//if err := m.Broadcast(msg); err != nil {
	//	log.Infof("Failed to broadcast message: %v", err)
	//}
}
