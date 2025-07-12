package system

type Unit interface {
	SetConfig(config map[string]string)
	Start()
	GetValue(key string) string
	SetValue(key, value string)
	Stop()
}
