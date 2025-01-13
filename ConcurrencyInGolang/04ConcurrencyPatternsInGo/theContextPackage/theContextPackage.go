package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func printGreeting(ctx context.Context) error {
	greeting, err := genGreeting(ctx)
	if err != nil {
		return err
	}
	fmt.Printf("%s world!\n", greeting)
	return nil
}

func printFarewell(ctx context.Context) error {
	farewell, err := genFarewell(ctx)
	if err != nil {
		return err
	}
	fmt.Printf("%s world!\n", farewell)
	return nil
}

func genGreeting(ctx context.Context) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	switch locale, err := locale(ctx); {
	case err != nil:
		return "", err
	case locale == "EN/US":
		return "hello", nil
	}
	return "", fmt.Errorf("unsupported locale")
}

func genFarewell(ctx context.Context) (string, error) {

	switch locale, err := locale(ctx); {
	case err != nil:
		return "", err
	case locale == "EN/US":
		return "goodbye", nil
	}
	return "", fmt.Errorf("unsupported locale")
}

func locale(ctx context.Context) (string, error) {
	if deadline, ok := ctx.Deadline(); ok {
		if deadline.Sub(time.Now().Add(1*time.Minute)) <= 0 {
			return "", context.DeadlineExceeded // DeadlineExceeded บ่งบอกว่า Context หยุดทำงานเพราะถึงเวลาที่กำหนด
		}
	}

	select {
	case <-ctx.Done(): // หยุดทำงานเมื่อ done ถูกปิด
		return "", ctx.Err()
	case <-time.After(1 * time.Minute): // สร้าง channel ที่จะส่งค่าเมื่อเวลาที่กำหนดผ่านไป
	}
	return "EN/US", nil
}

type ctxKey int

const (
	ctxUserId ctxKey = iota
	ctxAuthToken
)

func UserID(ctx context.Context) string {
	return ctx.Value(ctxUserId).(string)
}

func AuthToken(ctx context.Context) string {
	return ctx.Value(ctxAuthToken).(string)
}

func ProcessRequest(userId, authToken string) {
	ctx := context.WithValue(context.Background(), ctxUserId, userId)
	ctx = context.WithValue(ctx, ctxAuthToken, authToken)
	HandleResponse(ctx)
}

func HandleResponse(ctx context.Context) {
	fmt.Printf("handling response for %v (%v)",
		UserID(ctx),
		AuthToken(ctx))
}

func main() {
	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	wg.Add(1)
	go func() {
		defer wg.Done()

		if err := printGreeting(ctx); err != nil {
			fmt.Printf("cannot print greeting %v\n", err)
			cancel()
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := printFarewell(ctx); err != nil {
			fmt.Printf("cannot print farewell: %v\n", err)
		}
	}()

	wg.Wait()

	fmt.Println("----- ProcessRequest -----")
	ProcessRequest("jane", "abc123")
}
