package v2

import (
	"fmt"
	"net/http"
	"runtime"

	"github.com/ayatk/gns3-on-k8s/pkg"
	"github.com/labstack/echo/v4"
)

type Version struct {
	Local bool `json:"local"`
	// GNS3 Version
	Version       string `json:"version"`
	ServerVersion string `json:"server_version"`
	GoVersion     string `json:"go_version"`
	GitCommit     string `json:"git_commit"`
	Compiler      string `json:"compiler"`
	Platform      string `json:"platform"`
}

type CheckVersion struct {
	Version string `json:"version"`
}

func VersionHandler(c echo.Context) error {
	v := Version{
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

func CheckVersionHandler(c echo.Context) error {
	body := new(CheckVersion)
	if err := c.Bind(body); err != nil {
		return err
	}

	if body.Version != pkg.Gns3Version {
		return echo.NewHTTPError(
			http.StatusConflict,
			fmt.Sprintf("Client version %s is not the same as server version %s", body.Version, pkg.Gns3Version))
	}

	return c.JSONPretty(http.StatusOK, map[string]interface{}{"version": body.Version}, "  ")
}
