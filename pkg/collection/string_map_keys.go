package collection

// StringKeys transform the map<string,string> keys -> []string
func StringKeys(ctx map[string]string) []string {
	keys := make([]string, 0, len(ctx))
	for k := range ctx {
		keys = append(keys, k)
	}

	return keys
}
