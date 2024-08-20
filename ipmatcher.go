package main

import (
    "os"
    "net"
    "bufio"
    "log"
)

func main() {
    ip_match := os.Args[1]
    filepath := os.Args[2]

    // read whitelist subnets from file

    file, err := os.Open(filepath)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    scanner.Split(bufio.ScanLines)

    var entries []string
    for scanner.Scan() {
        entries = append(entries, scanner.Text())
    }

    // ip matching

    ipaddr_match := net.ParseIP(ip_match)
    if ipaddr_match == nil {
        log.Fatalf("Argument 1 '%s' is not IP address", ip_match)
    }

    for i, entry := range entries {
        _, ipnet, err := net.ParseCIDR(entry)
        if err != nil {
            ipaddr := net.ParseIP(entry)
            if ipaddr == nil {
                log.Printf("Line %d '%s' of file '%s' is not valid", i + 1, entry, filepath)
            } else {
                if ip_match == entry {
                    os.Exit(0)
                }
            }
        } else {
            if ipnet.Contains(ipaddr_match) {
                os.Exit(0)
            }
        }
    }

    os.Exit(214)
}
