package sound

import (
	"math/rand"
	"testing"
	"time"
)

func TestMiniMp3(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	s := InitStreamer()
	go func() {
		for {
			p := s.Ctx.NewPlayer()
			p.Write(s.Sounds[rand.Intn(len(s.Sounds))])
			p.Close()
			time.Sleep(time.Second)
		}
	}()

	go func() {
		for {
			p := s.Ctx.NewPlayer()
			p.Write(s.Sounds[rand.Intn(len(s.Sounds))])
			p.Close()
			time.Sleep(time.Millisecond * 500)
		}
	}()

	select {}
}
