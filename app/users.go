package models

type User struct {
	ID       int
	UUID string
	Name string
	Email string
	PassWord string
	CreateAt time.time
}

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