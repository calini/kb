package kilburn

import "testing"

func TestMapMachineToLab(t *testing.T) {
	lab := MapMachineToLab("e-c07ki225a01")
	if lab != "MSc" {
		t.Errorf("That computer is in MSc; Reported: %v", lab)
	}

	lab = MapMachineToLab("e-c07ki225b01")
	if lab != "MSc" {
		t.Errorf("That computer is in MSc; Reported %v", lab)
	}

	lab = MapMachineToLab("e-c07kig2301")
	if lab != "G23" {
		t.Errorf("That computer is in G23; Reported %v", lab)
	}

	lab = MapMachineToLab("e-c07kilf3901")
	if lab != "Collab 1" {
		t.Errorf("That computer is in Collab 1; Reported %v", lab)
	}

	lab = MapMachineToLab("e-c07kilf3908")
	if lab != "Collab 1" {
		t.Errorf("That computer is in Collab 1; Reported %v", lab)
	}

	lab = MapMachineToLab("e-c07kilf3909")
	if lab != "Collab 2" {
		t.Errorf("That computer is in Collab 2; Reported: %v", lab)
	}

	lab = MapMachineToLab("e-c07kilf3912")
	if lab != "Collab 2" {
		t.Errorf("That computer is in Collab 2; Reported %v", lab)
	}

	lab = MapMachineToLab("e-c07kilf3101")
	if lab != "LF31" {
		t.Errorf("That computer is in LF31; Reported %v", lab)
	}

	lab = MapMachineToLab("retina")
	if lab != "none" {
		t.Errorf("That computer is not in Kilburn; Reported %v", lab)
	}
}
