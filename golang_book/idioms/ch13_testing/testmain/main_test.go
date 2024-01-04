package testmain

import (
	"fmt"
	"os"
	"testing"
	"time"
)

var testTime time.Time

func TestMain(m *testing.M) {
	fmt.Println("set up stuff for tests")
	testTime = time.Now()
	exitVal := m.Run()
	fmt.Println("clean up stuff after tests")
	os.Exit(exitVal)
}

func TestFirst(t *testing.T) {
	fmt.Println("TestFirst uses stuff set up in TestMain", testTime)
}

func TestSecond(t *testing.T) {
	fmt.Println("TestSecond uses stuff set up in TestMain", testTime)
}

// вспомогательная функция, вызываемая из нескольких тестов
func createFile(t *testing.T) (string, error) {
	f, err := os.Create("tmpFile")
	if err != nil {
		return "", err
	}
	t.Cleanup(func() {
		os.Remove(f.Name())
	})
	return f.Name(), nil
}

func TestFileProcess(t *testing.T) {
	_, err := createFile(t)
	if err != nil {
		t.Fatal(err)
	}
	// no worries about resources
}

