package db

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	_ "modernc.org/sqlite"
)

func TestSQLite_Exec(t *testing.T) {
	var err error
	Info = DBInfo{
		DatabaseType: SQLITE,
		DatabaseName: "books",
		TableName:    "books",
		FilePath:     "../books.db",
	}

	type args struct {
		sql       string
		colValues []interface{}
		options   string
	}
	tests := []struct {
		name    string
		dsn     string
		d       *SQLite
		args    args
		want    int64
		want1   int64
		wantErr bool
	}{
		{
			name: "SQLITE",
			dsn:  Info.FilePath,
			d:    &SQLite{},
			args: args{
				sql:       "INSERT INTO " + Info.TableName + " (TITLE,AUTHOR) VALUES (?,?)",
				colValues: []interface{}{"test2", "test3"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.d.dsn = tt.dsn
			Con, err = tt.d.connect()
			if err != nil {
				os.Remove(tt.dsn)
				t.Error(err)
				return
			}
			defer os.Remove(tt.dsn)
			defer Con.Close()

			err = tt.d.CreateTable()
			if err != nil {
				os.Remove(tt.dsn)
				t.Error(err)
				return
			}

			count, _, err := tt.d.Exec(tt.args.sql, tt.args.colValues, tt.args.options)
			if err != nil {
				t.Error(err)
				return
			}

			require.Equal(t, int64(1), count)
		})
	}
}