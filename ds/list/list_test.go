package list_test

import (
	"sort"
	"testing"

	ll "github.com/solpipe/solpipe-util/ds/list"
	"github.com/stretchr/testify/assert"
)

func isEqual(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestSingleDelete(t *testing.T) {
	r := []int{3, 4, 5, 1, 5, 7}

	// test middle and edges
	test_cut := []int{0, 2, len(r) - 2, len(r) - 1}
	for j := 0; j < len(test_cut); j++ {
		cut_i := test_cut[j]

		q := ll.CreateGeneric[int]()
		q.AppendArray(r)
		assert.Equal(t, isEqual(r, q.Array()), true, "failed to append")

		r_reduced := append(r[0:cut_i], r[cut_i+1:]...)

		assert.Nil(t, q.Iterate(func(obj int, index uint32, deleteNode func()) error {
			if index == uint32(cut_i) {
				deleteNode()
			}
			return nil
		}), "failed to iterate")

		t.Logf("r_reduced=%+v vs q=%+v", r_reduced, q.Array())
		assert.Equal(t, len(r_reduced), int(q.Size), "failed to delete node - 1")

		assert.Equal(t, isEqual(r_reduced, q.Array()), true, "failed to delete node - 2")
	}
}

func scb(a int, b int) int {
	if a < b {
		return -1
	} else if a == b {
		return 0
	} else {
		return 1
	}
}

func TestMultiSort(t *testing.T) {
	all := [][]int{
		{3, 4, 5, 1, 5, 7},
		{5, 4, 5, 1, 5, 5, 7},
		{0, -1, 10, -1, -1, 20, 21, 18},
		{18, 21, 20, -1, -1, 10, -1, 0},
		{9},
		{},
	}
	for k := 0; k < len(all); k++ {
		r := make([]int, len(all[k]))
		copy(r, all[k][:])
		r_sorted := make([]int, len(r))
		copy(r_sorted, r)
		sort.Ints(r_sorted)

		q := ll.CreateGeneric[int]()
		for i := 0; i < len(r); i++ {
			assert.NotNil(t, q.InsertSorted(r[i], scb), "failed to insert")
		}
		q.Iterate(func(obj int, index uint32, deleteNode func()) error {
			t.Logf("i=%d; v=%d size=%d", index, obj, q.Size)
			return nil
		})
		t.Logf("r_sorted=%+v vs q=%+v", r_sorted, q.Array())
		assert.Equal(t, len(r_sorted), int(q.Size), "failed to sort - 1")
		assert.Equal(t, isEqual(r_sorted, q.Array()), true, "failed to sort - 2")

		test_cut := []int{0, 2, len(r) - 2, len(r) - 1}
		if 2 < len(r) {
			for j := 0; j < len(test_cut); j++ {
				cut_i := test_cut[j]

				q2 := ll.CreateGeneric[int]()
				q2.AppendArray(r)
				assert.Equal(t, isEqual(r, q2.Array()), true, "failed to append")

				r_reduced := append(r[0:cut_i], r[cut_i+1:]...)

				assert.Nil(t, q2.Iterate(func(obj int, index uint32, deleteNode func()) error {
					if index == uint32(cut_i) {
						deleteNode()
					}
					return nil
				}), "failed to iterate")

				t.Logf("r_reduced=%+v vs q=%+v", r_reduced, q2.Array())
				assert.Equal(t, len(r_reduced), int(q2.Size), "failed to delete node - 1")

				assert.Equal(t, isEqual(r_reduced, q2.Array()), true, "failed to delete node - 2")
			}
		}

	}
}
