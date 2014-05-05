package qomp

type Version struct {
  Major uint8
  Minor uint8
  Patch uint8
}

var ProtocolV = Version{0,1,0}
