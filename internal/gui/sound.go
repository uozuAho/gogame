package gui

import (
	"bytes"
	"encoding/binary"
	"math"
)

// GenerateShootWavBytes produces a short (mono) PCM16 WAV containing a
// simple decaying sine burst suitable as a bullet-shot sound. No external
// asset required.
func GenerateShootWavBytes() []byte {
	const sampleRate = 44100
	const durationSec = 0.12
	const freq = 880.0
	samples := int(float64(sampleRate) * durationSec)
	numChannels := 1
	bitsPerSample := 16

	buf := &bytes.Buffer{}

	// RIFF header
	// ChunkID "RIFF"
	buf.WriteString("RIFF")
	// Placeholder for chunk size
	_ = binary.Write(buf, binary.LittleEndian, uint32(0))
	buf.WriteString("WAVE")

	// fmt chunk
	buf.WriteString("fmt ")
	_ = binary.Write(buf, binary.LittleEndian, uint32(16)) // Subchunk1Size for PCM
	_ = binary.Write(buf, binary.LittleEndian, uint16(1))  // AudioFormat PCM
	_ = binary.Write(buf, binary.LittleEndian, uint16(numChannels))
	_ = binary.Write(buf, binary.LittleEndian, uint32(sampleRate))
	byteRate := uint32(sampleRate * numChannels * bitsPerSample / 8)
	_ = binary.Write(buf, binary.LittleEndian, byteRate)
	blockAlign := uint16(numChannels * bitsPerSample / 8)
	_ = binary.Write(buf, binary.LittleEndian, blockAlign)
	_ = binary.Write(buf, binary.LittleEndian, uint16(bitsPerSample))

	// data chunk header
	buf.WriteString("data")
	dataSize := uint32(samples * numChannels * bitsPerSample / 8)
	_ = binary.Write(buf, binary.LittleEndian, dataSize)

	// write samples (PCM16 little endian)
	maxAmp := float64(1<<15 - 1)
	for i := 0; i < samples; i++ {
		t := float64(i) / float64(sampleRate)
		// simple decaying sine
		env := math.Exp(-6.0 * t) // fast decay
		v := math.Sin(2.0*math.Pi*freq*t) * env * 0.6
		s := int16(v * maxAmp)
		_ = binary.Write(buf, binary.LittleEndian, s)
	}

	// fill in RIFF chunk size (file size - 8)
	data := buf.Bytes()
	chunkSize := uint32(len(data) - 8)
	binary.LittleEndian.PutUint32(data[4:8], chunkSize)

	return data
}
