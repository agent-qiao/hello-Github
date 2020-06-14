package popcount

// 当前函数只是一个初始化函数只能被导入其他包时使用
//pc[i] is the pouplation count of i.
var pc [256]byte

func init()  {
    for i, _ := range pc {
        pc[i] = pc[i/2] + byte(i&1)
    }
}

// popcpunt returns the population count (number of set bits) of x

func popcount( x uint64) int {
    return int(pc[byte(x>>0*8)]+
        pc[byte(x>>1*8)]+
        pc[byte(x>>2*8)]+
        pc[byte(x>>3*8)]+
        pc[byte(x>>4*8)]+
        pc[byte(x>>5*8)]+
        pc[byte(x>>6*8)]+
        pc[byte(x>>7*8)])
}