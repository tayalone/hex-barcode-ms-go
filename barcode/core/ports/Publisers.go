package ports

// BarcodePb is Definition of Behavior Publisher for Domain
type BarcodePb interface {
	PushMessage(string, bool, uint, string) error
}
