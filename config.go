package streamdeck

import (
	"github.com/bearsh/hid"
)

type Config struct {
	ProductID        uint16 // ProductID is the USB ProductID
	NumButtonColumns int
	NumButtonRows    int
	Spacer           int // Spacer is the spacing distance (in pixel) of two buttons on the Stream Deck.
	ButtonSize       int
	ImageFormat      string
	ImageRotate      bool
	TransposeImage   bool
	ConvertKey       bool
	SimpleKeyEvents  bool
	MiniProtocol     bool
}

func (c Config) NumButtons() int {
	return c.NumButtonRows * c.NumButtonColumns
}

// PanelWidth is the total screen width of the Stream Deck (including spacers).
func (c *Config) PanelWidth() int {
	return c.NumButtonColumns*c.ButtonSize + c.Spacer*(c.NumButtonColumns-1)
}

// PanelHeight is the total screen height of the stream deck (including spacers).
func (c *Config) PanelHeight() int {
	return c.NumButtonRows*c.ButtonSize + c.Spacer*(c.NumButtonRows-1)
}

func (c *Config) fixKey(key int) int {
	if c.ConvertKey {
		keyCol := key % c.NumButtonColumns
		return 1 + ((key - keyCol) + ((c.NumButtonColumns - 1) - keyCol))
	}
	return key
}

// Model 20GAA9901
var Original = Config{
	ProductID:        0x60,
	NumButtonColumns: 5,
	NumButtonRows:    3,
	Spacer:           19,
	ButtonSize:       72,
	ImageFormat:      "bmp",
	ConvertKey:       true,
}

// Model 20GAA9902
var OriginalMk1 = Config{
	ProductID:        0x6d,
	NumButtonColumns: 5,
	NumButtonRows:    3,
	Spacer:           19,
	ButtonSize:       72,
	ImageFormat:      "jpg",
	ImageRotate:      true,
}

var Original2 = Config{
	ProductID:        0x80,
	NumButtonColumns: 5,
	NumButtonRows:    3,
	Spacer:           19,
	ButtonSize:       72,
	ImageFormat:      "jpg",
	ImageRotate:      true,
}

var Plus = Config{
	ProductID:        0x0084,
	NumButtonColumns: 4,
	NumButtonRows:    2,
	Spacer:           19,
	ButtonSize:       120,
	ImageFormat:      "jpg",
}

var Mini = Config{
	ProductID:        0x0063,
	NumButtonColumns: 3,
	NumButtonRows:    2,
	Spacer:           26,
	ButtonSize:       80,
	ImageFormat:      "bmp",
	TransposeImage:   true,
	SimpleKeyEvents:  true,
	MiniProtocol:     true,
}

var XL = Config{
	ProductID:        0x008f,
	NumButtonColumns: 8,
	NumButtonRows:    4,
	Spacer:           38,
	ButtonSize:       96,
	ImageFormat:      "jpg",
	ImageRotate:      true,
}

var Neo = Config{
	ProductID:        0x009a,
	NumButtonColumns: 4,
	NumButtonRows:    2,
	Spacer:           38,
	ButtonSize:       96,
	ImageFormat:      "jpg",
	ImageRotate:      true,
}

var AllConfigs = []Config{Original, OriginalMk1, Original2, Plus, Mini, XL, Neo}

func FindConnectedConfig() (Config, bool) {
	for _, c := range AllConfigs {
		devices := hid.Enumerate(VendorID, c.ProductID)
		if len(devices) > 0 {
			return c, true
		}
	}
	return Config{}, false
}
