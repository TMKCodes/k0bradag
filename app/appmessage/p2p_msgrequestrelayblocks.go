package appmessage

import (
	"github.com/kobradag/kobrad/domain/consensus/model/externalapi"
)

// MaxRequestRelayBlocksHashes is the maximum number of hashes that can
// be in a single RequestRelayBlocks message.
const MaxRequestRelayBlocksHashes = MaxInvPerMsg

// MsgRequestRelayBlocks implements the Message interface and represents a kobra
// RequestRelayBlocks message. It is used to request blocks as part of the block
// relay protocol.
type MsgRequestRelayBlocks struct {
	baseMessage
	Hashes []*externalapi.DomainHash
}

// Command returns the protocol command string for the message. This is part
// of the Message interface implementation.
func (msg *MsgRequestRelayBlocks) Command() MessageCommand {
	return CmdRequestRelayBlocks
}

// NewMsgRequestRelayBlocks returns a new kobra RequestRelayBlocks message that conforms to
// the Message interface. See MsgRequestRelayBlocks for details.
func NewMsgRequestRelayBlocks(hashes []*externalapi.DomainHash) *MsgRequestRelayBlocks {
	return &MsgRequestRelayBlocks{
		Hashes: hashes,
	}
}
