package badger_test

import (
	"testing"

	"github.com/elojah/badger"
	"github.com/elojah/services"
)

func TestUp(t *testing.T) {
	s := &badger.Service{}
	l := s.NewLauncher(badger.Namespaces{
		Badger: "badger",
	}, "badger")
	ls := services.Launchers{}
	ls = append(ls, l)
	if err := ls.Up("config_test.json"); err != nil {
		t.Error(err)
	}
}
