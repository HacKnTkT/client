// Auto-generated by avdl-compiler v1.3.17 (https://github.com/keybase/node-avdl-compiler)
//   Input file: avdl/keybase1/common.avdl

package keybase1

import (
	"github.com/keybase/go-framed-msgpack-rpc/rpc"
)

type Time int64

func (o Time) DeepCopy() Time {
	return o
}

type DurationSec float64

func (o DurationSec) DeepCopy() DurationSec {
	return o
}

type StringKVPair struct {
	Key   string `codec:"key" json:"key"`
	Value string `codec:"value" json:"value"`
}

func (o StringKVPair) DeepCopy() StringKVPair {
	return StringKVPair{
		Key:   o.Key,
		Value: o.Value,
	}
}

type Status struct {
	Code   int            `codec:"code" json:"code"`
	Name   string         `codec:"name" json:"name"`
	Desc   string         `codec:"desc" json:"desc"`
	Fields []StringKVPair `codec:"fields" json:"fields"`
}

func (o Status) DeepCopy() Status {
	return Status{
		Code: o.Code,
		Name: o.Name,
		Desc: o.Desc,
		Fields: (func(x []StringKVPair) []StringKVPair {
			var ret []StringKVPair
			for _, v := range x {
				vCopy := v.DeepCopy()
				ret = append(ret, vCopy)
			}
			return ret
		})(o.Fields),
	}
}

type UID string

func (o UID) DeepCopy() UID {
	return o
}

type DeviceID string

func (o DeviceID) DeepCopy() DeviceID {
	return o
}

type SigID string

func (o SigID) DeepCopy() SigID {
	return o
}

type LeaseID string

func (o LeaseID) DeepCopy() LeaseID {
	return o
}

type KID string

func (o KID) DeepCopy() KID {
	return o
}

type LinkID string

func (o LinkID) DeepCopy() LinkID {
	return o
}

type BinaryKID []byte

func (o BinaryKID) DeepCopy() BinaryKID {
	return (func(x []byte) []byte {
		if x == nil {
			return nil
		}
		return append([]byte(nil), x...)
	})(o)
}

type TLFID string

func (o TLFID) DeepCopy() TLFID {
	return o
}

type TeamID string

func (o TeamID) DeepCopy() TeamID {
	return o
}

type UserOrTeamID string

func (o UserOrTeamID) DeepCopy() UserOrTeamID {
	return o
}

type Seqno int64

func (o Seqno) DeepCopy() Seqno {
	return o
}

type SeqType int

const (
	SeqType_PUBLIC      SeqType = 1
	SeqType_PRIVATE     SeqType = 2
	SeqType_SEMIPRIVATE SeqType = 3
)

func (o SeqType) DeepCopy() SeqType { return o }

var SeqTypeMap = map[string]SeqType{
	"PUBLIC":      1,
	"PRIVATE":     2,
	"SEMIPRIVATE": 3,
}

var SeqTypeRevMap = map[SeqType]string{
	1: "PUBLIC",
	2: "PRIVATE",
	3: "SEMIPRIVATE",
}

func (e SeqType) String() string {
	if v, ok := SeqTypeRevMap[e]; ok {
		return v
	}
	return ""
}

type Bytes32 [32]byte

func (o Bytes32) DeepCopy() Bytes32 {
	var ret Bytes32
	copy(ret[:], o[:])
	return ret
}

type Text struct {
	Data   string `codec:"data" json:"data"`
	Markup bool   `codec:"markup" json:"markup"`
}

func (o Text) DeepCopy() Text {
	return Text{
		Data:   o.Data,
		Markup: o.Markup,
	}
}

type PGPIdentity struct {
	Username string `codec:"username" json:"username"`
	Comment  string `codec:"comment" json:"comment"`
	Email    string `codec:"email" json:"email"`
}

func (o PGPIdentity) DeepCopy() PGPIdentity {
	return PGPIdentity{
		Username: o.Username,
		Comment:  o.Comment,
		Email:    o.Email,
	}
}

type PublicKey struct {
	KID               KID           `codec:"KID" json:"KID"`
	PGPFingerprint    string        `codec:"PGPFingerprint" json:"PGPFingerprint"`
	PGPIdentities     []PGPIdentity `codec:"PGPIdentities" json:"PGPIdentities"`
	IsSibkey          bool          `codec:"isSibkey" json:"isSibkey"`
	IsEldest          bool          `codec:"isEldest" json:"isEldest"`
	ParentID          string        `codec:"parentID" json:"parentID"`
	DeviceID          DeviceID      `codec:"deviceID" json:"deviceID"`
	DeviceDescription string        `codec:"deviceDescription" json:"deviceDescription"`
	DeviceType        string        `codec:"deviceType" json:"deviceType"`
	CTime             Time          `codec:"cTime" json:"cTime"`
	ETime             Time          `codec:"eTime" json:"eTime"`
	IsRevoked         bool          `codec:"isRevoked" json:"isRevoked"`
}

