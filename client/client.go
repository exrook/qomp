package main

import (
  "github.com/exrook/qomp"
  "log"
  "net"
  "encoding/json"
)

func main() {
  raddr, err := net.ResolveTCPAddr("tcp","127.0.0.1:7248")
  c, err := net.DialTCP("tcp", nil, raddr)
  if err != nil {
    log.Fatalln("Error connecting:", err)
  }
  e := json.NewEncoder(c)
  d := json.NewDecoder(c)
  handshake(e,d)
  prog := progGet(e,d)
  bench := benchGet(e,d)
  getFunc(prog)(bench)
  benchSend(e,10)
  for {
    sendData(e,getFunc(prog)(getWork(d)))
  }
}

func handshake(e *json.Encoder, d *json.Decoder) {
  e.Encode(qomp.Packet{ID: 0x01, Ver: qomp.ProtocolV})
  var p qomp.Packet
  d.Decode(&p)
  if p.ID != 0x02 {
    log.Fatalln("Protocol Error")
  }
}

func progGet(e *json.Encoder, d *json.Decoder) qomp.Program {
  e.Encode(qomp.Packet{ID:0x03})
  var p qomp.Packet
  d.Decode(&p)
  if p.ID != 0x04 {
    log.Fatalln("Protocol Error")
  }
  return p.Prog
}

func benchGet(e *json.Encoder, d *json.Decoder) qomp.WorkUnit {
  e.Encode(qomp.Packet{ID:0x05})
  var p qomp.Packet
  d.Decode(&p)
  if p.ID != 0x06 {
    log.Fatalln("Protocol Error")
  }
  return p.Work
}

func getFunc(prog qomp.Program) func(qomp.WorkUnit)(qomp.DataUnit) {
  switch prog.ID {
    default:
      return square
  }
}

func square(w qomp.WorkUnit) qomp.DataUnit {
  out := make([]int32,w.End-w.Start)
  for i:= w.Start;i < w.End;i=i+1 {
    out[i-w.Start] = int32(i*(i/2))
  }
  return qomp.DataUnit{ID:w.ID, Values: out}
}

func benchSend(e *json.Encoder, rate uint32) {
  e.Encode(qomp.Packet{ID: 0x07, Rate: rate})
}

func getWork(d *json.Decoder) qomp.WorkUnit {
  var p qomp.Packet
  d.Decode(&p)
  if p.ID != 0x0A {
    log.Fatalln("Protocol Error")
  }
  return p.Work
}

func sendData(e *json.Encoder, data qomp.DataUnit) {
  e.Encode(qomp.Packet{ID:0x0B, Data: data})
}
