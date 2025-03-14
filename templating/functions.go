package templating

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
	"text/template"

	"github.com/fatih/color"
)

func FuncMap() template.FuncMap {
	return template.FuncMap{
		"include":   Include,
		"dump":      DumpArgs,
		"blue":      Blue,
		"cyan":      Cyan,
		"green":     Green,
		"magenta":   Magenta,
		"purple":    Magenta,
		"red":       Red,
		"yellow":    Yellow,
		"white":     White,
		"hiblue":    HighBlue,
		"hicyan":    HighCyan,
		"higreen":   HighGreen,
		"himagenta": HighMagenta,
		"hipurple":  HighMagenta,
		"hired":     HighRed,
		"hiyellow":  HighYellow,
		"hiwhite":   HighWhite,
	}
}

// Include is the function that implements inclusion of subfiles with
// an optional padding; when used without padding it is roughly equivalent
// to "template"; padding provides a way to prepend a constant string to
// each line in the output. The usage is as follows:
// {{ include <template> [<pipeline>] [<padding>] }}
func Include(args ...any) (string, error) {
	var (
		file    string
		padding string
		dynamic map[string]any
	)
	if args == nil {
		return "", errors.New("include: at least the template path must be specified")
	}
	var pipelineFound bool
	for i, arg := range args {
		var ok bool

		if i == 0 {
			if file, ok = arg.(string); !ok {
				return "", errors.New("include: the first argument (template) must be of type string")
			}
		} else if i == 1 {
			if dynamic, ok = arg.(map[string]any); !ok {
				if padding, ok = arg.(string); !ok {
					return "", errors.New("include: the second argument must either the pipeline or the padding")
				}
			} else {
				pipelineFound = true
			}
		} else if i == 2 {
			if !pipelineFound {
				return "", errors.New("include: the pipeline has not been provided")
			}
			if padding, ok = arg.(string); !ok {
				return "", errors.New("include: the third argument (padding) must be of type string")
			}
		}
	}

	// load the template
	t, err := template.ParseFiles(file)
	if err != nil {
		return "", err
	}

	var buffer strings.Builder
	if err = t.Execute(&buffer, dynamic); err != nil {
		return "", err
	}

	text := buffer.String()

	// apply padding only if necessary
	if padding != "" {
		var output strings.Builder
		scanner := bufio.NewScanner(strings.NewReader(text))
		for scanner.Scan() {
			output.WriteString(padding)
			output.WriteString(scanner.Text())
			output.WriteString("\n")
		}
		if scanner.Err() != nil {
			return "", err
		}
		return output.String(), nil
	}

	return text, nil
}

func DumpArgs(args ...any) (string, error) {
	result := ""
	if args != nil {
		for i, arg := range args {
			result += fmt.Sprintf("%d => '%v' (%T)\n", i, arg, arg)
		}
		fmt.Fprintln(os.Stderr, result)
		return result, nil
	} else {
		return "<empty>", nil
	}
}

var (
	blue      = color.New(color.FgBlue).SprintFunc()
	cyan      = color.New(color.FgCyan).SprintFunc()
	green     = color.New(color.FgGreen).SprintFunc()
	magenta   = color.New(color.FgMagenta).SprintFunc()
	red       = color.New(color.FgRed).SprintFunc()
	yellow    = color.New(color.FgYellow).SprintFunc()
	white     = color.New(color.FgWhite).SprintFunc()
	grey      = color.RGB(10, 10, 10).SprintFunc()
	hiblue    = color.New(color.FgHiBlue).SprintFunc()
	hicyan    = color.New(color.FgHiCyan).SprintFunc()
	higreen   = color.New(color.FgHiGreen).SprintFunc()
	himagenta = color.New(color.FgHiMagenta).SprintFunc()
	hired     = color.New(color.FgHiRed).SprintFunc()
	hiyellow  = color.New(color.FgHiYellow).SprintFunc()
	hiwhite   = color.New(color.FgHiWhite).SprintFunc()
)

func HighBlue(v any) string {
	return hiblue(v)
}

func HighCyan(v any) string {
	return hicyan(v)
}

func HighGreen(v any) string {
	return higreen(v)
}

func HighMagenta(v any) string {
	return himagenta(v)
}

func HighRed(v any) string {
	return hired(v)
}

func HighYellow(v any) string {
	return hiyellow(v)
}

func HighWhite(v any) string {
	return hiwhite(v)
}

func Blue(v any) string {
	return blue(v)
}

func Cyan(v any) string {
	return cyan(v)
}

func Green(v any) string {
	return green(v)
}

func Magenta(v any) string {
	return magenta(v)
}

func Red(v any) string {
	return red(v)
}

func Yellow(v any) string {
	return yellow(v)
}

func White(v any) string {
	return white(v)
}

func Grey(v any) string {
	return grey(v)
}
