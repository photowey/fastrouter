package collection

// StringValues transform the map<string,string> values -> []string
func StringValues(ctx map[string]string) []string {
	values := make([]string, 0, len(ctx))
	for _, v := range ctx {
		values = append(values, v)
	}

	return values
}
