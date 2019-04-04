package cover


//通过偏移量计算出值
func BitMapGetValue(b [1]byte, offset uint, len uint) int8 {
    deflection := 8 - offset - len

    i := int8(b[0])
    i =  i >> offset

    if deflection > 0 {
        i = i << deflection
        i = i >> deflection
    }

    return i
}