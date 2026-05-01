.PHONY: new img2webp help go py

# Default target
help:
	@echo "LeetCode Solution Generator & Documentation Tools"
	@echo ""
	@echo "Usage:"
	@echo "  make new go              - Create new Go solution"
	@echo "  make new py              - Create new Python solution"
	@echo "  make img2webp            - Convert PNG/JPG images to WebP in docs/"
	@echo ""
	@echo "Example:"
	@echo "  make new go"
	@echo "  > Please input LeetCode problem number: 1"
	@echo ""
	@echo "  make new py"
	@echo "  > Please input LeetCode problem number: 42"
	@echo ""
	@echo "New solution will:"
	@echo "  1. Fetch problem metadata from LeetCode API"
	@echo "  2. Create directory: <language>/solutions/<number>_<title_slug>"
	@echo "  3. Generate solution and test files"

# Create new LeetCode solution structure
new:
	@LANG=$(filter-out new,$(MAKECMDGOALS)); \
	if [ -z "$$LANG" ]; then \
		echo "Error: Please specify language (go or py)"; \
		echo "Usage: make new go  OR  make new py"; \
		exit 1; \
	fi; \
	if [ "$$LANG" != "go" ] && [ "$$LANG" != "py" ]; then \
		echo "Error: Language must be 'go' or 'py'"; \
		echo "Usage: make new go  OR  make new py"; \
		exit 1; \
	fi; \
	read -p "Please input LeetCode problem number: " problem_num; \
	if command -v uv >/dev/null 2>&1; then \
		cd python && uv run python ../scripts/create_solution.py --language $$LANG $$problem_num; \
	elif command -v python3 >/dev/null 2>&1; then \
		python3 scripts/create_solution.py --language $$LANG $$problem_num; \
	elif command -v python >/dev/null 2>&1; then \
		python scripts/create_solution.py --language $$LANG $$problem_num; \
	else \
		echo "Error: Neither uv nor python is available in PATH"; \
		exit 1; \
	fi

# Prevent make from treating 'go' and 'py' as targets
go py:
	@:

# Convert images to WebP
img2webp:
	@echo "🖼️  Converting all images in docs/ to WebP..."
	@if command -v uv >/dev/null 2>&1; then \
		cd python && uv run python ../scripts/img2webp.py --all; \
	elif command -v python3 >/dev/null 2>&1; then \
		python3 scripts/img2webp.py --all; \
	elif command -v python >/dev/null 2>&1; then \
		python scripts/img2webp.py --all; \
	else \
		echo "Error: Neither uv nor python is available in PATH"; \
		exit 1; \
	fi
