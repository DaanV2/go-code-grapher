package golang

import (
	"path/filepath"
	"strings"
)

// QualifyPackageName returns the fully qualified package name using the module path
//	- moduleDir: Absolute path to the module directory
//	- moduleName: The module name from go.mod, eg "github.com/<org>/<repo>"
//	- pkgDir: The directory of the package
//	- pkgName: The name of the package
func QualifyPackageName(moduleDir, moduleName, pkgDir, pkgName string) (string, error) {
	pkgDir, err := filepath.Abs(pkgDir)
	if err != nil {
		return "", err
	}
	modulepath, err := filepath.Rel(moduleDir, pkgDir)
	if err != nil {
		return "", err
	}

	var built string
	if modulepath == "." {
		built = moduleName
	} else {
		built = moduleName + "/" + modulepath
	}

	if filepath.Separator != '/' {
		built = strings.ReplaceAll(built, string(filepath.Separator), "/")
	}

	return built, nil
}