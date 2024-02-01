# loadbalancer-go

Simple load balancer in go

## To run
1) run in terminal /src
    ```
    go run main.go
    ```
    
    output:
    ```
    Load Balancer listening on port: 8000
    ```

2) open localhost:8000 in browser
    
    output:
    ```
    Proxying request to: http://www.facebook.com
    Proxying request to: http://www.google.com
    Proxying request to: http://www.linkedin.com
    Proxying request to: http://www.facebook.com
    Proxying request to: http://www.google.com
    Proxying request to: http://www.linkedin.com
    Proxying request to: http://www.facebook.com
    Proxying request to: http://www.google.com
    Proxying request to: http://www.linkedin.com
    Proxying request to: http://www.facebook.com
    ```