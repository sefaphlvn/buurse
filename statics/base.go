package statics

type Paths string

const (
	BootstrapFolderPath  Paths = "/usr/suubar/svc/%s"
	FullHotRestarterPath Paths = "/usr/suubar/svc/%s/hotrestarter.py"
	FullBootstrapPath    Paths = "/usr/suubar/svc/%s/bootstrap.json"
	FullScriptPath       Paths = "/usr/suubar/svc/%s/script.sh"
	FullSystemDPath      Paths = "/usr/lib/systemd/system/svc_%s.service"
	LogFolderPath        Paths = "/var/log/suubar/"
)

func (kt Paths) String() string {
	return string(kt)
}
