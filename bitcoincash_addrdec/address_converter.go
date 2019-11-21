package bitcoincash_addrdec

import (
	"errors"
	"github.com/blocktree/go-owcdrivers/addressEncoder"
	"github.com/blocktree/go-owcrypt"
	"strings"
)

func ConvertToLegacy(address string) (string, error) {

	if address == "" {
		return "", errors.New("Invalid address!")
	}

	// p2pkh legacy
	if strings.Index(address, "1") == 0 {
		hash, err := Decode(address, BitcoinAlphabet)
		if err != nil || hash[0] != 0x00 {
			return "", err
		}

		check := owcrypt.Hash(hash[:21], 0, owcrypt.HASh_ALG_DOUBLE_SHA256)[:4]
		for key, value := range check {
			if value != hash[21+key] {
				return "", errors.New("Invalid address!")
			}
		}
		return address, nil
	}

	if strings.Index(address, "bitcoincash:q") == 0 || strings.Index(address, "q") == 0 {
		if strings.Index(address, "q") == 0 {
			address = "bitcoincash:"+address
		}

		hash, err := addressEncoder.AddressDecode(address, addressEncoder.BCH_mainnetAddressCash)
		if err != nil || hash[0] != 0x00 {
			return "", err
		}

		return Default.AddressEncode(hash[1:], BCH_mainnetAddressP2PKH)
	}

	if strings.Index(address, "C") == 0 {
		hash, err := Decode(address, BitcoinAlphabet)
		if err != nil || hash[0] != 0x1C {
			return "", err
		}

		check := owcrypt.Hash(hash[:21], 0, owcrypt.HASh_ALG_DOUBLE_SHA256)[:4]
		for key, value := range check {
			if value != hash[21+key] {
				return "", errors.New("Invalid address!")
			}
		}
		return Default.AddressEncode(hash[1:21], BCH_mainnetAddressP2PKH)
	}

	if strings.Index(address, "3") == 0 {
		hash, err := Decode(address, BitcoinAlphabet)
		if err != nil || hash[0] != 0x05 {
			return "", err
		}

		check := owcrypt.Hash(hash[:21], 0, owcrypt.HASh_ALG_DOUBLE_SHA256)[:4]
		for key, value := range check {
			if value != hash[21+key] {
				return "", errors.New("Invalid address!")
			}
		}
		return address, nil
	}

	if strings.Index(address, "bitcoincash:p") == 0 || strings.Index(address, "p") == 0 {
		if strings.Index(address, "p") == 0 {
			address = "bitcoincash:" + address
		}
		hash, err := addressEncoder.AddressDecode(address, addressEncoder.BCH_mainnetAddressCash)
		if err != nil || hash[0] != 0x08 {
			return "", err
		}

		return Default.AddressEncode(hash[1:], BCH_mainnetAddressP2SH)
	}

	if strings.Index(address, "H") == 0 {
		hash, err := Decode(address, BitcoinAlphabet)
		if err != nil || hash[0] != 0x28 {
			return "", err
		}

		check := owcrypt.Hash(hash[:21], 0, owcrypt.HASh_ALG_DOUBLE_SHA256)[:4]
		for key, value := range check {
			if value != hash[21+key] {
				return "", errors.New("Invalid address!")
			}
		}
		return Default.AddressEncode(hash[1:21], BCH_mainnetAddressP2SH)
	}

	return "", errors.New("Invalid address!")
}