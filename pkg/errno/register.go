package errorn

func init() {
	register(ErrSuccess, 200, "OK")
	register(ErrDatabase, 500, "database errorn")
}
