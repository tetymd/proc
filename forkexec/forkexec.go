package main

import (
    "log"
    "syscall"
    "time"
)

func main() {
    argv := "forkexec/script"
    args := []string{}
    attr := &syscall.ProcAttr{}

    pid, err := syscall.ForkExec(argv, args, attr)
    if err != nil{
        log.Println(err)
    }
    log.Println(pid)
    time.Sleep(1 * time.Second)
}
