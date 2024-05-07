## State-Based CRDT

### Overview

This project implements a state-based Conflict-Free Replicated Data Type (CRDT) using Go and Protocol Buffers for the backend, and React.js for the frontend interface. The backend is deployed across a distributed cluster of nodes, while the frontend runs on a local machine.
Architecture
The project consists of two major components:
1. Central Node - Keeps track of all nodes and port numbers for documents
Acts as a registry for collaborative nodes
2. Collaborative Node - Collaborates with other peers to simultaneously edit documents
Implemented using Last-Write-Wins (LWW) Register
###  LWW Register
The LWW Register is a core component of the Collaborative Node, implemented as follows:
```
class LWWRegister<T> {
  readonly id: string;
  state: {
    [peer: string]: {
      timestamp: number;
      value: T;
    };
  };
}
```
### Demo
A [demo] (https://drive.google.com/file/d/1KuCn3-mpacxnuBlkWBf-a_HYmugaVws4/view?usp=sharing) of the project is available, showcasing the real-time collaboration capabilities of the CRDT implementation.

### High-Level Design


### References
[An Interactive Intro to CRDTs by Jake lazaroff] (https://jakelazaroff.com/words/an-interactive-intro-to-crdts/)


### Getting Started
Clone the repository and run the backend on a distributed cluster of nodes
Run the frontend on a local machine and access it through a web browser
Use the REST APIs and port forwarding to communicate between the backend and frontend

Start central node
```
sh central.sh
```

Start node that owns the document
```
sh collabOwner.sh
```

Start other collaborative nodes
```
sh collabContributer.sh
```
