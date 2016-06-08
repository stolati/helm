package chartutil

import (
	"io/ioutil"

	"github.com/ghodss/yaml"

	"k8s.io/helm/pkg/proto/hapi/chart"
)

// UnmarshalChartfile takes raw Chart.yaml data and unmarshals it.
func UnmarshalChartfile(data []byte) (*chart.Metadata, error) {
	y := &chart.Metadata{}
	err := yaml.Unmarshal(data, y)
	if err != nil {
		return nil, err
	}
	return y, nil
}

// LoadChartfile loads a Chart.yaml file into a *chart.Metadata.
func LoadChartfile(filename string) (*chart.Metadata, error) {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return UnmarshalChartfile(b)
}

// SaveChartfile saves the given metadata as a Chart.yaml file at the given path.
//
// 'filename' should be the complete path and filename ('foo/Chart.yaml')
func SaveChartfile(filename string, cf *chart.Metadata) error {
	out, err := yaml.Marshal(cf)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filename, out, 0755)
}