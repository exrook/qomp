QOMP Protocol
=============
-------------

Types
-----
 * `Version`  - A protocol version
 * `Program`  - Tells the client how to compute work units
 * `WorkUnit` - Describres a chunk of work to be computed by the client
 * `DataUnit` - The result of running a WorkUnit

Packet Structure
----------------
All packets are transmitted as JSON encoded strings over TCP (default port 7248). Any invalid packets will be ignored while malformed packets will result in the connection being dropped

### Common Fields ###
Fields present in every packet

* `ID` - Packet ID, determines the rest of the packet's contents

### Specific Fields ###
Fields that are only sent if they are used by the specific packet type

* `Ver`  - Version,  used by the handshake packets
* `Prog` - Program,  specifies which program/computation to run
* `Rate` - uint32,   used when returning benchmark results, work per minute
* `Work` - WorkUnit, used when sending loads to clients
* `Data` - DataUnit, used when sending data from clients

### Packet Types ###
In general, even IDs are sent by the client, odd are sent by the server

 * [`0x00`](#0x00) - Reserved
 * [`0x01`](#0x01) - Initiate Handshake
 * [`0x02`](#0x02) - Accept Handshake
 * [`0x03`](#0x03) - Program Request
 * [`0x04`](#0x04) - Program Response
 * [`0x05`](#0x05) - Benchmark Request
 * [`0x06`](#0x06) - Benchmark Response
 * [`0x07`](#0x07) - Benchmark Data
 * [`0x08`](#0x08) - Reserved
 * [`0x09`](#0x09) - Work Request
 * [`0x0A`](#0x0A) - Work Response
 * [`0x0B`](#0x0B) - Work Result
 * [`0x0C`](#0x0C) - Job Complete
 * [`0x0D`](#0x0D) - Work Unit Rejected

### Packet Overviews ###
#### <a name="0x00"></a>0x00 ####
Reserved for future use
#### <a name="0x01"></a>0x01 ####
 * Ver Version

Initiates a connection with the server
#### <a name="0x02"></a>0x02 ####
 * Ver Version

Indicates the server is ready and finishes the handshake
#### <a name="0x03"></a>0x03 ####
Requests the server to provide information regarding the current program/computation
#### <a name="0x04"></a>0x04 ####
 * Prog Program

Responds to the client with information about the current program/computation
#### <a name="0x05"></a>0x05 ####
Indicates that the client would like to begin the benchmarking process so that it can then begin recieving workunits
#### <a name="0x06"></a>0x06 ####
 * Work WorkUnit

Responds to the client with a workunit that is to be used as a benchmark
#### <a name="0x07"></a>0x07 ####
 * Rate uint32

Returns the results of the benchmark to the server, which then begins sending apropriately sized workunits
#### <a name="0x08"></a>0x08 ####
Reserved for future use
#### <a name="0x09"></a>0x09 ####
Sent if the client has been previously benchmarked to indicate it would like to begin recieving packets
#### <a name="0x0A"></a>0x0A ####
 * Work WorkUnit

Sends a workunit to the client
#### <a name="0x0B"></a>0x0B ####
 * Data DataUnit

Sends results of a computation back to the server
#### <a name="0x0C"></a>0x0C ####
Sent to client once no more work units are avalible
#### <a name="0x0D"></a>0x0D ####
Sent when the client disconnects

Connection and Handshake
------------------------
     Client     Server  
      0x01   |   0x02     - Handshake
      0x03   |   0x04     - Program request/response
      0x05   |   0x06     - Benchmark request/response
      0x07   |   0x0A     - Benchmark Data/ First work unit
    While Working:
      0x0B   |   0x0A     - Returns results/ Sends next work unit
    Client Rejects Work/Disconnects:
      0x0D   |   N/A      - Client rejects work, server closes connection
    All work finished:
      0x0B   |   0x0C     - Client sends work, server sends work finished and closes connection
