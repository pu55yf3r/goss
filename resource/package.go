package resource

import "github.com/aelsabbahy/goss/system"

type Package struct {
	Name      string   `json:"name"`
	Installed bool     `json:"installed"`
	Versions  []string `json:"versions,omitempty"`
}

func (p *Package) Validate(sys *system.System) []TestResult {
	sysPkg := sys.NewPackage(p.Name, sys)

	var results []TestResult

	results = append(results, ValidateValue(p.Name, "installed", p.Installed, sysPkg.Installed))
	if !p.Installed {
		return results
	}
	results = append(results, ValidateValues(p.Name, "version", p.Versions, sysPkg.Versions))

	return results
}

func NewPackage(sysPackage system.Package) *Package {
	name := sysPackage.Name()
	versions, _ := sysPackage.Versions()
	installed, _ := sysPackage.Installed()
	return &Package{
		Name:      name,
		Versions:  versions,
		Installed: installed.(bool),
	}
}