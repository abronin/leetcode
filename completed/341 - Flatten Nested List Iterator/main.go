/**
 * // This is the interface that allows for creating nested lists.
 * // You should not implement it, or speculate about its implementation
 * type NestedInteger struct {
 * }
 *
 * // Return true if this NestedInteger holds a single integer, rather than a nested list.
 * func (this NestedInteger) IsInteger() bool {}
 *
 * // Return the single integer that this NestedInteger holds, if it holds a single integer
 * // The result is undefined if this NestedInteger holds a nested list
 * // So before calling this method, you should have a check
 * func (this NestedInteger) GetInteger() int {}
 *
 * // Set this NestedInteger to hold a single integer.
 * func (n *NestedInteger) SetInteger(value int) {}
 *
 * // Set this NestedInteger to hold a nested list and adds a nested integer to it.
 * func (this *NestedInteger) Add(elem NestedInteger) {}
 *
 * // Return the nested list that this NestedInteger holds, if it holds a nested list
 * // The list length is zero if this NestedInteger holds a single integer
 * // You can access NestedInteger's List element directly if you want to modify it
 * func (this NestedInteger) GetList() []*NestedInteger {}
 */

type NestedIterator struct {
	Current  int
	Iterator *NestedIterator
	Len      int
	List     []*NestedInteger
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
	return &NestedIterator{
		Current:  0,
		Iterator: nil,
		Len:      len(nestedList),
		List:     nestedList,
	}
}

func (this *NestedIterator) Next() int {
	if this.List[this.Current].IsInteger() {
		this.Current++
		return this.List[this.Current-1].GetInteger()
	} else {
		if this.Iterator.HasNext() {
			next := this.Iterator.Next()
			if !this.Iterator.HasNext() {
				this.Iterator = nil
				this.Current++
			}
			return next
		} else {
			this.Iterator = nil
			this.Current++
			return this.Next()
		}
	}

}

func (this *NestedIterator) HasNext() bool {
	if this.Current >= this.Len {
		return false
	}
	if this.List[this.Current].IsInteger() {
		return true
	}
	if this.Iterator == nil {
		this.Iterator = Constructor(this.List[this.Current].GetList())
	}
	if this.Iterator.HasNext() {
		return true
	}

	this.Iterator = nil
	this.Current++
	return this.HasNext()
}