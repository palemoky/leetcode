.PHONY: new help

# Default target
help:
	@echo "LeetCode Solution Generator"
	@echo ""
	@echo "Usage:"
	@echo "  make new PROBLEM=<problem_number>"
	@echo ""
	@echo "Example:"
	@echo "  make new PROBLEM=1"
	@echo "  make new PROBLEM=42"
	@echo ""
	@echo "This will:"
	@echo "  1. Fetch problem metadata from LeetCode API"
	@echo "  2. Create directory: go/solutions/<number>_<title_slug>"
	@echo "  3. Generate solution file: <title_slug>.go"
	@echo "  4. Generate test file: <title_slug>_test.go"

# Create new LeetCode solution structure
new:
	@if [ -z "$(PROBLEM)" ]; then \
		read -p "Please input LeetCode problem number: " problem_num; \
		python3 scripts/create_solution.py $$problem_num; \
	else \
		echo "Creating solution structure for problem $(PROBLEM)..."; \
		python3 scripts/create_solution.py $(PROBLEM); \
	fi
