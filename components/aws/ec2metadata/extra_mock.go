package ec2metadata

// Iface3PartyEC2Metadata interface for mock github
//go:generate mockery --name Iface3PartyEC2Metadata --output ../mocks
type Iface3PartyEC2Metadata interface {
	GetMetadata(p string) (string, error)
}
