if rand.Float64() < 0.01 {
	event.CPU = 90 + rand.Float64()*10
	event.Status = "warning"
}

if rand.Float64() < 0.001 {
	event.Status = "offline"
	event.CPU = 0
	event.Memory = 0
	event.Latency = 0
}