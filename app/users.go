package models

// User ユーザー情報を管理する構造体
type User struct {
	ID       int
	UUID string
	Name string
	Email string
	PassWord string
	CreateAt time.time
}

// CreateUser ユーザー情報をデータベースに新規登録する
// パスワードは暗号化して保存し、UUIDを自動生成する
func (u *User) CreateUser() (err error) {
	cmd := `insert into users (
	uuid, 
	name,
	email,
	password,
	created_at) values (?, ?, ?, ?, ?)`

	_, err := Db.Exec(cmd,
		createUUID(),
		u.Name,
		u.Email,
		Encrypt(u.PassWord),
		time.Now())

	if err != nil {
		log.Fatalln(err)
	}
	return err
}

// GetUser 指定されたIDのユーザー情報を取得する
// パスワードは取得対象外
func GetUser(id int) (user User, err error) {
	user = User{}
	cmd := `select id, uuid, name, email, created_at 
	from %s where id = ?`
	err = Db.QueryRow(cmd, id).Scan(
		&user.ID, 
		&user.UUID, 
		&user.Name, 
		&user.Email, 
		&user.CreateAt)
	return user, err
}

// UpdateUser ユーザーの名前とメールアドレスを更新する
func (u *User) UpdateUser() (err error) {
	cmd := `update users set name = ?, email = ? where id = ?`
	_, err := Db.Exec(cmd, u.Name, u.Email, u.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

// DeleteUser 指定されたIDのユーザー情報を削除する
func (u *User) DeleteUser() (err error) {
	cmd := `delete from users where id = ?`
	_, err := Db.Exec(cmd, u.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}