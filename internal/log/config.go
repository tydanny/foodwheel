package log

// Config holds details necessary for logging.
type Config struct {
	// Format specifies the output log format.
	// Accepted values are: json, console.
	Format string

	// Level is the minimum log level that should appear on the output.
	Level string
}
