service "greeter" {
  host = "greeter.service"

  endpoint {
    method = "GET"
    path = "/hello"
  }
}
