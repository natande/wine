package wine

// Default favicon.ico, a bottle of wine
var faviconBytes = []byte{
	0x00, 0x00, 0x01, 0x00, 0x01, 0x00, 0x10, 0x10, 0x00, 0x00, 0x01, 0x00,
	0x20, 0x00, 0x68, 0x04, 0x00, 0x00, 0x16, 0x00, 0x00, 0x00, 0x28, 0x00,
	0x00, 0x00, 0x10, 0x00, 0x00, 0x00, 0x20, 0x00, 0x00, 0x00, 0x01, 0x00,
	0x20, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x04, 0x00, 0x00, 0x12, 0x0b,
	0x00, 0x00, 0x12, 0x0b, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x2e, 0x35, 0x75, 0x00, 0x2e, 0x35, 0x75, 0x00, 0x2e, 0x35,
	0x75, 0x00, 0x2e, 0x35, 0x75, 0x00, 0x2e, 0x35, 0x75, 0x00, 0x0a, 0x0a,
	0x0a, 0xff, 0x0a, 0x0a, 0x0a, 0xff, 0x0a, 0x0a, 0x0a, 0xff, 0x0a, 0x0a,
	0x0a, 0xff, 0x0a, 0x0a, 0x0a, 0xff, 0x0a, 0x0a, 0x0a, 0xff, 0x0a, 0x0a,
	0x0a, 0xff, 0x2e, 0x35, 0x75, 0x00, 0x2e, 0x35, 0x75, 0x00, 0x2e, 0x35,
	0x75, 0x00, 0x2e, 0x35, 0x75, 0x00, 0x2e, 0x35, 0x75, 0x00, 0x2e, 0x35,
	0x75, 0x00, 0x2e, 0x35, 0x75, 0x00, 0x2e, 0x35, 0x75, 0x00, 0x2e, 0x35,
	0x75, 0x00, 0x2e, 0x35, 0x75, 0x00, 0x2e, 0x35, 0x75, 0x00, 0x2e, 0x35,
	0x75, 0x00, 0x0a, 0x0a, 0x0a, 0xff, 0x2e, 0x35, 0x75, 0x00, 0x2e, 0x35,
	0x75, 0x00, 0x2e, 0x35, 0x75, 0x00, 0x2e, 0x35, 0x75, 0x00, 0x2e, 0x35,
	0x75, 0x00, 0x2e, 0x35, 0x75, 0x00, 0x2e, 0x35, 0x75, 0x00, 0x2e, 0x35,
	0x75, 0x00, 0x2e, 0x35, 0x75, 0x00, 0x2e, 0x35, 0x75, 0x00, 0x2e, 0x35,
	0x75, 0x00, 0x2e, 0x35, 0x75, 0x00, 0x2e, 0x35, 0x75, 0x00, 0x2e, 0x35,
	0x75, 0x00, 0x2e, 0x35, 0x75, 0x00, 0x0a, 0x0a, 0x0a, 0xff, 0x2e, 0x35,
	0x75, 0x00, 0x2e, 0x35, 0x75, 0x00, 0x2e, 0x35, 0x75, 0x00, 0x2e, 0x35,
	0x75, 0x00, 0x2e, 0x35, 0x75, 0x00, 0x2e, 0x35, 0x75, 0x00, 0x2e, 0x35,
	0x75, 0x00, 0x2e, 0x35, 0x75, 0x00, 0x2e, 0x35, 0x75, 0x00, 0x2e, 0x35,
	0x75, 0x00, 0x2e, 0x35, 0x75, 0x00, 0x2e, 0x35, 0x75, 0x00, 0x2e, 0x35,
	0x75, 0x00, 0x2e, 0x35, 0x75, 0x00, 0x2e, 0x35, 0x75, 0x00, 0x0a, 0x0a,
	0x0a, 0xff, 0x2e, 0x35, 0x75, 0x00, 0x2e, 0x35, 0x75, 0x00, 0x2e, 0x35,
	0x75, 0x00, 0x2e, 0x35, 0x75, 0x00, 0x2e, 0x35, 0x75, 0x00, 0x2e, 0x35,
	0x75, 0x00, 0x2e, 0x35, 0x75, 0x00, 0x2e, 0x35, 0x75, 0x00, 0x2e, 0x35,
	0x75, 0x00, 0x2e, 0x35, 0x75, 0x00, 0x2e, 0x35, 0x75, 0x00, 0x2e, 0x35,
	0x75, 0x00, 0x2e, 0x35, 0x75, 0x00, 0x2e, 0x35, 0x75, 0x00, 0x2e, 0x35,
	0x75, 0x00, 0x0a, 0x0a, 0x0a, 0xff, 0x2e, 0x35, 0x75, 0x00, 0x2e, 0x35,
	0x75, 0x00, 0x2e, 0x35, 0x75, 0x00, 0x2e, 0x35, 0x75, 0x00, 0x2e, 0x35,
	0x75, 0x00, 0x2e, 0x35, 0x75, 0x00, 0x2e, 0x35, 0x75, 0x00, 0x2e, 0x35,
	0x75, 0x00, 0x2e, 0x35, 0x75, 0x00, 0x2e, 0x35, 0x75, 0x00, 0x2e, 0x35,
	0x75, 0x00, 0x2e, 0x35, 0x75, 0x00, 0x2e, 0x35, 0x75, 0x00, 0x2e, 0x35,
	0x75, 0x00, 0x2e, 0x35, 0x75, 0x00, 0x0a, 0x0a, 0x0a, 0xff, 0x2e, 0x35,
	0x75, 0x00, 0x2e, 0x35, 0x75, 0x00, 0x2e, 0x35, 0x75, 0x00, 0x2e, 0x35,
	0x75, 0x00, 0x2e, 0x35, 0x75, 0x00, 0x2e, 0x35, 0x75, 0x00, 0x2e, 0x35,
	0x75, 0x00, 0x2e, 0x35, 0x75, 0x00, 0x2e, 0x35, 0x75, 0x00, 0x2e, 0x35,
	0x75, 0x00, 0x2e, 0x35, 0x75, 0x00, 0x2e, 0x35, 0x75, 0x00, 0x2e, 0x35,
	0x75, 0x00, 0x2e, 0x35, 0x75, 0x00, 0x0a, 0x0a, 0x0a, 0xff, 0x0a, 0x0a,
	0x0a, 0xff, 0x0a, 0x0a, 0x0a, 0xff, 0x2e, 0x35, 0x75, 0x00, 0x2e, 0x35,
	0x75, 0x00, 0x2e, 0x35, 0x75, 0x00, 0x2e, 0x35, 0x75, 0x00, 0x2e, 0x35,
	0x75, 0x00, 0x2e, 0x35, 0x75, 0x00, 0x2e, 0x35, 0x75, 0x00, 0x2e, 0x35,
	0x75, 0x00, 0x2e, 0x35, 0x75, 0x00, 0x2e, 0x35, 0x75, 0x00, 0x2e, 0x35,
	0x75, 0x00, 0x0a, 0x0a, 0x0a, 0xff, 0x0a, 0x0a, 0x0a, 0xff, 0x4f, 0x26,
	0xef, 0xff, 0x4f, 0x26, 0xef, 0xff, 0x4f, 0x26, 0xef, 0xff, 0x0a, 0x0a,
	0x0a, 0xff, 0x0a, 0x0a, 0x0a, 0xff, 0x2e, 0x35, 0x75, 0x00, 0x2e, 0x35,
	0x75, 0x00, 0x2e, 0x35, 0x75, 0x00, 0x2e, 0x35, 0x75, 0x00, 0x2e, 0x35,
	0x75, 0x00, 0x2e, 0x35, 0x75, 0x00, 0x2e, 0x35, 0x75, 0x00, 0x2e, 0x35,
	0x75, 0x00, 0x0a, 0x0a, 0x0a, 0xff, 0x4f, 0x26, 0xef, 0xff, 0x4f, 0x26,
	0xef, 0xff, 0x4f, 0x26, 0xef, 0xff, 0x4f, 0x26, 0xef, 0xff, 0x4f, 0x26,
	0xef, 0xff, 0x4f, 0x26, 0xef, 0xff, 0x4f, 0x26, 0xef, 0xff, 0x0a, 0x0a,
	0x0a, 0xff, 0x2e, 0x35, 0x75, 0x00, 0x2e, 0x35, 0x75, 0x00, 0x2e, 0x35,
	0x75, 0x00, 0x2e, 0x35, 0x75, 0x00, 0x2e, 0x35, 0x75, 0x00, 0x2e, 0x35,
	0x75, 0x00, 0x2e, 0x35, 0x75, 0x00, 0x0a, 0x0a, 0x0a, 0xff, 0x4f, 0x26,
	0xef, 0xff, 0x4f, 0x26, 0xef, 0xff, 0x4f, 0x26, 0xef, 0xff, 0x4f, 0x26,
	0xef, 0xff, 0x4f, 0x26, 0xef, 0xff, 0x4f, 0x26, 0xef, 0xff, 0x4f, 0x26,
	0xef, 0xff, 0x0a, 0x0a, 0x0a, 0xff, 0x2e, 0x35, 0x75, 0x00, 0x2e, 0x35,
	0x75, 0x00, 0x2e, 0x35, 0x75, 0x00, 0x2e, 0x35, 0x75, 0x00, 0x2e, 0x35,
	0x75, 0x00, 0x2e, 0x35, 0x75, 0x00, 0x2e, 0x35, 0x75, 0x00, 0x0a, 0x0a,
	0x0a, 0xff, 0x4f, 0x26, 0xef, 0xff, 0x4f, 0x26, 0xef, 0xff, 0x4f, 0x26,
	0xef, 0xff, 0x4f, 0x26, 0xef, 0xff, 0x4f, 0x26, 0xef, 0xff, 0x4f, 0x26,
	0xef, 0xff, 0x4f, 0x26, 0xef, 0xff, 0x0a, 0x0a, 0x0a, 0xff, 0x2e, 0x35,
	0x75, 0x00, 0x2e, 0x35, 0x75, 0x00, 0x2e, 0x35, 0x75, 0x00, 0x2e, 0x35,
	0x75, 0x00, 0x2e, 0x35, 0x75, 0x00, 0x2e, 0x35, 0x75, 0x00, 0x2e, 0x35,
	0x75, 0x00, 0x0a, 0x0a, 0x0a, 0xff, 0x2e, 0x35, 0x75, 0x00, 0x2e, 0x35,
	0x75, 0x00, 0x2e, 0x35, 0x75, 0x00, 0x2e, 0x35, 0x75, 0x00, 0x2e, 0x35,
	0x75, 0x00, 0x2e, 0x35, 0x75, 0x00, 0x2e, 0x35, 0x75, 0x00, 0x0a, 0x0a,
	0x0a, 0xff, 0x2e, 0x35, 0x75, 0x00, 0x2e, 0x35, 0x75, 0x00, 0x2e, 0x35,
	0x75, 0x00, 0x2e, 0x35, 0x75, 0x00, 0x2e, 0x35, 0x75, 0x00, 0x2e, 0x35,
	0x75, 0x00, 0x2e, 0x35, 0x75, 0x00, 0x0a, 0x0a, 0x0a, 0xff, 0x2e, 0x35,
	0x75, 0x00, 0x2e, 0x35, 0x75, 0x00, 0x2e, 0x35, 0x75, 0x00, 0x2e, 0x35,
	0x75, 0x00, 0x2e, 0x35, 0x75, 0x00, 0x2e, 0x35, 0x75, 0x00, 0x2e, 0x35,
	0x75, 0x00, 0x0a, 0x0a, 0x0a, 0xff, 0x2e, 0x35, 0x75, 0x00, 0x2e, 0x35,
	0x75, 0x00, 0x2e, 0x35, 0x75, 0x00, 0x2e, 0x35, 0x75, 0x00, 0x2e, 0x35,
	0x75, 0x00, 0x2e, 0x35, 0x75, 0x00, 0x2e, 0x35, 0x75, 0x00, 0x06, 0x06,
	0x06, 0xff, 0x2e, 0x35, 0x75, 0x00, 0x2e, 0x35, 0x75, 0x00, 0x2e, 0x35,
	0x75, 0x00, 0x2e, 0x35, 0x75, 0x00, 0x2e, 0x35, 0x75, 0x00, 0x2e, 0x35,
	0x75, 0x00, 0x2e, 0x35, 0x75, 0x00, 0x06, 0x06, 0x06, 0xff, 0x2e, 0x35,
	0x75, 0x00, 0x2e, 0x35, 0x75, 0x00, 0x2e, 0x35, 0x75, 0x00, 0x2e, 0x35,
	0x75, 0x00, 0x2e, 0x35, 0x75, 0x00, 0x2e, 0x35, 0x75, 0x00, 0x2e, 0x35,
	0x75, 0x00, 0x2e, 0x35, 0x75, 0x00, 0x2e, 0x35, 0x75, 0x00, 0x2e, 0x35,
	0x75, 0x00, 0x2e, 0x35, 0x75, 0x00, 0x2e, 0x35, 0x75, 0x00, 0x2e, 0x35,
	0x75, 0x00, 0x2e, 0x35, 0x75, 0x00, 0x2e, 0x35, 0x75, 0x00, 0x2e, 0x35,
	0x75, 0x00, 0x2e, 0x35, 0x75, 0x00, 0x2e, 0x35, 0x75, 0x00, 0x2e, 0x35,
	0x75, 0x00, 0x2e, 0x35, 0x75, 0x00, 0x2e, 0x35, 0x75, 0x00, 0x2e, 0x35,
	0x75, 0x00, 0x2e, 0x35, 0x75, 0x00, 0x2e, 0x35, 0x75, 0x00, 0x2e, 0x35,
	0x75, 0x00, 0x2e, 0x35, 0x75, 0x00, 0x2e, 0x35, 0x75, 0x00, 0x2e, 0x35,
	0x75, 0x00, 0x2e, 0x35, 0x75, 0x00, 0x2e, 0x35, 0x75, 0x00, 0x2e, 0x35,
	0x75, 0x00, 0x2e, 0x35, 0x75, 0x00, 0x2e, 0x35, 0x75, 0x00, 0x2e, 0x35,
	0x75, 0x00, 0x2e, 0x35, 0x75, 0x00, 0xf8, 0x0f, 0x00, 0x00, 0xff, 0x7f,
	0x00, 0x00, 0xff, 0x7f, 0x00, 0x00, 0xff, 0x7f, 0x00, 0x00, 0xff, 0x7f,
	0x00, 0x00, 0xff, 0x7f, 0x00, 0x00, 0xfe, 0x3f, 0x00, 0x00, 0xf8, 0x0f,
	0x00, 0x00, 0xf0, 0x07, 0x00, 0x00, 0xf0, 0x07, 0x00, 0x00, 0xf0, 0x07,
	0x00, 0x00, 0xf7, 0xf7, 0x00, 0x00, 0xf7, 0xf7, 0x00, 0x00, 0xf7, 0xf7,
	0x00, 0x00, 0xff, 0xff, 0x00, 0x00, 0xff, 0xff, 0x00, 0x00,
}
