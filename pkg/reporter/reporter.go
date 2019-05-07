package reporter

import (
	"base/pkg/cerror"
	"fmt"
	"github.com/getsentry/raven-go"
	"math/rand"
	"runtime"
	"time"
)

type Config struct {
	DSN          string
	OnFailed     func(err error)
	ProjectName  string
	Compilation  string
	Version      string
	BuildDateStr string
	Environment  string
}

var _cfg Config
var isDisabled = true

func Init(cfg Config) {
	if cfg.DSN == "" {
		return
	}
	_cfg = cfg
	rand.Seed(time.Now().UTC().UnixNano())
	err := raven.SetDSN(_cfg.DSN)
	if err != nil {
		if _cfg.OnFailed != nil {
			return
		}
		_cfg.OnFailed(err)
	}
	isDisabled = false
}

func ReportError(err cerror.Error, extra map[string]interface{}) chan error {
	if isDisabled {
		return nil
	}
	p := raven.Packet{
		Message:  err.Text,
		Project:  _cfg.ProjectName,
		Level:    raven.Severity(cerror.ERROR),
		Platform: runtime.GOOS,
		Release: fmt.Sprintf(
			"Version: %s | Compilation: %s | TimeCompilation: %s",
			_cfg.Version, _cfg.Compilation, _cfg.BuildDateStr,
		),
		Environment: string(_cfg.Environment),
		Extra:       map[string]interface{}{"extra": extra, "error": err},
	}
	_, errOut := raven.Capture(&p, err.Meta)
	return errOut
}
