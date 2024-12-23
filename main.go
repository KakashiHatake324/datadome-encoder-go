package main

import (
	"datadome-encoder-go/encoder"
	"log"
	"time"

	"github.com/KakashiHatake324/mockjs"
)

const (
	datadomeCID = "xkswwJ_dA85U08xIGrVPNqovB9aYgBW~BL5vgcLRDA59HA5rKxKTXupKCz8gHPLaQKnkyz95bCzXdBrsGeUnDhLXW1zNEXLEPoPgIW~eil6U0IXNo8PH7eu~Hy7KpejE"
	websiteHash = "A55FBF4311ED6F1BF9911EB71931D5"
	testCount   = 10000
)

func main() {
	var tests = make(map[string]int)
	testStart := time.Now()
	for n := 0; n < testCount; n++ {
		var currentTime = int32(time.Now().UnixMilli())
		var randomValue1 = int32(mockjs.Math.Floor(mockjs.Math.Random() * 1000))
		var randomValue2 = int32(mockjs.Math.Floor(mockjs.Math.Random() * 1000))

		enc := &encoder.PseudoRandom{}
		enc.Constructor(encoder.CreateHash(datadomeCID, websiteHash, currentTime, randomValue1, randomValue2), currentTime, randomValue1, randomValue2)
		enc.AddSignal("hi", 1)
		enc.AddSignal("hi", false)
		enc.AddSignal("hi", "hoe")
		enc.AddSignal("hi", 1234)
		enc.AddSignal("hi", 123)
		enc.AddSignal("hi", map[string]interface{}{"yo": 12345678})
		enc.BuildPayload()
		tests[enc.Payload]++
	}
	testFinish := time.Since(testStart).Milliseconds()
	log.Println(tests)
	log.Println("Test Took:", testFinish, "ms..", "for", testCount, "gens")
}
