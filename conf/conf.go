package conf

import (
	"path/filepath"
	"time"

	"github.com/iyidan/goutils/config"
	"github.com/iyidan/goutils/mise"
)

var (
	defConf *config.Config
)

// Startup parse config file depend on the ENV
func Startup() {
	var err error
	cluster := GetCluster()
	if len(cluster) == 0 {
		cluster = "dev"
	}
	filename := filepath.Join(mise.GetRootPath(), "runtime", "config_"+cluster+".json")
	defConf, err = config.ParseFromFile(filename)
	if err != nil {
		mise.PanicOnError(err, "config.Startup")
	}
}

// Get get a config with original value
// if k not exists, return nil
func Get(k string) interface{} {
	return defConf.Get(k)
}

// StrictGet get a config with original value
// if k not exists, panic
func StrictGet(k string) interface{} {
	return defConf.StrictGet(k)
}

// String get a config with k, if k not exists or parse error, panic
func String(k string) string {
	return defConf.String(k)
}

// GetTime same as defConf.String method
func GetTime(k string) time.Time {
	return defConf.GetTime(k)
}

// GetDuration same as defConf.String method
func GetDuration(k string) time.Duration {
	return defConf.GetDuration(k)
}

// Int same as defConf.String method
func Int(k string) int {
	return defConf.Int(k)
}

// Int64 same as defConf.String method
func Int64(k string) int64 {
	return defConf.Int64(k)
}

// Float same as defConf.String method
func Float(k string) float64 {
	return defConf.Float(k)
}

// Bool same as defConf.String method
func Bool(k string) bool {
	return defConf.Bool(k)
}

// Slice same as defConf.String method
func Slice(k string) []interface{} {
	return defConf.Slice(k)
}

// SliceString same as defConf.String method
func SliceString(k string) []string {
	return defConf.SliceString(k)
}

// SliceInt same as defConf.String method
func SliceInt(k string) []int {
	return defConf.SliceInt(k)
}

// SliceFloat same as defConf.String method
func SliceFloat(k string) []float64 {
	return defConf.SliceFloat(k)
}

// SliceBool same as defConf.String method
func SliceBool(k string) []bool {
	return defConf.SliceBool(k)
}

// MapString same as defConf.String method
func MapString(k string) map[string]interface{} {
	return defConf.MapString(k)
}

// MapStringString same as defConf.String method
func MapStringString(k string) map[string]string {
	return defConf.MapStringString(k)
}

// MapStringInt same as defConf.String method
func MapStringInt(k string) map[string]int {
	return defConf.MapStringInt(k)
}

// MapStringFloat same as defConf.String method
func MapStringFloat(k string) map[string]float64 {
	return defConf.MapStringFloat(k)
}

// MapStringBool same as defConf.String method
func MapStringBool(k string) map[string]bool {
	return defConf.MapStringBool(k)
}

// MapStringSliceString same as defConf.String method
func MapStringSliceString(k string) map[string][]string {
	return defConf.MapStringSliceString(k)
}

// MapStringSliceInt same as defConf.String method
func MapStringSliceInt(k string) map[string][]int {
	return defConf.MapStringSliceInt(k)
}

// MapStringSliceFloat same as defConf.String method
func MapStringSliceFloat(k string) map[string][]float64 {
	return defConf.MapStringSliceFloat(k)
}

// MapStringSliceBool same as defConf.String method
func MapStringSliceBool(k string) map[string][]bool {
	return defConf.MapStringSliceBool(k)
}

// Unmarshal config k into v
func Unmarshal(k string, v interface{}) error {
	return defConf.Unmarshal(k, v)
}

// MustUnmarshal config k into v, if error, panic
func MustUnmarshal(k string, v interface{}) {
	defConf.MustUnmarshal(k, v)
}
