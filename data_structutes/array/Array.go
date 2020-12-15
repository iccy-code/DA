/**
 * @Author:    iccy
 * @Mail:      iccy.fun@outlook.com
 * @Data:      2020/12/15 0:24
 * @License:   $
 */

package array

// Array is a fixed(确定的) size slice
type Array struct {
	values []interface{}
}

/*
 * @Description: Creates(创建) a new array with passed(通过) size
 * @Param: size int
 * @Return: *Array
 * @Author: iccy
 * @Date: 2020/12/15
 */
func New(size int) *Array {
	return &Array{values:make([]interface{}, size, size)}
}

/*
 * @Description: Creates a new array from(从) another(另一个) array, and copy its values
 * @Param: other *Array
 * @Return: *Array
 * @Author: iccy
 * @Date: 2020/12/15
 */
func NewFormArray(other *Array) *Array {
	newArray := New(other.Size())

	for i, _ := range other.values {
		newArray.values[i] = other.values[i]
	}
	return newArray
}

/*
 * @Description: Fills(填充) array a with value val
 * @Param: val interface{}
 * @Return: void
 * @Author: iccy
 * @Date: 2020/12/15
 */
func (a *Array) Fill(val interface{}) {
	for i, _ := range a.values {
		a.values[i] = val
	}
}

/*
 * @Description: Sets(置, 设置) value val to the position(位置, 方位) pos of the array
 * @Param: pos int, val interface{}
 * @Return: void
 * @Author: iccy
 * @Date: 2020/12/15
 */
func (a *Array) Set(pos int, val interface{}) {
	if pos < 0 || pos >= a.Size() {
		return
	}
	a.values[pos] = val
}

/*
 * @Description: Returns the value at position pos in the array
 * @Param: pos int
 * @Return: interface{}
 * @Author: iccy
 * @Date: 2020/12/15
 */
func (a *Array) At(pos int) interface{} {
	if pos < 0 || pos >= a.Size() {
		return nil
	}
	return a.values[pos]
}

/*
 * @Description: Returns the first(第一) value in the array
 * @Param: void
 * @Return: interface{}
 * @Author: iccy
 * @Date: 2020/12/15
 */
func (a *Array) Front() interface{} {
	return a.At(0)
}

/*
 * @Description: Returns the last(最后) value in the array
 * @Param: void
 * @Return: interface{}
 * @Author: iccy
 * @Date: 2020/12/15
 */
func (a *Array) Back() interface{} {
	return a.At(a.Size() - 1)
}

/*
 * @Description: Returns number of elements(元素, 基础) within(在...之内) the array
 * @Param: void
 * @Return: int
 * @Author: iccy
 * @Date: 2020/12/15
 */
func (a *Array) Size() int {
	return len(a.values)
}

/*
 * @Description: Returns whether(是否) the array is empty(空的, 无意义) or not
 * @Param:
 * @Return:
 * @Author: iccy
 * @Date: 2020/12/16
 */
func (a *Array) Empty() bool {
	return a.Size() == 0
}

/*
 * @Description: Swaps the values of two elements at the specified position in array
 * @Param: pos1 int, pos2 int
 * @Return: void
 * @Author: iccy
 * @Date: 2020/12/16
 */
func (a *Array) Swap(pos1 int, pos2 int) {
	if pos1 < 0 || pos1 >= a.Size() || pos2 < 0 || pos2 >= a.Size() {
		return
	}
	a.values[pos1], a.values[pos2] = a.values[pos2], a.values[pos1]
}
