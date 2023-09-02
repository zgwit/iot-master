package config

type Options interface {
	Load() error
	Store() error
	FromEnv() error
	ToEnv() []string
}
