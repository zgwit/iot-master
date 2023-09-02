package db

func init() {
	err := Load()
	if err != nil {
		_ = Store()
	}
}
