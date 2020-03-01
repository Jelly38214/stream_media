package dbops

import (
	"database/sql"
	"log"
	"stream_media/api/defs"
	"stream_media/api/utils"
	"time"

	// Blank import in non-main package
	_ "github.com/go-sql-driver/mysql"
)

// AddUserCredential func
func AddUserCredential(loginName string, pwd string) error {
	stmtIns, err := dbConn.Prepare(`
		INSERT INTO users (login_name, pwd) VALUES (?, ?)		
	`)
	if err != nil {
		return err
	}

	_, err = stmtIns.Exec(loginName, pwd)

	if err != nil {
		return err
	}

	defer stmtIns.Close() // make sure closing connection
	return nil
}

// GetUserCredential func
func GetUserCredential(loginName string) (string, error) {
	stmtOut, err := dbConn.Prepare(`
		SELECT pwd FROM users WHERE login_name = ?	
	`)

	if err != nil {
		log.Panicf("%s", err)
		return "", err
	}

	var pwd string
	err = stmtOut.QueryRow(loginName).Scan(&pwd)

	// sql.ErrNoRows没有结果，按错误返回
	if err != nil && err != sql.ErrNoRows {
		return "", err
	}

	defer stmtOut.Close()

	return pwd, nil
}

// DeleteUser func
func DeleteUser(loginName string, pwd string) error {
	stmtDel, err := dbConn.Prepare(`
		DELETE FROM users WHERE login_name=? AND pwd=?
	`)

	if err != nil {
		return err
	}

	_, err = stmtDel.Exec(loginName, pwd)
	if err != nil {
		return err
	}

	defer stmtDel.Close()
	return nil
}

// AddNewVideo func
func AddNewVideo(aid int, name string) (*defs.VideoInfo, error) {
	vid, err := utils.NewUUID()

	if err != nil {
		return nil, err
	}

	t := time.Now()
	ctime := t.Format("Jan 02 2006, 15:04:05") // 固定字符串，Go正式推出时间

	stmtIns, err := dbConn.Prepare(`
		INSERT INTO video_info (id, author_id, name, display_ctime) VALUES(?, ?, ?, ?)
	`)
	if err != nil {
		return nil, err
	}

	_, err = stmtIns.Exec(vid, aid, name, ctime)

	if err != nil {
		return nil, err
	}

	res := &defs.VideoInfo{ID: vid, AuthodID: aid, Name: name, DisplayCtime: ctime}

	defer stmtIns.Close()
	return res, nil
}

// GetVideoInfo func
func GetVideoInfo(vid string) (*defs.VideoInfo, error) {
	stmtOut, err := dbConn.Prepare(`
			SELECT author_id, name, display_ctime FROM video_info WHERE id = ? 
	 `)

	var (
		aid  int
		dct  string
		name string
	)

	err = stmtOut.QueryRow(vid).Scan(&aid, &name, &dct)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	if err == sql.ErrNoRows {
		return nil, nil
	}

	res := &defs.VideoInfo{
		ID:           vid,
		AuthodID:     aid,
		Name:         name,
		DisplayCtime: dct,
	}

	defer stmtOut.Close()
	return res, nil
}

// DeleteVideoInfo func
func DeleteVideoInfo(vid string) error {
	stmtDel, err := dbConn.Prepare(`
		DELETE FROM video_info WHERE id = ?
	`)

	if err != nil {
		return err
	}

	_, err = stmtDel.Exec(vid)

	if err != nil {
		return err
	}

	defer stmtDel.Close()
	return nil
}

// AddNewComments func
func AddNewComments(vid string, aid int, content string) error {
	id, err := utils.NewUUID()
	if err != nil {
		return err
	}

	stmtIns, err := dbConn.Prepare(`
		INSERT INTO comments (id, video_id, author_id, content) values(?, ?, ?, ?)	
	`)

	if err != nil {
		return err
	}

	_, err = stmtIns.Exec(id, vid, aid, content)

	if err != nil {
		return err
	}
	defer stmtIns.Close()
	return nil
}

// ListComments func, *加struct是什么意思？？
// 这里需要合表，user+commens, FROM_UNIXTIME 精确到秒
func ListComments(vid string, from, to int) ([]*defs.Comment, error) {
	stmtOut, err := dbConn.Prepare(`
		SELECT comments.id, users.Login_name, comments.content FROM comments
		INNER JOIN users ON comments.author_id = users.id 
		WHERE comments.video_id = ? AND comments.time > FROM_UNIXTIME(?) AND comments.time <= FROM_UNIXTIME(?)
	`)

	var res []*defs.Comment

	rows, err := stmtOut.Query(vid, from, to)
	if err != nil {
		return res, err
	}

	for rows.Next() {
		var id, name, content string
		if err := rows.Scan(&id, &name, &content); err != nil {
			return res, err
		}

		c := &defs.Comment{
			ID:      id,
			VideoID: vid,
			Author:  name,
			Content: content,
		}

		res = append(res, c)
	}

	defer stmtOut.Close()

	return res, nil
}
