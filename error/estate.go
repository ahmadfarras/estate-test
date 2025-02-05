package error

import "fmt"

var (
	ErrEstateNotFound            = fmt.Errorf("estate not found")
	ErrTreePositionOutOfBoundary = fmt.Errorf("tree position is out of estate boundary")
	ErrTreePositionNegative      = fmt.Errorf("tree position must be positive")
	ErrTreeHeightNegative        = fmt.Errorf("tree height must be positive")
	ErrTreeAlreadyPlanted        = fmt.Errorf("tree already planted at this position")
)
