package procfs

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestSysctlInts(t *testing.T) {
	fs := getProcFixtures(t)

	for _, tc := range []struct {
		sysctl string
		want   []int
	}{
		{"kernel.random.entropy_avail", []int{3943}},
		{"vm.lowmem_reserve_ratio", []int{256, 256, 32, 0, 0}},
	} {
		t.Run(tc.sysctl, func(t *testing.T) {
			got, err := fs.SysctlInts(tc.sysctl)
			if err != nil {
				t.Fatal(err)
			}
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Fatalf("unexpected syscall value(-want +got):\n%s", diff)
			}
		})
	}
}

func TestSysctlStrings(t *testing.T) {
	fs := getProcFixtures(t)

	for _, tc := range []struct {
		sysctl string
		want   []string
	}{
		{"kernel.seccomp.actions_avail", []string{"kill_process", "kill_thread", "trap", "errno", "trace", "log", "allow"}},
	} {
		t.Run(tc.sysctl, func(t *testing.T) {
			got, err := fs.SysctlStrings(tc.sysctl)
			if err != nil {
				t.Fatal(err)
			}
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Fatalf("unexpected syscall value(-want +got):\n%s", diff)
			}
		})
	}
}

func TestSysctlIntsError(t *testing.T) {
	fs := getProcFixtures(t)

	for _, tc := range []struct {
		sysctl string
		want   string
	}{
		{"kernel.seccomp.actions_avail", "field 0 in sysctl kernel.seccomp.actions_avail is not a valid int: strconv.ParseInt: parsing \"kill_process\": invalid syntax"},
	} {
		t.Run(tc.sysctl, func(t *testing.T) {
			_, err := fs.SysctlInts(tc.sysctl)
			if err == nil {
				t.Fatal("expected error")
			}
			if diff := cmp.Diff(tc.want, err.Error()); diff != "" {
				t.Fatalf("unexpected syscall value(-want +got):\n%s", diff)
			}
		})
	}
}
