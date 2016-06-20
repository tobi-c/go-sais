package suffixarray

// S-Type : true
// L-Type : false
type sType []bool

func (sTy sType) isLMSType(i int) bool {
	return !sTy[i-1] && sTy[i]
}

type bucket []int

func (bkt bucket) setSARight(sa suffixArray, index int, bktIndex int) {
	sa[bkt[bktIndex]] = index
	bkt[bktIndex]--
}

func (bkt bucket) setSALeft(sa suffixArray, index int, bktIndex int) {
	sa[bkt[bktIndex]] = index
	bkt[bktIndex]++
}

type suffixArray []int

func (sa suffixArray) init() {
	for i := range sa {
		sa[i] = -1
	}
}

func (sa suffixArray) lmsPrefixToLMSSubstrings(sTy sType) suffixArray {
	lmsLen := 0
	for i := 0; i < len(sa); i++ {
		if sa[i] != 0 && sTy.isLMSType(sa[i]) {
			index := sa[i]
			sa[i] = -1
			sa[lmsLen] = index
			lmsLen++
		} else {
			sa[i] = -1
		}
	}
	return sa[:lmsLen]
}

type byteData []byte
type intData []int

func sais(data []byte) []int {
	sa := make(suffixArray, len(data))
	(byteData(data)).sais(sa, 256)
	return sa
}
