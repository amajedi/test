# HTTP Byte-Range Proxy Server in GO

# Questions
* Should / be the target path?
* Which URL verbs need to be supported? (GET/POST/HEAD/...)
* Which HTTP Headers need to be supported? (Range, ...)
* Send 200-OK or 206-Partial-Content responses?

# TODO
* Add caching 
* Add logging 
* Add tests (unit, functional)
* Document code in godoc
* Web framework
* Concurrency
* URL Signing
* Containerize (docker)
* Documentation (Expand README, ...)