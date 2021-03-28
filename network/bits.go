package network

type Bits uint64

const (
	F0 Bits = 1 << iota
	F1
	F2
	F3
	F4
	F5
	F6
	F7
	F8
	F9
	F10
	F11
	F12
	F13
	F14
	F15
	F16
	F17
	F18
	F19
	F20
	F21
	F22
	F23
	F24
	F25
	F26
	F27
	F28
	F29
	F30
	F31
	F32
	F33
	F34
	F35
	F36
	F37
	F38
	F39
	F40
	F41
	F42
	F43
	F44
	F45
	F46
	F47
	F48
	F49
	F50
	F51
	F52
	F53
	F54
	F55
	F56
	F57
	F58
	F59
	F60
	F61
	F62
	F63
)

func SetBits(b, flag Bits) Bits    { return b | flag }
func ClearBits(b, flag Bits) Bits  { return b &^ flag }
func ToggleBits(b, flag Bits) Bits { return b ^ flag }
func HasBits(b, flag Bits) bool    { return b&flag != 0 }

func LookupBit(pos byte) Bits {
	if pos == 0 {
		return F0
	} else if pos == 1 {
		return F1
	} else if pos == 2 {
		return F2
	} else if pos == 3 {
		return F3
	} else if pos == 4 {
		return F4
	} else if pos == 5 {
		return F5
	} else if pos == 6 {
		return F6
	} else if pos == 7 {
		return F7
	} else if pos == 8 {
		return F8
	} else if pos == 9 {
		return F9
	} else if pos == 10 {
		return F10
	} else if pos == 11 {
		return F11
	} else if pos == 12 {
		return F12
	} else if pos == 13 {
		return F13
	} else if pos == 14 {
		return F14
	} else if pos == 15 {
		return F15
	} else if pos == 16 {
		return F16
	} else if pos == 17 {
		return F17
	} else if pos == 18 {
		return F18
	} else if pos == 19 {
		return F19
	} else if pos == 20 {
		return F20
	} else if pos == 21 {
		return F21
	} else if pos == 22 {
		return F22
	} else if pos == 23 {
		return F23
	} else if pos == 24 {
		return F24
	} else if pos == 25 {
		return F25
	} else if pos == 26 {
		return F26
	} else if pos == 27 {
		return F27
	} else if pos == 28 {
		return F28
	} else if pos == 29 {
		return F29
	} else if pos == 30 {
		return F30
	} else if pos == 31 {
		return F31
	} else if pos == 32 {
		return F32
	} else if pos == 33 {
		return F33
	} else if pos == 34 {
		return F34
	} else if pos == 35 {
		return F35
	} else if pos == 36 {
		return F36
	} else if pos == 37 {
		return F37
	} else if pos == 38 {
		return F38
	} else if pos == 39 {
		return F39
	} else if pos == 40 {
		return F30
	} else if pos == 41 {
		return F41
	} else if pos == 42 {
		return F42
	} else if pos == 43 {
		return F43
	} else if pos == 44 {
		return F44
	} else if pos == 45 {
		return F45
	} else if pos == 46 {
		return F46
	} else if pos == 47 {
		return F47
	} else if pos == 48 {
		return F48
	} else if pos == 49 {
		return F49
	} else if pos == 50 {
		return F50
	} else if pos == 51 {
		return F51
	} else if pos == 52 {
		return F52
	} else if pos == 53 {
		return F53
	} else if pos == 54 {
		return F54
	} else if pos == 55 {
		return F55
	} else if pos == 56 {
		return F56
	} else if pos == 57 {
		return F57
	} else if pos == 58 {
		return F58
	} else if pos == 59 {
		return F59
	} else if pos == 60 {
		return F60
	} else if pos == 61 {
		return F61
	} else if pos == 62 {
		return F62
	}

	return F63
}
