package postgres

import (
	"context"
	"log"
	"misis_baholar/models"
	repoi "misis_baholar/storage/repoI"

	"github.com/jackc/pgx/v5"
)

type grade struct {
	conn *pgx.Conn
}

func NewGrade(conn *pgx.Conn) repoi.GradeRepoI {
	return &grade{conn: conn}
}

func (g *grade) CreateGrade(ctx context.Context, req *models.Grade) error {
	query := `
		INSERT INTO grades(
			grade_id,
			grade_name,
			grade_value,
			grade_date,
			subject_id,
			group_id,
			student_id,
			created_at,
			updated_at
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8, $9
		)
	`
	_, err := g.conn.Exec(
		ctx, query,
		req.GradeID,
		req.GradeName,
		req.GradeValue,
		req.GradeDate,
		req.SubjectID,
		req.GroupID,
		req.StudentID,
		req.CreatedAt,
		req.UpdatedAt,
	)
	if err != nil {
		log.Println("Error in CreateGrade:", err)
		return err
	}
	log.Println("CreateGrade successfully")
	return nil
}

func (g *grade) GetGradeList(ctx context.Context, req *models.GetListReq) (*models.GetGradesList, error) {
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
			grade_id,
			grade_name,
			grade_value,
			grade_date,
			subject_id,
			group_id,
			student_id,
			created_at,
			updated_at
		FROM
			grades
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
		log.Println("Error in GetGradeList:", err)
		return nil, err
	}
	defer rows.Close()

	var grades []*models.Grade
	for rows.Next() {
		var grade models.Grade
		err := rows.Scan(
			&grade.GradeID,
			&grade.GradeName,
			&grade.GradeValue,
			&grade.GradeDate,
			&grade.SubjectID,
			&grade.GroupID,
			&grade.StudentID,
			&grade.CreatedAt,
			&grade.UpdatedAt,
		)
		if err != nil {
			log.Println("Error in scanning grade:", err)
			return nil, err
		}
		grades = append(grades, &grade)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error in rows:", err)
		return nil, err
	}

	return &models.GetGradesList{
		Grades: grades,
		Count:  len(grades),
	}, nil
}

func (g *grade) GetGradeByID(ctx context.Context, id string) (*models.Grade, error) {
	var grade models.Grade
	query := `
		SELECT 
			grade_id,
			grade_name,
			grade_value,
			grade_date,
			subject_id,
			group_id,
			student_id,
			created_at,
			updated_at
		FROM
			grades
		WHERE
			grade_id = $1
	`
	err := g.conn.QueryRow(ctx, query, id).Scan(
		&grade.GradeID,
		&grade.GradeName,
		&grade.GradeValue,
		&grade.GradeDate,
		&grade.SubjectID,
		&grade.GroupID,
		&grade.StudentID,
		&grade.CreatedAt,
		&grade.UpdatedAt,
	)
	if err != nil {
		log.Println("Error in GetGradeByID:", err)
		return nil, err
	}

	return &grade, nil
}

func (g *grade) UpdateGrade(ctx context.Context, req *models.Grade) error {
	query := `
		UPDATE
			grades
		SET
			grade_name = $1,
			grade_value = $2,
			grade_date = $3,
			subject_id = $4,
			group_id = $5,	
			student_id = $6,
			created_at = $7,
			updated_at = $8
		WHERE
			grade_id = $9
	`
	_, err := g.conn.Exec(
		ctx, query,
		req.GradeName,
		req.GradeValue,
		req.GradeDate,
		req.SubjectID,
		req.GroupID,
		req.StudentID,
		req.CreatedAt,
		req.UpdatedAt,
		req.GradeID,
	)
	if err != nil {
		log.Println("Error in UpdateGrade:", err)
		return err
	}
	return nil
}

func (g *grade) DeleteGrade(ctx context.Context, id string) error {
	query := `
		DELETE FROM 
			grades
		WHERE
			grade_id = $1
	`
	_, err := g.conn.Exec(ctx, query, id)
	if err != nil {
		log.Println("Error in DeleteGrade:", err)
		return err
	}
	return nil
}
