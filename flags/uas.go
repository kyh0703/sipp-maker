package flags

import "github.com/spf13/pflag"

type Uas struct {
	UriFormat string
	Code      int
}

func (f *Uas) SetFlags(cmd *pflag.FlagSet) {
	if code, err := cmd.GetInt("code"); err != nil {
		f.Code = code
	}
}
