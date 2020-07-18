package logging

import (
	"fmt"
	"strconv"
	"strings"
)

// Level for logging
// Levels are defaulted to that of syslog severity level https://en.wikipedia.org/wiki/Syslog
//
// Interpretation of the Levels is up to the appliation that is using them, the comments provided are only recommendations
type Level uint8

var (
	levelStringMap = make(map[string]Level)
)

func init() {
	// build out a map of all our string based names to actual levels, easier to just iterate
	for i := uint8(0); i < uint8(LevelNone); i++ {
		if Level(i).String() == "" {
			continue
		}
		levelStringMap[strings.ToLower(Level(i).String())] = Level(i)
		levelStringMap[strings.ToLower(Level(i).Short())] = Level(i)
	}
}

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

	// Level 9-254 not handled

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

// Is will return if the provided Level is verbose enough from the current level. This is suitable for checking before outputing logs
//
// i.e. entry.Level.Is(LevelDebug) => write it when Debug-Emergency Level
func (l Level) Is(lvl Level) bool {
	if lvl == LevelNone {
		return false
	}

	return l <= lvl
}

// MarshalSetting provides the custom string output for the log level in config
func (l Level) MarshalSetting() string {
	return l.String()
}

// UnmarshalSetting provides the custom parsing of log levels in config
func (l *Level) UnmarshalSetting(v string) error {
	if v == "" {
		return nil
	}

	lm, found := levelStringMap[strings.ToLower(v)]
	if !found {
		return fmt.Errorf("value %q is not a valid log level", v)
	}

	*l = lm

	return nil
}

// Equals returns if the current value matches the input
func (l Level) Equals(v string) bool {
	lm, found := levelStringMap[strings.ToLower(v)]
	if !found {
		return false
	}

	return strings.EqualFold(l.String(), lm.String())
}
