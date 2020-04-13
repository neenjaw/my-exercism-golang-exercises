// +build !bonus

package robotname

import (
	"testing"
)

func TestNameValid(t *testing.T) {
	n := New().getName(t, false)
	if !namePat.MatchString(n) {
		t.Errorf(`Invalid robot name %q, want form "AA###".`, n)
	}
}

func TestNameSticks(t *testing.T) {
	r := New()
	n1 := r.getName(t, false)
	n2 := r.getName(t, true)
	if n2 != n1 {
		t.Errorf(`Robot name changed.  Now %s, was %s.`, n2, n1)
	}
}

func TestSuccessiveRobotsHaveDifferentNames(t *testing.T) {
	n1 := New().getName(t, false)
	n2 := New().getName(t, false)
	if n1 == n2 {
		t.Errorf(`Robots with same name.  Two %s's.`, n1)
	}
}

func TestResetName(t *testing.T) {
	r := New()
	n1 := r.getName(t, false)
	r.Reset()
	n2 := r.getName(t, false)
	if n2 == n1 {
		t.Errorf(`Robot name not cleared on reset.  Still %s.`, n1)
	}
}

// // Note if you go for bonus points, this benchmark likely won't be
// // meaningful.  Bonus thought exercise, why won't it be meaningful?
// func BenchmarkName(b *testing.B) {
// 	// Benchmark combined time to create robot and name.
// 	for i := 0; i < b.N; i++ {
// 		New().getName(b, false)
// 	}
// }
// var maxNames = 26 * 26 * 10 * 10 * 10

// func TestCollisions(t *testing.T) {
// 	// Test uniqueness for new robots.
// 	for i := len(seen); i <= maxNames-600000; i++ {
// 		_ = New().getName(t, false)
// 	}

// 	// Test that names aren't recycled either.
// 	r := New()
// 	for i := len(seen); i < maxNames; i++ {
// 		r.Reset()
// 		_ = r.getName(t, false)
// 	}

// 	// Test that name exhaustion is handled more or less correctly.
// 	_, err := New().Name()
// 	if err == nil {
// 		t.Fatalf("should return error if namespace is exhausted")
// 	}
// }
