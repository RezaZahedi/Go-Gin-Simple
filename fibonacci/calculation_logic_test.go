package fibonacci

import (
	"github.com/stretchr/testify/assert"
	"runtime"
	"testing"
)

func TestNewFiboGeneratorSequential(t *testing.T) {
	fiboGenerator := NewFiboGenerator()
	defer fiboGenerator.Close()
	tests := []struct {
		name string
		args int
		want string
		wantErr error
	}{
		{"1", 1, "1", nil},
		{"2", 2, "1", nil},
		{"8", 8, "21", nil},
		{"70", 70, "190392490709135", nil},
		{"10", 10, "55", nil},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ans, err := fiboGenerator.GenerateNumber(test.args)
			assert.Equal(t, test.wantErr, err)
			assert.Equal(t, test.want, ans)
		})
	}
}

func TestNewFiboGeneratorParallel(t *testing.T) {
	t.Parallel()
	fiboGenerator := NewFiboGenerator()

	// to make sure that we release the resources
	runtime.SetFinalizer(fiboGenerator,
		func(f *FiboGenerator) {
			f.Close()
			t.Log("resource released!")
		})

	tests := []struct {
		name string
		args int
		want string
		wantErr error
	}{
		{"1", 1, "1", nil},
		{"2", 2, "1", nil},
		{"8", 8, "21", nil},
		{"70", 70, "190392490709135", nil},
		{"10", 10, "55", nil},
	}


	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			ans, err := fiboGenerator.GenerateNumber(test.args)
			assert.Equal(t, test.wantErr, err)
			assert.Equal(t, test.want, ans)
		})
	}
}

func TestNewFiboGeneratorPanicCase(t *testing.T) {
	fiboGenerator := NewFiboGenerator()
	defer fiboGenerator.Close()
	tests := []struct {
		name string
		args int
		want string
		wantErr error
	}{
		{"-10", -10, "Nothing", nil},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r == nil {
					t.Errorf("The code did not panic")
				}
			}()

			ans, err := fiboGenerator.GenerateNumber(test.args)
			assert.Equal(t, test.wantErr, err)
			assert.Equal(t, test.want, ans)
		})
	}
}
