package rate

// SI unit multiples.
const (
	Kilo float64 = 1000
	Mega         = Kilo * 1000
	Giga         = Mega * 1000
	Tera         = Giga * 1000
	Peta         = Tera * 1000
	Exa          = Peta * 1000
)

// symbols Prefix
var symbols = []string{"b", "B", "Hz", "bps", "Bps"}

/*
SI_PREFIX for base10
*/
const (
	Kilobyte float64 = 1000
	KB               = Kilobyte
	K                = Kilobyte
	Megabyte         = Kilobyte * 1000
	MB               = Megabyte
	M                = Megabyte
	Gigabyte         = Megabyte * 1000
	GB               = Gigabyte
	G                = Gigabyte
	Terabyte         = Gigabyte * 1000
	TB               = Terabyte
	T                = Terabyte
	Petabyte         = Terabyte * 1000
	PB               = Petabyte
	P                = Petabyte
	Exabyte          = Petabyte * 1000
	EB               = Exabyte
	E                = Exabyte
)

// siPrefix MAP
var siPrefix = map[string]float64{
	"K": Kilobyte,
	"M": Megabyte,
	"G": Gigabyte,
	"T": Terabyte,
	"P": Petabyte,
	"E": Exabyte,
	"d": decibyte,
	"c": centibyte,
	"m": millibyte,
	"µ": microbyte,
	"n": nanobyte,
	"p": picobyte,
}

/*
BINARY_PREFIX for base2
*/
const (
	Kibibyte float64 = 1024
	KiB              = Kibibyte
	Ki               = Kibibyte
	Mebibyte         = Kibibyte * 1024
	MiB              = Mebibyte
	Mi               = Megabyte
	Gibibyte         = Mebibyte * 1024
	GiB              = Gibibyte
	Gi               = Gigabyte
	Tebibyte         = Gibibyte * 1024
	TiB              = Tebibyte
	Ti               = Tebibyte
	Pebibyte         = Tebibyte * 1024
	PiB              = Pebibyte
	Pi               = Pebibyte
	Exbibyte         = Pebibyte * 1024
	EiB              = Exbibyte
	Ei               = Exbibyte
)

// binaryPrefix MAP
var binaryPrefix = map[string]float64{
	"Ki": Kibibyte,
	"Mi": Mebibyte,
	"Gi": Gibibyte,
	"Ti": Tebibyte,
	"Pi": Pebibyte,
	"Ei": Exbibyte,
}

//Rational
var (
	decibyte  float64 = 1 / 10
	dB                = decibyte
	d                 = decibyte
	centibyte float64 = 1 / 100
	cB                = centibyte
	c                 = centibyte
	millibyte float64 = 1 / 1000
	mB                = millibyte
	m                 = millibyte
	microbyte         = millibyte * (1 / 1000)
	µB                = microbyte
	µ                 = microbyte
	nanobyte          = microbyte * (1 / 1000)
	nB                = nanobyte
	n                 = nanobyte
	picobyte          = nanobyte * (1 / 1000)
	pB                = picobyte
	p                 = picobyte
)

// allPrefix MAP
var allPrefix = map[string]float64{
	"K":  Kilobyte,
	"M":  Megabyte,
	"G":  Gigabyte,
	"T":  Terabyte,
	"P":  Petabyte,
	"E":  Exabyte,
	"d":  decibyte,
	"c":  centibyte,
	"m":  millibyte,
	"µ":  microbyte,
	"n":  nanobyte,
	"p":  picobyte,
	"Ki": Kibibyte,
	"Mi": Mebibyte,
	"Gi": Gibibyte,
	"Ti": Tebibyte,
	"Pi": Pebibyte,
	"Ei": Exbibyte,
}
