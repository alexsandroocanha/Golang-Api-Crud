package db

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetDBConnection(t *testing.T) {
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_PORT", "5432")
	os.Setenv("DB_USER", "postgres")
	os.Setenv("DB_PASSWORD", "secret")
	os.Setenv("DB_NAME", "testdb")
	os.Setenv("DB_SSLMODE", "disable")
	defer os.Clearenv()

	conn := GetDBConnection()

	assert.Contains(t, conn, "host=localhost")
	assert.Contains(t, conn, "dbname=testdb")
	assert.Contains(t, conn, "sslmode=disable")
}
