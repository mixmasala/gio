// SPDX-License-Identifier: Unlicense OR MIT

//go:build !android
// +build !android

package app

// Start is a no-op on platforms other than android
func Start(title, text string) (stop func(), err error) {
	return func() {}, nil
}
