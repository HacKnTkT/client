package teams

import (
	"context"
	"fmt"
	"strings"

	"github.com/keybase/client/go/libkb"
	"github.com/keybase/client/go/protocol/keybase1"
)

// Things TeamLoader uses that are mocked out for tests.
type LoaderContext interface {
	// Get new links from the server.
	GetNewLinksFromServer(ctx context.Context,
		teamID keybase1.TeamID, lowSeqno keybase1.Seqno, lowGen keybase1.PerTeamKeyGeneration,
		readSubteamID *keybase1.TeamID) (*rawTeam, error)
	// Get full links from the server.
	// Does not guarantee that the server returned the correct links, nor that they are unstubbed.
	GetLinksFromServer(ctx context.Context,
		teamID keybase1.TeamID, requestSeqnos []keybase1.Seqno,
		readSubteamID *keybase1.TeamID) (*rawTeam, error)
	GetMe(context.Context) (keybase1.UserVersion, error)
	// Lookup the eldest seqno of a user. Can use the cache.
	LookupEldestSeqno(context.Context, keybase1.UID) (keybase1.Seqno, error)
	ResolveNameToIDUntrusted(context.Context, keybase1.TeamName) (keybase1.TeamID, error)
	// Get the current user's per-user-key's derived encryption key (full).
	PerUserEncryptionKey(ctx context.Context, userSeqno keybase1.Seqno) (*libkb.NaclDHKeyPair, error)
	MerkleLookup(ctx context.Context, teamID keybase1.TeamID) (r1 keybase1.Seqno, r2 keybase1.LinkID, err error)
	MerkleLookupLeafAtHashMeta(ctx context.Context, leafID keybase1.UserOrTeamID, hm keybase1.HashMeta) (leaf *libkb.MerkleGenericLeaf, err error)
	ForceLinkMapRefreshForUser(ctx context.Context, uid keybase1.UID) (linkMap map[keybase1.Seqno]keybase1.LinkID, err error)
	LoadKeyV2(ctx context.Context, uid keybase1.UID, kid keybase1.KID) (*keybase1.UserPlusKeysV2, *keybase1.PublicKeyV2NaCl, map[keybase1.Seqno]keybase1.LinkID, error)
}

// The main LoaderContext is G.
type LoaderContextG struct {
	libkb.Contextified
}

var _ LoaderContext = (*LoaderContextG)(nil)

func NewLoaderContextFromG(g *libkb.GlobalContext) LoaderContext {
	return &LoaderContextG{
		Contextified: libkb.NewContextified(g),
	}
}

func (l *LoaderContextG) GetNewLinksFromServer(ctx context.Context,
	teamID keybase1.TeamID, lowSeqno keybase1.Seqno, lowGen keybase1.PerTeamKeyGeneration,
	readSubteamID *keybase1.TeamID) (*rawTeam, error) {

	arg := libkb.NewRetryAPIArg("team/get")
	arg.NetContext = ctx
	arg.SessionType = libkb.APISessionTypeREQUIRED
	arg.Args = libkb.HTTPArgs{
		"id":               libkb.S{Val: teamID.String()},
		"low":              libkb.I{Val: int(lowSeqno)},
		"per_team_key_low": libkb.I{Val: int(lowGen)},
	}
	if readSubteamID != nil {
		arg.Args["read_subteam_id"] = libkb.S{Val: readSubteamID.String()}
	}

	var rt rawTeam
	if err := l.G().API.GetDecode(arg, &rt); err != nil {
		return nil, err
	}
	if !rt.ID.Eq(teamID) {
		return nil, fmt.Errorf("server returned wrong team ID: %v != %v", rt.ID, teamID)
	}
	return &rt, nil
}

func (l *LoaderContextG) GetLinksFromServer(ctx context.Context,
	teamID keybase1.TeamID, requestSeqnos []keybase1.Seqno, readSubteamID *keybase1.TeamID) (*rawTeam, error) {

	var seqnoStrs []string
	for _, seqno := range requestSeqnos {
		seqnoStrs = append(seqnoStrs, fmt.Sprintf("%d", int(seqno)))
	}
	seqnoCommas := strings.Join(seqnoStrs, ",")

	arg := libkb.NewRetryAPIArg("team/get")
	arg.NetContext = ctx
	arg.SessionType = libkb.APISessionTypeREQUIRED
	arg.Args = libkb.HTTPArgs{
		"id":     libkb.S{Val: teamID.String()},
		"seqnos": libkb.S{Val: seqnoCommas},
	}
	if readSubteamID != nil {
		arg.Args["read_subteam_id"] = libkb.S{Val: readSubteamID.String()}
	}

	var rt rawTeam
	if err := l.G().API.GetDecode(arg, &rt); err != nil {
		return nil, err
	}
	if !rt.ID.Eq(teamID) {
		return nil, fmt.Errorf("server returned wrong team ID: %v != %v", rt.ID, teamID)
	}
	return &rt, nil
}

