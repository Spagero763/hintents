// Copyright 2025 Erst Users
// SPDX-License-Identifier: Apache-2.0

package visualizer

// ANSI SGR escape codes used for terminal colour output.
// These mirror the constants in internal/terminal but are needed here
// because the visualizer package constructs colour strings directly.
const (
	sgrReset   = "\033[0m"
	sgrRed     = "\033[31m"
	sgrGreen   = "\033[32m"
	sgrYellow  = "\033[33m"
	sgrBlue    = "\033[34m"
	sgrMagenta = "\033[35m"
	sgrCyan    = "\033[36m"
	sgrDim     = "\033[2m"
	sgrBold    = "\033[1m"
)
