package postgres

import (
	"context"
	"log"
	"misis_baholar/models"
	repoi "misis_baholar/storage/repoI"
	"github.com/jackc/pgx/v5"
)

type subject struct {
	conn *pgx.Conn
}

func NewSubject(conn *pgx.Conn) repoi.SubjectRepoI {
	return &subject{conn: conn}
}

func (s *subject) CreateSubject(ctx context.Context, req *models.Subject) error {
	query := `
		INSERT INTO subjects(
			subject_id,
			subject_name,
			group_id,
			teacher_id,
			created_at,
			updated_at			
		)VALUES(
			$1,$2,$3,$4,$5,$6	
		)
	`
	_, err := s.conn.Exec(
		ctx, query,
		req.SubjectID,
		req.SubjectName,
		req.GroupID,
		req.TeacherID,
		req.CreatedAt,
		req.UpdatedAt,
	)
	if err != nil {
		log.Println("error with CreateSubjectRepo:", err)
		return err
	}
	return nil
}
func (s *subject) GetSubjectList(ctx context.Context, req *models.GetListReq) (*models.GetSubjectsList, error) {
	
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
			subject_id,
			subject_name,
			group_id,
			teacher_id,
			created_at,
			updated_at
		FROM
			subjects
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
		log.Println("error with GetSubjectListRepo:", err)
		return nil, err
	}
	defer rows.Close()

	var subjects []*models.Subject
	for rows.Next() {
		var subject models.Subject
		err := rows.Scan(
			&subject.SubjectID,
			&subject.SubjectName,
			&subject.GroupID,
			&subject.TeacherID,
			&subject.CreatedAt,
			&subject.UpdatedAt,
		)
		if err != nil {
			log.Println("error with subjectScan:", err)
			return nil, err
		}
		subjects = append(subjects, &subject)
	}
	return &models.GetSubjectsList{
		Subjects: subjects,
		Count:    len(subjects),
	}, nil
}
func (s *subject) GetSubjectByID(ctx context.Context, id string) (*models.Subject, error) {
	query := `
		SELECT 
			subject_id,
			subject_name,
			group_id,
			teacher_id,
			created_at,
			updated_at
		FROM
			subjects
		WHERE 
			subject_id = $1

	`
	var subject models.Subject
	err := s.conn.QueryRow(ctx, query, id).Scan(
		&subject.SubjectID,
		&subject.SubjectName,
		&subject.GroupID,
		&subject.TeacherID,
		&subject.CreatedAt,
		&subject.UpdatedAt,
	)
	if err != nil {
		log.Println("error With GetSubjectByIdScan:", err)
		return nil, err
	}

	return &subject, nil
}
func (s *subject) UpdateSubject(ctx context.Context, req *models.Subject) error {
	query := `
		UPDATE
			subjects
		SET
			subject_name = $1,
			group_id = $2,
			teacher_id = $3,
			updated_at = $4
		WHERE 
			subject_id = $5
	`
	_, err := s.conn.Exec(
		ctx, query,
		req.SubjectName,
		req.GroupID,
		req.TeacherID,
		req.UpdatedAt,
		req.SubjectID,
	)
	if err != nil {
		log.Println("error with UpdateSubjectRepo:", err)
		return err
	}

	return nil
}
func (s *subject) DeleteSubject(ctx context.Context, id string) error {
	query := `
		DELETE FROM 
			subjects
		WHERE
			subject_id = $1
	`
	_, err := s.conn.Exec(ctx, query, id)
	if err != nil {
		log.Println("error with DeleteSubjectRepo:", err)
		return err
	}
	return nil
}
