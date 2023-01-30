package main

import (
	"fmt"

	"github.com/saucesteals/radar"
)

func main() {
	extractors := radar.ComposeExtractors().
		CookieEnabled(true, 0).
		DoNotTrack(false, 0).
		Languages([]string{"en-US", "en"}, 0).
		Platform("MacIntel", 0).
		Plugins([]radar.Plugin{}, 0).
		Screen(radar.Screen{
			AvailWidth:       1920,
			AvailHeight:      1080,
			ColorDepth:       30,
			DevicePixelRatio: 2,
		}, 0).
		TimezoneOffset(480, 0).
		HasDocumentTouch(false, 0).
		StorageState(radar.StorageStateEnabled, radar.StorageStateEnabled, 0).
		FontFingerprint("0100111100000000100000111100000000000010001000010111111", 0).
		UserAgent("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36", 0).
		ShockwaveFlashVersion("", 0).
		Adblocker(false, 0).
		CanvasFingerprint("ebcba3d055b476bc2d6bca490ad777cb", 0)

	fmt.Println(extractors.Values())
}
