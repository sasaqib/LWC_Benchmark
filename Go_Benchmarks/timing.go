package main

import (
	"encoding/binary"
	"machine"
	"time"
)

// #define FrameBitsIV 0x10
const FrameBitsIV uint32 = 0x10
const FrameBitsAD uint32 = 0x30
const FrameBitsPC uint32 = 0x50 //Framebits for plaintext/ciphertext
const FrameBitsFinalization uint32 = 0x70

const NROUND1 uint32 = 128 * 5
const NROUND2 uint32 = 128 * 8

func uint32touint8(input []uint32, position uint32) uint8 {
	r := uint8(0)

	//input = []uint32{1, 2, 3, 4, 4}
	r = uint8(input[position/4] >> 8 * (position % 4))

	return r

}
func i64tob(val uint64) []byte {
	r := make([]byte, 8)
	for i := uint64(0); i < 8; i++ {
		r[i] = byte((val >> (i * 8)) & 0xff)
	}
	return r
}

func btoi64(val []byte) uint64 {
	r := uint64(0)
	for i := uint64(0); i < 8; i++ {
		r |= uint64(val[i]) << (8 * i)
	}
	return r
}

func i32tob(val uint32) []uint8 {
	r := make([]uint8, 4)
	for i := uint32(0); i < 4; i++ {
		r[i] = byte((val >> (8 * i)) & 0xff)
	}
	return r
}

func btoi32(val []byte) uint32 {
	r := uint32(0)
	for i := uint32(0); i < 4; i++ {
		r |= uint32(val[i]) << (8 * i)
	}
	return r
}
func uintsToBytes(vs []uint32) []byte {
	buf := make([]byte, len(vs)*4)
	for i, v := range vs {
		binary.LittleEndian.PutUint32(buf[i*4:], v)
	}
	return buf
}

func bytesToUints(vs []byte) []uint32 {
	out := make([]uint32, (len(vs))/4)
	for i := range out {
		out[i] = binary.LittleEndian.Uint32(vs[i*4:])
	}
	return out
}
func bytesToUintsi(vs []byte, i int) []uint32 {
	out := make([]uint32, len(vs)/4)
	for i := range out {
		out[i] = binary.LittleEndian.Uint32(vs[i*4:])
	}
	return out
}
func bytesToUint64(vs []byte) []uint64 {
	out := make([]uint64, len(vs)/8)
	for i := range out {
		out[i] = binary.LittleEndian.Uint64(vs[i*8:])
	}
	return out
}

func uint32ToUint64(arg []uint32) []uint64 {
	out := make([]uint64, len(arg)/2)
	for i := range out {
		out[i] = binary.LittleEndian.Uint64(uintsToBytes(arg[i*2:]))
	}
	return out
}

var (
	uart = machine.UART1
	tx   = machine.UART1_TX_PIN
	rx   = machine.UART1_RX_PIN
)

func configure() {
	uart.Configure(machine.UARTConfig{
		BaudRate: 115200,
		TX:       tx,
		RX:       rx,
	})
}

// func uint32tobytes(source []uint32) byte{

// 	buf := new(bytes.Buffer)
//     source = []uint32{1, 2, 3}
//     err := binary.Write(buf, binary.LittleEndian, source)
//     if err != nil {
//         fmt.Println("binary.Write failed:", err)
//     }
//     fmt.Printf("Encoded: % x\n", buf.Bytes())

//		return
//	}

//fmt.Print(f)

// for i := 0; uint64(i) < clen; i++ {
// 	test1 := uintsToBytes(c)
// 	fmt.Print(test1[i]) // printing ciphertext iterated
// 	fmt.Print("\n")
// 	// fmt.Print(test1)
// 	// fmt.Print("\n")
// }
// fmt.Print("reversed: ")
// crypto_aead_decrypt(m1, mlen, nsec, c1, clen, ad, adlen, npub, key)
// for i := 0; i < 16; i++ {
// 	fmt.Print("\n")
// 	fmt.Printf("%d", bits.ReverseBytes32(m1[i]))
// 	//fmt.Print("\n")
// }
// fmt.Print("not reversed: ")
// for i := 0; i < 16; i++ {
// 	fmt.Print("\n")
// 	fmt.Printf("%d", (m1[i]))
// 	//fmt.Print("\n")
// }
// fmt.Print("\n")

