package main

import (
	"fmt"
	"math/rand"
	"time"
)

func logMessages(chEmails, chSms chan string) {
	for {
		select {
			case email, ok := <-chEmails:
				if !ok {
					return
				}
				logEmail(email)
			case sms, ok := <-chSms:
				if !ok {
					return
				}
				logSms(sms)
		}
	}
}

func logSms(sms string) {
	fmt.Println("SMS:", sms)
}

func logEmail(email string) {
	fmt.Println("Email:", email)
}

func test(sms []string, emails []string) {
	fmt.Println("Starting...")

	chSms, chEmails := sendToLogger(sms, emails)

	logMessages(chEmails, chSms)
	fmt.Println("===============================")
}
// 13.8
func saveBackups(snapshotTicker, saveAfter <-chan time.Time) {
	for {
		select {
			case <- snapshotTicker:
				takeSnapshot()
			case <- saveAfter:
				saveSnapshot()
				return
			default:
				waitForData()
				time.Sleep(time.Microsecond * 500)
		}

	}
}

// don't touch below this line

func takeSnapshot() {
	fmt.Println("Taking a backup snapshot...")
}

func saveSnapshot() {
	fmt.Println("All backups saved!")
}

func waitForData() {
	fmt.Println("Nothing to do, waiting...")
}

func test2() {
	snapshotTicker := time.Tick(800 * time.Millisecond)
	saveAfter := time.After(2800 * time.Millisecond)
	saveBackups(snapshotTicker, saveAfter)
	fmt.Println("===========================")
}


func main() {
	// 13.8
	test2()

	rand.Seed(0)
	test(
		[]string{
			"hi friend",
			"What's going on?",
			"Welcome to the business",
			"I'll pay you to be my friend",
		},
		[]string{
			"Will you make your appointment?",
			"Let's be friends",
			"What are you doing?",
			"I can't believe you've done this.",
		},
	)
	test(
		[]string{
			"this song slaps hard",
			"yooo hoooo",
			"i'm a big fan",
		},
		[]string{
			"What do you think of this song?",
			"I hate this band",
			"Can you believe this song?",
		},
	)
}

func sendToLogger(sms, emails []string) (chSms, chEmails chan string) {
	chSms = make(chan string)
	chEmails = make(chan string)
	go func() {
		for i := 0; i < len(sms) && i < len(emails); i++ {
			done := make(chan struct{})
			s := sms[i]
			e := emails[i]
			t1 := time.Millisecond * time.Duration(rand.Intn(1000))
			t2 := time.Millisecond * time.Duration(rand.Intn(1000))
			go func() {
				time.Sleep(t1)
				chSms <- s
				done <- struct{}{}
			}()
			go func() {
				time.Sleep(t2)
				chEmails <- e
				done <- struct{}{}
			}()
			<-done
			<-done
			time.Sleep(10 * time.Millisecond)
		}
		close(chSms)
		close(chEmails)
	}()
	return chSms, chEmails
}
