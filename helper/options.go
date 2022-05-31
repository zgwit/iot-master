package helper

type Options map[string]interface{}

func (o Options) GetDefaultBool(name string, value bool) bool {
	if val, has := o[name]; has {
		if v, ok := val.(bool); ok {
			return v
		}
	}
	return value
}

func (o Options) GetDefaultInt(name string, value int) int {
	if val, has := o[name]; has {
		if v, ok := val.(int); ok {
			return v
		}
	}
	return value
}

func (o Options) GetDefaultString(name string, value string) string {
	if val, has := o[name]; has {
		if v, ok := val.(string); ok {
			return v
		}
	}
	return value
}

func (o Options) GetDefaultFloat(name string, value float64) float64 {
	if val, has := o[name]; has {
		if v, ok := val.(float64); ok {
			return v
		}
	}
	return value
}

func (o Options) GetBool(name string) (bool, bool) {
	if val, has := o[name]; has {
		if v, ok := val.(bool); ok {
			return v, true
		}
	}
	return false, false
}

func (o Options) GetInt(name string) (int, bool) {
	if val, has := o[name]; has {
		if v, ok := val.(int); ok {
			return v, true
		}
	}
	return 0, false
}

func (o Options) GetString(name string) (string, bool) {
	if val, has := o[name]; has {
		if v, ok := val.(string); ok {
			return v, true
		}
	}
	return "", false
}

func (o Options) GetFloat(name string) (float64, bool) {
	if val, has := o[name]; has {
		if v, ok := val.(float64); ok {
			return v, true
		}
	}
	return 0, false
}