func (o PublicKey) DeepCopy() PublicKey {
	return PublicKey{
		KID:            o.KID.DeepCopy(),
		PGPFingerprint: o.PGPFingerprint,
		PGPIdentities: (func(x []PGPIdentity) []PGPIdentity {
			var ret []PGPIdentity
			for _, v := range x {
				vCopy := v.DeepCopy()
				ret = append(ret, vCopy)
			}
			return ret
		})(o.PGPIdentities),
		IsSibkey:          o.IsSibkey,
		IsEldest:          o.IsEldest,
		ParentID:          o.ParentID,
		DeviceID:          o.DeviceID.DeepCopy(),
		DeviceDescription: o.DeviceDescription,
		DeviceType:        o.DeviceType,
		CTime:             o.CTime.DeepCopy(),
		ETime:             o.ETime.DeepCopy(),
		IsRevoked:         o.IsRevoked,
	}
}

type KeybaseTime struct {
	Unix  Time  `codec:"unix" json:"unix"`
	Chain Seqno `codec:"chain" json:"chain"`
}

func (o KeybaseTime) DeepCopy() KeybaseTime {
	return KeybaseTime{
		Unix:  o.Unix.DeepCopy(),
		Chain: o.Chain.DeepCopy(),
	}
}

type RevokedKey struct {
	Key  PublicKey   `codec:"key" json:"key"`
	Time KeybaseTime `codec:"time" json:"time"`
	By   KID         `codec:"by" json:"by"`
}

func (o RevokedKey) DeepCopy() RevokedKey {
	return RevokedKey{
		Key:  o.Key.DeepCopy(),
		Time: o.Time.DeepCopy(),
		By:   o.By.DeepCopy(),
	}
}

type User struct {
	Uid      UID    `codec:"uid" json:"uid"`
	Username string `codec:"username" json:"username"`
}

func (o User) DeepCopy() User {
	return User{
		Uid:      o.Uid.DeepCopy(),
		Username: o.Username,
	}
}

type Device struct {
	Type         string   `codec:"type" json:"type"`
	Name         string   `codec:"name" json:"name"`
	DeviceID     DeviceID `codec:"deviceID" json:"deviceID"`
	CTime        Time     `codec:"cTime" json:"cTime"`
	MTime        Time     `codec:"mTime" json:"mTime"`
	LastUsedTime Time     `codec:"lastUsedTime" json:"lastUsedTime"`
	EncryptKey   KID      `codec:"encryptKey" json:"encryptKey"`
	VerifyKey    KID      `codec:"verifyKey" json:"verifyKey"`
	Status       int      `codec:"status" json:"status"`
}

func (o Device) DeepCopy() Device {
	return Device{
		Type:         o.Type,
		Name:         o.Name,
		DeviceID:     o.DeviceID.DeepCopy(),
		CTime:        o.CTime.DeepCopy(),
		MTime:        o.MTime.DeepCopy(),
		LastUsedTime: o.LastUsedTime.DeepCopy(),
		EncryptKey:   o.EncryptKey.DeepCopy(),
		VerifyKey:    o.VerifyKey.DeepCopy(),
		Status:       o.Status,
	}
}

type DeviceType int

const (
	DeviceType_DESKTOP DeviceType = 0
	DeviceType_MOBILE  DeviceType = 1
)

func (o DeviceType) DeepCopy() DeviceType { return o }

var DeviceTypeMap = map[string]DeviceType{
	"DESKTOP": 0,
	"MOBILE":  1,
}

var DeviceTypeRevMap = map[DeviceType]string{
	0: "DESKTOP",
	1: "MOBILE",
}

func (e DeviceType) String() string {
	if v, ok := DeviceTypeRevMap[e]; ok {
		return v
	}
	return ""
}

type Stream struct {
	Fd int `codec:"fd" json:"fd"`
}

func (o Stream) DeepCopy() Stream {
	return Stream{
		Fd: o.Fd,
	}
}

type LogLevel int

const (
	LogLevel_NONE     LogLevel = 0
	LogLevel_DEBUG    LogLevel = 1
	LogLevel_INFO     LogLevel = 2
	LogLevel_NOTICE   LogLevel = 3
	LogLevel_WARN     LogLevel = 4
	LogLevel_ERROR    LogLevel = 5
	LogLevel_CRITICAL LogLevel = 6
	LogLevel_FATAL    LogLevel = 7
)

func (o LogLevel) DeepCopy() LogLevel { return o }

