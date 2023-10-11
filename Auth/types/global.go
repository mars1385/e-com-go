package types

type Config struct {
	Port string
}

type Application struct {
	Config Config
}

type SuccessResponse struct {
	message string
	data    interface{}
}
