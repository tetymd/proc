package main

import (
    "log"
    "syscall"
)

func main() {
    pp := []int{0,0}
    // pipeを作成。
    // 配列の中に書き込み用と読み込み用のfdが代入される
    err := syscall.Pipe(pp)
    if err != nil {
        println(err)
    }

    log.Println("length:", len(pp))
    log.Println("read fd:", pp[0])
    log.Println("write fd:", pp[1])

    wbuf := make([]byte, 16)
    rbuf := make([]byte, 16)

    wbuf = []byte("hello!")
    _, _ = syscall.Write(pp[1], wbuf)
    n, _ := syscall.Read(pp[0], rbuf)
    log.Println(string(rbuf[:n]))
}
