package calcserver

import (
	"fmt"
	"testing"

	cpb "../../calcpb"
)

func TestProcessPrime(t *testing.T) {
	var testTable = map[int64][]cpb.PrimeNoDecompResponse{
		34: {
			cpb.PrimeNoDecompResponse{PrimeFactor: 2, Power: 1},
			cpb.PrimeNoDecompResponse{PrimeFactor: 17, Power: 1},
		},
		3500: {
			cpb.PrimeNoDecompResponse{PrimeFactor: 2, Power: 2},
			cpb.PrimeNoDecompResponse{PrimeFactor: 5, Power: 3},
			cpb.PrimeNoDecompResponse{PrimeFactor: 7, Power: 1},
		},
		2: {
			cpb.PrimeNoDecompResponse{PrimeFactor: 2, Power: 1},
		},
		8: {
			cpb.PrimeNoDecompResponse{PrimeFactor: 2, Power: 3},
		},
	}

	for n, assertTable := range testTable {
		testIdx := 0

		mockSend := func(r *cpb.PrimeNoDecompResponse) error {
			expectedFact := assertTable[testIdx].PrimeFactor
			actualFact := r.GetPrimeFactor()

			if actualFact != expectedFact {
				t.Fatalf("The prime factor is incorrect: actual: %d expected: %d", actualFact, expectedFact)
			}

			expectedPow := assertTable[testIdx].Power
			actualPow := r.GetPower()

			if actualPow != expectedPow {
				t.Fatalf("The prime factor is incorrect: actual: %d expected: %d", actualPow, expectedPow)
			}

			testIdx++

			return nil
		}

		fmt.Printf("Testing number %d \n", n)
		processPrimes(n, mockSend)

		if testIdx != len(assertTable) {
			t.Fatalf("The decomposition len is wrong for number %d: %d expected: %d", n, testIdx, len(assertTable))
		}
	}

}
