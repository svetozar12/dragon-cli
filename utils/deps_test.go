package utils

import (
	"testing"
)

func TestGetDeps(t *testing.T) {
	t.Run("TestGetDeps correct behavior", func(t *testing.T) {
		if len(deps) > 0 || len(devDeps) > 0 {
			t.Fatalf("TestGetDeps() both deps and devDeps should be empty arrays")
		}
		newDeps := "new deps"
		SetDeps([]string{newDeps})
		SetDevDeps([]string{newDeps})
		if deps[0] != newDeps && devDeps[0] != newDeps {
			t.Fatalf("GetDeps() doesn't get correct dependencies")
		}
	})
}

func TestRemoveDuplicates(t *testing.T) {
	t.Run("TestRemoveDuplicates correct behavior", func(t *testing.T) {
		result := removeDuplicates([]string{"Same String 1234 !!", "Same String 1234 !!"})
		if len(result) > 1 {
			t.Fatalf("removeDuplicates() should remove strings that are the same")
		}
	})
	t.Run("TestRemoveDuplicates correct behavior with empty array", func(t *testing.T) {
		result := removeDuplicates([]string{})
		if len(result) > 0 {
			t.Fatalf("removeDuplicates() should return empty array")
		}
	})
}

func TestSetDeps(t *testing.T) {
	t.Run("TestSetDeps correct behavior", func(t *testing.T) {
		newDeps := "new deps"
		SetDeps([]string{newDeps})
		if deps[0] != newDeps {
			t.Fatalf("SetDeps() should set new dependency in array")
		}
	})
	t.Run("TestSetDeps shouldn't have two deps with same name", func(t *testing.T) {
		newDeps := "new deps"
		SetDeps([]string{newDeps, newDeps})
		if len(deps) > 1 {
			t.Fatalf("SetDeps() shouldn't have two deps with same name")
		}
	})
}

func TestSetDevDeps(t *testing.T) {
	t.Run("TestSetDevDeps correct behavior", func(t *testing.T) {
		newDeps := "new deps"
		SetDevDeps([]string{newDeps})
		if deps[0] != newDeps {
			t.Fatalf("SetDevDeps() should set new dependency in array")
		}
	})
	t.Run("TestSetDevDeps shouldn't have two deps with same name", func(t *testing.T) {
		newDeps := "new deps"
		SetDevDeps([]string{newDeps, newDeps})
		if len(deps) > 1 {
			t.Fatalf("SetDevDeps() shouldn't have two deps with same name")
		}
	})
}
