package test

import (
	"encoding/json"
	"fmt"
	"sync"

	orderedmap "github.com/wk8/go-ordered-map/v2"

	"github.com/dop251/goja"
)

var GenPayload func(any, any) string
var encodingMu sync.Mutex

type TClass struct {
	N int64
	T int64
	I int64
	G int64
	O float64
	E float64
	S func(q string) int64
	A func(q int64) int64
	L *LClass
}

type LClass struct {
	AddSignal    func(string, interface{})
	signals      *orderedmap.OrderedMap[string, interface{}]
	BuildPayload func(dd interface{}) string
	Payload      string
	L            [][]int64
	P            [][]int64
	K            []int64
}

func init() {
	vm := goja.New()
	if _, err := vm.RunString(genScript()); err != nil {
		fmt.Println("Error executing JavaScript:", err)
		return
	}

	err := vm.ExportTo(vm.Get("genPayload"), &GenPayload)
	if err != nil {
		fmt.Println("Error vm.ExportTo:", err)
		return
	}
}

// build a new payload instance
func NewT() *TClass {
	newT := &TClass{}
	newT.L = func() *LClass {
		lClass := &LClass{}
		lClass.signals = orderedmap.New[string, interface{}]()
		lClass.AddSignal = func(q string, p interface{}) {
			lClass.signals.Set(q, p)
		}
		lClass.BuildPayload = func(dd interface{}) string {
			jsonInfo, _ := lClass.signals.MarshalJSON()
			ddInfo, _ := json.Marshal(dd)

			var object1 map[string]interface{}
			json.Unmarshal(jsonInfo, &object1)

			var object2 map[string]interface{}
			json.Unmarshal(ddInfo, &object2)
			encodingMu.Lock()
			defer encodingMu.Unlock()
			return GenPayload(object1, object2)
		}
		return lClass
	}()

	return newT
}

func (*TClass) New(q int64, r, h string) *TClass {
	newT := &TClass{}
	newT.L = func() *LClass {
		lClass := &LClass{}
		lClass.AddSignal = func(q string, p interface{}) {
			lClass.signals.Set(q, p)
		}
		lClass.BuildPayload = func(dd interface{}) string {
			jsonInfo, _ := lClass.signals.MarshalJSON()
			ddInfo, _ := json.Marshal(dd)

			var object1 map[string]interface{}
			json.Unmarshal(jsonInfo, &object1)

			var object2 map[string]interface{}
			json.Unmarshal(ddInfo, &object2)
			encodingMu.Lock()
			defer encodingMu.Unlock()
			return GenPayload(object1, object2)
		}
		return lClass
	}()

	return newT
}

