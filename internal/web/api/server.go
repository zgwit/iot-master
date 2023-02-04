package api

func afterServerCreate(data interface{}) error {
	//server := data.(*model.Server)

	//TODO start server
	return nil
}

func afterServerUpdate(data interface{}) error {
	//server := data.(*model.Server)

	//TODO restart server
	return nil
}

func afterServerDelete(id interface{}) error {
	//gid := id.(string)

	//todo stop server
	return nil
}