// fmt.Print("message: ", m1)
// fmt.Print("\n")

// for i := 0; i < 1; i++ {
// 	//fmt.Printf("%x", bits.ReverseBytes32(m1[i]))

// 	// finalm := bits.ReverseBytes32(m1[i])
// 	fmt.Printf("%d", bits.ReverseBytes32(m1[i]))

// 	// fmt.Printf("%d\n", finalm)
// 	// fmt.Printf("%02X\n, finalm")
// 	//fmt.Printf("%d", bits.ReverseBytes32(m1[i]))

// 	// // full array whole array
// 	// println("Printing m[i] reversebytes32")
// 	// println((bits.ReverseBytes32(m1[i])))
// 	// println("Printing m[i]")
// 	// println(m1[i])
// 	// println("Printing m")
// 	// println(m1)
// 	//}
// }

// duration := time.Since(start)
// fmt.Println(duration.Nanoseconds())

// fmt.Println("in nanoseconds: ", duration.Nanoseconds())
// fmt.Println("in milliseconds: ", duration.Milliseconds())
// fmt.Println("in seconds: ", duration.Seconds())
func timing_functions() {
	//m := []uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	//4-byte message
	m := []uint8{97, 98, 99, 100}
	mlen := uint64(4)
	c1 := []uint8{38, 108, 231, 63, 176, 17, 130, 1, 62, 140, 89, 58} //abcd

	//8-byte message
	// m := []uint8{97, 98, 99, 100, 101, 102, 103, 104}
	// mlen := uint64(8)
	// c1 := []uint8{38, 108, 231, 63, 31, 7, 57, 89, 151, 238, 183, 66, 129, 18, 180, 31} //abcdefgh

	// //16-byte message
	// m := []uint8{97, 98, 99, 100, 101, 102, 103, 104, 105, 106, 107, 108, 109, 110, 111, 112}
	// mlen := uint64(16)
	//c1 := []uint8{38, 108, 231, 63, 31, 7, 57, 89, 138, 230, 8, 48, 223, 26, 242, 78, 239, 209, 195, 134, 56, 130, 166, 180} //abcdefghijklmnop

	c := []uint32{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	clen := mlen + 8
	ad := []uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	adlen := uint64(0)
	npub := []uint8{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	nsec := []uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	key := []uint8{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	m1 := []uint32{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	for i := 0; i < 250; i++ {
		crypto_aead_encrypt(c, clen, m, mlen, ad, adlen, nsec, npub, key)
		crypto_aead_decrypt(m1, mlen, nsec, c1, clen, ad, adlen, npub, key)
	}

}

// func test_encrypt() {

// 	//m := []uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
// 	m := []uint8{0x61, 0x62, 0x63, 0x64}
// 	//m := []uint8{ 0x65, 0x66, 0x67, 0x68}
// 	//m := []uint8{0x61, 0x62, 0x63, 0x64, 0x65, 0x66, 0x67, 0x68}
// 	//  m := []uint8{0x61, 0x62, 0x63, 0x64, 0x65, 0x66, 0x67, 0x68, 0x69, 0x70, 0x71, 0x72, 0x73, 0x74, 0x75, 0x76}
// 	c := []uint32{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
// 	mlen := uint64(4)
// 	clen := mlen + 8
// 	ad := []uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
// 	adlen := uint64(0)
// 	npub := []uint8{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
// 	nsec := []uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
// 	key := []uint8{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}

// 	crypto_aead_encrypt(c, clen, m, mlen, ad, adlen, nsec, npub, key)

// }
// func test_decrypt() {

// 	c1 := []uint8{38, 108, 231, 63, 176, 17, 130, 1, 62, 140, 89, 58} //abcd
// 	//c1 := []uint8{38, 108, 231, 63, 31, 7, 57, 89, 151, 238, 183, 66, 129, 18, 180, 31} //abcdefgh
// 	//c1 := []uint8{38, 108, 231, 63, 31, 7, 57, 89, 138, 230, 8, 48,223, 26, 242, 78, 239, 209, 195, 134, 56, 130, 166, 180} //abcdefghijklmnop
// 	m1 := []uint32{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
// 	mlen := uint64(4)
// 	nsec := []uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
// 	npub := []uint8{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
// 	ad := []uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
// 	adlen := uint64(0)
// 	key := []uint8{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
// 	clen := uint64(12)

// 	crypto_aead_decrypt(m1, mlen, nsec, c1, clen, ad, adlen, npub, key)

// }
func main() {
	//Timing Encrypt & Decrypt

	for i := 0; i < 10; i++ {
		println("%d", i)
		time.Sleep(500 * time.Millisecond)
	}

	// println("Start timing now for (4-bytes): ")
	println("Start timing now for (16-bytes): ")
	// println("Start timing now for (16-bytes): ")
	timing_functions()
	println("Finished Program!!!")

}

func state_update(state []uint32, key []uint8, number_of_steps uint32) {

	var t1, t2, t3, t4 uint32
	//var data = binary.BigEndian.Uint32(key)
	for i := 0; uint32(i) < number_of_steps; i += 128 {

		t1 = (state[1] >> 15) | (state[2] << 17) // 47 = 1*32+15
		t2 = (state[2] >> 6) | (state[3] << 26)  // 47 + 23 = 70 = 2*32 + 6
		t3 = (state[2] >> 21) | (state[3] << 11) // 47 + 23 + 15 = 85 = 2*32 + 21
		t4 = (state[2] >> 27) | (state[3] << 5)  // 47 + 23 + 15 + 6 = 91 = 2*32 + 27
		// var z uint32 = uint32(key[0])
		state[0] = state[0] ^ ((t1) ^ (^(t2 & t3)) ^ (t4) ^ (btoi32(key[0:4])))

		t1 = (state[2] >> 15) | (state[3] << 17)
		t2 = (state[3] >> 6) | (state[0] << 26)
		t3 = (state[3] >> 21) | (state[0] << 11)
		t4 = (state[3] >> 27) | (state[0] << 5)
		// z = uint(key[1])
		state[1] ^= t1 ^ (^(t2 & t3)) ^ t4 ^ (btoi32(key[4:8]))

		t1 = (state[3] >> 15) | (state[0] << 17)
		t2 = (state[0] >> 6) | (state[1] << 26)
		t3 = (state[0] >> 21) | (state[1] << 11)
		t4 = (state[0] >> 27) | (state[1] << 5)
		// z = uint(key[2])
		state[2] ^= t1 ^ (^(t2 & t3)) ^ t4 ^ (btoi32(key[8:12]))

		t1 = (state[0] >> 15) | (state[1] << 17)
		t2 = (state[1] >> 6) | (state[2] << 26)
		t3 = (state[1] >> 21) | (state[2] << 11)
		t4 = (state[1] >> 27) | (state[2] << 5)
		// z = uint(key[3])
		state[3] ^= t1 ^ (^(t2 & t3)) ^ t4 ^ (btoi32(key[12:16]))
	}
}

func initialization(key []uint8, iv []byte, state []uint32) {

	// var i int

	for i := 0; i < 4; i++ {
		state[i] = 0

	}

	state_update(state, key, NROUND2)

	for i := 0; i < 3; i++ {

		state[1] ^= FrameBitsIV
		state_update(state, key, NROUND1)
		state[3] ^= (btoi32(iv[(i * 4) : (i+1)*4]))

	}

}

func process_ad(k []uint8, ad []uint8, adlen uint64, state []uint32) {

	var i int
	//var j uint32

	for i := 0; uint64(i) < (adlen >> 2); i++ {
		state[1] ^= FrameBitsAD
		state_update(state, k, NROUND1)
		// state[3] ^= (btoi32(ad[i*4:(i+1)*4]))
		state[3] ^= (btoi32(ad[(i * 4) : (i+1)*4]))
		//fmt.Printf("main block", state)

	}

	if (adlen & 3) > 0 {

		// fmt.Print("If statement ad")

		state[1] ^= FrameBitsAD
		state_update(state, k, NROUND1)
		for j := 0; uint64(j) < (adlen & 3); j++ {
			(uintsToBytes(state))[12+j] ^= ad[(i<<2)+(j)]
		}
		state[1] ^= uint32(adlen & 3)
		//fmt.Printf("partial SECONDLOOP")

	}
	//fmt.Printf("hi")

}

func crypto_aead_encrypt(c []uint32, clen uint64, m []uint8, mlen uint64, ad []uint8, adlen uint64, nsec []uint8, npub []uint8, k []uint8) {
	//c := []uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	// var i uint64
	var j uint64
	mac := []uint32{0, 0, 0, 0, 0, 0, 0, 0}
	//var state []uint32 = []uint32{0, 0, 0, 0}
	state := []uint32{0, 0, 0, 0}
	//test := uint32(m[0])
	//check := uint32(0)

	//initialization stage
	initialization(k, npub, state)

	//process the associated data
	process_ad(k, ad, adlen, state)

	//process the plaintext
	for l := 0; uint64(l) < (mlen >> 2); l++ {
		state[1] ^= FrameBitsPC
		state_update(state, k, NROUND2)
		//state[3] ^= btoi32(m[l*4 : (l+1)*4])
		state[3] ^= btoi32(m[l*4 : (l+1)*4])
		//state[3] ^=
		//z := state[2]
		(c)[l] = state[2] ^ btoi32(m[(l*4):(l+1)*4])
		// fmt.Print(z)
		// fmt.Print("\n")
		//fmt.Printf("%X", c[i])

	}
	// fmt.Printf("\n")
	// fmt.Printf("\n")
	// for i := 0; i < int(clen); i++ {
	// 	fmt.Printf("%d", i32tob(c[i]))
	// }
	// // //if mlen is not a multiple of 4, we process the remaining bytes
	// if (mlen & 3) > 0 {

	// 	state[1] ^= FrameBitsPC
	// 	state_update(state, k, NROUND2)
	// 	for j = 0; j < (mlen & 3); j++ {
	// 		(uintsToBytes(state))[12+j] ^= m[(i<<2)+j]
	// 		(c)[(i<<2)+j] = (state)[8+j] ^ btoi32(m[((i<<2)+j)*4:(i+1)*4])
	// 	}
	// 	uint32ToUint64(state)[1] ^= mlen & 3
	// }

	//finalization stage, we assume that the tag length is 8 bytes
	state[1] ^= FrameBitsFinalization
	state_update(state, k, NROUND2)
	(mac)[0] = state[2]

	state[1] ^= FrameBitsFinalization
	state_update(state, k, NROUND1)
	(mac)[1] = state[2]

	//clen = (mlen + 8) //True C_length

	//clen = (mlen + 8)

	for j = 0; j < 8; j++ {
		c[1+j] = mac[j]
	}

	//Print statements

	// fmt.Printf("\n")
	// fmt.Printf("\n")

	// fmt.Printf("\n")
	// fmt.Printf("\n")
	// for i := 0; i < int(clen); i++ {
	// 	fmt.Printf("%X", i32tob(c[i]))
	// }
	// fmt.Printf("\n")
	// fmt.Printf("\n")

	// for i := 0; i < int(clen); i++ {
	// 	fmt.Printf("%X", uintsToBytes(c[i*4:(i+1)*4]))
	// }

}

func crypto_aead_decrypt(m []uint32, mlen uint64, nsec []uint8, c []uint8, clen uint64, ad []uint8, adlen uint64, npub []uint8, k []uint8) {

	var i uint64
	var j uint64
	//check := uint32(0)
	//m = []uint8{0, 0, 0, 0}
	mac := []uint32{0, 0, 0, 0, 0, 0, 0, 0}
	state := []uint32{0, 0, 0, 0}

	// mlen = 16
	// clen = mlen + 8
	// fmt.Println("\nShow me ciphertext")
	// fmt.Println(c)

	// fmt.Println("Mlen is: ", mlen)
	// fmt.Println("Clen is: ", clen)

	//initialization stage
	initialization(k, npub, state)

	//process the associated data
	process_ad(k, ad, adlen, state)

	//fmt.Printf("\nShow me mlen %d\n", mlen)

	//process the ciphertext
	for i = 0; i < (mlen >> 2); i++ {
		state[1] ^= FrameBitsPC
		state_update(state, k, NROUND2)
		(m)[i] = state[2] ^ bytesToUints(c)[i]
		// fmt.Printf("\n Show me m please %d", bytesToUints(m))
		//fmt.Println(m)
		state[3] ^= (m)[i]
		// fmt.Printf("\n Show me m in for loop %d", m)
		// fmt.Printf("\n Show me state[2] %d", state[2])
		// fmt.Printf("\n Show me whats in c %d", c[i])
	}
	if (mlen & 3) > 0 {
		state[1] ^= FrameBitsPC
		state_update(state, k, NROUND2)
		for j = 0; j < (mlen & 3); j++ {
			m[(i<<2)+j] = bytesToUints(c)[((i<<2)+j)] ^ (state)[(8+j)]

			(state)[12+j] ^= m[(i<<2)+j]
		}
		uint32ToUint64(state)[1] ^= mlen & 3
	}
	state[1] ^= FrameBitsFinalization
	state_update(state, k, NROUND2)
	//(bytesToUints(mac))[0] = state[2]
	(mac)[0] = state[2]

	// fmt.Print("\n")
	// fmt.Print("\n")

	// fmt.Print("Length of message: ", mlen)
	state[1] ^= FrameBitsFinalization
	state_update(state, k, NROUND1)
	//(bytesToUints(mac))[1] = state[2]
	(mac)[1] = state[2]

	// mac[2] = 0
	// mac[3] = 0
	// mac[4] = 0
	// mac[5] = 0
	// mac[6] = 0
	// mac[7] = 0

	//clen = (mlen + 8) //True C_length
	//verification of the authentication tag
	// fmt.Print("\n")
	// fmt.Println("Clen is: ", clen)

	// fmt.Println("Printing Message")

	// fmt.Println(byteArraytoUintArray(m, mlen))

	// for i := 0; uint64(i) < mlen; i++ {

	// 	println(m[i])
	// 	//println(bytetoUintArray(m[i]))

	// }
	// for i := 0; uint64(i) < mlen; i++ {

	// 	println(bytetoUintArray(m[i], mlen))
	// 	//println(bytetoUintArray(m[i]))

	// }

	//println("Done printing message")

	// for l := 0; l < 8; l++ {
	// 	check |= mac[i] ^ bytesToUints(c)[(clen-8+j)]
	// }
	// if check == 0 {
	// 	return 0
	// } else {
	// 	return -1
	// }

}

// Video on TinyGO and UART: https://youtu.be/Oug_CYgs45E
// Video on starting a TinyGO Project: https://youtu.be/Fl5eFIYU1Xg

//useful commands to remember:
// to flash the pico in terminal: tinygo flash -target=pico timing.go
// to obtain the uf2: tinygo build -o golang_4byte.uf2 -target=pico timing.go

// Manual Stopwatch w/ Human Error
// 1st Trial: 11:07:271
// 2nd Trial: 11:06:000
// 3rd Trial: 11:07:395

// Manual Stopwatch, Russel Code
// Encrypt: 6:44:868 (1 million times)
//Decrypt: ~4:45 (100,000 times)

//To-DO list
// 1. Run rust on the same pico
// 2. Use GO on one of the decrypt or encrypt functions

// CT for 8-bytes "abcdefgh" is 38 108 231 63 31 7 57 89 151 238 183 66 129 18 180 31
// CT for 16-byte "a-p" is: {38, 108, 231, 63, 31, 7, 57, 89, 138, 230, 8, 48, 223, 26, 242, 78, 239, 209, 195, 134, 56, 130, 166, 180}

// 2/7/2023 Tests
// Running 500,000 rounds

// Round1:
// 04-bytes time: 2:16:25
// 08-bytes time: 2:46:60
// 16-bytes time: 4:10:73

// Round2:
// 04-bytes time: 2:16:28
// 08-bytes time: 2:46:57
// 16-bytes time: 4:10:96

// Round3:
// 04-bytes time: 2:16:27
// 08-bytes time: 2:46:61
// 16-bytes time: 4:10:99
