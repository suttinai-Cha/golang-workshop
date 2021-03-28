package confReading

type IConfigReader interface {
	LoadConfig()
	GetConf() *Configuration
}
