// Package spike is a throwaway module for the linkclean M2 Dependabot spike.
package spike

import _ "embed"

//go:embed data/clean-urls.json
var Data []byte
