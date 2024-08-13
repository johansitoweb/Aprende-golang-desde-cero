import (
    "database/sql"
    _ "github.com/lib/pq" // Driver para PostgreSQL
)

func conectarDB() (*sql.DB, error) {
    db, err := sql.Open("postgres", "user=usuario dbname=mi_db sslmode=disable")
    if err != nil {
        return nil, err
    }
    return db, nil
}
