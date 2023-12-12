package flags

import (
	"github.com/spf13/pflag"
)

type Uac struct {
	UriFormat string
	IsSdp     bool
}

func (f *Uac) SetFlags(cmd *pflag.FlagSet) {
	if uriFormat, err := cmd.GetString("uri"); err == nil {
		f.UriFormat = uriFormat
	}
	if isSdp, err := cmd.GetBool("sdp"); err == nil {
		f.IsSdp = isSdp
	}
}
