package main

import (
	"fmt"
	"time"
)

func orChannels() {
	var or func(channels ...<-chan interface{}) <-chan interface{}
	or = func(channels ...<-chan interface{}) <-chan interface{} { // เรามีฟังก์ชัน or ที่รับฟิลด์ตัวแปรและส่งคืนฟิลด์เดียว
		switch len(channels) {
		case 0:
			return nil
		case 1:
			return channels[0]
		}

		orDone := make(chan interface{})
		go func() { // ส่วนหลักของฟังก์ชันและเป็นจุดที่เกิดการเรียกซ้ำ
			defer close(orDone)

			switch len(channels) {
			case 2: // การเรียกซ้ำแต่ละครั้งจะต้องมีอย่างน้อย 2 ฟิลด์ (index 0 , index 1)
				select {
				case <-channels[0]:
				case <-channels[1]:
				}
			default:
				select {
				case <-channels[0]:
				case <-channels[1]:
				case <-channels[2]:
				case <-or(append(channels[3:], orDone)...):
				}
			}
		}()
		return orDone
	}

	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	start := time.Now()
	<-or(
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(1*time.Second),
		sig(1*time.Hour),
		sig(1*time.Minute),
	)
	fmt.Printf("done after %v", time.Since(start))
}

func main() {
	orChannels()
}