func (l *LoaderContextG) GetMe(ctx context.Context) (res keybase1.UserVersion, err error) {
	return loadUserVersionByUID(ctx, l.G(), l.G().Env.GetUID())
}

func (l *LoaderContextG) LookupEldestSeqno(ctx context.Context, uid keybase1.UID) (keybase1.Seqno, error) {
	// Lookup the latest eldest seqno for that uid.
	// This value may come from a cache.
	upak, err := loadUPAK2(ctx, l.G(), uid, false /*forcePoll */)
	if err != nil {
		return keybase1.Seqno(1), err
	}
	return upak.Current.EldestSeqno, nil
}

// Resolve a team name to a team ID.
// Will always hit the server for subteams. The server can lie in this return value.
func (l *LoaderContextG) ResolveNameToIDUntrusted(ctx context.Context, teamName keybase1.TeamName) (id keybase1.TeamID, err error) {
	// For root team names, just hash.
	if teamName.IsRootTeam() {
		return teamName.ToTeamID(), nil
	}

	arg := libkb.NewRetryAPIArg("team/get")
	arg.NetContext = ctx
	arg.SessionType = libkb.APISessionTypeREQUIRED
	arg.Args = libkb.HTTPArgs{
		"name":        libkb.S{Val: teamName.String()},
		"lookup_only": libkb.B{Val: true},
	}

	var rt rawTeam
	if err := l.G().API.GetDecode(arg, &rt); err != nil {
		return id, err
	}
	id = rt.ID
	if !id.Exists() {
		return id, fmt.Errorf("could not resolve team name: %v", teamName.String())
	}
	return id, nil
}

func (l *LoaderContextG) PerUserEncryptionKey(ctx context.Context, userSeqno keybase1.Seqno) (*libkb.NaclDHKeyPair, error) {
	kr, err := l.G().GetPerUserKeyring()
	if err != nil {
		return nil, err
	}
	// Try to get it locally, if that fails try again after syncing.
	encKey, err := kr.GetEncryptionKeyBySeqno(ctx, userSeqno)
	if err == nil {
		return encKey, err
	}
	if err := kr.Sync(ctx); err != nil {
		return nil, err
	}
	encKey, err = kr.GetEncryptionKeyBySeqno(ctx, userSeqno)
	return encKey, err
}

func (l *LoaderContextG) MerkleLookup(ctx context.Context, teamID keybase1.TeamID) (r1 keybase1.Seqno, r2 keybase1.LinkID, err error) {
	leaf, err := l.G().GetMerkleClient().LookupTeam(ctx, teamID)
	if err != nil {
		return r1, r2, err
	}
	if !leaf.TeamID.Eq(teamID) {
		return r1, r2, fmt.Errorf("merkle returned wrong leaf: %v != %v", leaf.TeamID.String(), teamID.String())
	}
	if leaf.Private == nil {
		return r1, r2, fmt.Errorf("merkle returned nil leaf")
	}
	return leaf.Private.Seqno, leaf.Private.LinkID.Export(), nil
}

func (l *LoaderContextG) MerkleLookupLeafAtHashMeta(ctx context.Context, leafID keybase1.UserOrTeamID, hm keybase1.HashMeta) (leaf *libkb.MerkleGenericLeaf, err error) {
	return l.G().MerkleClient.LookupLeafAtHashMeta(ctx, leafID, hm)
}

func (l *LoaderContextG) ForceLinkMapRefreshForUser(ctx context.Context, uid keybase1.UID) (linkMap map[keybase1.Seqno]keybase1.LinkID, err error) {
	arg := libkb.NewLoadUserArgBase(l.G()).WithNetContext(ctx).WithUID(uid).WithForcePoll()
	upak, _, err := l.G().GetUPAKLoader().LoadV2(*arg)
	if err != nil {
		return nil, err
	}
	return upak.SeqnoLinkIDs, nil
}

func (l *LoaderContextG) LoadKeyV2(ctx context.Context, uid keybase1.UID, kid keybase1.KID) (*keybase1.UserPlusKeysV2, *keybase1.PublicKeyV2NaCl, map[keybase1.Seqno]keybase1.LinkID, error) {
	return l.G().GetUPAKLoader().LoadKeyV2(ctx, uid, kid)
}
