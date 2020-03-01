package session

import (
	"stream_media/api/dbops"
	"stream_media/api/defs"
	"stream_media/api/utils"
	"sync"
	"time"
)

var sessionMap *sync.Map

func init() {
	sessionMap = &sync.Map{}
}

// deleteExpiredSession func
func deleteExpiredSession(sid string) {
	sessionMap.Delete(sid)
	dbops.DeleteSession(sid)
}

// LoadSessionsFromDB func
func LoadSessionsFromDB() error {
	r, err := dbops.RetrieveAllSessions()
	if err != nil {
		return err
	}

	r.Range(func(k, v interface{}) bool {
		ss := v.(*defs.SimpleSession)
		sessionMap.Store(k, ss)
		return true
	})

	return nil
}

// GenerateNewSessionID func
func GenerateNewSessionID(un string) string {
	id, _ := utils.NewUUID()
	ct := time.Now().UnixNano() / 1000000
	ttl := ct + 30*60*1000 // 30min avalable

	ss := &defs.SimpleSession{Username: un, TTL: ttl}
	sessionMap.Store(id, ss)

	dbops.InsertSessions(id, ttl, un)
	return id
}

// IsSessionExpired func
func IsSessionExpired(sid string) (string, bool) {
	if ss, ok := sessionMap.Load(sid); ok == true {
		ct := time.Now().UnixNano() / 1000000
		// expired
		if ss.(*defs.SimpleSession).TTL < ct {
			deleteExpiredSession(sid)
			return "", true
		}

		return ss.(*defs.SimpleSession).Username, false
	}

	return "", true
}
