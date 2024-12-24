package encoder

import (
	"math"
	"strings"
	"time"

	"github.com/KakashiHatake324/mockjs"
)

var (
	_hsv                   = "window._hsv"
	interstitialSeed int32 = func() int32 {
		seed := 4046101435
		return int32(seed)
	}()
)

type PseudoRandom struct {
	currentStep     int32
	initialTime     int32
	seed            int32
	timeByte        int32
	keys            [][]byte
	values          [][]byte
	bytes           []byte
	bytesUsedForXor []byte
	Payload         string
}

func (p *PseudoRandom) Constructor(createdHash, initialTime, randomValue1, randomValue2 int32) {
	p.initialTime = initialTime
	p.initialTime += randomValue2
	p.initialTime += randomValue1 + randomValue2
	p.timeByte = int32(time.Now().UnixMilli()) & 255
	p.seed = createdHash ^ p.initialTime
	p.currentStep = -1
}

func (p *PseudoRandom) GetByte() int32 {
	p.currentStep++
	if p.currentStep > 2 {
		p.currentStep = 0
		p.seed = (func(o int32) int32 {
			o ^= o << 13
			o ^= o >> 17
			o ^= o << 5
			return o
		})(p.seed)
	}

	var s int32 = 16 - p.currentStep*int32(8)

	var t int32 = 0
	t |= p.timeByte << 24
	t |= p.timeByte << 16
	t |= p.timeByte << 8
	t |= p.timeByte << 0
	return ((t ^ p.seed) >> s) & 255

}

func (p *PseudoRandom) AddSignal(itemKey string, itemValue any) {
	key := mockjs.InitWindow().JSON.Stringify(itemKey)
	value := mockjs.InitWindow().JSON.Stringify(itemValue)
	p.bytes = append(p.bytes, byte(p.GetByte()))
	p.keys = append(p.keys, p.char(key))
	p.bytes = append(p.bytes, byte(p.GetByte()))
	p.values = append(p.values, p.char(value))
}

func (s *PseudoRandom) xorValues(charCodesRaw []int32) []byte {
	var charCodesXor = []byte{}
	var bytesXor = []byte{}
	//if !strings.Contains(charCodesRaw, "\"") {
	//	charCodesRaw = "\"" + charCodesRaw + "\""
	//}

	//charCodesRaw = []int32{34, 104, 105, 34}
	for i := 0; i < len(charCodesRaw); i++ {
		var newByte = s.GetByte()
		bytesXor = append(bytesXor, byte(newByte))
		charCodesXor = append(charCodesXor, byte(charCodesRaw[i]^newByte))
	}

	s.bytesUsedForXor = append(s.bytesUsedForXor, bytesXor...)
	return charCodesXor
}

func (p *PseudoRandom) char(key string) []byte {
	codePoints := []int32{}
	if math.IsNaN(mockjs.Math.ToFloat64(key)) {
		key = mockjs.InitWindow().JSON.Stringify(key)
	}
	for _, char := range key {
		var charString = string(char)
		codePoints = append(codePoints, int32(mockjs.InitWindow().CharcodeAt(charString, 0)))
	}
	return p.xorValues(codePoints)
}

func (p *PseudoRandom) BuildPayload() {
	var output = []byte{}

	for idx := 0; idx < len(p.keys); idx++ {
		// seperate = 123 { || 44 ,
		seperate := func(idx int) byte {
			if idx == 0 {
				return 123
			}
			return 44
		}(idx)
		output = append(output, seperate^p.bytes[2*idx])
		output = append(output, p.keys[idx]...)
		output = append(output, 58^p.bytes[2*idx+1])
		output = append(output, p.values[idx]...)
	}
	var signature = []byte{func(idx int) byte {
		if idx == 0 {
			return 123
		}
		return 44
	}(len(output)) ^ byte(p.GetByte())}
	r3n := mockjs.InitWindow().JSON.Stringify("r3n")
	r33 := mockjs.InitWindow().JSON.Stringify(_hsv)
	signature = append(signature, p.char(r3n)...)
	signature = append(signature, byte(58^p.GetByte()))
	signature = append(signature, p.char(r33)...)
	signature = append(signature, byte(125^p.GetByte()))
	output = append(output, signature...)

	var finalPayload = []string{}
	for t := 0; t < len(output); {
		var r int32
		var r1 int32
		var r2 int32
		var r3 int32

		r1 = ((int32(output[t]) ^ p.timeByte) << 16)
		t++
		if t < len(output) {
			r2 = ((int32(output[t]) ^ p.timeByte) << 8)
		}
		t++
		if t < len(output) {
			r3 = (int32(output[t]) ^ p.timeByte)
		}
		t++
		r = r1 | r2 | r3
		finalPayload = append(finalPayload,
			mockjs.InitWindow().FromCharcode([]any{determineNumber((r >> 18) & 63)}),
			mockjs.InitWindow().FromCharcode([]any{determineNumber((r >> 12) & 63)}),
			mockjs.InitWindow().FromCharcode([]any{determineNumber((r >> 6) & 63)}),
			mockjs.InitWindow().FromCharcode([]any{determineNumber(63 & r)}),
		)
	}
	var i = len(output) % 3
	if i != 0 {
		var popQty = 3 - i
		for p := 0; p < popQty; p++ {
			finalPayload = mockjs.Array.PopString(finalPayload)
		}
	}
	p.Payload = strings.Join(finalPayload, "")
}

func determineNumber(n int32) int32 {
	if n > 37 {
		return 59 + n
	} else if n > 11 {
		return 53 + n
	} else if n > 1 {
		return 46 + n
	}
	return 50*n + 45
}

func hashString(s string) int32 {
	if s == "" {
		return 1789537805
	}
	var t int32 = 0
	for e := 0; e < len(s); e++ {
		t = (t << 5) - t + int32(mockjs.InitWindow().CharcodeAt(s, e))
	}
	if t == 0 {
		return 1789537805
	}
	return t
}

func CreateHash(cid, sitehash string, initialTime, randomValue1, randomValue2 int32) int32 {
	return int32(hashString(cid)) ^ int32(initialTime+randomValue1+randomValue2*2) ^ int32(hashString(sitehash)) ^ int32(interstitialSeed)
}
