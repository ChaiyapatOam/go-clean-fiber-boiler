package domain

type Env struct {
	MYSQL_URI string `validate:"required"`
	PORT      string `validate:"required"`
}
