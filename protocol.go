package qomp

type Packet struct {
  ID   uint8
  Ver  Version  `json:,"omitempty"`
  Prog Program  `json:,"omitempty"`
  Rate uint32   `json:,"omitempty"`
  Work WorkUnit `json:,"omitempty"`
  Data DataUnit `json:,"omitempty"`
}

