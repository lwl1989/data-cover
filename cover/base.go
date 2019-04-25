package cover


func coverToByte(b []interface{}) []byte {
    bt := make([]byte, len(b))
    for k,v := range b {
        bt[k] = v.(byte)
    }
    return bt
}