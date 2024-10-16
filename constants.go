package goscii

var (
	CharsetBlock   = Charset("█")
	CharsetDefault = Charset(" .,:;i1tfLCG08@")
)

const (
	FitWidth Fit = iota + 1
	FitHeight
	FitCover
	FitContain
	FitFill
)
