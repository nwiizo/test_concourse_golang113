package fizzbuzz

import "testing"

func TestFizzbuzz(t *testing.T) {
	cases := []struct {
		nu   int
		want string
	}{
		{1, "1"},
		{2, "2"},
		{3, "Fizz"},
		{4, "4"},
		{5, "Buzz"},
		{15, "FizzBuzz"},
		{30, "FizzBuzz"},
	}
	for _, fb := range cases {
		actual := Fizzbuzz(fb.nu)
		if actual != fb.want {
			t.Errorf("Fizzbuzz(%d): expected %q, actual %q", fb.nu, fb.want, actual)
		}
	}
}
