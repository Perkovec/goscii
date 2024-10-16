package main

import (
	"errors"
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"strconv"

	"github.com/Perkovec/goscii"
	"github.com/Perkovec/tutils"
	"github.com/spf13/cobra"
)

type options struct {
	input   string
	fit     string
	columns string
	rows    string
}

var (
	ErrInput          = errors.New("input cannot be empty")
	ErrInputNotExists = errors.New("input file not exists")
	ErrInvalidFit     = errors.New("invalid fit")
	ErrInvalidColumns = errors.New("invalid columns")
	ErrInvalidRows    = errors.New("invalid rows")

	fitMap = map[string]goscii.Fit{
		"contain": goscii.FitContain,
		"cover":   goscii.FitCover,
		"height":  goscii.FitHeight,
		"width":   goscii.FitWidth,
		"fill":    goscii.FitFill,
	}
)

func main() {
	cmdOptions := &options{}

	rootCmd := &cobra.Command{
		Use: "goscii",
		Run: func(cmd *cobra.Command, args []string) {
			err := cmdOptions.Validate()
			if err != nil {
				panic(err)
			}

			err = cmdOptions.exec()

			if err != nil {
				panic(err)
			}
		},
	}

	rootCmd.PersistentFlags().StringVarP(&cmdOptions.input, "input", "i", "", "Input image file")
	rootCmd.PersistentFlags().StringVarP(&cmdOptions.fit, "fit", "f", "contain", "Resize algorithm")
	rootCmd.PersistentFlags().StringVarP(&cmdOptions.columns, "columns", "c", "0", "Output art columns (0 - terminal columns)")
	rootCmd.PersistentFlags().StringVarP(&cmdOptions.rows, "rows", "r", "0", "Output art rows (0 - terminal rows)")

	rootCmd.Execute()
}

func (o *options) Validate() error {
	if len(o.input) == 0 {
		return ErrInput
	}

	if _, ok := fitMap[o.fit]; !ok {
		return ErrInvalidFit
	}

	if o.getColumns() < 0 {
		return ErrInvalidColumns
	}

	if o.getRows() < 0 {
		return ErrInvalidRows
	}

	// Check if input file is exists
	if _, err := os.Stat(o.input); os.IsNotExist(err) {
		return ErrInputNotExists
	}

	return nil
}

func (o *options) getFit() goscii.Fit {
	return fitMap[o.fit]
}

func (o *options) getColumns() int {
	cols, err := strconv.Atoi(o.columns)
	if err != nil {
		return -1
	}
	return cols
}

func (o *options) getRows() int {
	rows, err := strconv.Atoi(o.rows)
	if err != nil {
		return -1
	}
	return rows
}

func getImageFromFilePath(filePath string) (image.Image, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	image, _, err := image.Decode(f)
	return image, err
}

func (o *options) exec() error {
	imageFile, err := getImageFromFilePath(o.input)
	if err != nil {
		return err
	}

	// Get terminal size
	size, err := tutils.GetSize()
	if err != nil {
		return err
	}

	cols := o.getColumns()
	rows := o.getRows()

	if cols == 0 {
		cols = size.Columns
	}

	if rows == 0 {
		rows = size.Rows
	}

	converter, err := goscii.NewConverter(goscii.GOSCIIConverterOptions{
		Columns: cols,
		Rows:    rows,
		Fit:     o.getFit(),
	})
	if err != nil {
		return err
	}

	artRows := converter.Convert(imageFile)
	for _, artRow := range artRows {
		fmt.Println(artRow)
	}

	return nil
}
