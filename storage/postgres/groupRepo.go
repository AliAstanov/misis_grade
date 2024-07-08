package postgres

import (
	"context"
	"log"
	"misis_baholar/models"
	repoi "misis_baholar/storage/repoI"

	"github.com/jackc/pgx/v5"
)

type group struct {
	conn *pgx.Conn
}

func NewGroup(conn *pgx.Conn) repoi.GroupRepoI {
	return &group{conn: conn}
}
func (g *group) CreateGroup(ctx context.Context, req *models.Group) error {
	query := `
		INSERT INTO groups(
			group_id,
			group_name,
			course_id,
			created_at,
			updated_at		
		)VALUES(
			$1,$2,$3,$4,$5
			)
	`
	_, err := g.conn.Exec(
		ctx, query,
		req.GroupID,
		req.GroupName,
		req.CourseID,
		req.CreatedAt,
		req.UpdatedAt,
	)
	if err != nil {
		log.Println("error with CreateGroup on postgres package:", err)
		return err
	}
	return nil
}

func (g *group) GetGroupList(ctx context.Context, req *models.GetListReq) (*models.GetGroupsList, error) {

	limit := req.Limit
	if limit == 0 {
		limit = DefaultLimit
	}

	page := req.Page
	if page == 0 {
		page = DefaultPage
	}
	offset := (page - 1) * limit
	query := `
		SELECT 
			group_id,
			group_name,
			course_id,
			created_at,
			updated_at
		FROM
			groups
		LIMIT 
			$1
		OFFSET 
			$2
	`
	rows, err := g.conn.Query(
		ctx, query,
		limit,
		offset,
	)
	if err != nil {
		log.Println("error with GetGroupList on postgres package:", err)
		return nil, err
	}
	defer rows.Close()
	var groups []*models.Group
	for rows.Next() {
		var group models.Group
		if err := rows.Scan(
			&group.GroupID,
			&group.GroupName,
			&group.CourseID,
			&group.CreatedAt,
			&group.UpdatedAt,
		); err != nil {
			log.Println("error with GetGroupList on scan:", err)
			return nil, err
		}
		groups = append(groups, &group)
	}
	if err := rows.Err(); err != nil {
		log.Println("error with GetGrouplist", err)
		return nil, err
	}
	return &models.GetGroupsList{
		Groups: groups,
		Count:  len(groups),
	}, nil
}
func (g *group) GetGroupByID(ctx context.Context, id string) (*models.Group, error) {
	query := `
		SELECT 
			group_id,
			group_name,
			course_id,
			created_at,
			updated_at
		FROM
			groups
		WHERE
			group_id = $1	
	`
	var group models.Group
	err := g.conn.QueryRow(ctx, query, id).Scan(
		&group.GroupID,
		&group.GroupName,
		&group.CourseID,
		&group.CreatedAt,
		&group.UpdatedAt,
	)
	if err != nil {
		log.Println("error with GetGroupById:", err)
		return nil, err
	}
	return &group, nil
}

func (g *group) UpdateGroup(ctx context.Context, req *models.Group) error {
	query := `
		UPDATE
			groups
		SET
			group_name = $1,
			course_id = $2
		WHERE
		 group_id = $3
	`
	_, err := g.conn.Exec(
		ctx, query,
		req.GroupName,
		req.CourseID,
		req.GroupID,
	)
	if err != nil {
		log.Println("error with UpdategroupRepo:", err)
		return err
	}
	return nil
}

func (g *group) DeleteGroup(ctx context.Context, id string) error {
	query := `
		DELETE FROM 
			groups
		WHERE 
			group_id = $1
	`
	_, err := g.conn.Exec(ctx, query, id)
	if err != nil {
		log.Println("error with DeletegroupRepo:", err)
		return err
	}
	return nil
}
