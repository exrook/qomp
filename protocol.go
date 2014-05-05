package qomp


type Packet struct {
  ID uint8
}

type Version struct {
  Major uint8
  Minor uint8
  Patch uint8
}

var ProtocolV = Version{0,1,0}

type P0x00 struct { // reserved
  *Packet
}

// Initiate Handshake
type P0x01 struct {
  *Packet
  V Version
}

// Accept Handshake
type P0x02 struct {
  *Packet
  V Version
}

// Program Request
type P0x03 struct {
  *Packet
}

// Program Response
type P0x04 struct {
  *Packet
}

// Benchmark request
type P0x05 struct {
  *Packet
}

// Benchmark Response
type P0x06 struct {
  *Packet
}

// Benchmark Data
type P0x07 struct {
  *Packet
}

// Reserved
type P0x08 struct { // reserved
  *Packet
}

// Work Request
type P0x09 struct {
  *Packet
}

// Work Response
type P0x0A struct {
  *Packet
}

// Work Results
type P0x0B struct {
  *Packet
}

// No Work Remaining
type P0x0C struct {
  *Packet
}

// Work Rejected
type P0x0D struct {
  *Packet
}
