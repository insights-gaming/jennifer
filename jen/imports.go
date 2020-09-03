package jen

import "strings"

type ImportGroupingPolicy func(map[string]importdef) []map[string]importdef

func NoImportGrouping(imports map[string]importdef) []map[string]importdef {
	return []map[string]importdef{imports}
}

func SeparateStdlib(imports map[string]importdef) []map[string]importdef {
	stdlib := make(map[string]importdef, len(imports))
	external := make(map[string]importdef, len(imports))

	for path, def := range imports {
		if strings.ContainsRune(strings.Split(path, "/")[0], '.') {
			external[path] = def
		} else {
			stdlib[path] = def
		}
	}

	return []map[string]importdef{stdlib, external}
}
