#!/usr/bin/env python3
"""
LeetCode Solution Generator
Automatically creates Go solution directory and files based on problem number.
"""

import json
import re
import sys
from pathlib import Path

import requests


# Always use project cache
PROJECT_ROOT = Path(__file__).parent.parent
CACHE_FILE = PROJECT_ROOT / ".github" / "leetcode_cache.json"


def load_cache() -> dict:
    """Load problem cache from local file."""
    if not CACHE_FILE.exists():
        return {}

    try:
        with open(CACHE_FILE, "r", encoding="utf-8") as f:
            return json.load(f)
    except (json.JSONDecodeError, IOError):
        return {}


def save_cache(cache: dict) -> None:
    """Save problem cache to local file."""
    try:
        with open(CACHE_FILE, "w", encoding="utf-8") as f:
            json.dump(cache, f, ensure_ascii=False, indent=2)
    except IOError as e:
        print(f"Warning: Failed to save cache: {e}")


def fetch_all_problems() -> dict:
    """Fetch all problems from LeetCode API and build cache."""
    url = "https://leetcode.cn/graphql/"
    list_query = """
    query problemsetQuestionList($categorySlug: String, $limit: Int, $skip: Int, $filters: QuestionListFilterInput) {
        problemsetQuestionList(categorySlug: $categorySlug limit: $limit skip: $skip filters: $filters) {
            questions {
                difficulty
                frontendQuestionId
                paidOnly
                title
                titleSlug
            }
        }
    }
    """

    headers = {"Content-Type": "application/json"}
    all_problems = {}
    skip = 0
    page_size = 100

    print("Building cache from LeetCode API (this may take a moment)...")

    try:
        while True:
            payload = {
                "query": list_query,
                "variables": {
                    "categorySlug": "algorithms",
                    "skip": skip,
                    "limit": page_size,
                },
                "operationName": "problemsetQuestionList"
            }

            response = requests.post(url, headers=headers, json=payload, timeout=10)
            response.raise_for_status()
            data = response.json()

            questions = data.get("data", {}).get("problemsetQuestionList", {}).get("questions", [])

            if not questions:
                break

            # Add to cache
            for question in questions:
                problem_id = question["frontendQuestionId"]
                all_problems[problem_id] = question

            print(f"  Cached {len(all_problems)} problems...")

            if len(questions) < page_size:
                break

            skip += page_size

        print(f"✓ Cache built successfully: {len(all_problems)} problems")
        return all_problems

    except requests.exceptions.RequestException as e:
        print(f"Error fetching problems: {e}")
        return {}


def fetch_problem_from_api(problem_number: str) -> dict | None:
    """Fetch a single problem from API by calculating its page."""
    url = "https://leetcode.cn/graphql/"
    list_query = """
    query problemsetQuestionList($categorySlug: String, $limit: Int, $skip: Int, $filters: QuestionListFilterInput) {
        problemsetQuestionList(categorySlug: $categorySlug limit: $limit skip: $skip filters: $filters) {
            questions {
                difficulty
                frontendQuestionId
                paidOnly
                title
                titleSlug
            }
        }
    }
    """

    headers = {"Content-Type": "application/json"}

    try:
        # Calculate the page directly based on problem number
        problem_num = int(problem_number)
        page_size = 100
        estimated_page = (problem_num - 1) // page_size
        skip = estimated_page * page_size

        payload = {
            "query": list_query,
            "variables": {
                "categorySlug": "algorithms",
                "skip": skip,
                "limit": page_size,
            },
            "operationName": "problemsetQuestionList"
        }

        response = requests.post(url, headers=headers, json=payload, timeout=10)
        response.raise_for_status()
        data = response.json()

        questions = data.get("data", {}).get("problemsetQuestionList", {}).get("questions", [])

        # Search in the calculated page
        for question in questions:
            if question["frontendQuestionId"] == problem_number:
                return question

        return None

    except requests.exceptions.RequestException as e:
        print(f"Error fetching problem data: {e}")
        return None


def fetch_problem_info(problem_number: str) -> dict | None:
    """Fetch problem information from cache or LeetCode API."""
    # Try cache first
    cache = load_cache()

    if problem_number in cache:
        print(f"✓ Found problem {problem_number} in cache")
        return cache[problem_number]

    # Step 2: Cache miss - try querying the estimated page (fast)
    print(f"Problem {problem_number} not in cache, querying estimated page...")
    problem_info = fetch_problem_from_api(problem_number)

    if problem_info:
        # Found it! Add to cache
        cache[problem_number] = problem_info
        save_cache(cache)
        print(f"✓ Added problem {problem_number} to cache")
        return problem_info

    # Step 3: Still not found - build full cache (slow, but comprehensive)
    print(f"Problem {problem_number} not found in estimated page")
    print("Building full cache to ensure we have all problems...")
    cache = fetch_all_problems()

    if cache:
        save_cache(cache)
        if problem_number in cache:
            print(f"✓ Found problem {problem_number} in full cache")
            return cache[problem_number]

    print(f"Error: Problem {problem_number} not found")
    return None


def sanitize_title_slug(title_slug: str) -> str:
    """Ensure title slug uses underscores and is properly formatted."""
    # Replace hyphens with underscores
    slug = title_slug.replace("-", "_")
    # Remove any non-alphanumeric characters except underscores
    slug = re.sub(r"[^a-z0-9_]", "", slug.lower())
    return slug


