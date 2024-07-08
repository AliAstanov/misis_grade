package postgres

import (
	"context"
	"log"
	"misis_baholar/models"
	repoi "misis_baholar/storage/repoI"

	"github.com/jackc/pgx/v5"
)

const (
	DefaultLimit = 10
	DefaultPage  = 1
)

type studentRepo struct {
	conn *pgx.Conn
}

func NewStudentRepo(conn *pgx.Conn) repoi.StudentRepoI {
	return &studentRepo{conn: conn}
}

func (s *studentRepo) CreateStudent(ctx context.Context, req *models.Student) error {
	query := `
		INSERT INTO students(
			student_id,
			name,  
			surname,
			group_id,
			created_at,
			updated_at			
		) VALUES (
			$1, $2, $3, $4, $5, $6
		)
	`
	_, err := s.conn.Exec(
		ctx, query,
		req.StudentID,
		req.Name,
		req.Surname,
		req.GroupID,
		req.CreatedAt,
		req.UpdatedAt,
	)
	if err != nil {
		log.Println("Error in CreateStudent:", err)
		return err
	}

	return nil
}

func (s *studentRepo) GetStudentList(ctx context.Context, req *models.GetListReq) (*models.GetStudentsList, error) {
	var students []*models.Student

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
			student_id,
			name,  	
			surname,
			group_id,	
			created_at,
			updated_at
		FROM 
			students
		LIMIT 
			$1
		OFFSET
			$2				
	`
	rows, err := s.conn.Query(
		ctx, query,
		limit,
		offset,
	)
	if err != nil {
		log.Println("Error in GetStudentList:", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var student models.Student
		if err := rows.Scan(
			&student.StudentID,
			&student.Name,
			&student.Surname,
			&student.GroupID,
			&student.CreatedAt,
			&student.UpdatedAt,
		); err != nil {
			log.Println("Error in scanning student row:", err)
			return nil, err
		}

		students = append(students, &student)
	}
	if err := rows.Err(); err != nil {
		log.Println("Error on rows in GetStudentList:", err)
		return nil, err
	}

	return &models.GetStudentsList{
		Students: students,
		Count:    len(students),
	}, nil
}

func (s *studentRepo) GetStudentByID(ctx context.Context, id string) (*models.Student, error) {
	query := `
		SELECT 
			student_id,
			name,  	
			surname,
			group_id,	
			created_at,
			updated_at
		FROM
			students
		WHERE
			student_id = $1
	`
	var student models.Student

	err := s.conn.QueryRow(
		ctx, query, id,
	).Scan(
		&student.StudentID,
		&student.Name,
		&student.Surname,
		&student.GroupID,
		&student.CreatedAt,
		&student.UpdatedAt,
	)
	if err != nil {
		log.Println("Error in GetStudentByID:", err)
		return nil, err
	}
	return &student, nil
}

func (s *studentRepo) UpdateStudent(ctx context.Context, req *models.Student) error {
	query := `
		UPDATE 
			students
		SET
			name = $1,  	
			surname = $2,	
			group_id = $3,	
			created_at = $4,
			updated_at = $5
		WHERE
			student_id = $6			
	`
	_, err := s.conn.Exec(
		ctx, query,
		req.Name,
		req.Surname,
		req.GroupID,
		req.CreatedAt,
		req.UpdatedAt,
		req.StudentID,
	)
	if err != nil {
		log.Println("Error in UpdateStudent:", err)
		return err
	}
	return nil
}

func (s *studentRepo) DeleteStudent(ctx context.Context, id string) error {
	query := `
		DELETE FROM
			students
		WHERE
			student_id = $1
	`
	_, err := s.conn.Exec(ctx, query, id)
	if err != nil {
		log.Println("Error in DeleteStudent:", err)
		return err
	}
	return nil
}
