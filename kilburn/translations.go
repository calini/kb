package kilburn

import "strings"

// MapMachineToLab returns the lab name from a
func MapMachineToLab(machineName string) string {
	// Special case for Collab 1/2
	if  machineName == "e-c07kilf3909" || strings.HasPrefix(machineName,"e-c07kilf391") {
		return "Collab 2"
	} else if strings.HasPrefix(machineName,"e-c07kilf390") {
		return "Collab 1"
	}

	labsPrefixes := map[string]string{
		"e-c07kig23":    "G23",
		"e-c07kilf9":    "Tootil 0",
		"e-c07kilf16":   "Tootil 1",
		"e-c07kilf31":   "LF31",
		"e-c07kilf17":   "LF17",
		"e-c07ki18":     "Quiet",
		"e-c07ki225":    "MSc",
	}
	for prefix, lab := range labsPrefixes {
		if strings.HasPrefix(machineName, prefix) {
			return lab
		}
	}
	return "none"
}

// LabsRegex is a translation table from a computer hostname to a lab, in case it's ever needed
//var LabsRegex = map[string]string{
//	"e-c07kig23(0[1-9]|[1-7][0-9]|80)":      "G23",      // c07kig23{01..80}
//	"e-c07kilf9(0[1-9]|[1-2][0-9]|3[0-6])":  "Tootil 0", // c07kilf9{01..36}
//	"e-c07kilf16(0[1-9]|[1-4][0-9]|5[0-3])": "Tootil 1", // c07kilf16{01..53}
//	"e-c07kilf31(0[1-9]|[1-6][0-9]|7[0-5])": "LF31",     // c07kilf31{01..75}
//	"e-c07kilf17(0[1-9]|1[0-3])":            "LF17",     // c07kilf17{01..13}
//	"e-c07kilf39(0[1-8])":                   "Collab 1", // c07kilf39{01..08}
//	"e-c07kilf39(09|1[0-2])":                "Collab 2", // c07kilf39{09..12}
//	"e-c07ki18(0[1-9]|[1-5][0-9]|60)":       "Quiet",    // c07ki18{01..60}
//	"e-c07ki225a(0[1-9]|[1-2][0-9]|3[0-6])": "MSC",      // c07ki225a{01..36}
//	"e-c07ki225b(0[1-9]|[1-2][0-9]|3[0-7])": "MSC",      // c07ki225b{01..37}
//	"e-c07ki225c(0[1-9]|[1-3][0-9]|40|41)":  "MSC",      // c07ki225c{01..41}
//}