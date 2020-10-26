package pathbuilder

const (
	directorySeparator = "/"
	windowsSeparator   = "\\"
)

// Custom concatenates a path with a custom separator
func Custom(separator string, components []string) string {
	var path string

	// TODO: Build path

	return path
}

// URL concatenates a path with unix style path separators
func URL(components []string) string {
	return Custom(directorySeparator, components)
}

// System concatenates a path with directory separators
func System(components []string) string {
	return Custom(directorySeparator, components)
}

// Unix concatenates a path with unix style path separators
func Unix(components []string) string {
	return Custom(directorySeparator, components)
}

// Windows concatenates a path with windows style path separators
func Windows(components []string) string {
	return Custom(windowsSeparator, components)
}
