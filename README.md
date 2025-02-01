
# GoLogger 

**GoLogger** is an HTTP server which listens and logs all incoming requests to stdout. 

GoLogger was designed as a troubleshooting tool easy to deploy in Kubernetes environments. It aims to help at debugging microservices communication scenarios where provided logs are not enough.

## Configuration
GoLogger listens on `0.0.0.0:8080` by default. This can be changed by passing `GOLOGGER_LISTEN_ADDR` and `GOLOGGER_LISTEN_PORT` environment variables to the
 
## Provided logs
For each request the following information is written to stdout:
- Arrival timestamp
- HTTP Method 
- URL
- Protocol version
- Headers 

