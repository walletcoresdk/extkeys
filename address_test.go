package extkeys

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"
	"testing"

	"github.com/walletcoresdk/extkeys/base58"
	"github.com/walletcoresdk/extkeys/chaincfg"
	"golang.org/x/crypto/ripemd160"
)

type CustomParamStruct struct {
	PubKeyHashAddrID byte
	ScriptHashAddrID byte
}

var CustomParams = CustomParamStruct{
	PubKeyHashAddrID: 0x30, // starts with L
	ScriptHashAddrID: 0x32, // starts with M
}

// We use this function to be able to test functionality in DecodeAddress for
// defaultNet addresses
func applyCustomParams(params chaincfg.Params, customParams CustomParamStruct) chaincfg.Params {
	params.PubKeyHashAddrID = customParams.PubKeyHashAddrID
	params.ScriptHashAddrID = customParams.ScriptHashAddrID
	return params
}

var customParams = applyCustomParams(chaincfg.MainNetParams, CustomParams)

func TestAddresses(t *testing.T) {
	tests := []struct {
		name    string
		addr    string
		encoded string
		valid   bool
		net     *chaincfg.Params
	}{

		// Unsupported witness versions (version 0 only supported at this point)
		{
			name:  "segwit mainnet witness v1",
			addr:  "bc1pw508d6qejxtdg4y5r3zarvary0c5xw7kw508d6qejxtdg4y5r3zarvary0c5xw7k7grplx",
			valid: false,
			net:   &chaincfg.MainNetParams,
		},
		{
			name:  "segwit mainnet witness v2",
			addr:  "bc1zw508d6qejxtdg4y5r3zarvaryvg6kdaj",
			valid: false,
			net:   &chaincfg.MainNetParams,
		},

		{
			name:  "segwit invalid checksum",
			addr:  "bc1qw508d6qejxtdg4y5r3zarvary0c5xw7kv8f3t5",
			valid: false,
			net:   &chaincfg.MainNetParams,
		},
		{
			name:  "segwit invalid witness version",
			addr:  "BC13W508D6QEJXTDG4Y5R3ZARVARY0C5XW7KN40WF2",
			valid: false,
			net:   &chaincfg.MainNetParams,
		},
	}

	for _, test := range tests {
		// Decode addr and compare error against valid.
		decoded, err := DecodeAddress(test.addr, test.net)
		if (err == nil) != test.valid {
			t.Errorf("%v: decoding test failed: %v", test.name, err)
			return
		}

		if err == nil {
			fmt.Printf("DecodeAddress: %s", decoded.String())
			// Ensure the stringer returns the same address as the
			// original.
			if decodedStringer, ok := decoded.(fmt.Stringer); ok {
				addr := test.addr

				// For Segwit addresses the string representation
				// will always be lower case, so in that case we
				// convert the original to lower case first.
				if strings.Contains(test.name, "segwit") {
					addr = strings.ToLower(addr)
				}

				if addr != decodedStringer.String() {
					t.Errorf("%v: String on decoded value does not match expected value: %v != %v",
						test.name, test.addr, decodedStringer.String())
					return
				}
			}

			// Encode again and compare against the original.
			encoded := decoded.EncodeAddress()
			if test.encoded != encoded {
				t.Errorf("%v: decoding and encoding produced different addressess: %v != %v",
					test.name, test.encoded, encoded)
				return
			}

			// Ensure the address is for the expected network.
			if !decoded.IsForNet(test.net) {
				t.Errorf("%v: calculated network does not match expected",
					test.name)
				return
			}
		}

	}
}

func TestBTCAddresses(t *testing.T) {
	// Public key (uncompressed) starts with '04' and is 65 bytes long
	pubkeyHex := "04ae31c31bf91278d99b8377a35bbce5b27d9fff15456839e919453fc7b3f721f0ba403ff96c9deeb680e5fd341c0fc3a7b90da4631ee39560639db462e9cb850f"
	pubkey, err := hex.DecodeString(pubkeyHex)
	if err != nil {
		log.Fatal(err)
	}

	// Convert public key to Bitcoin address
	address := pubkeyToAddress(pubkey)
	fmt.Println("Bitcoin Address:", address)

	// 15yN7NPEpu82sHhB6TzCW5z5aXoamiKeGy
}

// 将公钥转换为比特币地址
func pubkeyToAddress(pubkey []byte) string {
	// Step 1: Perform SHA-256 and then RIPEMD-160 to get public key hash (PKH)
	sha256Hash := sha256.New()
	sha256Hash.Write(pubkey)
	sha256Result := sha256Hash.Sum(nil)

	ripemd160Hash := ripemd160.New()
	ripemd160Hash.Write(sha256Result)
	pubKeyHash := ripemd160Hash.Sum(nil)

	// Step 2: Add version byte (0x00 for mainnet P2PKH address)
	versionedPayload := append([]byte{0x00}, pubKeyHash...)

	// Step 3: Calculate the checksum (first 4 bytes of double SHA-256 of versioned payload)
	sha256Hash.Reset()
	sha256Hash.Write(versionedPayload)
	sha256First := sha256Hash.Sum(nil)

	sha256Hash.Reset()
	sha256Hash.Write(sha256First)
	checksum := sha256Hash.Sum(nil)[:4]

	// Step 4: Add checksum to the end of the versioned payload
	addressBytes := append(versionedPayload, checksum...)

	// Step 5: Encode in Base58
	address := base58.Encode(addressBytes)

	return address
}
