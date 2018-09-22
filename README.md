# s7
Golang implementation of Siemens S7 communication protocol

Functions
-------------------
AG:
*   Read/Write Data Block (DB)
*   Read/Write Merkers(MB)
*   Read/Write IPI (EB)
*   Read/Write IPU (AB)
*   Read/Write Timer (TM)
*   Read/Write Counter (CT)
*   Get Block Info

PG:
*   Hot start/Cold start / Stop PLC
*   Get CPU of PLC status
*   List available blocks in PLC
*   Set/Clear password for session
*   Get CPU protection and CPU Order code
*   Get CPU/CP Information
*   Read/Write clock for the PLC

Supported communication
-----------------
*   TCP
*   Serial (PPI, MPI) (under construction)

References
----------
- libnodave http://libnodave.sourceforge.net/
- snap7 http://snap7.sourceforge.net/
- tarm serial library https://github.com/tarm/serial
- Simatic Open TCP/IP Communication via Industrial Ethernet from Siemens(doku)
- SIMATIC NET FDL-Programmierschnittstelle (doku)

Original implementation
----------
[https://github.com/robinson/gos7](https://github.com/robinson/gos7)