def create_solution_file(dir_path: Path, package_name: str, title: str) -> None:
    """Create the main solution Go file."""
    file_path = dir_path / f"{package_name}.go"

    content = f"""package {package_name}

// Solution 1:
// Time: O(), Space: O()
func solution() {{
\t// TODO: Implement solution
}}
"""

    with open(file_path, "w", encoding="utf-8") as f:
        f.write(content)

    print(f"  ✓ Created {file_path.name}")


def create_python_solution_file(dir_path: Path, title: str) -> None:
    """Create the main solution Python file."""
    file_path = dir_path / "solution.py"

    content = f"""# Solution 1:
# Time: O(), Space: O()
def solution():
    # TODO: Implement solution
    pass
"""

    with open(file_path, "w", encoding="utf-8") as f:
        f.write(content)

    print(f"  ✓ Created {file_path.name}")


def create_python_test_file(dir_path: Path) -> None:
    """Create the test Python file."""
    file_path = dir_path / "test_solution.py"

    content = """import pytest
from solution import solution


class TestSolution:
    def test_example_1(self):
        # TODO: Add test cases
        assert solution() is not None
"""

    with open(file_path, "w", encoding="utf-8") as f:
        f.write(content)

    print(f"  ✓ Created {file_path.name}")


def create_test_file(dir_path: Path, package_name: str) -> None:
    """Create the test Go file."""
    file_path = dir_path / f"{package_name}_test.go"

    content = f"""package {package_name}

import (
\t"testing"

\t"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {{
\tt.Parallel()
\ttestCases := []struct {{
\t\tname     string
\t\tinput    any
\t\texpected any
\t}}{{
\t\t{{
\t\t\tname:     "example 1",
\t\t\tinput:    nil,
\t\t\texpected: nil,
\t\t}},
\t}}

\tfuncsToTest := map[string]func([]int) int{{
\t\t"": solution,
\t}}

\tfor fnName, fn := range funcsToTest {{
\t\tt.Run(fnName, func(t *testing.T) {{
\t\t\tfor _, tc := range testCases {{
\t\t\t\tt.Run(tc.name, func(t *testing.T) {{
\t\t\t\t\tt.Parallel()
\t\t\t\t\tresult := fn(tc.input)
\t\t\t\t\tassert.Equal(t, tc.expected, result)
\t\t\t\t}})
\t\t\t}}
\t\t}})
\t}}
}}
"""

    with open(file_path, "w", encoding="utf-8") as f:
        f.write(content)

    print(f"  ✓ Created {file_path.name}")


def create_solution_structure(problem_number: str, language: str = "go") -> None:
    """Main function to create the complete solution structure."""
    print(f"Fetching problem {problem_number} from LeetCode API...")

    problem_info = fetch_problem_info(problem_number)
    if not problem_info:
        sys.exit(1)

    # Extract problem details
    title = problem_info["title"]
    title_slug = sanitize_title_slug(problem_info["titleSlug"])
    difficulty = problem_info["difficulty"]
    is_paid = problem_info["paidOnly"]

    # Format problem number with leading zeros (4 digits)
    formatted_number = problem_number.zfill(4)

    # Create directory name: 0001_two_sum
    dir_name = f"{formatted_number}_{title_slug}"

    # Get the project root and solutions directory
    project_root = Path(__file__).parent.parent

    # Determine language directory
    if language == "py":
        solutions_dir = project_root / "python" / "solutions"
    else:  # default to go
        solutions_dir = project_root / "go" / "solutions"

    target_dir = solutions_dir / dir_name

    # Check if directory already exists
    if target_dir.exists():
        print(f"Error: Directory {dir_name} already exists")
        sys.exit(1)

    # Display problem information
    print(f"\nProblem Information:")
    print(f"  Number: {formatted_number}")
    print(f"  Title: {title}")
    print(f"  Slug: {title_slug}")
    print(f"  Difficulty: {difficulty}")
    print(f"  Paid Only: {'Yes' if is_paid else 'No'}")
    print(f"\nCreating directory: {dir_name}")

    # Create directory
    target_dir.mkdir(parents=True, exist_ok=True)
    print(f"  ✓ Created directory")

    # Create solution and test files based on language
    if language == "py":
        create_python_solution_file(target_dir, title)
        create_python_test_file(target_dir)
    else:  # go
        create_solution_file(target_dir, title_slug, title)
        create_test_file(target_dir, title_slug)

    print(f"\n✅ Successfully created solution structure for problem {problem_number}")
    print(f"📁 Location: {target_dir.relative_to(project_root)}")
    print(f"\nNext steps:")

    if language == "py":
        print(f"  1. Open solution.py and implement the solution")
        print(f"  2. Update test_solution.py with test cases")
        print(f"  3. Run: cd {target_dir.relative_to(project_root)} && uv run pytest -v")
    else:  # go
        print(f"  1. Open {title_slug}.go and implement the solution")
        print(f"  2. Update {title_slug}_test.go with test cases")
        print(f"  3. Run: cd {target_dir.relative_to(project_root)} && go test -v")


def main():
    """Entry point."""
    import argparse

    parser = argparse.ArgumentParser(description="Create LeetCode solution structure")
    parser.add_argument("--language", choices=["go", "py"], required=True,
                        help="Programming language (go or py)")
    parser.add_argument("problem_number", help="LeetCode problem number")

    args = parser.parse_args()

    # Validate problem number
    if not args.problem_number.isdigit():
        print("Error: Problem number must be a positive integer")
        sys.exit(1)

    create_solution_structure(args.problem_number, args.language)


if __name__ == "__main__":
    main()
