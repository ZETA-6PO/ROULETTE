package main

import (
	"fmt"
	"log"
	crypto_rand "crypto/rand"
    "encoding/binary"
    math_rand "math/rand"
)

func init() {
    var b [8]byte
    _, err := crypto_rand.Read(b[:])
    if err != nil {
        panic("cannot seed math/rand package with cryptographically secure random number generator")
    }
    math_rand.Seed(int64(binary.LittleEndian.Uint64(b[:])))
}


func belongToInterval(x int, min int, max int) bool {
	if x >= min && x <= max {
		return true
	}
	return false
}

func wichColumn(x int) int {
	if !belongToInterval(x,1,36) {
		return 0
	}
	return x%3
}

//1 = red;
//2 = black
func wichColor(x int) int {
	if !belongToInterval(x,1,36) {
		return -1
	}
	return x%3+1
}

func play(first_12 float32, second_12 float32,
	third_12 float32, zero float32, red float32, black float32,
	pair float32, odd float32, c_1_18 float32, c_19_36 float32, first_2_1 float32,
	second_2_1 float32, third_2_1 float32) (float32, int) {
	//select random number
	var number = math_rand.Intn(38)
	var returned float32 = 0
	log.Println("Number picked : ", number)
	
	if belongToInterval(number, 1, 12) && first_12 > 0 {
		returned += first_12 * 3
	}
	if belongToInterval(number, 13, 24) && second_12 > 0 {
		returned += second_12 * 3
	}
	if belongToInterval(number, 25, 36) && third_12 > 0{
		returned += third_12 * 3
		
	}
	if number == 0 && zero >= 0 {
		returned += zero * 35
	}
	if wichColumn(number)!=0 {
		if wichColumn(number) == 1 && first_2_1 > 0 {
			returned += first_2_1 * 3
		}else if wichColumn(number) == 2 && second_2_1 > 0{
			returned += second_2_1 * 3
		}else if wichColumn(number) == 3 && third_2_1 > 0{
			returned += third_2_1*3
		}
	}
	if c_1_18 > 0 && belongToInterval(number, 1, 18) {
		returned += c_1_18*2
	}
	if c_19_36 > 0 && belongToInterval(number, 19, 36) {
		returned += c_19_36*2
	}
	if red > 0 && wichColor(number) > 0 {
		returned += red*2
	}
	if black > 0 && wichColor(number) > 0 {
		returned += black*2
	}
	return float32(returned), number
}

func main() {
	test()
}

func test() {
	var start float32 = 2000
	picked := make([]int, 0)
	//scramble the number to create a number historic
	for i := 0; i < 10; i++ {
		_, num := play(0,0,0,0,0,0,0,0,0,0,0,0,0)
		picked = append(picked, num)
	}
	for i := 0; i < 1000; i++ {
		if belongToInterval(picked[len(picked)-1], 1, 12) && belongToInterval(picked[len(picked)-2], 1, 12) {
			var mise float32 = 100
			start -= mise
			rt, num:=play(0,float32(mise/2),float32(mise/2),0,0,0,0,0,0,0,0,0,0)
			picked = append(picked, num)
			start += rt
		}else if belongToInterval(picked[len(picked)-1], 13, 24) && belongToInterval(picked[len(picked)-2], 13, 24){
			var mise float32 = 100
			start -= mise
			rt, num:=play(float32(mise/2),0,float32(mise/2),0,0,0,0,0,0,0,0,0,0)
			picked = append(picked, num)
			start += rt
		}else if belongToInterval(picked[len(picked)-1], 25, 36) && belongToInterval(picked[len(picked)-2], 25, 36){
			var mise float32 = 100
			start -= mise
			rt, num:=play(float32(mise/2),float32(mise/2),0,0,0,0,0,0,0,0,0,0,0)
			picked = append(picked, num)
			start += rt
		}else{
			_, num := play(0,0,0,0,0,0,0,0,0,0,0,0,0)
			picked = append(picked, num)
		}
	};
	fmt.Println("ended:",start)
}