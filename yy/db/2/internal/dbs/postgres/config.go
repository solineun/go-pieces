package postgres

type Config struct {
	Addr               string
	Port               uint16
	User, Password, Db string
}
