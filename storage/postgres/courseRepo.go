package postgres

import (
	"context"
	"log"
	"misis_baholar/models"
	repoi "misis_baholar/storage/repoI"

	"github.com/jackc/pgx/v5"
)

type courseRepo struct {
	conn *pgx.Conn
}

func NewCourseRepo(db *pgx.Conn) repoi.CourseRepoI {
	return &courseRepo{conn: db}
}

func (c *courseRepo) CreateCourse(ctx context.Context, req *models.Course) error {

	query := `
		INSERT INTO courses(
			course_id,
			course_name,
			created_at,
			updated_at
		)VALUES(
			$1,$2,$3,$4
		)
	`
	_, err := c.conn.Exec(
		ctx, query,
		req.CourseID,
		req.CourseName,
		req.CreatedAt,
		req.UpdatedAt,
	)
	if err != nil {
		log.Printf("error on CreatteCourseRepo: %v", err)
		return err
	}

	return nil
}
func (c *courseRepo) GetCourseList(ctx context.Context, req *models.GetListReq) (*models.GetCoursesList, error) {

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
			course_id,
			course_name,
			created_at,
			updated_at
		FROM
			courses
		LIMIT
			$1
		OFFSET
			$2
	`
	rows, err := c.conn.Query(ctx, query, limit, offset)
	if err != nil {
		log.Default().Println("error on GetCourseListRepo:", err)
		return nil, err
	}
	defer rows.Close()
	var courses []*models.Course
	for rows.Next() {
		var course models.Course
		if err := rows.Scan(
			&course.CourseID,
			&course.CourseName,
			&course.CreatedAt,
			&course.UpdatedAt,
		); err != nil {
			log.Println("error on GetCourseList:", err)
			return nil, err
		}
		courses = append(courses, &course)
	}
	if err := rows.Err(); err != nil {
		log.Println("error on rows in GetCourseList:", err)
		return nil, err
	}
	return &models.GetCoursesList{
		Courses: courses,
		Count:   len(courses),
	}, nil

}
func (c *courseRepo) GetCourseByID(ctx context.Context, id string) (*models.Course, error) {
	query := `
		SELECT 
			course_id,
			course_name,
			created_at,
			updated_at
		FROM
			courses
		WHERE 
			course_id = $1
	`
	var course models.Course
	if err := c.conn.QueryRow(ctx, query, id).Scan(
		&course.CourseID,
		&course.CourseName,
		&course.CreatedAt,
		&course.UpdatedAt,
	); err != nil {
		log.Println("error on GetCourseById:", err)
		return nil, err
	}
	return &course, nil
}
func (c *courseRepo) UpdateCourse(ctx context.Context, req *models.Course) error {

	query := `
		UPDATE 
			courses
		SET
			course_name = $1,
		WHERE
			course_id = $2
	`
	_, err := c.conn.Exec(
		ctx, query,
		req.CourseName,
		req.CourseID,
	)
	if err != nil {
		log.Println("error on UpdateCourseById", err)
		return err
	}
	return nil
}
func (c *courseRepo) DeleteCourse(ctx context.Context, id string) error {
	query := `
		DELETE FROM
			courses
		WHERE 
			course_id = $1
	`
	_, err := c.conn.Exec(ctx, query, id)
	if err != nil {
		log.Println("error on DeleteCourse:", err)
		return err
	}
	return nil
}
