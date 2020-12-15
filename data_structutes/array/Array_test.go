/**
 * @Author:    iccy
 * @Mail:      iccy.fun@outlook.com
 * @Data:      2020/12/15 0:24
 * @License:   $
 */

package array

import "testing"

func TestArray(t *testing.T) {
	array := New(10)

	array.Fill(18)
	for i := 0; i < 10; i++ {
		t.Log(array.At(i))
		array.Set(i, i)
	}

	array.Swap(0, array.Size() - 1)

	t.Log(array.Front())
	t.Log(array.Back())

	t.Log(array.Empty())
}