package main

import (
	"datadome-encoder-go/encoder"
	"datadome-encoder-go/test"
	"log"
	"time"

	orderedmap "github.com/wk8/go-ordered-map/v2"

	"github.com/KakashiHatake324/mockjs"
)

const (
	datadomeCID = "xkswwJ_dA85U08xIGrVPNqovB9aYgBW~BL5vgcLRDA59HA5rKxKTXupKCz8gHPLaQKnkyz95bCzXdBrsGeUnDhLXW1zNEXLEPoPgIW~eil6U0IXNo8PH7eu~Hy7KpejE"
	websiteHash = "A55FBF4311ED6F1BF9911EB71931D5"
	testCount   = 1
)

func main() {
	var tests = make(map[string]int)
	testStart := time.Now()
	var currentTime = int32(time.Now().UnixMilli())
	var randomValue1 = int32(mockjs.Math.Floor(mockjs.Math.Random() * 1000))
	var randomValue2 = int32(mockjs.Math.Floor(mockjs.Math.Random() * 1000))

	var payload = orderedmap.New[string, interface{}]()
	payload.Set("hi", 832984732)
	payload.Set("yo", false)
	payload.Set("ye", true)
	payload.Set("summer", map[string]interface{}{"yo": 12345678})
	payload.Set("datadome", "sucks")

	for n := 0; n < testCount; n++ {
		enc := &encoder.PseudoRandom{}
		enc.Constructor(encoder.CreateHash(datadomeCID, websiteHash, currentTime, randomValue1, randomValue2), currentTime, randomValue1, randomValue2)
		enc.AddSignal("hi", 832984732)
		enc.AddSignal("yo", false)
		enc.AddSignal("ye", true)
		enc.AddSignal("summer", map[string]interface{}{"yo": 12345678})
		enc.AddSignal("datadome", "sucks")
		enc.BuildPayload()
		log.Println(enc.Payload)
		tests[enc.Payload]++
	}

	for n := 0; n < testCount; n++ {
		enc := test.NewT()
		enc.L.AddSignal("hi", 832984732)
		enc.L.AddSignal("yo", false)
		enc.L.AddSignal("ye", true)
		enc.L.AddSignal("summer", map[string]interface{}{"yo": 12345678})
		enc.L.AddSignal("datadome", "sucks")
		enc.L.Payload = enc.L.BuildPayload(map[string]interface{}{"cid": "xkswwJ_dA85U08xIGrVPNqovB9aYgBW~BL5vgcLRDA59HA5rKxKTXupKCz8gHPLaQKnkyz95bCzXdBrsGeUnDhLXW1zNEXLEPoPgIW~eil6U0IXNo8PH7eu~Hy7KpejE", "hash": "A55FBF4311ED6F1BF9911EB71931D5"})
		log.Println(enc.L.Payload)
		tests[enc.L.Payload]++
	}

	testFinish := time.Since(testStart).Milliseconds()
	log.Println(tests)
	log.Println("Test Took:", testFinish, "ms..", "for", testCount, "gens")
}
