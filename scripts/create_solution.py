#!/usr/bin/env python3
"""
LeetCode Solution Generator
Automatically creates Go solution directory and files based on problem number.
"""

import json
import os
import re
import sys
from pathlib import Path

import requests


def fetch_problem_info(problem_number: str) -> dict | None:
    """Fetch problem information from LeetCode GraphQL API."""
    url = "https://leetcode.cn/graphql/"

    # Query to get problem details by frontend question ID
    query = """
    query questionData($titleSlug: String!) {
        question(titleSlug: $titleSlug) {
            questionId
            questionFrontendId
            title
            titleSlug
            difficulty
            isPaidOnly
        }
    }
    """

    # First, we need to get the titleSlug from the problem list
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

    # Fetch problem list to find the titleSlug
    payload = {
        "query": list_query,
        "variables": {
            "categorySlug": "algorithms",
            "skip": 0,
            "limit": 3500,  # Fetch enough to cover most problems
        },
        "operationName": "problemsetQuestionList"
    }

    headers = {"Content-Type": "application/json"}

    try:
        response = requests.post(url, headers=headers, json=payload, timeout=10)
        response.raise_for_status()
        data = response.json()

        questions = data.get("data", {}).get("problemsetQuestionList", {}).get("questions", [])

        # Find the problem by frontend ID
        for question in questions:
            if question["frontendQuestionId"] == problem_number:
                return question

        print(f"Error: Problem {problem_number} not found")
        return None

    except requests.exceptions.RequestException as e:
        print(f"Error fetching problem data: {e}")
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

// {title}
// 解法一：
// Time: O(), Space: O()
func solution() {{
	// TODO: Implement solution
}}
"""

    with open(file_path, "w", encoding="utf-8") as f:
        f.write(content)

    print(f"  ✓ Created {file_path.name}")


def create_test_file(dir_path: Path, package_name: str) -> None:
    """Create the test Go file."""
    file_path = dir_path / f"{package_name}_test.go"

    content = f"""package {package_name}

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {{
	t.Parallel()
	testCases := []struct {{
		name     string
		input    interface{{}}
		expected interface{{}}
	}}{{
		{{
			name:     "example 1",
			input:    nil,
			expected: nil,
		}},
	}}

	funcsToTest := map[string]func([]int) int{{
		"": solution,
	}}

	for fnName, fn := range funcsToTest {
		t.Run(fnName, func(t *testing.T) {
			for _, tc := range testCases {
				t.Run(tc.name, func(t *testing.T) {
					t.Parallel()
					result := fn(tc.input)
					assert.Equal(t, tc.expected, result)
				})
			}
		})
	}
}}
"""

    with open(file_path, "w", encoding="utf-8") as f:
        f.write(content)

    print(f"  ✓ Created {file_path.name}")


def create_solution_structure(problem_number: str) -> None:
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

    # Create solution and test files
    create_solution_file(target_dir, title_slug, title)
    create_test_file(target_dir, title_slug)

    print(f"\n✅ Successfully created solution structure for problem {problem_number}")
    print(f"📁 Location: {target_dir.relative_to(project_root)}")
    print(f"\nNext steps:")
    print(f"  1. Open {title_slug}.go and implement the solution")
    print(f"  2. Update {title_slug}_test.go with test cases")
    print(f"  3. Run: cd {target_dir.relative_to(project_root)} && go test -v")


def main():
    """Entry point."""
    if len(sys.argv) != 2:
        print("Usage: python3 create_solution.py <problem_number>")
        print("Example: python3 create_solution.py 1")
        sys.exit(1)

    problem_number = sys.argv[1]

    # Validate problem number
    if not problem_number.isdigit():
        print("Error: Problem number must be a positive integer")
        sys.exit(1)

    create_solution_structure(problem_number)


if __name__ == "__main__":
    main()
