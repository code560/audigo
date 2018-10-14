package app

import (
	"flag"
	"os"
	"runtime/trace"
	"sync"
	"time"

	"github.com/code560/audigo/player"
	"github.com/code560/audigo/util"
)

var log = util.GetLogger()

func SoundPlay() {
	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		return
	}

	// trace
	f, err := os.Create("trace.out")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	trace.Start(f)
	defer trace.Stop()
	// trace end

	w := sync.WaitGroup{}
	list := playFiles(args, &w)
	w.Wait()

	time.Sleep(time.Second * 3)
	stopPlayers(list)
}

func playFiles(files []string, w *sync.WaitGroup) []player.Proxy {
	plist := make([]player.Proxy, len(files))
	for i, arg := range files {
		p := player.NewProxy()
		plist[i] = p

		w.Add(1)
		go func(p player.Proxy, name string) {
			p.GetChannel() <- &player.Action{
				Act:  player.Play,
				Args: &player.PlayArgs{Src: name, Loop: false}}
			w.Done()
		}(p, arg)
	}
	return plist
}

func stopPlayers(plist []player.Proxy) {
	time.Sleep(time.Second * 2)
	for _, p := range plist {
		p.GetChannel() <- &player.Action{Act: player.Stop}
	}
}
