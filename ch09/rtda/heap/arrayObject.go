package heap

func (self *Object) Bytes() []int8 {
	return self.data.([]int8)
}
func (self *Object) Shorts() []int16 {
	return self.data.([]int16)
}
func (self *Object) Ints() []int32 {
	return self.data.([]int32)
}
func (self *Object) Longs() []int64 {
	return self.data.([]int64)
}
func (self *Object) Chars() []uint16 {
	return self.data.([]uint16)
}
func (self *Object) Floats() []float32 {
	return self.data.([]float32)
}
func (self *Object) Doubles() []float64 {
	return self.data.([]float64)
}
func (self *Object) Refs() []*Object {
	return self.data.([]*Object)
}

func ArrayCopy(src, dst *Object, srcPos, dstPos, length int32) {
	switch src.data.(type) {
	case []int8:
		inSrc := src.data.([]int8)[srcPos: srcPos+length]
		inDst := dst.data.([]int8)[dstPos: dstPos+length]
		copy(inDst, inSrc)
	case []int16:
		inSrc := src.data.([]int16)[srcPos: srcPos+length]
		inDst := dst.data.([]int16)[dstPos: dstPos+length]
		copy(inDst, inSrc)
	case []int32:
		inSrc := src.data.([]int32)[srcPos: srcPos+length]
		inDst := dst.data.([]int32)[dstPos: dstPos+length]
		copy(inDst, inSrc)
	case []int64:
		inSrc := src.data.([]int64)[srcPos: srcPos+length]
		inDst := dst.data.([]int64)[dstPos: dstPos+length]
		copy(inDst, inSrc)
	case []uint16:
		inSrc := src.data.([]uint16)[srcPos: srcPos+length]
		inDst := dst.data.([]uint16)[dstPos: dstPos+length]
		copy(inDst, inSrc)
	case []float32:
		inSrc := src.data.([]float32)[srcPos: srcPos+length]
		inDst := dst.data.([]float32)[dstPos: dstPos+length]
		copy(inDst, inSrc)
	case []float64:
		inSrc := src.data.([]float64)[srcPos: srcPos+length]
		inDst := dst.data.([]float64)[dstPos: dstPos+length]
		copy(inDst, inSrc)
	case []*Object:
		inSrc := src.data.([]*Object)[srcPos: srcPos+length]
		inDst := dst.data.([]*Object)[dstPos: dstPos+length]
		copy(inDst, inSrc)
	default:
		panic("Not array!")
	}
}

func (self *Object) ArrayLength() int32 {
	switch self.data.(type) {
	case []int8:
		return int32(len(self.data.([]int8)))
	case []int16:
		return int32(len(self.data.([]int16)))
	case []int32:
		return int32(len(self.data.([]int32)))
	case []int64:
		return int32(len(self.data.([]int64)))
	case []uint16:
		return int32(len(self.data.([]uint16)))
	case []float32:
		return int32(len(self.data.([]float32)))
	case []float64:
		return int32(len(self.data.([]float64)))
	case []*Object:
		return int32(len(self.data.([]*Object)))
	default:
		panic("Not array!")
	}
}
