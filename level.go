package logging

import "strconv"

// Level for logging
// Levels are defaulted to that of syslog severity level https://en.wikipedia.org/wiki/Syslog
//
// Interpretation of the Levels is up to the appliation that is using them, the comments provided are only recommendations
type Level int

const (
	// LevelEmergency - A panic condition.
	LevelEmergency Level = iota

	// LevelAlert - A condition that should be corrected immediately, such as a corrupted system database.
	LevelAlert

	// LevelCritical - Hard application/device/hardware failures
	LevelCritical

	// LevelError - non-recoverable application errors
	LevelError

	// LevelWarning - recoverable application errors
	LevelWarning

	// LevelNotice - Conditions that are not error conditions, but may require special handling
	LevelNotice

	// LevelInformational - Messages that contain information about the application operation
	LevelInformational

	// LevelDebug - Messages that contain information normally of use only when debugging a program
	LevelDebug

	// LevelTrace - Messages that are likely only using while developing applications
	//
	// This is not part of the syslog spec, and might not be supported by all logging.Writer implementations
	LevelTrace

	// LevelNone - Not inteded to have output
	LevelNone Level = 255
)

func (l Level) String() string {
	switch l {
	case LevelEmergency:
		return "Emergency"
	case LevelAlert:
		return "Alert"
	case LevelCritical:
		return "Critical"
	case LevelError:
		return "Error"
	case LevelWarning:
		return "Warning"
	case LevelNotice:
		return "Notice"
	case LevelInformational:
		return "Informational"
	case LevelDebug:
		return "Debug"
	case LevelTrace:
		return "Trace"
	default:
		return "None"
	}
}

// Short returns the short representation of the level
func (l Level) Short() string {
	switch l {
	case LevelEmergency:
		return "emerg"
	case LevelAlert:
		return "alert"
	case LevelCritical:
		return "crit"
	case LevelError:
		return "err"
	case LevelWarning:
		return "warning"
	case LevelNotice:
		return "notice"
	case LevelInformational:
		return "info"
	case LevelDebug:
		return "debug"
	case LevelTrace:
		return "trace"
	default:
		return "none"
	}
}

// Prefix returns the syslog/journald prefix format of <#>
func (l Level) Prefix() string {
	return "<" + strconv.Itoa(int(l)) + ">"
}
