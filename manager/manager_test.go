package manager

import "testing"

type mockOutputStrategy struct{}

var processInput string

func (mos mockOutputStrategy) Process(values string) {
	processInput = values
}

var mockStrategy = mockOutputStrategy{}
var sut = NewManager(mockStrategy)

func TestOpenInBrowser(t *testing.T) {
	calledWithParameter := ""

	openBrowser = func(url string) error {
		calledWithParameter = url

		return nil
	}

	sut.OpenInBrowser("http://batikansenemoglu.com")

	if calledWithParameter != "http://batikansenemoglu.com" {
		t.Errorf("Expected `http://batikansenemoglu.com` but given `%s`.", calledWithParameter)
	}
}

func TestFinishProcess(t *testing.T) {
	ch := make(chan string)

	go func() {
		ch <- "dummy value"
	}()

	sut.FinishProcess(ch)

	if processInput != "dummy value" {
		t.Errorf("Expected `dummy value` but given `%s`.", processInput)
	}
}
