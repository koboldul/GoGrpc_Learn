package calcserver

import (
	"fmt"
	"io"
	"testing"

	cpb "../../calcpb"
	grpc "google.golang.org/grpc"
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

func TestComputeAverageHappy(t *testing.T) {
	testTable := map[float64][]int32{
		2:   []int32{2, 2},
		2.5: []int32{1, 2, 3, 4},
	}

	for assert, input := range testTable {
		//Arrange
		s := server{}
		idx := int32(0)
		fmt.Printf("Index address : %v \n", &idx)
		c := ComputeAverageServerMock{
			t:        t,
			expected: assert,
			numbers:  input,
			current:  &idx,
		}

		//Act
		err := s.ComputeAverage(c)

		if err != nil {
			t.Fatalf("Should not be errored")
		}
	}
}

func TestComputeMaxHappy(t *testing.T) {
	testTable := map[int32][]maxTestCasePair{
		0: {{2, 2}, {1, 2}, {9, 9}},
		1: {{0, 0}},
		2: {},
	}

	for _, input := range testTable {
		//Arrange
		s := server{}
		idx := int32(0)
		fmt.Printf("Index address : %v \n", &idx)
		c := ComputeMaxServerMock{
			t:       t,
			input:   input,
			current: &idx,
		}

		//Act
		err := s.ComputeMax(c)

		if err != nil {
			t.Fatalf("Should not be errored")
		}
	}
}

//ComputeAverage tests
type ComputeAverageServerMock struct {
	grpc.ServerStream

	t        *testing.T
	current  *int32
	numbers  []int32
	expected float64
}

func (m ComputeAverageServerMock) SendAndClose(rsp *cpb.ComputeAverageResponse) error {
	avg := rsp.GetAverage()

	if avg != m.expected {
		m.t.Fatalf("The average expected for testcase number %d should be %v but was %v", *m.current, m.expected, avg)
	}

	return nil
}

func (m ComputeAverageServerMock) Recv() (*cpb.ComputeAverageRequest, error) {
	fmt.Printf("Was called request, %v \n", *m.current)

	if *m.current == int32(len(m.numbers)) {
		return nil, io.EOF
	}

	rq := &cpb.ComputeAverageRequest{
		Number: m.numbers[*m.current],
	}
	*m.current++

	return rq, nil
}

type maxTestCasePair struct {
	in       int32
	expected int32
}

//Compute Max tests
type ComputeMaxServerMock struct {
	grpc.ServerStream

	t       *testing.T
	current *int32
	input   []maxTestCasePair
}

func (m ComputeMaxServerMock) Send(rsp *cpb.ComputeMaxResponse) error {
	max := rsp.GetMax()

	expected := m.input[*m.current].expected

	if max != expected {
		m.t.Fatalf("The average expected for testcase number %d should be %v but was %v", *m.current, expected, max)
	}

	*m.current++

	return nil
}

func (m ComputeMaxServerMock) Recv() (*cpb.ComputeMaxRequest, error) {
	fmt.Printf("Was called request, %v \n", *m.current)

	if *m.current == int32(len(m.input)) {
		return nil, io.EOF
	}

	rq := &cpb.ComputeMaxRequest{
		Number: m.input[*m.current].in,
	}

	return rq, nil
}
