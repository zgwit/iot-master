package device

func Open() error {

	subscribeProperty()

	subscribeEvent()

	return nil
}

func Close() {

}
