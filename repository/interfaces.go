// This file contains the interfaces for the repository layer.
// The repository layer is responsible for interacting with the database.
// For testing purpose we will generate mock implementations of these
// interfaces using mockgen. See the Makefile for more information.
package repository

import "context"

type RepositoryInterface interface {
	GetTestById(ctx context.Context, input GetTestByIdInput) (output GetTestByIdOutput, err error)

	CreateEstate(ctx context.Context, input CreateEstateInput) error
	GetEstateById(ctx context.Context, input GetEstateByIdInput) (output *GetEstateByIdOutput, err error)
	CreateTree(ctx context.Context, input CreateTreeInput) error
	GetAllTreesByEstateID(ctx context.Context, input GetAllTreesByEstateIDInput) (output GetAllTreesByEstateIDOutput, err error)
	GetTreeByEstateIDAndCoordinate(ctx context.Context, input GetTreeByEstateIDAndCoordinateInput) (output *GetTreeByEstateIDAndCoordinateOutput, err error)
}
