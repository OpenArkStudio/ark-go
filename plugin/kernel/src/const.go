package kernelSrc

type AFGUID int64
type ID_TYPE uint32

type ArkMaskType uint16

const (
	PF_SYNC_VIEW  ArkMaskType = iota // send to others
	PF_SYNC_SELF                     // send to self
	PF_REAL_TIME                     // send real-time when changed
	PF_SAVE                          // save to database
	PF_SYNC_TEAM                     // sync to team member
	PF_SYNC_GUILD                    // sync to guild member
	PF_SYNC_MAP                      // sync to all player in same map
	PF_LOG                           // log when changed
)

type ArkDataMask [16]rune

func (mask *ArkDataMask) HaveMask(t ArkMaskType) bool {
	return mask[int(t)] == 1
}
