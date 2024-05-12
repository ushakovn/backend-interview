package main

import log "github.com/sirupsen/logrus"

func main() {
	defer func() {
		if r := recover(); r != nil {
			log.Errorf("panic recovered: %v", r)
		}
	}()

	var m map[string]struct{} // init nil map

	log.Info(m[""]) // no panic: zero value

	m[""] = struct{}{} // panic: call recover function
}
