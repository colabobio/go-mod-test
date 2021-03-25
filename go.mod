module main

go 1.15

require (
	github.com/google/uuid v1.2.0 // indirect
	internal/test v1.0.0
)

replace internal/test => ./internal/test
