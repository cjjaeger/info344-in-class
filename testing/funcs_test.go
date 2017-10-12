package testing

import (
	"testing"
)

func TestReverse(t *testing.T) {
	cases := []struct {
        input          string
        expectedOutput string
    }{
        {"", ""},
        {"a", "a"},
        {"ab", "ba"},
        {"abc", "cba"},
        {"abcd", "dcba"},
		{"aibohphobia", "aibohphobia"},
		{"Hello, 世界", "界世 ,olleH"},		
    }

    for _, c := range cases {
        if output := Reverse(c.input); output != c.expectedOutput {
            t.Errorf("incorrect output for `%s`: expected `%s` but got `%s`", c.input, c.expectedOutput, output)
        }
    }
}

func TestGetGreeting(t *testing.T) {
	cases := []struct {
		input string
		expectedOutput string
	}{
		{
			input:"stuff",
			expectedOutput: "Hello, stuff!",
		},
		{
			input:"",
			expectedOutput: "Hello, World!",
		},
	}


	for _, c := range cases {
		output := GetGreeting(c.input)
		if output != c.expectedOutput{
			t.Errorf("Got %s but expected %s", output, c.expectedOutput)
		}
	}

}

func TestParseSize(t *testing.T) {
	cases := []struct {
		input string
		expectedOutput *Size
	}{
		{
			input:"60x60",
			expectedOutput: &Size{60, 60},
		},
		{
			input:"foox60",
			expectedOutput: &Size{0, 60},
		},
		{
			input:"",
			expectedOutput: &Size{},
		},
	}

	for _, c := range cases {
		output := ParseSize(c.input)
		if output.Height != c.expectedOutput.Height || output.Width != c.expectedOutput.Width {
			t.Errorf("Got %v but expected %v", output, c.expectedOutput)
		}
	}

}

func TestLateDaysConsume(t *testing.T) {
	 ld := NewLateDays()
	 for i := 4; i > -10 ; i--{
		 expectedLateDays := i;
		 if expectedLateDays < 0{
			expectedLateDays = 0
		 }
		 output := ld.Consume("test")
		 if output != expectedLateDays {
			 t.Errorf("iteration %d: got %d but expected %d", i, output, expectedLateDays)
		 }
	 }
}
