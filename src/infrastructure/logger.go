package infrastructure

import "log"

func Log(args ...interface{}) {
	log.Println(args...)
}
