package main

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/sync/errgroup"
)

func doSlowStuff(name string, ctx context.Context, duration time.Duration) error {
	fmt.Println("Starting doSlowStuff --", name)
	valDur := int64(duration)
	shorDur := time.Duration(valDur / 10)
	for i := 0; i < 10; i++ {
		if ctx.Err() != nil {
			fmt.Println(ctx.Err(), "--", name)
			return ctx.Err()
		}
		fmt.Println("... sleeping", i, "--", name)
		time.Sleep(shorDur)
	}
	fmt.Println("Did stuff", "--", name)
	return nil
}

func doSlowStuffNoCtx(duration time.Duration) error {
	name := "NoCtx"
	fmt.Println("Starting doSlowStuff --", name)
	valDur := int64(duration)
	shorDur := time.Duration(valDur / 10)
	for i := 0; i < 10; i++ {
		fmt.Println("... sleeping", i, "--", name)
		time.Sleep(shorDur)
	}
	fmt.Println("Did stuff", "--", name)
	return nil
}

func main() {
	bgCtx := context.Background()

	func() {
		fmt.Println("\n*** Early cancellation of error group")
		timeout := time.Duration(100) * time.Millisecond

		toCtx, cancel := context.WithTimeout(bgCtx, timeout)
		defer cancel()
		eg, egCtx := errgroup.WithContext(toCtx)
		eg.Go(func() error {
			return doSlowStuff("bgCtx", bgCtx, time.Duration(200)*time.Millisecond)
		})
		eg.Go(func() error {
			return doSlowStuff("egCtx", egCtx, time.Duration(200)*time.Millisecond)
		})
		eg.Go(func() error {
			return doSlowStuff("toCtx", toCtx, time.Duration(200)*time.Millisecond)
		})
		eg.Go(func() error {
			return doSlowStuffNoCtx(time.Duration(200)*time.Millisecond)
		})
		fmt.Println("eg.Wait =", eg.Wait())
	}()
}
