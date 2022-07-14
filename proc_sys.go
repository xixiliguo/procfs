package procfs

import (
	"fmt"
	"strings"

	"github.com/prometheus/procfs/internal/util"
)

func sysctlToPath(sysctl string) string {
	return strings.Replace(sysctl, ".", "/", -1)
}
func (fs FS) SysctlStrings(sysctl string) ([]string, error) {
	value, err := util.SysReadFile(fs.proc.Path("sys", sysctlToPath(sysctl)))
	if err != nil {
		return nil, err
	}
	return strings.Fields(value), nil

}
func (fs FS) SysctlInts(sysctl string) ([]int, error) {
	fields, err := fs.SysctlStrings(sysctl)
	if err != nil {
		return nil, err
	}

	values := make([]int, len(fields))
	for i, f := range fields {
		vp := util.NewValueParser(f)
		values[i] = vp.Int()
		if err := vp.Err(); err != nil {
			return nil, fmt.Errorf("field %d in sysctl %s is not a valid int: %w", i, sysctl, err)
		}
	}
	return values, nil
}
