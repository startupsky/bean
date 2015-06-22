package util

type BeanSlice []interface{}

func (this BeanSlice) Remove(ele interface{}) BeanSlice {
	for i, e := range this {
		if e == ele {
			this[i], this[len(this)-1] = this[len(this)-1], this[i]
			this[len(this)-1] = nil
			return this[:len(this)-1]
		}
	}
	return this
}
