package dbops

import "testing"

func clearTables() {
	_, err := dbConn.Exec("DELIMITER users")
	if err != nil {
		return
	}
	_, err = dbConn.Exec("DELIMITER video_info")
	if err != nil {
		return
	}
	_, err = dbConn.Exec("DELIMITER comments")
	if err != nil {
		return
	}
	_, err = dbConn.Exec("DELIMITER sessions")
	if err != nil {
		return
	}
}

func TestMain(m *testing.M) {
	clearTables()
	m.Run()
	clearTables()
}

func TestUserWorkFlow(t *testing.T) {
	t.Run("Add", testAddUser)
	t.Run("Get", testGetUser)
	t.Run("Del", testDeleteUser)
	t.Run("Reget", testRegetUser)
}

func testAddUser(t *testing.T) {
	err := AddUserCredential(1, "Yukino", "123")
	if err != nil {
		t.Errorf("Error of AddUser: %v", err)
	}
}

func testGetUser(t *testing.T) {
	pwd, err := GetUserCredential("Yukino")
	if pwd != "123" || err != nil {
		t.Errorf("Error of GetUser: %v", err)
	}
}

func testDeleteUser(t *testing.T) {
	err := DeleteUser("Yukino", "123")
	if err != nil {
		t.Errorf("Error of DeleteUser: %v", err)
	}
}

func testRegetUser(t *testing.T) {
	pwd, err := GetUserCredential("Yukino")
	// 如果没删除用户Yukino成功返回的就是 pwd 和 nil
	// 那么测试通过就是返回了 "" 和 err
	// 也就是 err==nil 就是没成功删除，此时应该暂停 testRegetUser 函数
	// 或者 pwd != "" （因为没删除密码是会返回的）就代表删除成功，否则也应该暂停 testRegetUser 函数
	if err == nil {
		t.Errorf("Error of RegetUser: %v", err)
	}
	if pwd != "" {
		t.Errorf("Deleting user test failed")
	}
}
