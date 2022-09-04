package wrapper

import (
	"log"
	"os"
	"os/exec"
)

func (w *pluginWrapper) Run() Wrapper {
	if !w.IsPrepared() {
		log.Println("wrapper", "not prepared")
		return w
	}
	if w.IsProcessed() {
		log.Println("wrapper", "already processed")
		return w
	}

	log.Println("wrapper", "process", "start")
	w.proc = exec.Command(w.cmd, w.args...)
	w.proc.Stdout = os.Stdout
	w.proc.Stderr = os.Stderr
	w.procErr = w.proc.Run()
	log.Println("wrapper", "process", "finish")

	log.Println("wrapper", "callback", "start")
	w.plugin.Run()
	log.Println("wrapper", "callback", "finish")

	return w
}
