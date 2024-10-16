package goscii

// Converter options
type GOSCIIConverterOptions struct {
	// Charset to use
	Charset Charset
	// Output columns
	Columns int
	// Output rows
	Rows int
	// Output resize algorithm
	Fit Fit
	// Font aspect ratio
	FontAspectRatio float64
}

var defaultOptions = GOSCIIConverterOptions{
	Charset:         CharsetDefault,
	Columns:         80,
	Rows:            24,
	Fit:             FitContain,
	FontAspectRatio: FontAspectRatio,
}

func (dest *GOSCIIConverterOptions) Merge(source GOSCIIConverterOptions) {
	if len(source.Charset) > 0 {
		dest.Charset = source.Charset
	}

	if source.Columns > 0 {
		dest.Columns = source.Columns
	}

	if source.Rows > 0 {
		dest.Rows = source.Rows
	}

	if source.FontAspectRatio > 0 {
		dest.FontAspectRatio = source.FontAspectRatio
	}

	if source.Fit > 0 {
		dest.Fit = source.Fit
	}
}
