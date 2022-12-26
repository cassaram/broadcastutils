package broadcastutils

/*
Represents a camera FStop

Ex: f2.3
*/
type FStop float32

/*
Returns a FStop as an int with 1 decimal precision

Ex: f11.4 -> 114
*/
func FStopToInt(fs FStop) int {
	return int(fs * 10)
}

/*
Converts an int to a FStop with 1 decimal precision

Ex: 114 -> f11.4
*/
func IntToFStop(fs int) FStop {
	return FStop(fs / 10)
}