var LogLevelMap = map[string]LogLevel{
	"NONE":     0,
	"DEBUG":    1,
	"INFO":     2,
	"NOTICE":   3,
	"WARN":     4,
	"ERROR":    5,
	"CRITICAL": 6,
	"FATAL":    7,
}

var LogLevelRevMap = map[LogLevel]string{
	0: "NONE",
	1: "DEBUG",
	2: "INFO",
	3: "NOTICE",
	4: "WARN",
	5: "ERROR",
	6: "CRITICAL",
	7: "FATAL",
}

func (e LogLevel) String() string {
	if v, ok := LogLevelRevMap[e]; ok {
		return v
	}
	return ""
}

type ClientType int

const (
	ClientType_NONE       ClientType = 0
	ClientType_CLI        ClientType = 1
	ClientType_GUI_MAIN   ClientType = 2
	ClientType_KBFS       ClientType = 3
	ClientType_GUI_HELPER ClientType = 4
)

func (o ClientType) DeepCopy() ClientType { return o }

var ClientTypeMap = map[string]ClientType{
	"NONE":       0,
	"CLI":        1,
	"GUI_MAIN":   2,
	"KBFS":       3,
	"GUI_HELPER": 4,
}

var ClientTypeRevMap = map[ClientType]string{
	0: "NONE",
	1: "CLI",
	2: "GUI_MAIN",
	3: "KBFS",
	4: "GUI_HELPER",
}

type UserVersionVector struct {
	Id       int64 `codec:"id" json:"id"`
	SigHints int   `codec:"sigHints" json:"sigHints"`
	SigChain int64 `codec:"sigChain" json:"sigChain"`
	CachedAt Time  `codec:"cachedAt" json:"cachedAt"`
}

func (o UserVersionVector) DeepCopy() UserVersionVector {
	return UserVersionVector{
		Id:       o.Id,
		SigHints: o.SigHints,
		SigChain: o.SigChain,
		CachedAt: o.CachedAt.DeepCopy(),
	}
}

type PerUserKeyGeneration int

func (o PerUserKeyGeneration) DeepCopy() PerUserKeyGeneration {
	return o
}

type PerUserKey struct {
	Gen         int   `codec:"gen" json:"gen"`
	Seqno       Seqno `codec:"seqno" json:"seqno"`
	SigKID      KID   `codec:"sigKID" json:"sigKID"`
	EncKID      KID   `codec:"encKID" json:"encKID"`
	SignedByKID KID   `codec:"signedByKID" json:"signedByKID"`
}

func (o PerUserKey) DeepCopy() PerUserKey {
	return PerUserKey{
		Gen:         o.Gen,
		Seqno:       o.Seqno.DeepCopy(),
		SigKID:      o.SigKID.DeepCopy(),
		EncKID:      o.EncKID.DeepCopy(),
		SignedByKID: o.SignedByKID.DeepCopy(),
	}
}

type UserPlusKeys struct {
	Uid               UID               `codec:"uid" json:"uid"`
	Username          string            `codec:"username" json:"username"`
	EldestSeqno       Seqno             `codec:"eldestSeqno" json:"eldestSeqno"`
	DeviceKeys        []PublicKey       `codec:"deviceKeys" json:"deviceKeys"`
	RevokedDeviceKeys []RevokedKey      `codec:"revokedDeviceKeys" json:"revokedDeviceKeys"`
	PGPKeyCount       int               `codec:"pgpKeyCount" json:"pgpKeyCount"`
	Uvv               UserVersionVector `codec:"uvv" json:"uvv"`
	DeletedDeviceKeys []PublicKey       `codec:"deletedDeviceKeys" json:"deletedDeviceKeys"`
	PerUserKeys       []PerUserKey      `codec:"perUserKeys" json:"perUserKeys"`
}

func (o UserPlusKeys) DeepCopy() UserPlusKeys {
	return UserPlusKeys{
		Uid:         o.Uid.DeepCopy(),
		Username:    o.Username,
		EldestSeqno: o.EldestSeqno.DeepCopy(),
		DeviceKeys: (func(x []PublicKey) []PublicKey {
			var ret []PublicKey
			for _, v := range x {
				vCopy := v.DeepCopy()
				ret = append(ret, vCopy)
			}
			return ret
		})(o.DeviceKeys),
		RevokedDeviceKeys: (func(x []RevokedKey) []RevokedKey {
			var ret []RevokedKey
			for _, v := range x {
				vCopy := v.DeepCopy()
				ret = append(ret, vCopy)
			}
			return ret
		})(o.RevokedDeviceKeys),
		PGPKeyCount: o.PGPKeyCount,
		Uvv:         o.Uvv.DeepCopy(),
		DeletedDeviceKeys: (func(x []PublicKey) []PublicKey {
			var ret []PublicKey
			for _, v := range x {
				vCopy := v.DeepCopy()
				ret = append(ret, vCopy)
			}
			return ret
		})(o.DeletedDeviceKeys),
		PerUserKeys: (func(x []PerUserKey) []PerUserKey {
			var ret []PerUserKey
			for _, v := range x {
				vCopy := v.DeepCopy()
				ret = append(ret, vCopy)
			}
			return ret
		})(o.PerUserKeys),
	}
}

