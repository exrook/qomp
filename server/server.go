package main

import (
  "github.com/exrook/qomp"
  "net"
  "fmt"
  "encoding/json"
)

func main() {
  addr := net.TCPAddr{Port:7248}
  l, err := net.ListenTCP("tcp", &addr)
  defer l.Close()
  if err != nil {
    fmt.Println("Error listening on port 7248:", err)
  }
  for {
    c, err := l.AcceptTCP()
    if err != nil {
      fmt.Println("Error while accepting connection:", err)
      continue
    }
    go cHandle(c)
  }
}

func cHandle(c *net.TCPConn) {
  defer c.Close()
  d := json.NewDecoder(c)
  e := json.NewEncoder(c)
  var p qomp.Packet
  var wRate uint32
  if err := d.Decode(&p); err != nil {
    fmt.Println("Error decoding packet:", err)
    return    
  }
  if p.ID == 0x01 {
    if err := e.Encode(qomp.Packet{ID:0x02,Ver:qomp.ProtocolV}); err != nil {
      fmt.Println("Error decoding packet:", err)
      return
    }
    fmt.Println("Sucessful Connection")
  } else {
    return
  }
  for {
    var p qomp.Packet
    var err error
    if err := d.Decode(&p); err != nil {
      fmt.Println("Error decoding packet:", err)
      return
    }
    switch p.ID {
      default:
        continue
      case 0x03:
        err = e.Encode(qomp.Packet{ID:0x04, Prog:qomp.Program{0}})
      case 0x05:
        err = e.Encode(qomp.Packet{ID:0x06, Work:qomp.WorkUnit{0,map[string]interface{}{"Start": 1,"End": 2000}}})
      case 0x07:
        if p.Rate > 0 {
          wRate = p.Rate
        }
        fmt.Println(wRate)
        err = e.Encode(qomp.Packet{ID:0x0A, Work:qomp.WorkUnit{0,map[string]interface{}{"Start": 1,"End": wRate}}})
      case 0x0B:
        err = e.Encode(qomp.Packet{ID:0x0A, Work:qomp.WorkUnit{0,map[string]interface{}{"Start": 1,"End": 2000}}})
      case 0x0D:
        return
    }
    if err != nil {
      fmt.Println("Error encoding packet:", err)
      return
    }
  }
}
