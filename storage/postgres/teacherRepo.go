package postgres

import (
	"context"
	"log"
	"misis_baholar/models"
	repoi "misis_baholar/storage/repoI"

	"github.com/jackc/pgx/v5"
)

type teacherRepo struct {
	conn *pgx.Conn
}

func NewTeacherRepo(db *pgx.Conn) repoi.TeacherRepoI {
	return &teacherRepo{conn: db}
}

func (t *teacherRepo) CreateTeacher(ctx context.Context, req *models.Teacher) error {

	query := `
		INSERT INTO teachers(
			teacher_id,
			name,
			surname,
			email,
			password,
			created_at,
			updated_at
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7
		)
	`
	_, err := t.conn.Exec(
		ctx, query,
		req.TeacherID,
		req.Name,
		req.Surname,
		req.Email,
		req.Password,
		req.CreatedAt,
		req.UpdatedAt,
	)
	if err != nil {
		log.Println("Error in CreateTeacher:", err)
		return err
	}
	return nil
}

func (t *teacherRepo) GetTeacherList(ctx context.Context, req *models.GetListReq) (*models.GetTeachersList, error) {

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
			teacher_id,
			name,
			surname,
			email,
			created_at,
			updated_at
		FROM 
			teachers
		LIMIT 
			$1 
		OFFSET 
			$2
	`
	rows, err := t.conn.Query(ctx, query, limit, offset)
	if err != nil {
		log.Println("Error in GetTeacherList:", err)
		return nil, err
	}
	defer rows.Close()

	var teachers []*models.Teacher
	for rows.Next() {
		var teacher models.Teacher
		if err := rows.Scan(
			&teacher.TeacherID,
			&teacher.Name,
			&teacher.Surname,
			&teacher.Email,
			&teacher.CreatedAt,
			&teacher.UpdatedAt,
		); err != nil {
			log.Println("Error in scanning teacher:", err)
			return nil, err
		}
		teachers = append(teachers, &teacher)
	}

	if err = rows.Err(); err != nil {
		log.Println("Error in rows:", err)
		return nil, err
	}

	return &models.GetTeachersList{
		Teachers: teachers,
		Count:    len(teachers),
	}, nil
}

func (t *teacherRepo) GetTeacherByID(ctx context.Context, id string) (*models.Teacher, error) {
	query := `
		SELECT 
			teacher_id, 
			name, 
			surname, 
			email, 
			created_at, 
			updated_at
		FROM 
			teachers
		WHERE 
			teacher_id = $1
	`
	var teacher models.Teacher
	err := t.conn.QueryRow(ctx, query, id).Scan(
		&teacher.TeacherID,
		&teacher.Name,
		&teacher.Surname,
		&teacher.Email,
		&teacher.CreatedAt,
		&teacher.UpdatedAt,
	)
	if err != nil {
		log.Println("Error in GetTeacherByID:", err)
		return nil, err
	}
	return &teacher, nil
}

func (t *teacherRepo) UpdateTeacher(ctx context.Context, req *models.Teacher) error {

	query := `
		UPDATE 
			teachers
		SET 
			name = $1, 
			surname = $2, 
			email = $3, 
			password = $4, 
			updated_at = $5
		WHERE 
			teacher_id = $6
	`
	_, err := t.conn.Exec(
		ctx, query,
		req.Name,
		req.Surname,
		req.Email,
		req.Password,
		req.UpdatedAt,
		req.TeacherID)
	if err != nil {
		log.Println("Error in UpdateTeacher:", err)
		return err
	}
	return nil
}

func (t *teacherRepo) DeleteTeacher(ctx context.Context, id string) error {
	query := `
		DELETE FROM 
			teachers
		WHERE 
			teacher_id = $1
	`
	_, err := t.conn.Exec(ctx, query, id)
	if err != nil {
		log.Println("Error in DeleteTeacher:", err)
		return err
	}
	return nil
}
