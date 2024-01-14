package internal

func Boot() error {

	subscribeProperty()

	subscribeEvent()

	subscribeOnline()

	return nil
}
