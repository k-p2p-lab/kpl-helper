package distribution_test

import (
	"math"
	"math/rand"
	"testing"
	"time"

	. "github.com/k-p2p-lab/kpl-helper/distribution"
)

// --- Helper for floating-point comparison ---
func approxEqual(a, b, tol float64) bool {
	return math.Abs(a-b) < tol
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

// --- Tests ---

func TestPoissonRandom(t *testing.T) {
	lambda := 4.0
	nSamples := 10000
	sum := 0.0

	for i := 0; i < nSamples; i++ {
		v := PoissonRandom(lambda)
		if v < 0 {
			t.Fatalf("PoissonRandom produced negative value: %d", v)
		}
		sum += float64(v)
	}

	mean := sum / float64(nSamples)
	if !approxEqual(mean, lambda, 0.3) {
		t.Errorf("Poisson mean mismatch: got %.2f, expected ~%.2f", mean, lambda)
	}
}

func TestExponentialRandom(t *testing.T) {
	lambda := 2.0
	nSamples := 10000
	sum := 0.0

	for i := 0; i < nSamples; i++ {
		v := ExponentialRandom(lambda)
		if v < 0 {
			t.Fatalf("ExponentialRandom produced negative value: %f", v)
		}
		sum += v
	}

	mean := sum / float64(nSamples)
	expected := 1.0 / lambda
	if !approxEqual(mean, expected, 0.1) {
		t.Errorf("Exponential mean mismatch: got %.2f, expected ~%.2f", mean, expected)
	}
}

func TestNormalRandom(t *testing.T) {
	mu := 5.0
	sigma := 2.0
	nSamples := 10000
	sum := 0.0

	for i := 0; i < nSamples; i++ {
		sum += NormalRandom(mu, sigma)
	}

	mean := sum / float64(nSamples)
	if !approxEqual(mean, mu, 0.2) {
		t.Errorf("Normal mean mismatch: got %.2f, expected ~%.2f", mean, mu)
	}
}

func TestBinomialRandom(t *testing.T) {
	n := 10
	p := 0.3
	nSamples := 10000
	sum := 0.0

	for i := 0; i < nSamples; i++ {
		v := BinomialRandom(n, p)
		if v < 0 || v > n {
			t.Fatalf("BinomialRandom produced out-of-range value: %d", v)
		}
		sum += float64(v)
	}

	mean := sum / float64(nSamples)
	expected := float64(n) * p
	if !approxEqual(mean, expected, 0.3) {
		t.Errorf("Binomial mean mismatch: got %.2f, expected ~%.2f", mean, expected)
	}
}

func TestUniformRandom(t *testing.T) {
	a, b := 10.0, 20.0
	nSamples := 10000

	for i := 0; i < nSamples; i++ {
		v := UniformRandom(a, b)
		if v < a || v >= b {
			t.Fatalf("UniformRandom out of range: %f", v)
		}
	}
}

func TestParetoRandom(t *testing.T) {
	xm, alpha := 1.0, 2.0
	nSamples := 10000
	sum := 0.0

	for i := 0; i < nSamples; i++ {
		v := ParetoRandom(xm, alpha)
		if v < xm {
			t.Fatalf("ParetoRandom below xm: %f", v)
		}
		sum += v
	}

	mean := sum / float64(nSamples)
	expected := (alpha * xm) / (alpha - 1)
	if !approxEqual(mean, expected, expected*0.2) {
		t.Errorf("Pareto mean mismatch: got %.2f, expected ~%.2f", mean, expected)
	}
}
