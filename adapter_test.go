package fsadapter

import (
	"io/fs"
	"os"
	"testing"

	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	"github.com/stretchr/testify/assert"
)

func TestCasbinFsAdapter_LoadPolicy(t *testing.T) {
	t.Run("TestCasbinFsAdapter_LoadPolicy_ShouldIgnoreComment", func(t *testing.T) {
		fsys := os.DirFS("examples/config/")
		adapter := NewAdapter(fsys, "policy.csv")

		m, err := NewModel(fsys, "model.conf")
		assert.NoError(t, err)

		e, err := casbin.NewEnforcer(m, adapter)
		assert.NoError(t, err)

		// Check if the policy is loaded
		loaded := e.HasPolicy("test_comment", "/comment", "GET", "allow")
		assert.True(t, loaded)
	})
	t.Run("TestCasbinFsAdapter_LoadPolicy_EmptyShouldReturnError", func(t *testing.T) {
		var err error
		fsys := os.DirFS("examples/config/")
		adapter := NewAdapter(fsys, "")

		m := model.NewModel()
		err = adapter.LoadPolicy(m)
		assert.Error(t, err)
		assert.Equal(t, err.Error(), errInvalidFilePath)
	})

	t.Run("TestCasbinFsAdapter_LoadPolicy_MissingShouldReturnError", func(t *testing.T) {
		var err error
		fsys := os.DirFS("examples/config/")
		adapter := NewAdapter(fsys, "missing.csv")

		m := model.NewModel()
		err = adapter.LoadPolicy(m)
		assert.Error(t, err)
		assert.ErrorIs(t, err, fs.ErrNotExist)
	})
}

func TestCasbinFsAdapter_SavePolicy(t *testing.T) {
	t.Run("TestCasbinFsAdapter_SavePolicy_ShouldReturnNotImplement", func(t *testing.T) {
		var err error
		fsys := os.DirFS("examples/config/")
		adapter := NewAdapter(fsys, "policy.csv")

		m := model.NewModel()
		err = adapter.SavePolicy(m)
		assert.Error(t, err)
		assert.Equal(t, errNotImplemented, err.Error())
	})
}

func TestCasbinFsAdapter_AddPolicy(t *testing.T) {
	t.Run("TestCasbinFsAdapter_AddPolicy_ShouldReturnNotImplement", func(t *testing.T) {
		var err error
		fsys := os.DirFS("examples/config/")
		adapter := NewAdapter(fsys, "policy.csv")

		sec := "foo"
		ptype := "bar"
		var rule []string
		err = adapter.AddPolicy(sec, ptype, rule)
		assert.Error(t, err)
		assert.Equal(t, errNotImplemented, err.Error())
	})
}

func TestCasbinFsAdapter_RemovePolicy(t *testing.T) {
	t.Run("TestCasbinFsAdapter_RemovePolicy_ShouldReturnNotImplement", func(t *testing.T) {
		var err error
		fsys := os.DirFS("examples/config/")
		adapter := NewAdapter(fsys, "policy.csv")

		sec := "foo"
		ptype := "bar"
		var rule []string
		err = adapter.RemovePolicy(sec, ptype, rule)
		assert.Error(t, err)
		assert.Equal(t, errNotImplemented, err.Error())
	})
}

func TestCasbinFsAdapter_RemoveFilteredPolicy(t *testing.T) {
	t.Run("TestCasbinFsAdapter_RemoveFilteredPolicy_ShouldReturnNotImplement", func(t *testing.T) {
		var err error
		fsys := os.DirFS("examples/config/")
		adapter := NewAdapter(fsys, "policy.csv")

		sec := "foo"
		ptype := "bar"
		fieldIndex := 0
		fieldValue := "filter"
		err = adapter.RemoveFilteredPolicy(sec, ptype, fieldIndex, fieldValue)
		assert.Error(t, err)
		assert.Equal(t, errNotImplemented, err.Error())
	})
}
