package species

import (
	"pokecalc/domain/stats"
	"pokecalc/domain/types"
)

type Repository interface {
	Load(name string) (*stats.Species, *types.Types, error)
}
