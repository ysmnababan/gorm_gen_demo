package handler

import (
	"gorm_demo/src/query"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func newMockGorm(t *testing.T) (*gorm.DB, sqlmock.Sqlmock, func()) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	dialector := postgres.New(postgres.Config{
		Conn:       db,
		DriverName: "postgres",
	})
	gormDB, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}
	return gormDB, mock, func() { db.Close() }
}
func TestListNav(t *testing.T) {
	db, mock, close := newMockGorm(t)
	defer close()
	q := query.Use(db)

	n := q.Navigation
	// Prepare expected rows (columns must match what GORM scans into your struct)
	now := time.Now()
	rows := sqlmock.NewRows([]string{"id", "created_at"}).
		AddRow(1, now.UnixMilli())

	// Set expectation before executing the GORM query.
	// Use a regexp that roughly matches the SQL GORM will produce.
	mock.ExpectQuery(`SELECT \* FROM "navigation" WHERE "navigation"."created_at" IS NOT NULL`).
		WillReturnRows(rows)
	out, err := n.Where(n.CreatedAt.IsNotNull()).Find()
	_ = db
	_ = mock
	require.NoError(t, err)
	assert.NotNil(t, out)
}
