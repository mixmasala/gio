// SPDX-License-Identifier: Unlicense OR MIT

//+build !android

package foreground

func start(title, text string) error {
	return nil
}

func stop() error {
	return nil
}
