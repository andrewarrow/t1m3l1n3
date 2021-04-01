package network

type UniverseLink struct {
	Flavor    string
	U1        *Universe
	U2        *Universe
	Host      []string
	Viewer    string
	Following uint64
}

func NewLocalUniverseLink(u1, u2 *Universe) *UniverseLink {
	u := UniverseLink{}
	u.Flavor = "local"
	u.Viewer = "aa"
	u.U1 = u1
	u.U2 = u2

	return &u
}

func (ul *UniverseLink) ShouldDeliverFromViewerToUserInU2(otherUser string) bool {
	fromIndex := ul.U2.UsernameToIndex(otherUser)
	return HasBits(Bits(ul.Following), LookupBit(fromIndex))
}
