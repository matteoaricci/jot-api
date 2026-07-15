package version

import "runtime"

var (
	GitCommit = "unknown"
	Version   = "0.1.0-dev"
	BuildTime = "unknown"
)

type Info struct {
	GitCommit string `json:"gitCommit"`
	Version   string `json:"version"`
	BuildTime string `json:"buildTime"`
	GoVersion string `json:"goVersion"`
}

func GetInfo() Info {
	return Info{
		GitCommit: GitCommit,
		Version:   Version,
		BuildTime: BuildTime,
		GoVersion: runtime.Version(),
	}
}
