variable "aws_creds_secret" {
  type        = string
  description = "Name of the aws creds secret to use"
}

variable "volume_host_path" {
  type        = string
  description = "which path to mount into the container"
  default     = null
}

variable "volume_container_path" {
  type        = string
  description = "which path to use for volume inside the container"
  default     = null
}

variable "image" {
  type        = string
  description = "Image URL of the image including the tag"
}

variable "replicas" {
  type        = number
  description = "Amount of replicas to spawn"
  default     = 1
}

variable "env_vars" {
  type        = map(string)
  description = "additional env vars to set"
  default     = {}
}

variable "ports" {
  type        = map(string)
  description = "additional ports to publish"
  default     = {}
}