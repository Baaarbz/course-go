package postgres

type sqlBuilderCourse struct {
	ID          string `db:"id"`
	Name        string `db:"name"`
	Description string `db:"description"`
}

type sqlBuilderCreateCourse struct {
	Name        string `db:"name"`
	Description string `db:"description"`
}
