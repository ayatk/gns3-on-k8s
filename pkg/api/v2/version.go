package v2

import (
	"fmt"
	"net/http"
	"runtime"

	"github.com/ayatk/gns3-on-k8s/pkg"
	"github.com/labstack/echo/v4"
)

type VersionInfo struct {
	Local bool `json:"local"`
	// GNS3 Version
	Version       string `json:"version"`
	ServerVersion string `json:"server_version"`
	GoVersion     string `json:"go_version"`
	GitCommit     string `json:"git_commit"`
	Compiler      string `json:"compiler"`
	Platform      string `json:"platform"`
}

func Version(c echo.Context) error {
	v := VersionInfo{
		Local:         false,
		Version:       pkg.Gns3Version,
		ServerVersion: pkg.Version,
		GitCommit:     pkg.GitCommit,
		GoVersion:     runtime.Version(),
		Compiler:      runtime.Compiler,
		Platform:      fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH),
	}
	return c.JSONPretty(http.StatusOK, v, "  ")
}
