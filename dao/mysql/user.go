package mysql

//func Register(user *models.User) (err error) {
//	sqlStr := "select count(user_id) from user where username = ?"
//	var count int64
//	err = db.Get(&count, sqlStr, user.UserName)
//	if err != nil && err != sql.ErrNoRows {
//		return err
//	}
//	if count > 0 {
//		// 用户已存在
//		return ErrorUserExit
//	}
//	// 生成user_id
//	userID, err := snowflake.GetID()
//	if err != nil {
//		return ErrorGenIDFailed
//	}
//	// 生成加密密码
//	password := encryptPassword([]byte(user.Password))
//	// 把用户插入数据库
//	sqlStr = "insert into user(user_id, username, password) values (?,?,?)"
//	_, err = db.Exec(sqlStr, userID, user.UserName, password)
//	return
//}
