package table

// table 模型

type User struct {
	Id       int64  `db:"id"`
	Username string `db:"username"`
	Password string `db:"password"`
	Status   int    `db:"status"`
	Phone    string `db:"phone"`
}

func GetTableName() string {
	return "user"
}
