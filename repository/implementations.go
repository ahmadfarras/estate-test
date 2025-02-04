package repository

import (
	"context"

	"github.com/sirupsen/logrus"
)

func (r *Repository) GetTestById(ctx context.Context, input GetTestByIdInput) (output GetTestByIdOutput, err error) {
	err = r.Db.QueryRowContext(ctx, "SELECT name FROM test WHERE id = $1", input.Id).Scan(&output.Name)
	if err != nil {
		return
	}
	return
}

func (r *Repository) CreateEstate(ctx context.Context, input CreateEstateInput) error {
	_, err := r.Db.ExecContext(ctx, "INSERT INTO estate (id, length, width) VALUES ($1, $2, $3)", input.ID, input.Length, input.Width)
	if err != nil {
		logrus.WithContext(ctx).WithError(err).Error("failed to create estate")
		return err
	}
	return nil
}

func (r *Repository) GetEstateById(ctx context.Context, input GetEstateByIdInput) (output GetEstateByIdOutput, err error) {
	err = r.Db.QueryRowContext(ctx, "SELECT id, length, width, created_at, updated_at FROM estate WHERE id = $1", input.ID).
		Scan(&output.ID, &output.Length, &output.Width, &output.CreatedAt, &output.UpdatedAt)
	if err != nil {
		logrus.WithContext(ctx).WithError(err).Error("failed to get estate by id")
		return
	}
	return
}

func (r *Repository) CreateTree(ctx context.Context, input CreateTreeInput) error {
	_, err := r.Db.ExecContext(ctx, "INSERT INTO tree (id, estate_id, height, x, y) VALUES ($1, $2, $3, $4, $5)", input.ID, input.EstateID, input.Height, input.X, input.Y)
	if err != nil {
		logrus.WithContext(ctx).WithError(err).Error("failed to create tree")
		return err
	}
	return nil
}
