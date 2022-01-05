# ami
![Tests](https://github.com/ahollist/ami/actions/workflows/go.yml/badge.svg)

Am I Online?

ami is a simple module to check internet connectivity using the net.Dial function.

Simply ask: `ami.Online()` 

---
## Options

Some light configuration is possible using the Option functions:

```golang
WithNetwork(network string)
```
* To define a network protocol (one of: "tcp", "tcp4", "udp", udp4")
* This is defaulted to "tcp"
  
```golang
WithAddress(address string)
```
* To choose a specific address to attempt to dial the connection
* Hostnames that resolve to an IP should be handled by the underlying net package
* This is defaulted to "8.8.8.8", Google's DNS servers

```golang
WithPort(port string)
```
* If needed, a specific port can be defined 
* This defaults to "53", the default DNS resolver port

--- 
