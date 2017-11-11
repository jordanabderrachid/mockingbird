service "greeter" {
  host = "greeter.service"

  endpoint "greet" {
    method = "GET"
    path = "/hello"
  }
}

service "barista" {
  host = "barista.service"

  endpoint "order-beverage" {
    method = "POST"
    path = "/order"
  }
}
