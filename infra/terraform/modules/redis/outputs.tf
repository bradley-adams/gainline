output "host" {
  description = "Redis instance host IP"
  value       = google_redis_instance.redis.host
}

output "port" {
  description = "Redis instance port"
  value       = google_redis_instance.redis.port
}