package dbops

import (
	"testing"
)

/*
	init(login, truncate tables) => run tests => clear data(truncate, tables)
*/

var temVid = ""

func clearTables() {
	dbConn.Exec("truncate users")
	dbConn.Exec("truncate video_info")
	dbConn.Exec("truncate comments")
	dbConn.Exec("truncate sessions")
}

func TestMain(m *testing.M) {
	clearTables()
	m.Run()
	clearTables()
}

func TestUserWorkFlow(t *testing.T) {
	t.Run("Add", testAddUser)
	t.Run("Get", testGetUser)
	t.Run("Delete", testDeleteUser)
	t.Run("Reget", testRegetUser)
}

func TestVideoWorkFlow(t *testing.T) {
	clearTables()
	t.Run("PrepareUser", testAddUser)
	t.Run("AddVideo", testAddVideoInfo)
	t.Run("GetVideo", testGetVideoInfo)
	t.Run("DeleteVideo", testDeleteVideoInfo)
	t.Run("RegetVideo", testRegetVideoInfo)
}

func testAddUser(t *testing.T) {
	err := AddUserCredential("jelly", "c11090201")
	if err != nil {
		t.Errorf("Error of AddUser: %v", err)
	}
}

func testGetUser(t *testing.T) {
	pwd, err := GetUserCredential("jelly")
	if pwd != "c11090201" || err != nil {
		t.Errorf("Error of GetUser: %v", err)
	}
}

func testDeleteUser(t *testing.T) {
	err := DeleteUser("jelly", "c11090201")
	if err != nil {
		t.Errorf("Error of DeleteUser: %v", err)
	}
}

func testRegetUser(t *testing.T) {
	pwd, err := GetUserCredential("jelly")

	if err != nil {
		t.Errorf("Error of RegetUser: %v", err)
	}

	// 如果用户不存在，pwd应该为空
	if pwd != "" {
		t.Errorf("Deleting user test failed")
	}
}

func testAddVideoInfo(t *testing.T) {
	videoInfo, err := AddNewVideo(1, "Game")
	if err != nil {
		t.Errorf("Error of AddVideoInfo: %v", err)
	}

	if videoInfo.Name != "Game" {
		t.Errorf("Error of VideoInfo Name: %v", videoInfo.Name)
	}

	temVid = videoInfo.ID
}

func testGetVideoInfo(t *testing.T) {
	_, err := GetVideoInfo(temVid)
	if err != nil {
		t.Errorf("Error of GetVideoInfo: %v", err)
	}
}

func testDeleteVideoInfo(t *testing.T) {
	err := DeleteVideoInfo(temVid)
	if err != nil {
		t.Errorf("Error of DeleteVideoInfo: %v", err)
	}
}

func testRegetVideoInfo(t *testing.T) {
	videoInfo, err := GetVideoInfo(temVid)

	if err != nil || videoInfo != nil {
		t.Errorf("Error of RegetVideoInfo: %v", videoInfo)
	}
}
