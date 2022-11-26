package dbops

import (
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func AddUserCredential(id int, loginName string, pwd string) error {
	stmtIns, err := dbConn.Prepare("INSERT INTO users (id,login_name, pwd) VALUES (?,?,?)")
	if err != nil {
		return err
	}
	_, err = stmtIns.Exec(id, loginName, pwd)
	if err != nil {
		return err
	}
	err = stmtIns.Close()
	if err != nil {
		return err
	}

	return nil
}

func GetUserCredential(loginName string) (string, error) {
	stmtOut, err := dbConn.Prepare("SELECT pwd FROM users WHERE login_name = ?")
	if err != nil {
		log.Printf("%s", err)
		return "", err
	}

	var pwd string
	err = stmtOut.QueryRow(loginName).Scan(&pwd)
	if err != nil {
		log.Printf("%s", err)
		return "", err
	}

	err = stmtOut.Close()
	if err != nil {
		log.Printf("%s", err)
		return "", err
	}
	return pwd, nil
}

func DeleteUser(loginName string, pwd string) error {
	stmtDel, err := dbConn.Prepare("DELETE FROM users WHERE login_name= ? AND pwd= ?")
	if err != nil {
		log.Printf("DeleteUser error: %s", err)
		return err
	}

	_, err = stmtDel.Exec(loginName, pwd)
	if err != nil {
		return err
	}
	err = stmtDel.Close()
	if err != nil {
		return err
	}
	return nil
}
