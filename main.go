package main

import (
	"fmt"
	"image/png"
	"log"
	"net/http"
	"time"

	"github.com/cr0sh/ttfimg"
)

const width, height = 50, 50
const fsize, dpi = 15, 72

func main() {
	d, err := ttfimg.NewDrawer("/usr/share/fonts/malgun.ttf", width, height, fsize, dpi)
	if err != nil {
		log.Fatal(err)
	}
	http.HandleFunc("/now", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/png")
		t := time.Now()
		var buf string
		h, m, s := t.Hour(), t.Minute(), t.Second()
		buf += fmt.Sprintf("%02d:%02d\n %02d초", (h-1)%12+1, m, s)
		rgba := d.Draw(buf)
		if err := png.Encode(w, rgba); err != nil {
			log.Printf("PNG encode error: %#v", err)
		}
	})
	http.ListenAndServe(":8333", nil)
}
