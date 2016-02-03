package MTProto

import "errors"

// MTProto cloud chat
type CloudChat struct {
	Auth_key_id    Auth_key_id
	Msg_key        Msg_key
	Encrypted_data []byte
}

// MTProto Payload struct
type Payload struct {
	Time            int
	Length          int
	Sequence_number int
}

// MTProto struct
type MTProto struct {
	Salt       int64
	Session_id int64
	Payload    Payload
	Padding    [15]byte
}

// Message key (msg_key)
type Msg_key [16]byte

// Shared key
type Shared_key struct {
	Auth_key_id Auth_key_id
	Auth_key    Auth_key
}

// Auth key id
type Auth_key_id [8]byte

// Auth key
type Auth_key []byte

// AES key
type Aes_key [32]byte

// AES IGE IV
type Aes_ige_iv [32]byte

// Get message key
func (m *MTProto) GetMessageKey() (Msg_key, error) {
	// Generate SHA-1 of m.Salt, m.Session_id and m.Payload
	return Msg_key{}, errors.New("to be implemented")
}

// Sign message key
func (s *Shared_key) SignMessageKey(message_key Msg_key) (Aes_key, Aes_ige_iv, error) {
	// Gets multiple SHA-1
	return Aes_key{}, Aes_ige_iv{}, errors.New("to be implemented")
}

// Get encrypted MTProto
func (m *MTProto) GetEncrypted(aes_key Aes_key, aes_ige_iv Aes_ige_iv) ([]byte, error) {
	// Do AES IGE encryption
	return []byte{}, errors.New("to be implemented")
}
