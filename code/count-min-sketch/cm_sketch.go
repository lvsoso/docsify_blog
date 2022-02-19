package main

import (
	"math/rand"
	"time"
)

// https://hardcore.feishu.cn/docs/doccn12GOjgkMGZAXEzVJSlEQrb#

//一个BitMap的实现
type cmRow []byte //byte = uint8 = 0000，0000 = COUNTER 4BIT = 2 counter

//64 counter
//1 uint8 =  2counter
//32 uint8 = 64 counter
func newCmRow(numCounters int64) cmRow {
	return make(cmRow, numCounters/2)
}

func (r cmRow) get(n uint64) byte {
	return byte(r[n/2]>>((n&1)*4)) & 0x0f
}

// 0000,0000|0000,0000| 0000,0000 make([]byte, 3) = 6 counter

func (r cmRow) increment(n uint64) {
	//定位到第i个Counter
	i := n / 2 //r[i]
	//右移距离，偶数为0，奇数为4
	s := (n & 1) * 4
	//取前4Bit还是后4Bit
	v := (r[i] >> s) & 0x0f //0000, 1111
	//没有超出最大计数时，计数+1
	if v < 15 {
		r[i] += 1 << s
	}
}

//cmRow 100,
//保鲜
func (r cmRow) reset() {
	// 计数减半
	for i := range r {
		r[i] = (r[i] >> 1) & 0x77 //0111，0111
	}
}

func (r cmRow) clear() {
	// 清空计数
	for i := range r {
		r[i] = 0
	}
}

//快速计算最接近x的二次幂的算法
//比如x=5，返回8
//x = 110，返回128

//2^n
//1000000 (n个0）
//01111111（n个1） + 1
// x = 1001010 = 1111111 + 1 =10000000
func next2Power(x int64) int64 {
	x--
	x |= x >> 1
	x |= x >> 2
	x |= x >> 4
	x |= x >> 8
	x |= x >> 16
	x |= x >> 32
	x++
	return x
}

//cmSketch封装

const cmDepth = 4

type cmSketch struct {
	rows [cmDepth]cmRow
	seed [cmDepth]uint64
	mask uint64
}

//numCounter - 1 = next2Power() = 0111111(n个1）

//0000,0000|0000,0000|0000,0000
//0000,0000|0000,0000|0000,0000
//0000,0000|0000,0000|0000,0000
//0000,0000|0000,0000|0000,0000

func newCmSketch(numCounters int64) *cmSketch {
	if numCounters == 0 {
		panic("cmSketch: bad numCounters")
	}

	numCounters = next2Power(numCounters)
	sketch := &cmSketch{mask: uint64(numCounters - 1)}
	// Initialize rows of counters and seeds.
	source := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < cmDepth; i++ {
		sketch.seed[i] = source.Uint64()
		sketch.rows[i] = newCmRow(numCounters)
	}
	return sketch
}

func (s *cmSketch) Increment(hashed uint64) {
	for i := range s.rows {
		s.rows[i].increment((hashed ^ s.seed[i]) & s.mask)
	}
}

// 找到最小的计数值
func (s *cmSketch) Estimate(hashed uint64) int64 {
	min := byte(255)
	for i := range s.rows {
		val := s.rows[i].get((hashed ^ s.seed[i]) & s.mask)
		if val < min {
			min = val
		}
	}
	return int64(min)
}

// 让所有计数器都减半，保鲜机制
func (s *cmSketch) Reset() {
	for _, r := range s.rows {
		r.reset()
	}
}

// 清空所有计数器
func (s *cmSketch) Clear() {
	for _, r := range s.rows {
		r.clear()
	}
}
