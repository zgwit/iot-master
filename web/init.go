package web

func init() {
	err := Load()
	if err != nil {
		_ = Store()
	}
}
