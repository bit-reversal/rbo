package rbo

import(
 "testing"
 "fmt"
)

// Returns (t+min{d>0 : r1<= RevBits( (t+d)mod 2^k   ) <=r2}) mod 2^k,
// which is the next wake-up time slot modulo 2^k after t for the RBO receiver if the values of lb and ub are r1 and r2, respectively.
// We assume 0<=r1<=r2< 2^k.
// This is naive implementation of this function
func naiveNSI(k uint8, t uint64, r1 uint64, r2 uint64) uint64 {
	var mask uint64 = (1 << k) - 1 // 2^k-1 
	t = ((t + 1) & mask)           // (t+1) mod 2^k
	r := RevBits(k, t)
	for r < r1 || r2 < r {
		t = ((t + 1) & mask)
		r = RevBits(k, t)
	}
	return t
}

func TestNSI(test *testing.T) {

	const kMax uint8 = 9

	for k := uint8(1); k <= kMax; k++ {
                fmt.Printf("k = %v\n", k)
		twoToK := uint64(1 << k)
		for t := uint64(0); t < twoToK; t++ {
			for r1 := uint64(0); r1 < twoToK; r1++ {
				for r2 := r1; r2 < twoToK; r2++ {
					n1 := NSI(k, t, r1, r2)
					n2 := naiveNSI(k, t, r1, r2)
					if n1 != n2 {
						test.Errorf("for k=%v, t=%v, r1=%v, r2=%v, we have NSI=%v and naiveNSI=%v", k, t, r1, r2, n1, n2)
						return
					}
				}
			}
		}
	}

}
