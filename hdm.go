package nmea

const (
	// TypeHDM type for HDM sentences
	TypeHDM = "HDM"
)

// HDM is for the Magnetic Heading Message
type HDM struct {
	BaseSentence
	Heading float64 // Heading in degrees
	Magnetic    bool    // Heading is magnetic
}

// newHDM constructor
func newHDM(s BaseSentence) (HDM, error) {
	p := NewParser(s)
	p.AssertType(TypeHDM)
	m := HDM{
		BaseSentence: s,
		Heading:      p.Float64(0, "heading"),
		Magnetic:         p.EnumString(1, "true", "M") == "M",
	}
	return m, p.Err()
}
