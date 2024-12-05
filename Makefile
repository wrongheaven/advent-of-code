run:
	@day=$(day) && if [ -z "$$day" ]; then echo "Usage: make run day=X"; exit 1; fi
	@padded_day=$$(printf "%02d" $$day) && go run ./internal/day$${padded_day}/main.go