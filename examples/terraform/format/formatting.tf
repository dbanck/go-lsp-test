resource "aws_lambda_function" "f" {
    environment {
    variables = {
      a = "b"
    }
  }
}
