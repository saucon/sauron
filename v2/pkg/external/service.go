package external

import (
	notify_error "github.com/saucon/sauron/v2/pkg/external/gspace_chat"
	"github.com/saucon/sauron/v2/pkg/log/logconfig"
)

type External struct {
	Gchat *notify_error.Client
}

func ProvideExternalSvc(config *logconfig.GspaceChat) *External {
	return &External{
		Gchat: notify_error.New(config),
	}
}