func genScript() string {
	return `
			function d(q, p, U) {
				const window = {};
				var r = 1789537805,
				T = Date["now"]() & 255,
				e = Date["now"](),
				c = e,
				a = Math["floor"](Math["random"]() * 1000),
				n = Math["floor"](Math["random"]() * 1000);

			function f(q) {
				if (!q) return r;
				for (var p = 0, U = 0; U < q["length"]; U++)
				p = ((p << 5) - p + q["charCodeAt"](U)) << 0;
				return p == 0 ? r : p;
			}

			function t(q) {
				return q > 37 ? 59 + q : q > 11 ? 53 + q : q > 1 ? 46 + q : 50 * q + 45;
			}

			function b(q) {
				var p = q ^ c,
				U = -1,
				r = function () {
					p = (function (q) {
					for (var p = 188; true; ) {
						switch (p) {
						case 193:
							var U = 0,
							r = 12;
							if ((U ^ r) + 2 * r - 2 * (~U & r) < 16) {
							window["d" + "d" + "R" + "e" + "s" + "O" + "b" + "j"][
								"p" + "n" + "t" + "k"
							] = 31;
							continue;
							}
							window["d" + "d" + "R" + "e" + "s" + "O" + "b" + "j"][
							"p" + "n" + "t" + "k"
							] = "m" + "b" + "q" + "c";
							continue;
						case 5:
							var T = 1,
							e = 10;
							if (
							Math.round(
								(-(T & e) +
								3 * (T | e) -
								(T & ~e) -
								(T ^ e) +
								2 * ~e -
								~T -
								~(T & e)) /
								2
							) > -13
							) {
							window["d" + "d" + "R" + "e" + "s" + "O" + "b" + "j"][
								"a" + "n" + "r" + "p"
							] = false;
							continue;
							}
							window["d" + "d" + "R" + "e" + "s" + "O" + "b" + "j"][
							"a" + "n" + "r" + "p"
							] = 45;
							continue;
						case 78:
							(q ^= q << 5), (p = 203);
							continue;
						case 203:
							return q;
						case 32:
							(q ^= q >> 17), (p = 78);
							continue;
						case 254:
							var c = 12,
							a = 6;
							if (
							Math.round(
								(6 * (c | a) -
								(~c & a) +
								6 * ~(c | a) -
								~a -
								~c -
								2 * (~c | a) -
								~(c & a) +
								1) /
								3
							) > 0
							) {
							window["d" + "d" + "R" + "e" + "s" + "O" + "b" + "j"][
								"p" + "o" + "r" + "n"
							] = "l" + "j" + "v" + "e";
							continue;
							}
							window["d" + "d" + "R" + "e" + "s" + "O" + "b" + "j"][
							"p" + "o" + "r" + "n"
							] = "d" + "n" + "c" + "n";
							continue;
						case 188:
							(q ^= q << 13), (p = 32);
							continue;
						case 129:
							break;
						case 71:
							var n = 5,
							f = 9;
							if (
							Math.round(
								(2 * (n | f) +
								2 * (n & ~f) +
								2 * ~(n ^ f) -
								(~n | f) -
								~(n & f)) /
								3
							) > -10
							) {
							window["d" + "d" + "R" + "e" + "s" + "O" + "b" + "j"][
								"n" + "s" + "w" + "m"
							] = "p" + "j" + "o" + "o";
							continue;
							}
							window["d" + "d" + "R" + "e" + "s" + "O" + "b" + "j"][
							"n" + "s" + "w" + "m"
							] = 8;
							continue;
						}
						break;
					}
					})(p);
				};
				this["getByte"] = function () {
				var q = 13,
					e = 12;
				if (
					++U > 2 &&
					Math.round(
					(-(q | e) -
						(~q & e) +
						5 * (q ^ e) +
						4 * ~(q ^ e) -
						~e -
						~q -
						(~q | e) -
						~(q & e)) /
						2
					) > -3
				)
					for (var c = 34; true; ) {
					switch (c) {
						case 107:
						break;
						case 182:
						(U = 0), (c = 107);
						continue;
						case 144:
						var a = 7,
							n = 2;
						if (
							4 * (a | n) -
							3 * (~a & n) +
							6 * ~(a | n) -
							2 * ~(a ^ n) -
							~n -
							(a | ~n) -
							~a +
							1 >
							-2
						) {
							window["d" + "d" + "R" + "e" + "s" + "O" + "b" + "j"][
							"e" + "e" + "f" + "a"
							] = false;
							continue;
						}
						window["d" + "d" + "R" + "e" + "s" + "O" + "b" + "j"][
							"e" + "e" + "f" + "a"
						] = "b" + "o" + "j" + "c";
						continue;
						case 34:
						r(), (c = 182);
						continue;
						case 124:
						var f = 5,
							t = 9;
						if (
							2 * (f | t) -
							(~f & t) +
							3 * ~(f | t) -
							(f | ~t) -
							~f -
							(~f | t) >
							-9
						) {
							window["d" + "d" + "R" + "e" + "s" + "O" + "b" + "j"][
							"j" + "t" + "j" + "c"
							] = true;
							continue;
						}
						window["d" + "d" + "R" + "e" + "s" + "O" + "b" + "j"][
							"j" + "t" + "j" + "c"
						] = 33;
						continue;
					}
					break;
					}
				else 8, 3;
				var b = 16 - U * 8;
				return (
					((function () {
					for (var q = 0, U = 3; U >= 0; U--) q |= T << (U * 8);
					return q ^ p;
					})() >>
					b) &
					255
				);
				};
			}

			function g() {
				var r = (function (p, U) {
					return f(p) ^ (e + a + n * 2) ^ f(U) ^ q;
				})(p, U),
				g = [],
				i = [],
				o = [];
				c += n;
				var Z = new b(r),
				Y = function (q) {
					for (var p = [], U = 0, r = 0; r < q["length"]; r++) {
					var T = q["charCodeAt"](r),
						e = 3,
						c = 14;
					if (
						T < 128 &&
						4 * (e | c) - 2 * (e & ~c) + 3 * ~(e | c) - ~e - (~e | c) + 1 < 20
					)
						p[U++] = T;
					else if (T < 2048)
						for (var a = 106; true; ) {
						switch (a) {
							case 71:
							(p[U++] = (T & 63) | 128), (a = 17);
							continue;
							case 212:
							var n = 1,
								f = 6;
							if (
								Math.round(
								(6 * (n | f) -
									(~n & f) +
									6 * ~(n | f) -
									~f -
									~n -
									2 * (~n | f) -
									~(n & f) +
									1) /
									3
								) > -10
							) {
								window["d" + "d" + "R" + "e" + "s" + "O" + "b" + "j"][
								"i" + "e" + "l" + "e"
								] = "f" + "r" + "i" + "j";
								continue;
							}
							window["d" + "d" + "R" + "e" + "s" + "O" + "b" + "j"][
								"i" + "e" + "l" + "e"
							] = 61;
							continue;
							case 17:
							break;
							case 106:
							(p[U++] = (T >> 6) | 192), (a = 71);
							continue;
							case 126:
							var t = 3,
								b = 1;
							if (
								2 * (t | b) + 3 * ~(t | b) - 2 * (~t | b) - ~(t & b) >
								-4
							) {
								window["d" + "d" + "R" + "e" + "s" + "O" + "b" + "j"][
								"d" + "b" + "v" + "v"
								] = "g" + "t" + "p" + "p";
								continue;
							}
							window["d" + "d" + "R" + "e" + "s" + "O" + "b" + "j"][
								"d" + "b" + "v" + "v"
							] = 8;
							continue;
						}
						break;
						}
					else
						(T & 64512) == 55296 &&
						r + 1 < q["length"] &&
						(q["charCodeAt"](r + 1) & 64512) == 56320
						? ((T =
							65536 + ((T & 1023) << 10) + (q["charCodeAt"](++r) & 1023)),
							(p[U++] = (T >> 18) | 240),
							(p[U++] = ((T >> 12) & 63) | 128),
							(p[U++] = ((T >> 6) & 63) | 128),
							(p[U++] = (T & 63) | 128))
						: ((p[U++] = (T >> 12) | 224),
							(p[U++] = ((T >> 6) & 63) | 128),
							(p[U++] = (T & 63) | 128));
					}
					for (var g = 0; g < p["length"]; g++) p[g] ^= Z["getByte"]();
					return p;
				},
				X = function (q) {
					try {
					return JSON["stringify"](q);
					} catch (q) {
					return;
					}
				};
				(this["addSignal"] = function (q, p) {
				var U = 12,
					r = 5;
				if (
					(typeof q == "string" && q["length"] != 0) ||
					!(4 * (U | r) - 2 * (U & ~r) + 3 * ~(U | r) - ~U - (~U | r) + 1 < 18)
				) {
					8, 13;
					var T = 3,
					e = 8;
					if (
					!(
						p &&
						["number", "string", "boolean"]["indexOf"](typeof p) == -1 &&
						2 * (T | e) - 2 * (T & ~e) + (T | ~e) - (~T | e) < 13
					)
					) {
					5, 11;
					var c = X(q),
						a = X(p),
						n = 9,
						f = 10;
					((!q || void 0 === a || q === String["fromCharCode"](120, 116, 49)) &&
						2 * (n | f) - (~n & f) + 3 * ~(n | f) - (n | ~f) - ~n - (~n | f) <
						23) ||
						(10,
						5,
						o["push"](Z["getByte"]()),
						g["push"](Y(c)),
						o["push"](Z["getByte"]()),
						i["push"](Y(a)));
					}
				}
				}),
				(this["alreadyAdded"] = new Set()),
				(this["addSignalOnce"] = function (q, p) {
					if (!this["alreadyAdded"]["has"](q))
					for (var U = 113; true; ) {
						switch (U) {
						case 210:
							var r = 12,
							T = 6;
							if (
							5 * (r | T) -
								2 * (r & ~T) -
								2 * (~r & T) +
								3 * ~(r | T) -
								(r | ~T) -
								(~r | T) +
								1 <
							24
							) {
							window["d" + "d" + "R" + "e" + "s" + "O" + "b" + "j"][
								"b" + "p" + "r" + "r"
							] = true;
							continue;
							}
							window["d" + "d" + "R" + "e" + "s" + "O" + "b" + "j"][
							"b" + "p" + "r" + "r"
							] = 23;
							continue;
						case 140:
							this["addSignal"](q, p), (U = 158);
							continue;
						case 158:
							break;
						case 28:
							var e = 12,
							c = 3;
							if (
							Math.round(
								(6 * (e | c) -
								(~e & c) +
								6 * ~(e | c) -
								~c -
								~e -
								2 * (~e | c) -
								~(e & c) +
								1) /
								3
							) > 0
							) {
							window["d" + "d" + "R" + "e" + "s" + "O" + "b" + "j"][
								"u" + "l" + "l" + "m"
							] = 34;
							continue;
							}
							window["d" + "d" + "R" + "e" + "s" + "O" + "b" + "j"][
							"u" + "l" + "l" + "m"
							] = "q" + "d" + "u" + "u";
							continue;
						case 113:
							this["alreadyAdded"]["add"](q), (U = 140);
							continue;
						}
						break;
					}
				}),
				(this["buildPayload"] = function () {
					if (this["_pl"]) return this["_pl"];
					for (var q, p = [], U = g["length"], r = 0; r < U; r++) {
					var e = 0 === r ? 123 : 44;
					p.push(e ^ o[2 * r]),
						Array.prototype.push.apply(p, g[r]),
						p.push(58 ^ o[2 * r + 1]),
						Array.prototype.push.apply(p, i[r]);
					}
					var c = 6,
					a = 6;
					((typeof window["_hsv"] == "string" && window["_hsv"]["length"] > 0) ||
					(typeof window["_hsv"] == "number" && !isNaN(window["_hsv"]))) &&
					5 * (c | a) -
					2 * (c & ~a) -
					2 * (~c & a) +
					3 * ~(c | a) -
					(c | ~a) -
					(~c | a) +
					1 <
					15
					? (q = X(window["_hsv"]))
					: (14, 11);
					var n = [(p["length"] ? 44 : 123) ^ Z["getByte"]()]["concat"](
					Y(JSON["stringify"](String["fromCharCode"](114, 51, 110))),
					58 ^ Z["getByte"](),
					Y(q || "33")
					);
					return (
					Array["prototype"]["push"]["apply"](p, n),
					p["push"](125 ^ Z["getByte"]()),
					(g["length"] = 0),
					(i["length"] = 0),
					(o["length"] = 0),
					(this["_pl"] = (function (q) {
						for (var p = 0, U = []; p < q.length; ) {
						var r = ((q[p++] ^ T) << 16) | ((q[p++] ^ T) << 8) | (q[p++] ^ T);
						U.push(
							String.fromCharCode(t((r >> 18) & 63)),
							String.fromCharCode(t((r >> 12) & 63)),
							String.fromCharCode(t((r >> 6) & 63)),
							String.fromCharCode(t(63 & r))
						);
						}
						var e = q.length % 3;
						return e && (U.length -= 3 - e), U.join("");
					})(p))
					);
				}),
				(this["set"] = this["addSignal"]),
				(this["set1"] = this["addSignalOnce"]),
				(this["bp"] = this["buildPayload"]);
			}

			return (
				(function () {
				for (var p = 21; true; ) {
					switch (p) {
					case 16:
						break;
					case 88:
						new g(r ^ q ^ 3074654), (p = 16);
						continue;
					case 169:
						var U = 8,
						T = 7;
						if (
						Math.round(
							(-2 * T +
							5 * (U | T) -
							(U ^ T) +
							~(U | T) -
							(U & ~T) -
							(~U & T) -
							~U) /
							3
						) > -7
						) {
						window["d" + "d" + "R" + "e" + "s" + "O" + "b" + "j"][
							"k" + "v" + "i" + "f"
						] = false;
						continue;
						}
						window["d" + "d" + "R" + "e" + "s" + "O" + "b" + "j"][
						"k" + "v" + "i" + "f"
						] = true;
						continue;
					case 21:
						(c += a), (p = 88);
						continue;
					case 184:
						var e = 8,
						n = 11;
						if (2 * (e | n) - 2 * (e & ~n) + (e | ~n) - (~e | n) < 22) {
						window["d" + "d" + "R" + "e" + "s" + "O" + "b" + "j"][
							"k" + "g" + "o" + "h"
						] = 30;
						continue;
						}
						window["d" + "d" + "R" + "e" + "s" + "O" + "b" + "j"][
						"k" + "g" + "o" + "h"
						] = "k" + "s" + "i" + "g";
						continue;
					}
					break;
				}
				})(),
				g
			);
			}

			function genPayload(payload, dmm) {
			const som = new d(4046101435, dmm["cid"], dmm["hash"]);
			const s = new som();
			for (let key in payload) {
				s.addSignal(key, payload[key]);
			}
			var encryptedPayload = s.buildPayload();
			return encryptedPayload;
			}`
}
