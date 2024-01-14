package device

func Boot() error {

	subscribeProperty()

	subscribeEvent()

	subscribeOnline()

	return nil
}
