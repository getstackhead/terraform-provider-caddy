package caddy

import (
	"io/ioutil"
	"os"
	"path"

	"stackhead.io/terraform-caddy-provider/src/markers"
)

func CreateOrUpdateServerBlock(filename string, content string, m Config, markersMap map[string]interface{}, markersSplit map[string]interface{}) (string, error) {
	fullPath := path.Join(m.ConfigFolder, filename)
	content = markers.ReplaceMarkers(content, markers.ProcessMarkers(markersMap, markersSplit))

	if err := ioutil.WriteFile(fullPath, []byte(content), 0744); err != nil {
		return "", err
	}
	return fullPath, nil
}

func RemoveServerBlock(filename string, m Config) error {
	fullPath := path.Join(m.ConfigFolder, filename)
	// Remove configuration
	if err := os.Remove(fullPath); err != nil {
		return err
	}
	return nil
}

func ReadFile(filepath string) ([]byte, error) {
	return ioutil.ReadFile(filepath)
}
