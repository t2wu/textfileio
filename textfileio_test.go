package textfileio

import (
	"testing"
	// "github.com/t2wu/textfileio/textfileio"
	// "github.com/t2wu/textfileio"
)

func TestReadOneLine(t *testing.T) {
	f := Open("./testdata/testinput.txt")
	defer f.Close()

	line := f.Readline()
	if line != "This is the first line." {
		t.Errorf("Expect reading the first line, getting: %v", line)
	}

	line = f.Readline()
	if line != "This is the second line." {
		t.Errorf("Expect reading the second line, getting: %v", line)
	}

	line = f.Readline()
	if line != "This is the third line." {
		t.Errorf("Expect reading the third line, getting: %v", line)
	}
}

func TestReadIterator(t *testing.T) {
	f := Open("./testdata/testinput.txt")
	defer f.Close()

	i := 0
	for line := range f.Iter() {
		switch i {
		case 0:
			if line != "This is the first line." {
				t.Errorf("Expect reading the first line, getting: %v", line)
			}
		case 1:
			if line != "This is the second line." {
				t.Errorf("Expect reading the second line, getting: %v", line)
			}
		case 2:
			if line != "This is the third line." {
				t.Errorf("Expect reading the third line, getting: %v", line)
			}
		}
		i++
	}
}

// func TestTextFileHandler_Iter(t *testing.T) {
// 	type fields struct {
// 		f       *os.File
// 		scanner *bufio.Scanner
// 		yield   chan string
// 	}
// 	tests := []struct {
// 		name   string
// 		fields fields
// 		want   <-chan string
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			handler := &TextFileHandler{
// 				f:       tt.fields.f,
// 				scanner: tt.fields.scanner,
// 				yield:   tt.fields.yield,
// 			}
// 			if got := handler.Iter(); !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("TextFileHandler.Iter() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func TestTextFileHandler_Readline(t *testing.T) {
// 	type fields struct {
// 		f       *os.File
// 		scanner *bufio.Scanner
// 		yield   chan string
// 	}
// 	tests := []struct {
// 		name   string
// 		fields fields
// 		want   string
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			handler := &TextFileHandler{
// 				f:       tt.fields.f,
// 				scanner: tt.fields.scanner,
// 				yield:   tt.fields.yield,
// 			}
// 			if got := handler.Readline(); got != tt.want {
// 				t.Errorf("TextFileHandler.Readline() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
