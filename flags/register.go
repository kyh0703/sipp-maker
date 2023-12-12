package flags

import "github.com/spf13/pflag"

type Register struct {
	UriFormat string
}

func (f *Register) SetFlags(cmd *pflag.FlagSet) {
	if uriFormat, err := cmd.GetString("uri"); err == nil {
		f.UriFormat = uriFormat
	}
}
