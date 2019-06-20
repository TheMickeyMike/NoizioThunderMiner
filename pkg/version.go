package core

// Version is the Noizio version.
type Version string

const (
	// V20 is string for version 2.0
	V20 Version = "2.0"
)

// IsSupported returns nil if the version is supported,
// or ErrUnsupportedVersion if not supported.
func (v Version) IsSupported() error {
	switch v {
	case V20:
		return nil
	}
	return ErrUnsupportedVersion{v}
}