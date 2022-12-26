package broadcastutils

/*
Represents a decibel

Ex: 3.2 dB
*/
type Decibel float32

/*
Converts a Decibel to an int with 1 decimal precision

Ex: 3.2 dB -> 32
*/
func DecibelToInt(dB Decibel) int {
	return int(dB * 10)
}

/*
Converts an int to a Decibel with 1 decimal precision

Ex: 32 -> 3.2 dB
*/
func IntToDecibel(dB int) Decibel {
	return Decibel(dB / 10)
}
