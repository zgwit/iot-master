package db

var models []any

func CreateModel(m ...any) {
	models = append(models, m...)
}
