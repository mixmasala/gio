// SPDX-License-Identifier: Unlicense OR MIT

//go:build !android

package foreground

func start(title, text string) (func(), error) {
	return nil, nil
}
