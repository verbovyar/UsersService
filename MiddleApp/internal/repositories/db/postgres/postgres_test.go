package db

import (
	"github.com/pashagolub/pgxmock"
	"github.com/stretchr/testify/require"
	"regexp"
	"testing"
)

func TestUsersRepository_Delete(t *testing.T) {
	testTable := []struct {
		name        string
		id          uint32
		expectError error
		expectId    uint32
	}{
		{
			name:        "OK",
			id:          1,
			expectError: nil,
			expectId:    1,
		},
	}
	mock, err := pgxmock.NewPool()
	if err != nil {
		t.Fatal(err)
	}
	defer mock.Close()

	repo := New(mock)

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			mock.ExpectQuery(regexp.QuoteMeta("DELETE FROM Users WHERE Id = $1")).WithArgs().WillReturnRows()
			err, id := repo.Delete(1)
			require.NoError(t, err)
			require.Equal(t, uint32(1), id)
		})
	}
}
