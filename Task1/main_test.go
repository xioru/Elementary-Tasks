package main

import "testing"

func TestMakeLine(t *testing.T) {
	testTable := []struct {
		width    int
		expected string
	}{
		{
			width:    4,
			expected: "* * * *",
		},
		{
			width:    5,
			expected: "* * * * *",
		},
	}

	for _, testCase := range testTable {
		result := makeLine(testCase.width)

		t.Logf("Calling makeLine(%d), result %v\n", testCase.width, testCase.expected)

		if result != testCase.expected {
			t.Errorf("Incorrect expected. Expect expected %s but got %s", testCase.expected, result)
		}
	}
}

func TestDrawField(t *testing.T) {
	testTable := []struct {
		height   int
		line     string
		expected string
	}{
		{
			height:   4,
			line:     "* * * *",
			expected: "* * * * \n * * * *\n* * * * \n * * * *\n",
		},
	}

	for _, testCase := range testTable {
		result := drawField(testCase.height, testCase.line)

		t.Logf("Calling DrawField(%d, %s), result \n %v\n", testCase.height, testCase.line, testTable)

		if result != testCase.expected {
			t.Errorf("Incorrect expected. Expect expected \n%s\n but got\n%s", testCase.expected, result)
		}
	}
}
