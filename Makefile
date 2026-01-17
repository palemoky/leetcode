.PHONY: new optimize-images help

# Default target
help:
	@echo "LeetCode Solution Generator & Documentation Tools"
	@echo ""
	@echo "Usage:"
	@echo "  make new PROBLEM=<problem_number>"
	@echo "  make optimize-images"
	@echo ""
	@echo "Commands:"
	@echo "  new              - Create new LeetCode solution structure"
	@echo "  optimize-images  - Convert PNG/JPG images to WebP in core/"
	@echo ""
	@echo "Example:"
	@echo "  make new PROBLEM=1"
	@echo "  make new PROBLEM=42"
	@echo "  make optimize-images"
	@echo ""
	@echo "New solution will:"
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

# Optimize images: Convert PNG/JPG to WebP
optimize-images:
	@echo "🖼️  Converting images to WebP..."
	@python3 scripts/convert-images-to-webp.py
	@echo "✅ Done! Don't forget to commit the changes."
