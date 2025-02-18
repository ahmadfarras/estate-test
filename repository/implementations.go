package repository

import (
	"context"
	"database/sql"

	"github.com/SawitProRecruitment/UserService/model"
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

func (r *Repository) GetEstateById(ctx context.Context, input GetEstateByIdInput) (output *GetEstateByIdOutput, err error) {
	output = &GetEstateByIdOutput{}
	err = r.Db.QueryRowContext(ctx, "SELECT id, length, width, created_at, updated_at FROM estate WHERE id = $1", input.ID).
		Scan(&output.ID, &output.Length, &output.Width, &output.CreatedAt, &output.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		logrus.WithContext(ctx).WithError(err).Error("failed to get estate by id")
		return nil, err
	}
	return output, nil
}

func (r *Repository) CreateTree(ctx context.Context, input CreateTreeInput) error {
	_, err := r.Db.ExecContext(ctx, "INSERT INTO tree (id, estate_id, height, x, y) VALUES ($1, $2, $3, $4, $5)", input.ID, input.EstateID, input.Height, input.X, input.Y)
	if err != nil {
		logrus.WithContext(ctx).WithError(err).Error("failed to create tree")
		return err
	}
	return nil
}

func (r *Repository) GetAllTreesByEstateID(ctx context.Context, input GetAllTreesByEstateIDInput) (output GetAllTreesByEstateIDOutput, err error) {
	rows, err := r.Db.QueryContext(ctx, "SELECT id, estate_id, height, x, y FROM tree WHERE estate_id = $1", input.EstateID)
	if err != nil {
		logrus.WithContext(ctx).WithError(err).Error("failed to get all trees by estate id")
		return
	}
	defer rows.Close()

	for rows.Next() {
		var tree model.Tree
		err = rows.Scan(&tree.ID, &tree.EstateID, &tree.Height, &tree.X, &tree.Y)
		if err != nil {
			logrus.WithContext(ctx).WithError(err).Error("failed to scan tree")
			return
		}
		output.Trees = append(output.Trees, tree)
	}

	output.Total = len(output.Trees)
	return
}

func (r *Repository) GetTreeByEstateIDAndCoordinate(ctx context.Context, input GetTreeByEstateIDAndCoordinateInput,
) (output *GetTreeByEstateIDAndCoordinateOutput, err error) {
	output = &GetTreeByEstateIDAndCoordinateOutput{}
	err = r.Db.QueryRowContext(ctx, "SELECT id, estate_id, height, x, y FROM tree WHERE estate_id = $1 AND x = $2 AND y = $3", input.EstateID, input.X, input.Y).
		Scan(&output.Tree.ID, &output.Tree.EstateID, &output.Tree.Height, &output.Tree.X, &output.Tree.Y)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		logrus.WithContext(ctx).WithError(err).Error("failed to get tree by estate id and coordinate")
		return nil, err
	}

	return output, nil
}