type UserOrTeamLite struct {
	Id   UserOrTeamID `codec:"id" json:"id"`
	Name string       `codec:"name" json:"name"`
}

func (o UserOrTeamLite) DeepCopy() UserOrTeamLite {
	return UserOrTeamLite{
		Id:   o.Id.DeepCopy(),
		Name: o.Name,
	}
}

type RemoteTrack struct {
	Username string `codec:"username" json:"username"`
	Uid      UID    `codec:"uid" json:"uid"`
	LinkID   LinkID `codec:"linkID" json:"linkID"`
}

func (o RemoteTrack) DeepCopy() RemoteTrack {
	return RemoteTrack{
		Username: o.Username,
		Uid:      o.Uid.DeepCopy(),
		LinkID:   o.LinkID.DeepCopy(),
	}
}

type UserPlusAllKeys struct {
	Base         UserPlusKeys  `codec:"base" json:"base"`
	PGPKeys      []PublicKey   `codec:"pgpKeys" json:"pgpKeys"`
	RemoteTracks []RemoteTrack `codec:"remoteTracks" json:"remoteTracks"`
}

func (o UserPlusAllKeys) DeepCopy() UserPlusAllKeys {
	return UserPlusAllKeys{
		Base: o.Base.DeepCopy(),
		PGPKeys: (func(x []PublicKey) []PublicKey {
			var ret []PublicKey
			for _, v := range x {
				vCopy := v.DeepCopy()
				ret = append(ret, vCopy)
			}
			return ret
		})(o.PGPKeys),
		RemoteTracks: (func(x []RemoteTrack) []RemoteTrack {
			var ret []RemoteTrack
			for _, v := range x {
				vCopy := v.DeepCopy()
				ret = append(ret, vCopy)
			}
			return ret
		})(o.RemoteTracks),
	}
}

type MerkleTreeID int

const (
	MerkleTreeID_MASTER           MerkleTreeID = 0
	MerkleTreeID_KBFS_PUBLIC      MerkleTreeID = 1
	MerkleTreeID_KBFS_PRIVATE     MerkleTreeID = 2
	MerkleTreeID_KBFS_PRIVATETEAM MerkleTreeID = 3
)

func (o MerkleTreeID) DeepCopy() MerkleTreeID { return o }

var MerkleTreeIDMap = map[string]MerkleTreeID{
	"MASTER":           0,
	"KBFS_PUBLIC":      1,
	"KBFS_PRIVATE":     2,
	"KBFS_PRIVATETEAM": 3,
}

var MerkleTreeIDRevMap = map[MerkleTreeID]string{
	0: "MASTER",
	1: "KBFS_PUBLIC",
	2: "KBFS_PRIVATE",
	3: "KBFS_PRIVATETEAM",
}

// SocialAssertionService is a service that can be used to assert proofs for a
// user.
type SocialAssertionService string

func (o SocialAssertionService) DeepCopy() SocialAssertionService {
	return o
}

// SocialAssertion contains a service and username for that service, that
// together form an assertion about a user. Resolving an assertion requires
// that the user posts a Keybase proof on the asserted service as the asserted
// user.
type SocialAssertion struct {
	User    string                 `codec:"user" json:"user"`
	Service SocialAssertionService `codec:"service" json:"service"`
}

func (o SocialAssertion) DeepCopy() SocialAssertion {
	return SocialAssertion{
		User:    o.User,
		Service: o.Service.DeepCopy(),
	}
}

// UserResolution maps how an unresolved user assertion has been resolved.
type UserResolution struct {
	Assertion SocialAssertion `codec:"assertion" json:"assertion"`
	UserID    UID             `codec:"userID" json:"userID"`
}

func (o UserResolution) DeepCopy() UserResolution {
	return UserResolution{
		Assertion: o.Assertion.DeepCopy(),
		UserID:    o.UserID.DeepCopy(),
	}
}

type CommonInterface interface {
}

func CommonProtocol(i CommonInterface) rpc.Protocol {
	return rpc.Protocol{
		Name:    "keybase.1.Common",
		Methods: map[string]rpc.ServeHandlerDescription{},
	}
}

type CommonClient struct {
	Cli rpc.GenericClient
}
