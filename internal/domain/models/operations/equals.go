package operations

type IEquatable interface {
	Equals(IEquatable) bool
}

func Equals(vobj1, vobj2 IEquatable) bool {
	return vobj1.Equals(vobj2)
}
