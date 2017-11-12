service "greeter" {
  host = "greeter.service"

  endpoint "greet" {
    method = "GET"
    path = "/hello"

    behavior "ok" {
      request {
        content-type = "application/json"
      }

      response {
        code = 200
      }
    }
  }
}

service "barista" {
  host = "barista.service"

  endpoint "order-beverage" {
    method = "POST"
    path = "/order"

    behavior "not-enough-coffee" {
      request {
        content-type = "application/json"
      }

      response {
        code = 400
      }
    }
  }
}
