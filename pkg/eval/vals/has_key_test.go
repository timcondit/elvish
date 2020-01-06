package vals

import (
	"testing"

	. "github.com/elves/elvish/pkg/tt"
)

type hasKeyer struct{ key interface{} }

func (h hasKeyer) HasKey(k interface{}) bool { return k == h.key }

func TestHasKey(t *testing.T) {
	Test(t, Fn("HasKey", HasKey), Table{
		// Map
		Args(MakeMap("k", "v"), "k").Rets(true),
		Args(MakeMap("k", "v"), "bad").Rets(false),
		// StructMap
		Args(testStructMap{}, "name").Rets(true),
		Args(testStructMap{}, "bad").Rets(false),
		// HasKeyer
		Args(hasKeyer{"valid"}, "valid").Rets(true),
		Args(hasKeyer{"valid"}, "invalid").Rets(false),
		// Fallback to IterateKeys
		Args(keysIterator{vs("lorem")}, "lorem").Rets(true),
		Args(keysIterator{vs("lorem")}, "ipsum").Rets(false),
		// Fallback to Len
		Args(MakeList("lorem", "ipsum"), "0").Rets(true),
		Args(MakeList("lorem", "ipsum"), "0:").Rets(true),
		Args(MakeList("lorem", "ipsum"), "2").Rets(false),
	})
}
