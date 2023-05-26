package utils

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"testing"
	"time"
)

func TestExecCmdWithStdOut(t *testing.T) {
	pid := os.Getpid()
	process, err := os.FindProcess(pid)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(process)
	go func() {
		time.Sleep(time.Second)
		fmt.Println("killing...")
		ExecCmdWithStdOut("tskill", strconv.Itoa(pid))
	}()

	fmt.Println("enter sleep")
	time.Sleep(time.Second * 60)
}
