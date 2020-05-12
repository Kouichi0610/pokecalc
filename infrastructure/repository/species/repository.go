package species

import (
	"pokecalc/domain/stats"
	"pokecalc/domain/types"
)

type Repository interface {
	Load(id uint) (*stats.Species, *types.Types, error)
}
