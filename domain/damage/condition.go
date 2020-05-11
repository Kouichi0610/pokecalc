package damage

import (
	"pokecalc/domain/stats"
	"pokecalc/domain/types"
)

type Condition struct {
	Types *types.Types
	Stats *stats.Stats
}
