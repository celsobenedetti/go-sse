package assert

import "log"

func runAssert(msg string) {
	log.Fatalln(msg)
}

func Assert(cond bool, msg string) {
	if !cond {
		runAssert(msg)
	}
}

func Nil(x any, msg string) {
	if x != nil {
		runAssert(msg)
	}
}
