package bitcoincash_addrdec

import "testing"

func TestConvertToLegacy(t *testing.T) {
	p2pkhaddr_legacy := "1ArAfjvwXBiYhWcirLRc9CsvhuSgVqRQbc"
	p2pkhaddr_cashaddr := "bitcoincash:qpkqt9dxyae8c4r32m99lm24aw09vctc8ulss5czme"
	p2pkhaddr_cashaddr_short := "qpkqt9dxyae8c4r32m99lm24aw09vctc8ulss5czme"
	p2pkhaddr_copay := "CSK4EnH1QEh5beX9Y5kXiiVxL2f6Knung5"

	chk, err := ConvertToLegacy(p2pkhaddr_legacy)
	if err != nil || chk != p2pkhaddr_legacy{
		t.Error(err)
		return
	}

	chk, err = ConvertToLegacy(p2pkhaddr_cashaddr)
	if err != nil || chk != p2pkhaddr_legacy{
		t.Error(err)
		return
	}

	chk, err = ConvertToLegacy(p2pkhaddr_cashaddr_short)
	if err != nil || chk != p2pkhaddr_legacy{
		t.Error(err)
		return
	}

	chk, err = ConvertToLegacy(p2pkhaddr_copay)
	if err != nil || chk != p2pkhaddr_legacy{
		t.Error(err)
		return
	}

	p2shaddr_legacy := "3CXRBKkvzB1QPJz2AVJYavEBr228v6Hkqe"
	p2shaddr_cashaddr := "bitcoincash:ppmd3gkmjxedpexxxtvvx4vydasd8qceus83xaa7sn"
	p2shaddr_cashaddr_short := "ppmd3gkmjxedpexxxtvvx4vydasd8qceus83xaa7sn"
	p2shcopay := "HHMXe8C1qVE51Us42AxhZJkisg39khHwTy"

	chk, err = ConvertToLegacy(p2shaddr_legacy)
	if err != nil || chk != p2shaddr_legacy{
		t.Error(err)
		return
	}

	chk, err = ConvertToLegacy(p2shaddr_cashaddr)
	if err != nil || chk != p2shaddr_legacy{
		t.Error(err)
		return
	}

	chk, err = ConvertToLegacy(p2shaddr_cashaddr_short)
	if err != nil || chk != p2shaddr_legacy{
		t.Error(err)
		return
	}

	chk, err = ConvertToLegacy(p2shcopay)
	if err != nil || chk != p2shaddr_legacy{
		t.Error(err)
		return
	}
}
