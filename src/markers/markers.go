package markers

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
)

/// ReplaceMarkers replaces all markers in given content
func ReplaceMarkers(content string, markers map[string]string) string {
	var re *regexp.Regexp

	// Sort keys to avoid random order
	keys := make([]string, 0, len(markers))
	for k := range markers {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, key := range keys {
		value := markers[key]
		re, _ = regexp.Compile("{#\\s*" + key + "\\s*#}") // {#marker#}
		content = re.ReplaceAllString(content, value)
		re, _ = regexp.Compile("{~\\s*" + key + "\\s*~}") // {~marker~}
		content = re.ReplaceAllString(content, value)
		re, _ = regexp.Compile("{\\*\\s*" + key + "\\s*\\*}") // {*marker*}
		content = re.ReplaceAllString(content, value)
	}
	return content
}

/// ProcessMarkers resolves array values into single string replaces
func ProcessMarkers(markers map[string]interface{}, markersSplit map[string]interface{}) map[string]string {
	markersSplitKeys := make([]string, 0, len(markersSplit))
	for k := range markersSplit {
		markersSplitKeys = append(markersSplitKeys, k)
	}

	output := make(map[string]string)
	for key, value := range markers {
		stringValue := value.(string)
		splitChar := markersSplit[key]
		if splitChar == nil || splitChar.(string) == "" {
			output[key] = regexp.QuoteMeta(stringValue)
			continue
		}

		// Split value by character
		for i, slice := range strings.Split(stringValue, splitChar.(string)) {
			output[fmt.Sprintf(regexp.QuoteMeta("%s[%d]"), key, i)] = slice
		}
		output[fmt.Sprintf(regexp.QuoteMeta("%s[%s]"), key, "\\d+")] = ""
	}
	return output
}
