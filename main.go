package main

import (
    "net"
    "fmt"
    "os"
)

var create = []byte{0xf7, 0x32, 0x10, 0x00, 0x01, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x04, 0x00, 0x00, 0x00}
var connect = []byte{0xf7, 0x2f, 0x10, 0x00, 0x50, 0x58, 0x33, 0x57, 0x14, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}

func main() {
    var remote, local string
    var version byte
    fmt.Println("欢迎使用Warcraft III紫金港连接器")
    fmt.Println("目前仅支持有线和有线游戏，或者无线和无线游戏")
    fmt.Println("请先打开局域网对战窗口")
    fmt.Print("请输入您的IP地址，一般是10.171或者10.189开头：")
    fmt.Scan(&local)
    fmt.Print("请输入您要连接的主机的IP地址，一般是10.171或者10.189开头：")
    fmt.Scan(&remote)
    fmt.Print("请输入您的war3版本号，请使用以下数字对应：1.20e => 20, 1.21 => 21, 1.22 => 22, 后续还在测试中:")
    fmt.Scan(&version)

    
    rAddr, err := net.ResolveUDPAddr("udp", remote + ":6112")
    if err != nil {
        fmt.Println("出现了错误：" + err.Error() + "，可以尝试联系InsZVA@872673542(QQ)")
        os.Exit(-1)
    }
    lAddr, err := net.ResolveUDPAddr("udp", local + ":6112")
    if err != nil {
        fmt.Println("出现了错误：" + err.Error() + "，可以尝试联系InsZVA@872673542(QQ)")
        os.Exit(-1)
    }
    conn, err := net.DialUDP("udp", lAddr, rAddr)
    if err != nil {
        fmt.Println("出现了错误：" + err.Error() + "，可以尝试联系InsZVA@872673542(QQ)")
        os.Exit(-1)
    }

    connect[8] = version

    n, err := conn.Write(connect)
    
    if n == 16 && err == nil {
        fmt.Println("您应该已经可以看到房间了，如果没有看到，请尝试重新运行！")
    } else {
        fmt.Println("出现了错误：" + err.Error() + "，可以尝试联系InsZVA@872673542(QQ)")
    }
}