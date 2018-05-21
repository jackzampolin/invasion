package invasion

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

const (
	case1 = "city north=bar south=foo"
	case2 = `wKZFyuVMfdCDLIDI south=ITcQARQwwcRiEwAQ west=nkqHYeKJMzcNPiuw north=nkqHYeKJMzcNPiuw
  ITcQARQwwcRiEwAQ south=tSykJReNHmOKxKlk east=ITcQARQwwcRiEwAQ west=GERfnmgtNfwsoybs north=DNjAxNYKjaORKjva
  LMOoEgpJuDqvNupV north=wKZFyuVMfdCDLIDI
  oCvfzhCgTZUlbaFE north=MfJfuxBoUzaEvmys
  YaiRRsTalKofdgyE east=wKZFyuVMfdCDLIDI north=YaiRRsTalKofdgyE south=ITcQARQwwcRiEwAQ west=LMOoEgpJuDqvNupV
  DNjAxNYKjaORKjva north=tSykJReNHmOKxKlk east=GERfnmgtNfwsoybs
  MfJfuxBoUzaEvmys north=YaiRRsTalKofdgyE west=ITcQARQwwcRiEwAQ east=DNjAxNYKjaORKjva
  nkqHYeKJMzcNPiuw north=nkqHYeKJMzcNPiuw west=MfJfuxBoUzaEvmys south=DNjAxNYKjaORKjva east=MfJfuxBoUzaEvmys
  GERfnmgtNfwsoybs south=GERfnmgtNfwsoybs east=GERfnmgtNfwsoybs
  tSykJReNHmOKxKlk east=nkqHYeKJMzcNPiuw west=nkqHYeKJMzcNPiuw`
	case3 = "faopjp4owijtnionqaviao4howgow4ngai4whgoiah"
)

func writeTestMapFile(name, contents string) string {
	gopath, exists := os.LookupEnv("GOPATH")
	if !exists {
		fmt.Println("SET YOUR GOPATH TO RUN TESTS")
	}
	tmpFile := fmt.Sprintf("%s/src/github.com/jackzampolin/invasion/invasion/%s.map", gopath, name)
	if _, err := os.Stat("/path/to/whatever"); os.IsNotExist(err) {
		// path/to/whatever does not exist
		err := ioutil.WriteFile(tmpFile, []byte(contents), 0644)
		if err != nil {
			panic(err)
		}
	}
	return tmpFile
}

func TestNewMapFromFileCase1(t *testing.T) {
	tmpFile := writeTestMapFile("one", case1)
	m, err := NewMapFromFile(tmpFile)
	if err != nil {
		t.Fail()
	}
	if len(m.Cities) != 1 && len(m.Cities[0].Roads) != 2 {
		t.Fail()
	}
	os.Remove(tmpFile)
}

func TestNewMapFromFileCase2(t *testing.T) {
	tmpFile := writeTestMapFile("two", case2)
	m, err := NewMapFromFile(tmpFile)
	if err != nil {
		t.Fail()
	}
	if len(m.Cities) != 10 && len(m.Cities[5].Roads) != 2 {
		t.Fail()
	}
	os.Remove(tmpFile)
}

func TestNewMapFromFileCase3(t *testing.T) {
	tmpFile := writeTestMapFile("three", case3)
	m, err := NewMapFromFile(tmpFile)
	if err != nil {
		t.Fail()
	}
	if len(m.Cities) > 0 {
		t.Fail()
	}
	os.Remove(tmpFile)
}

func TestNewMapFromFileNoFile(t *testing.T) {
	_, err := NewMapFromFile("foobarbaz")
	if err == nil {
		t.Fail()
	}
}

func TestNewCities(t *testing.T) {
	m := NewMap()
	m.NewCities(10)
	if len(m.Cities) != 10 && len(m.Cities[3].Name) != 16 && len(m.Cities[3].Roads) > 0 {
		t.Fail()
	}
}
