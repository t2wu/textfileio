package textfileio

import (
	"bufio"
	"log"
	"os"
	"strings"
)

// Open opens file
// Mode can be "w" for writing or "r" for reading. Currently only supports text files.
// func Open(filepath, mode string) TextFileHandler {
func Open(filepath string, mode ...string) TextFileHandler {
	var f *os.File
	var err error

	if len(mode) > 0 && strings.Contains(mode[0], "w") {
		f, err = os.Create(filepath)
	} else {
		f, err = os.Open(filepath)
	}

	if err != nil {
		log.Fatal(err)
	}

	s := bufio.NewScanner(f)
	yield := make(chan string)
	outchannel := make(chan string)

	handler := TextFileHandler{f: f, scanner: s, yield: yield, outchannel: outchannel,
		isClosed: false, hasStarted: false}

	go func() {
		// fmt.Println("here")
		for s.Scan() {
			// handler.hasStarted = true
			yield <- s.Text()
		}

		// handler.hasStarted = true
		err = s.Err()
		if err != nil {
			log.Fatal(err)
		}

		for {
			defer handler.recoverFromSendOnClosedChannel() // Need this on test, I have no idea why.

			yield <- ""
		}
	}()

	return handler
}

// TextFileHandler is a user friendly scanner inspired by Python which handles EOF gracefully
// And handles everything in string
type TextFileHandler struct {
	f          *os.File
	scanner    *bufio.Scanner
	yield      chan string
	outchannel chan string
	isClosed   bool
	hasStarted bool
	started    chan bool
}

func (handler *TextFileHandler) closeYieldChanIfNotClosed() {
	defer handler.recoverFromSendOnClosedChannel()
	close(handler.yield)
}

func (handler *TextFileHandler) closeOutChanIfNotClosed() {
	defer handler.recoverFromSendOnClosedChannel()
	close(handler.outchannel)
}

func (handler *TextFileHandler) recoverFromSendOnClosedChannel() {
	recover()
}

// Close is to close the file
func (handler *TextFileHandler) Close() {
	defer handler.closeYieldChanIfNotClosed() // yes defer again
	defer handler.closeOutChanIfNotClosed()

	handler.isClosed = true
	if err := handler.f.Close(); err != nil {
		log.Fatal(err)
	}
}

// Iter returns a channel which can be used in for loops
func (handler *TextFileHandler) Iter() <-chan string {
	go func() {
		for {
			str := <-handler.yield // it will be constantly reading
			if str == "" {
				handler.closeOutChanIfNotClosed()
				break
			} else {
				handler.outchannel <- str // but if no one is reading the outchannel it waits here
			}
		}
	}()

	return handler.outchannel
}

// Readline reads the next line.
// This reads from a channel, when finished, it is done
func (handler *TextFileHandler) Readline() string {
	return <-handler.yield
}
