# Adaptive Replacement Cache Project in Golang

___

This is an implementation of the adaptive replacement cache (ARC) algorithm [developed by researchers at IBM](https://ieeexplore.ieee.org/stamp/stamp.jsp?arnumber=1297303&tag=1) written in the Go Language. The ARC algorithm builds on top of one of the most famous caching algorithms, Least Recently Used (LRU), and is able to outperform LRU until the cache reaches a certain size, for which ARC and LRU has equal hit rates.

### [Video Demo Here](https://youtu.be/PH0mPODbvUw)

### How to Run
- `git clone https://github.com/mrapi00/ARC-Cache
- `go test`
