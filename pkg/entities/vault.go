package entities

type Vault struct {
	Email			string `db:"email"`							// one account per user
	IDHash    string `db:"identity_hash"` 	// user email + passcode as PK
	Salt      []byte `db:"salt"`						// random salt used for Argon2 hashing password
	Nonce     []byte `db:"nonce"`						// Nonce used for AES-GCM
	VaultBlob []byte `db:"vault_blob"`			// Encrypted JSON Valut
}
