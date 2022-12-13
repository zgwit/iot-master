package broker

type Auth struct{}

func (a *Auth) Authenticate(user, password []byte) bool {
	//TODO 查询网关
	//TODO 查询插件
	//TODO 查询用户

	return true
}

func (a *Auth) ACL(user []byte, topic string, write bool) bool {
	return true
}
