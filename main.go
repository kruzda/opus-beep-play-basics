package main

import (
    "io"
    "os"
    "log"
    "bytes"
    "time"
    "github.com/faiface/beep"
    "github.com/faiface/beep/pcm"
    "github.com/faiface/beep/speaker"
    "gopkg.in/hraban/opus.v2"
    "encoding/binary"
)

func intarrtobytes(i []int16) ([]byte) {
    buf := new(bytes.Buffer)
    binary.Write(buf, binary.LittleEndian, i)
    return buf.Bytes()
}

func main() {
    const sampleRate = 48000
    const channels = 1

    reader, err := os.Open("test1.ogg")
    if err != nil {
        log.Fatal(err)
    }

    stream, err := opus.NewStream(reader)
    if err != nil {
        log.Fatal(err)
    }

    // this will hold the entire decoded PCM data
    var pcm_output []int16
    // this can be any length it seems, though it's best if it's close to the size of
    // what Opus stream is reading at a time - it seems it is 960 int16s)
    pcmbuf := make([]int16, 960)

    // the only precision value I found to be working is 2
    format := beep.Format{
		SampleRate:  beep.SampleRate(sampleRate),
		NumChannels: channels,
		Precision:   2,
	}

    // start reading the stream and save pieces in pcmbuf
    // each time append pcmbuf into pcm_output
    for {
		n, err := stream.Read(pcmbuf)
		if err == io.EOF {
			break
        } else if err != nil {
            log.Fatal(err)
        }
        pcm_output = append(pcm_output, pcmbuf[:n]...)
	}

    // creates an io.reader pointing at the PCM data (that is converted from []int16 to []byte)
    r := bytes.NewReader(intarrtobytes(pcm_output))
    // create a Beep streamer
    streamer := pcm.Decode(r, format)

    // stub function to help with quitting after playback finished
    done := make(chan bool)

    // create a Beep speaker and play the streamer
    sr := format.SampleRate
    speaker.Init(sr, sr.N(time.Second/10))
    speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		done <- true
	})))

    <-done
}
