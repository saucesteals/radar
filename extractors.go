package radar

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	cookieEnabled = iota
	doNotTrack
	languages
	platform
	plugins
	screen
	timezoneOffset
	hasDocumentTouch
	enabledStorage
	fontFingerprint
	_ // unused
	userAgent
	shockwaveFlashVersion
	adblocker
	canvasFingerprint

	extractorCount
)

type Extractors [extractorCount]extractorResult

func ComposeExtractors() *Extractors {
	return &Extractors{}
}

type extractorResult struct {
	value             string
	calculationTimeMs int
}

func (e *Extractors) CookieEnabled(value bool, time int) *Extractors {
	e[cookieEnabled] = extractorResult{strconv.FormatBool(value), time}
	return e
}

func (e *Extractors) DoNotTrack(value bool, time int) *Extractors {
	e[doNotTrack] = extractorResult{strconv.FormatBool(value), 0}
	return e
}

func (e *Extractors) Languages(value []string, time int) *Extractors {
	e[languages] = extractorResult{strings.Join(value, ","), time}
	return e
}

func (e *Extractors) Platform(value string, time int) *Extractors {
	e[platform] = extractorResult{value, time}
	return e
}

type PluginMimeType struct {
	Type     string
	Suffixes string
}

type Plugin struct {
	Name      string
	Filename  string
	MimeTypes []PluginMimeType
}

func (e *Extractors) Plugins(value []Plugin, time int) *Extractors {
	var pluginsString []string
	for _, p := range value {
		var mimeTypes []string
		for _, mt := range p.MimeTypes {
			mimeTypes = append(mimeTypes, mt.Type+","+mt.Suffixes)
		}
		pluginsString = append(pluginsString, p.Name+","+p.Filename+","+strings.Join(mimeTypes, "++"))
	}
	e[plugins] = extractorResult{strings.Join(pluginsString, ", "), time}
	return e
}

type Screen struct {
	AvailWidth, AvailHeight, ColorDepth, DevicePixelRatio int
}

func (e *Extractors) Screen(value Screen, time int) *Extractors {
	e[screen] = extractorResult{
		fmt.Sprintf("%dw_%dh_%dd_%dr", value.AvailWidth, value.AvailHeight, value.ColorDepth, value.DevicePixelRatio),
		time,
	}
	return e
}

func (e *Extractors) TimezoneOffset(value int, time int) *Extractors {
	e[timezoneOffset] = extractorResult{strconv.FormatFloat(float64(-value)/float64(60), 'f', -1, 64), time}
	return e
}

func (e *Extractors) HasDocumentTouch(value bool, time int) *Extractors {
	e[hasDocumentTouch] = extractorResult{strconv.FormatBool(value), time}
	return e
}

type storageState string

var (
	StorageStateEnabled     storageState = "enabled"     // enabled
	StorageStateUnavailable storageState = "unavailable" // null
	SessionStateDisabled    storageState = "disabled"    // exception
)

func (e *Extractors) StorageState(sessionStorage, localStorage storageState, time int) *Extractors {
	e[enabledStorage] = extractorResult{
		"sessionStorage-" + string(sessionStorage) + ", localStorage-" + string(localStorage),
		time,
	}
	return e
}

func (e *Extractors) FontFingerprint(value string, time int) *Extractors {
	e[fontFingerprint] = extractorResult{value, time}
	return e
}

func (e *Extractors) UserAgent(value string, time int) *Extractors {
	e[userAgent] = extractorResult{value, time}
	return e
}

func (e *Extractors) ShockwaveFlashVersion(value string, time int) *Extractors {
	e[shockwaveFlashVersion] = extractorResult{value, time}
	return e
}

func (e *Extractors) Adblocker(value bool, time int) *Extractors {
	e[adblocker] = extractorResult{strconv.FormatBool(value), time}
	return e
}

func (e *Extractors) CanvasFingerprint(value string, time int) *Extractors {
	e[canvasFingerprint] = extractorResult{value, time}
	return e
}

func (e *Extractors) Values() string {
	values := make([]string, 0, extractorCount)
	for _, extractor := range e {
		values = append(values, extractor.value)
	}
	return strings.Join(values, " ")
}
