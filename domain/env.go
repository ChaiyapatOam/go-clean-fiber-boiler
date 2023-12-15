package domain

type Env struct {
	GOOGLE_CLIENTID     string `validate:"required"`
	GOOGLE_CLIENTSECRET string `validate:"required"`
	GOOGLE_REDIRECT     string `validate:"required"`
	MYSQL_URI           string `validate:"required"`
	PORT                string `validate:"required"`
	SESSION_PREFIX      string `validate:"required"`
	SESSION_SECRET      string `validate:"required"`
	HASH_SECRET         string `validate:"required"`
	FRONTEND_URL        string `validate:"required"`
}
