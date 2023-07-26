package config

import (
	"gopkg.in/ini.v1"
)

func ReadInI(filePath string) (*ini.File, error) {
	optionObj := ini.LoadOptions{
		AllowPythonMultilineValues: true,
		ShortCircuit:               true,
	}

	return ini.LoadSources(optionObj, filePath)
}
