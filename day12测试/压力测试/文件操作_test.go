package main

import "testing"

func BenchmarkTruncfile(b *testing.B){
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		var blue =color{"blue",[4]int{1,54,65,24}}
		truncfile(blue)
	}
}


func BenchmarkLoadfile(b *testing.B){
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		var isBlue color
		loadfile(isBlue)
	}
}