run:
	@\
	year=$(year) && \
	day=$(day) && \
	if [ -z "$$year" ] || [ -z "$$day" ]; then \
		echo "Usage: make run year=YYYY day=X"; \
		exit 1; \
	fi
	@\
	padded_day=$$(printf "%02d" $$day) && \
	file=./$$year/day$${padded_day}/main.go && \
	if [ ! -f "$$file" ]; then \
		echo "Notice: No solution found for this day. Please try another."; \
		exit 0; \
	fi && \
	go run "$$file"