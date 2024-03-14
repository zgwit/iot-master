package device

func Open() error {

	subscribeProperty()

	subscribeEvent()

	subscribeOnline()

	return nil
}

func Close() {

}
