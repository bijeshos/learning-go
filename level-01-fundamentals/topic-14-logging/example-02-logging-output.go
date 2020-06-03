/*
A simple program to demonstrate logging (to bytes buffer) with different output levels
*/

package main

import (
	"bytes"
	"fmt"
	"log"
)

func main() {
	var (
		buf         bytes.Buffer
		infoLogger  = log.New(&buf, "INFO: ", log.Lshortfile)
		debugLogger = log.New(&buf, "DEBUG: ", log.Lshortfile)

		logInfo = func(message string) {
			infoLogger.Output(2, message)
		}

		logDebug = func(message string) {
			debugLogger.Output(2, message)
		}
	)

	logInfo("Hello world")
	logDebug("How are you doing")

	fmt.Print(&buf)
}
