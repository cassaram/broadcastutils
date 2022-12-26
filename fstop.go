package broadcastutils

/*
Represents a camera FStop

Ex: f2.3
*/
type FStop float32

/*
Returns an FStop as an int with 1 decimal precision

Ex: f11.4 -> 114
*/
func FStopToInt(fs FStop) int {
	return int(fs * 10)
}
