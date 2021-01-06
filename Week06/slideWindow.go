package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type SlideWindow struct {
	buckets  map[int64]*bucket
	interval int64
	mu       *sync.RWMutex
}

type bucket struct {
	Value float64
}

func NewSlideWindow(interval int64) *SlideWindow {
	c := &SlideWindow{
		buckets:  make(map[int64]*bucket),
		interval: interval,
		mu:       &sync.RWMutex{},
	}
	return c
}

func (c *SlideWindow) currentBucket() *bucket {
	now := time.Now().Unix()

	if b, ok := c.buckets[now]; ok {
		return b
	}

	b := &bucket{}
	c.buckets[now] = b
	return b
}

func (c *SlideWindow) removeOldBuckets() {
	t := time.Now().Unix() - c.interval
	for timestamp := range c.buckets {
		if timestamp <= t {
			delete(c.buckets, timestamp)
		}
	}
}

func (c *SlideWindow) Incr(i float64) {
	if i == 0 {
		return
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	b := c.currentBucket()
	b.Value += i
	c.removeOldBuckets()
}

func (c *SlideWindow) Sum() float64 {
	t := time.Now().Unix() - c.interval

	sum := float64(0)

	c.mu.RLock()
	defer c.mu.RUnlock()

	for timestamp, bucket := range c.buckets {
		if timestamp >= t {
			sum += bucket.Value
		}
	}

	return sum
}

func (c *SlideWindow) Max() float64 {
	t := time.Now().Unix() - c.interval

	var max float64

	c.mu.RLock()
	defer c.mu.RUnlock()

	for timestamp, bucket := range c.buckets {
		if timestamp >= t {
			if bucket.Value > max {
				max = bucket.Value
			}
		}
	}

	return max
}

func (c *SlideWindow) Min() float64 {
	t := time.Now().Unix() - c.interval

	var min float64

	c.mu.RLock()
	defer c.mu.RUnlock()

	for timestamp, bucket := range c.buckets {
		if timestamp >= t {
			if min == 0 {
				min = bucket.Value
				continue
			}
			if bucket.Value < min {
				min = bucket.Value
			}
		}
	}

	return min
}

func (c *SlideWindow) Avg() float64 {
	return c.Sum() / float64(c.interval)
}

func (c *SlideWindow) Stat(tickNum int) {
	tickChan := time.Tick(time.Duration(tickNum) * time.Second)
	for range tickChan {
		fmt.Printf("\n***** During %d second(s) ******\n\n", tickNum)

		// print map
		for timestamp, bucket := range c.buckets {
			fmt.Println("["+time.Unix(timestamp, 0).Format("2006-01-02 15:04:05")+"]", bucket.Value)
		}

		// print infos
		fmt.Println()
		fmt.Println("- Max:", c.Max())
		fmt.Println("- Min:", c.Min())
		fmt.Println("- Sum:", c.Sum())
		fmt.Println("- Avg:", c.Avg())
	}
}

// TestIncr
func (c *SlideWindow) TestIncr(times int) {
	randRange := 100
	for i := 0; i < times; i++ {
		n := rand.Intn(randRange)
		c.Incr(float64(n))
		time.Sleep(time.Duration(n) * 10 * time.Millisecond)
	}
}
