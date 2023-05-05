package icon

import "errors"

type Generator struct {
}

// GenerateFlag 生成用フラグ
type GenerateFlag struct {
	// 0. メガネの有無
	Glasses bool
	// 1. 口の形(7種類
	Mouth uint8
	// 2. 頬の種類 (3種類+なし
	Cheek uint8
	// 3. ひげの有無
	Whiskers bool
	// 4. リボン2種類/鈴/ネームタグ/なし
	Collar uint8
	// 5. 耳の形(5種類
	Ear uint8
	// 6. 目の形(6種類
	Eyes uint8
}

type UserIcon uint64

func NewGenerator() *Generator {
	return &Generator{}
}

func (g *Generator) NewIcon(f GenerateFlag) (UserIcon, error) {
	if f.Mouth > 7 {
		return 0, errors.New("invalid flag")
	}
	if f.Cheek > 3 {
		return 0, errors.New("invalid flag")
	}
	if f.Collar > 5 {
		return 0, errors.New("invalid flag")
	}
	if f.Ear > 4 {
		return 0, errors.New("invalid flag")
	}
	if f.Eyes > 5 {
		return 0, errors.New("invalid flag")
	}

	res := uint64(0)
	if f.Glasses {
		res += 1 << 60
	}

	res += uint64(f.Mouth) << 56
	res += uint64(f.Cheek) << 52
	if f.Whiskers {
		res += 1 << 48
	}

	res += uint64(f.Collar) << 44
	res += uint64(f.Ear) << 40
	res += uint64(f.Eyes) << 36

	return UserIcon(res), nil
}
