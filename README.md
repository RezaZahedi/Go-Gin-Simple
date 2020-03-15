# Go-Gin-Simple
This code illustrates the use of Gin framework to setup a simple server that computes and returns
fibonacci numbers. One can make a request to the server like this:

    curl 'localhost:8080/' -F 'number=8'
returning a JSON response of `21`.  
<br>This code has a test coverage of more than 70% in all files.  
<br>The function calls are cached and the cache is made concurrency safe.  
Function memoization code is at: <br> https://github.com/RezaZahedi/Go-Gin/tree/master/model/memo
<br> and for more information visit: <br> https://github.com/RezaZahedi/Go-Gin