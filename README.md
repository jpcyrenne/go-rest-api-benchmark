# go-rest-api-benchmark

I decided to see how GO would compare to this post : https://medium.com/@mihaigeorge.c/web-rest-api-benchmark-on-a-real-life-application-ebb743a5d7a3

I also wanted to learn more about Go by doing things in different ways. This first example of code is to test Fobonacci as (1) a regular loop, (2) as a recursive fucntion and (3) as a go routine (try to profit from multi processors).

Seems to give wrong numbes around the 95th number.  Not sure if it's a uint64 issue (bigger than unsigned 64 bit number)?  I will ty to undersand.

## How To
You need to have Golng insalled on your machine
1) Download the code and cd into go-rest-api-benchmark
2) Run : go run main.go
3) Open your browser and call the doffeent routes.
ex :
http://localhost:8000/
http://localhost:8000/fibonacci/20
http://localhost:8000/fibonacci2/20
http://localhost:8000/fibonacci3/20

You'll get:
- the number (20th here) in your browser
- and the time it took in your shell (µs = microseconds).
