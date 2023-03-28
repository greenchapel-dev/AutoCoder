# AutoCoder
An open source in car head unit

## Getting Started
To start the app run

~~~ 
go run autocoder-api/main.go
~~~

To start the OBD connection send a POST request to
~~~ 
127.0.0.1:8080/start-obd
~~~

To get car data send a GET request to
~~~
127.0.0.1:8080/all-car-data
~~~

### For OBD connection
This project uses a modified version of https://github.com/rzetterberg/elmobd
As the library for talking to the ELM OBD device

Can be tested by running elm emulator
https://github.com/Ircama/ELM327-emulator
~~~
python -m elm -s car -n 35000 -l
~~~
