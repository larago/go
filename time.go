package main

import (
    "time"
    "fmt"
)

var week time.Duration

func main() {
    t := time.Now()
    fmt.Println(t)
    fmt.Printf("%02d.%02d.%04d\n", t.Day(), t.Month(), t.Year())
    // 21.12.2011
    t = time.Now().UTC()
    fmt.Println(t)
    fmt.Println(time.Now())
    // calculating times:
    week = 60 * 60 * 24 * 7 * 1e9 // must be nanosec
    week_from_now := t.Add(week)
    fmt.Println(week_from_now)
    // formating times:
    fmt.Println(t.Format(time.RFC822))
    fmt.Println(t.Format(time.ANSIC))
    fmt.Println(t.Format("02 Jan 2006 15:04"))
    s := t.Format("20060102")
    fmt.Println(t ,"=>", s)
}