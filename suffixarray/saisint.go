package suffixarray

// sType
func (data intData) makeSType() sType {
	sTy := make(sType, len(data)) // all false
	// len(data)   : S-Type($)
	// len(data)-1 : L-Type
	for i := len(sTy) - 2; 0 <= i; i-- {
		t1 := data[i]
		t2 := data[i+1]
		if t1 < t2 || (t1 == t2 && sTy[i+1]) {
			sTy[i] = true
		}
	}

	return sTy
}

// bucket
func (data intData) countToBucket(bkt bucket) {
	for i := range bkt {
		bkt[i] = 0
	}

	for i := 0; i < len(data); i++ {
		bkt[data[i]]++
	}
}

func (data intData) setBucketRight(bkt bucket) {
	data.countToBucket(bkt)

	sum := -1
	for i := range bkt {
		sum += bkt[i]
		bkt[i] = sum
	}
}

func (data intData) setBucketLeft(bkt bucket) {
	data.countToBucket(bkt)

	sum := 0
	for i := range bkt {
		b := bkt[i]
		bkt[i] = sum
		sum += b
	}
}

// induced sorting
func (data intData) sortLType(sa suffixArray, sTy sType, bkt bucket) {
	data.setBucketLeft(bkt)
	// len(sa) : LMS-Type($)
	index := len(sa) - 1
	bkt.setSALeft(sa, index, int(data[index]))
	for i := range sa {
		index = sa[i] - 1
		if 0 <= index && !sTy[index] {
			bkt.setSALeft(sa, index, int(data[index]))
		}
	}
}

func (data intData) sortSType(sa suffixArray, sTy sType, bkt bucket) {
	data.setBucketRight(bkt)
	for i := len(sa) - 1; 0 <= i; i-- {
		index := sa[i] - 1
		if 0 <= index && sTy[index] {
			bkt.setSARight(sa, index, int(data[index]))
		}
	}
}

func (data intData) sortLMSTypeLMSPrefix(sa suffixArray, sTy sType, bkt bucket) {
	data.setBucketRight(bkt)
	for i := len(sa) - 1; 1 <= i; i-- {
		if sTy.isLMSType(i) {
			bkt.setSARight(sa, i, int(data[i]))
		}
	}
}

func (data intData) sortLMSType(sa, lmsIndices suffixArray, bkt bucket) {
	data.setBucketRight(bkt)
	for i := len(lmsIndices) - 1; i >= 0; i-- {
		index := lmsIndices[i]
		sa[i] = -1
		bkt.setSARight(sa, index, int(data[index]))
	}
}

func (data intData) inducedSortLMSPrefix(sa suffixArray, sTy sType, bkt bucket) {
	data.sortLMSTypeLMSPrefix(sa, sTy, bkt)
	data.sortLType(sa, sTy, bkt)
	data.sortSType(sa, sTy, bkt)
}

func (data intData) inducedSort(sa, lmsIndices suffixArray, sTy sType, bkt bucket) {
	data.sortLMSType(sa, lmsIndices, bkt)
	data.sortLType(sa, sTy, bkt)
	data.sortSType(sa, sTy, bkt)
}

func (data intData) substringNames(sa, lmsSubstrs suffixArray, sTy sType) (suffixArray, int) {
	if len(lmsSubstrs) == 0 {
		return sa[:0], 0
	}

	nameUniqCount := 0
	names := sa[len(sa)/2:]
	names[sa[0]/2] = nameUniqCount
	for i := 1; i < len(lmsSubstrs); i++ {
		pos := sa[i]
		prev := sa[i-1]

		for {
			if data[pos] != data[prev] {
				// not equal LMS-substring
				nameUniqCount++
				break
			}

			if (pos != sa[i]) && sTy.isLMSType(pos) && sTy.isLMSType(prev) {
				// equal LMS-substring
				break
			}

			pos++
			prev++
			if pos == len(data) || prev == len(data) {
				// not equal LMS-substring (last char == $)
				nameUniqCount++
				break
			}
		}

		names[sa[i]/2] = nameUniqCount
	}

	namesIndex := len(names) - 1
	for i := len(names) - 1; 0 <= i; i-- {
		if names[i] >= 0 {
			name := names[i]
			names[i] = -1
			names[namesIndex] = name
			namesIndex--
		}
	}
	names = names[namesIndex+1:]
	nameUniqCount++
	return names, nameUniqCount
}

func (data intData) nameToIndex(names, namesSA suffixArray, sTy sType) {
	for i, index := 1, 0; i < len(data); i++ {
		if sTy.isLMSType(i) {
			names[index] = i
			index++
		}
	}

	for i := 0; i < len(namesSA); i++ {
		namesSA[i] = names[namesSA[i]]
	}
}

func (data intData) sais(sa suffixArray, bktLen int) {
	if len(data) == 0 {
		return
	}

	sa.init()
	sTy := data.makeSType()
	bkt := make(bucket, bktLen)

	// induced sorting LMS-Substrings
	data.inducedSortLMSPrefix(sa, sTy, bkt)
	lmsSubstrs := sa.lmsPrefixToLMSSubstrings(sTy)

	// sort LMS-Type
	names, nameUniqCount := data.substringNames(sa, lmsSubstrs, sTy)
	if nameUniqCount < len(names) {
		namesSA := lmsSubstrs
		(intData(names)).sais(namesSA, nameUniqCount)

		data.nameToIndex(names, namesSA, sTy)
	}
	names.init()

	// induced sorting
	data.inducedSort(sa, lmsSubstrs, sTy, bkt)
}
