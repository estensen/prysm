package shuffle

import types "github.com/prysmaticlabs/eth2-types"

// TestCase --
type TestCase struct {
	Seed    string                 `yaml:"seed"`
	Count   uint64                 `yaml:"count"`
	Mapping []types.ValidatorIndex `yaml:"mapping"`
}
