package hardware

import (
	"github.com/shirou/gopsutil/v3/cpu"
	"golang.org/x/xerrors"
)

type Hardware struct {
	cpu []cpu.InfoStat
}

func NewHardware() (*Hardware, error) {
	c, err := cpu.Info()
	if err != nil {
		return nil, xerrors.Errorf("cpu info error: %w", err)
	}
	return &Hardware{
		cpu: c,
	}, nil
}

func (h *Hardware) GetCPUId() string {
	s := h.cpu[0].PhysicalID
	return s
}
