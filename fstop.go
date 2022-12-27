package broadcastutils

/*
Represents a camera FStop

Ex: f2.3
*/
type FStop float32

/*
Returns a FStop as an int with 2 decimal precision

Ex: f11.4 -> 1140
*/
func FStopToInt(fs FStop) int {
	return int(fs * 100)
}

/*
Converts an int to a FStop with 2 decimal precision

Ex: 1140 -> f11.4
*/
func IntToFStop(fs int) FStop {
	return FStop(fs / 100)
}
