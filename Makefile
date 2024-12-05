run:
	@\
	year=$(year) && \
	day=$(day) && \
	if [ -z "$$day" ]; then \
		echo "Usage: make run [year=YYYY] day=X"; \
		exit 0; \
	fi && \
	if [ -z "$$year" ]; then \
		year=$$(ls -d [0-9][0-9][0-9][0-9] | sort -r | head -n 1); \
	fi && \
	padded_day=$$(printf "%02d" $$day) && \
	file=./$$year/day$${padded_day}/main.go && \
	if [ ! -f "$$file" ]; then \
		echo "Notice: No solution found for this day. Please try another."; \
		exit 0; \
	fi && \
	go run "$$file"