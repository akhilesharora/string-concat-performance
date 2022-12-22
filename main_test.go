package string_concat_performance

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

const cStr = "string"

var (
	str     = "string"
	longStr = "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."
)

func BenchmarkNaivePlus(b *testing.B) {

	var s string
	for n := 0; n < b.N; n++ {
		s += str
	}
}

func BenchmarkLongNaivePlus(b *testing.B) {

	var s string
	for n := 0; n < b.N; n++ {
		s += longStr
	}
	_ = s
}

func BenchmarkConstNaivePlus(b *testing.B) {

	var s string
	for n := 0; n < b.N; n++ {
		s += cStr
	}
	_ = s
}

func BenchmarkJoin(b *testing.B) {

	var s string
	for n := 0; n < b.N; n++ {
		s = strings.Join([]string{s, str}, "")
	}
	_ = s
}

func BenchmarkLongJoin(b *testing.B) {

	var s string
	for n := 0; n < b.N; n++ {
		s = strings.Join([]string{s, longStr}, "")
	}
	_ = s
}

func BenchmarkConstJoin(b *testing.B) {

	var s string
	for n := 0; n < b.N; n++ {
		s = strings.Join([]string{s, cStr}, "")
	}
	_ = s
}

func BenchmarkSprintf(b *testing.B) {

	var s string
	for n := 0; n < b.N; n++ {
		s = fmt.Sprintf("%s %s", s, str)
	}
	_ = s
}

func BenchmarkLongSprintf(b *testing.B) {

	var s string
	for n := 0; n < b.N; n++ {
		s = fmt.Sprintf("%s %s", s, longStr)
	}
	_ = s
}

func BenchmarkConstSprintf(b *testing.B) {

	var s string
	for n := 0; n < b.N; n++ {
		s = fmt.Sprintf("%s %s", s, cStr)
	}
	_ = s
}

func BenchmarkStringBuilder(b *testing.B) {

	var sb strings.Builder
	for n := 0; n < b.N; n++ {
		sb.WriteString(str)
	}
	_ = sb
}

func BenchmarkLongStringBuilder(b *testing.B) {

	var sb strings.Builder
	for n := 0; n < b.N; n++ {
		sb.WriteString(longStr)
	}
	_ = sb
}

func BenchmarkConstStringBuilder(b *testing.B) {

	var sb strings.Builder
	for n := 0; n < b.N; n++ {
		sb.WriteString(cStr)
	}
	_ = sb
}

func BenchmarkBytesBuffer(b *testing.B) {

	var sf bytes.Buffer
	for n := 0; n < b.N; n++ {
		sf.WriteString(str)
	}
	_ = sf
}

func BenchmarkLongBytesBuffer(b *testing.B) {

	var sf bytes.Buffer
	for n := 0; n < b.N; n++ {
		sf.WriteString(longStr)
	}
	_ = sf
}

func BenchmarkConstBytesBuffer(b *testing.B) {

	var sf bytes.Buffer
	for n := 0; n < b.N; n++ {
		sf.WriteString(cStr)
	}
	_ = sf
}
