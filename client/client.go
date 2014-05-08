package main

import (
  "github.com/exrook/qomp"
  "log"
  "net"
  "encoding/json"
  "flag"
)

func main() {
  b := flag.Int("b",100,"How many numbers to grab at a time")
  rhost := flag.Arg(0)
  if rhost == "" {
    rhost = "127.0.0.1:7248"
  }
  raddr, err := net.ResolveTCPAddr("tcp",rhost)
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
  benchSend(e,uint32(*b))
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
  start, e := w.Data["Start"].(int32)
  end, e := w.Data["End"].(int32)
  if !e {
    start, end = 0,100
  }
  out := make([]int32,end-start)
  for i:= start;i < end;i=i+1 {
    out[i-start] = int32(i*i)
  }
  return qomp.DataUnit{ID:w.ID, Data: map[string]interface{}{"squares": out},}
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
