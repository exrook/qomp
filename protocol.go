package qomp

type Packet struct {
  ID uint8
  Ver Version
  Prog Program
  Rate uint32
  Work WorkUnit
  Data DataUnit
}

