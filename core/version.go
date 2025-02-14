package core

type VersionInfo struct {
	Version  string `json:"version"`
	Revision string `json:"revision"`
	Build    string `json:"build"`
}

var Version VersionInfo = VersionInfo{
	Version:  "Undefined",
	Revision: "Undefined",
	Build:    "Undefined",
}

func SetVersion(version, revision, build string) {
	Version = VersionInfo{
		Version:  version,
		Revision: revision,
		Build:    build,
	}
}
