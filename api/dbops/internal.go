package dbops

import (
	"database/sql"
	"strconv"
	"stream_media/api/defs"
	"sync"
)

// InsertSessions func
func InsertSessions(sid string, ttl int64, uname string) error {
	ttlstr := strconv.FormatInt(ttl, 10)
	stmtIns, err := dbConn.Prepare(`
		INSERT INTO sessions (session_id, TTL, login_name) VALUES(?, ?, ?)	
	`)

	if err != nil {
		return err
	}

	_, err = stmtIns.Exec(sid, ttlstr, uname)

	if err != nil {
		return err
	}

	defer stmtIns.Close()
	return nil
}

// RetrieveSession func
func RetrieveSession(sid string) (*defs.SimpleSession, error) {
	ss := &defs.SimpleSession{}
	stmtOut, err := dbConn.Prepare(`
		SELECT TTL, login_name FROM sessions WHERE session_id=?
	`)

	if err != nil {
		return nil, err
	}

	var ttl, uname string
	stmtOut.QueryRow(sid).Scan(&ttl, &uname)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	if res, err := strconv.ParseInt(ttl, 10, 64); err == nil {
		ss.TTL = res
		ss.Username = uname
	} else {
		return nil, err
	}

	defer stmtOut.Close()
	return ss, nil
}

// RetrieveAllSessions func
func RetrieveAllSessions() (*sync.Map, error) {
	m := &sync.Map{}
	stmtOut, err := dbConn.Prepare(`
		SELECT * FROM sessions	
	`)

	if err != nil {
		return nil, err
	}

	rows, err := stmtOut.Query()

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var id, ttlstr, loginName string
		if err := rows.Scan(&id, &ttlstr, &loginName); err != nil {
			break
		}

		if ttl, err1 := strconv.ParseInt(ttlstr, 10, 64); err1 == nil {
			ss := &defs.SimpleSession{Username: loginName, TTL: ttl}
			m.Store(id, ss)
		}

	}

	return m, nil
}

// DeleteSession func
func DeleteSession(sid string) error {
	stmtOut, err := dbConn.Prepare(`
		DELETE FROM sessions WHERE session_id = ?	
	`)

	if err != nil {
		return err
	}

	if _, err := stmtOut.Query(sid); err != nil {
		return err
	}

	return nil
}
