package stub

import (
	"testing"

	"github.com/hasura/graphql-engine/cli/migrate/source"
	st "github.com/hasura/graphql-engine/cli/migrate/source/testing"
)

func Test(t *testing.T) {
	s := &Stub{}
	d, err := s.Open("")
	if err != nil {
		t.Fatal(err)
	}

	m := source.NewMigrations()
	m.Append(&source.Migration{Version: 1, Direction: source.Up})
	m.Append(&source.Migration{Version: 1, Direction: source.Down})
	m.Append(&source.Migration{Version: 1, Direction: source.MetaUp})
	m.Append(&source.Migration{Version: 1, Direction: source.MetaDown})
	m.Append(&source.Migration{Version: 3, Direction: source.Up})
	m.Append(&source.Migration{Version: 4, Direction: source.MetaUp})
	m.Append(&source.Migration{Version: 5, Direction: source.Down})
	m.Append(&source.Migration{Version: 6, Direction: source.MetaDown})
	m.Append(&source.Migration{Version: 8, Direction: source.Up})
	m.Append(&source.Migration{Version: 8, Direction: source.Down})

	d.(*Stub).Migrations = m

	st.Test(t, d)
}

func TestWithEmptyMigration(t *testing.T) {
	s := &Stub{}
	d, err := s.Open("")
	if err != nil {
		t.Fatal(err)
	}

	m := source.NewMigrations()

	d.(*Stub).Migrations = m

	version, err := d.First()
	if err == nil {
		t.Fatalf("First: expected err not to be nil")
	}

	if version != 0 {
		t.Errorf("First: expected 0, got %v", version)
	}
}