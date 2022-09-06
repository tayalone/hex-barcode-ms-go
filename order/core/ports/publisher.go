package ports

// ReqBarcodePb is Definition of Behavior Publisher for Domain
type ReqBarcodePb interface {
	Publish(string, bool, uint) error
}
