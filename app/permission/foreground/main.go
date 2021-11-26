// SPDX-License-Identifier: Unlicense OR MIT

/*
Package foreground implements permissions to run a foreground service.
See https://developer.android.com/guide/components/foreground-services.

The following entries will be added to AndroidManifest.xml:

    <uses-permission android:name="android.permission.FOREGROUND_SERVICE"/>

*/

package foreground

// Start notifies the system that the program will perform
// background work and that it shouldn't be killed. It returns a channel
// that should be closed when the background work is complete.

// Start is a no-op on Linux, Windows, macOS; Android will
// display a notification during background work; iOS isn't supported.
func Start(title, text string) error {
	return start(title, text)
}

// Stop is a no-op on Linux, Windows, macOS; Android will stop the foreground
// service when the nubmer of calls to Stop equals the number of calls to
// Start.
func Stop() error {
	return stop()
}
