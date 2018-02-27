output "arn" {
  value = "${aws_lambda_function.function.arn}"
}

output "name" {
  value = "${aws_lambda_function.function.function_name}"
}

output "version" {
  value = "${aws_lambda_function.function.version}"
}
