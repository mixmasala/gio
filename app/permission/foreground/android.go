// SPDX-License-Identifier: Unlicense OR MIT

//go:build android

/*
Package foreground implements permissions to run a foreground service.
See https://developer.android.com/guide/components/foreground-services.

The following entries will be added to AndroidManifest.xml:

    <uses-permission android:name="android.permission.FOREGROUND_SERVICE"/>

*/

package foreground

import (
	"gioui.org/app"
)

func start(title, text string) (func(), error) {
	return app.Start(title, text)
}
