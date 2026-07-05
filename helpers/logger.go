package helpers

import "log"

func PanicErr(err error) {
	if err != nil {
		panic(err)
	}
}

func LogFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func LogFatalCaption(caption string, err error) {
	if err != nil {
		log.Fatal(caption, err)
	}
}
