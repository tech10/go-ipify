package ipify

import (
	"fmt"
	"runtime"
	"strings"
)

// The version of this library.
const VERSION = "2.0"

// The maximum amount of tries to attempt when making API calls.
const MAX_TRIES = 3

// This is the ipify service base URI.  This is where all API requests go.
// We will use all API requests available.
const API_URI_4 = "https://api.ipify.org"
const API_URI_64 = "https://api64.ipify.org"
const API_URI_6 = "https://api6.ipify.org"

// The user-agent string is provided so that I can (eventually) keep track of
// what libraries to support over time.  EG: Maybe the service is used
// primarily by Windows developers, and I should invest more time in improving
// those integrations.
var USER_AGENT = fmt.Sprintf("go-ipify/%s go/%s %s", VERSION, runtime.Version()[2:], strings.Title(runtime.GOOS))
