#!/usr/bin/env python3
"""
Update LeetCode cache file for GitHub Actions.
Fetches all problems from LeetCode API and saves to .github/leetcode_cache.json
"""

import json
import sys
from pathlib import Path

import requests


def fetch_all_problems() -> dict:
    """Fetch all problems from LeetCode API."""
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

    print("Fetching all problems from LeetCode API...")

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

            print(f"  Fetched {len(all_problems)} problems...")

            if len(questions) < page_size:
                break

            skip += page_size

        print(f"✓ Successfully fetched {len(all_problems)} problems")
        return all_problems

    except requests.exceptions.RequestException as e:
        print(f"Error fetching problems: {e}")
        sys.exit(1)


def main():
    """Main function."""
    # Get cache file path
    project_root = Path(__file__).parent.parent
    cache_file = project_root / ".github" / "leetcode_cache.json"

    # Fetch all problems
    problems = fetch_all_problems()

    if not problems:
        print("Error: No problems fetched")
        sys.exit(1)

    # Save to cache file
    print(f"Saving cache to {cache_file}...")
    cache_file.parent.mkdir(parents=True, exist_ok=True)

    with open(cache_file, "w", encoding="utf-8") as f:
        json.dump(problems, f, ensure_ascii=False, indent=2)

    print(f"✓ Cache saved successfully: {len(problems)} problems")


if __name__ == "__main__":
    main()
