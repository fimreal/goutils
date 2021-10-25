package ezap

var (
	Ezap *Logger
	Conf *Config
)

func init() {
	Ezap = New()
	Conf = Ezap.Config
	Ezap.ApplyConfig()
}
