package main

import (
	"flag"
	"fmt"
)

var mPrice int
var aveSuccess int
var rateSuccess float64

func init() {
	flag.IntVar(&mPrice, "mp", 5000, "price of m")
	flag.IntVar(&aveSuccess, "asg", 100000, "average Success gain")
	flag.Float64Var(&rateSuccess, "rs", 0.3, "rate Success")
	flag.Parse()
}

func main() {

	fmt.Printf("%f", gainGacha(mPrice, aveSuccess, rateSuccess))

}

func gainGacha(mp int, asg int, rs float64) float64 {
	return float64(-mp*4)*(1-rs) + float64(asg-mp*5)*rs
}
