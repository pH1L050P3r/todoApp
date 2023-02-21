package objects

type Environment struct {
	ENVIRONMENT string

	WEB_HOST string
	WEB_PORT string

	RDS_USERNAME       string
	RDS_PASSWORD       string
	RDS_HOSTNAME       string
	RDS_PORT           string
	RDS_DB_NAME        string
	RDS_CONNECT_STRING string
}
