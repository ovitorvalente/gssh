package clipboard

import "github.com/atotto/clipboard"

func Copy(text string) bool {
	return clipboard.WriteAll(text) == nil
}
