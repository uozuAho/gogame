package gui

import (
	"bytes"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
)

const SampleRate = 44100

type AudioPlayer struct {
	audioCtx *audio.Context
	shootWav []byte
	hitWav   []byte
}

func (p *AudioPlayer) Init() error {
	p.audioCtx = audio.NewContext(SampleRate)

	data, err := os.ReadFile("assets/audio/white-short.wav")
	if err != nil {
		return err
	}
	p.shootWav = data

	data, err = os.ReadFile("assets/audio/hit.wav")
	if err != nil {
		return err
	}
	p.hitWav = data

	return nil
}

func (p *AudioPlayer) PlayShootSound() {
	if p.audioCtx == nil || len(p.shootWav) == 0 {
		panic("no")
	}

	// play sound in a goroutine so we don't block
	go func() {
		r := bytes.NewReader(p.shootWav)
		s, err := wav.DecodeWithSampleRate(SampleRate, r)
		if err != nil {
			log.Printf("failed to decode shoot wav: %v", err)
			return
		}
		player, err := p.audioCtx.NewPlayer(s)
		if err != nil {
			log.Printf("failed to create audio player: %v", err)
			return
		}
		player.Play()
	}()
}

func (p *AudioPlayer) PlayHitSound() {
	if p.audioCtx == nil || len(p.hitWav) == 0 {
		panic("no")
	}

	go func() {
		r := bytes.NewReader(p.hitWav)
		s, err := wav.DecodeWithSampleRate(SampleRate, r)
		if err != nil {
			log.Printf("failed to decode hit wav: %v", err)
			return
		}
		player, err := p.audioCtx.NewPlayer(s)
		if err != nil {
			log.Printf("failed to create audio player: %v", err)
			return
		}
		player.Play()
	}()
}
