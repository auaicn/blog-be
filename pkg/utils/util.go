package utils

import "log"

// HandleError logs the error and terminates the program if it's critical
func HandleError(err error, critical bool) {
	if err != nil {
		log.Println("Error:", err)
		if critical {
			// In case of a critical error, terminate the program
			log.Fatal("Critical error encountered, terminating program.")
		}
	}
}
