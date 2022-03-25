package creational

import (
	"fmt"
	"testing"

	. "github.com/aldnav/godesignpatterns/creational/pkg/singleton"
)

func TestSingleton(t *testing.T) {
	log := GetLoggerInstance()

	log.SetLogLevel(1)
	log.Log("Log level set to 1")

	log = GetLoggerInstance()
	log.SetLogLevel(2)
	log.Log("Log level set to 2")

	log = GetLoggerInstance()
	log.SetLogLevel(3)
	log.Log("Log level set to 3")

	// creating several go routines
	for i := 0; i < 10; i++ {
		go GetLoggerInstance()
	}
	fmt.Scanln()
	fmt.Println("Done")
}
