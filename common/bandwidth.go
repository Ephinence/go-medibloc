package common

import (
	"github.com/medibloc/go-medibloc/util"
)

// Bandwidth is structure for cpu and net bandwidth
type Bandwidth struct {
	cpuUsage uint64
	netUsage uint64
}

// NewBandwidth returns new bandwidth
func NewBandwidth(cpu, net uint64) *Bandwidth {
	return &Bandwidth{
		cpuUsage: cpu,
		netUsage: net,
	}
}

// Clone clones existing bandwidth struct.
func (b *Bandwidth) Clone() *Bandwidth {
	return &Bandwidth{
		cpuUsage: b.cpuUsage,
		netUsage: b.netUsage,
	}
}

// CPUUsage returns cpuUsage
func (b *Bandwidth) CPUUsage() uint64 {
	return b.cpuUsage
}

// SetCPUUsage sets cpuUsage
func (b *Bandwidth) SetCPUUsage(cpuUsage uint64) {
	b.cpuUsage = cpuUsage
}

// NetUsage returns netUsage
func (b *Bandwidth) NetUsage() uint64 {
	return b.netUsage
}

// SetNetUsage sets netUsage
func (b *Bandwidth) SetNetUsage(netUsage uint64) {
	b.netUsage = netUsage
}

// Add add obj's bandwidth
func (b *Bandwidth) Add(obj *Bandwidth) {
	b.cpuUsage += obj.cpuUsage
	b.netUsage += obj.netUsage
}

// Sub subtract obj's bandwidth
func (b *Bandwidth) Sub(obj *Bandwidth) {
	b.cpuUsage -= obj.cpuUsage
	b.netUsage -= obj.netUsage
}

// CalcPoints multiply bandwidth and price
func (b *Bandwidth) CalcPoints(p Price) (points *util.Uint128, err error) {
	cpuPoints, err := p.CPUPrice.Mul(util.NewUint128FromUint(b.cpuUsage))
	if err != nil {
		return nil, err
	}
	netPoints, err := p.NetPrice.Mul(util.NewUint128FromUint(b.netUsage))
	if err != nil {
		return nil, err
	}
	return cpuPoints.Add(netPoints)
}

// IsZero return true when cpu and net value are both zero
func (b Bandwidth) IsZero() bool {
	return b.cpuUsage == 0 && b.netUsage == 0
}

// Price is structure for prices of cpu and net
type Price struct {
	CPUPrice *util.Uint128
	NetPrice *util.Uint128
}
