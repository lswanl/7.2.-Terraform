// Code generated by "stringer -type=DiffLanguage diff.go"; DO NOT EDIT.

package format

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[DiffLanguageProposedChange-80]
	_ = x[DiffLanguageDetectedDrift-68]
}

const (
	_DiffLanguage_name_0 = "DiffLanguageDetectedDrift"
	_DiffLanguage_name_1 = "DiffLanguageProposedChange"
)

func (i DiffLanguage) String() string {
	switch {
	case i == 68:
		return _DiffLanguage_name_0
	case i == 80:
		return _DiffLanguage_name_1
	default:
		return "DiffLanguage(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}