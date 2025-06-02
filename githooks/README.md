# githooks Directory

This directory contains custom Git hooks that help automate and enhance your development workflow. By using Git hooks, you can automatically perform a variety of actions at different points in the Git process—such as before committing, after merging, or before pushing—to enforce standards, run checks, or trigger custom scripts. This helps prevent common mistakes and ensures consistency across your team.

## How to Use

1. **Copy Hooks to .git/hooks**

   Copy the desired hook scripts (for example, `pre-commit`, `pre-push`, etc.) from this directory to your project's `.git/hooks` directory. For example:
   ```sh
   cp githooks/pre-commit .git/hooks/pre-commit
   ```

2. **Make Hooks Executable**

   Ensure the hook scripts have executable permissions:
   ```sh
   chmod +x .git/hooks/pre-commit
   ```

3. **Automatic Workflow**

   Once set up, the hooks will run automatically at the appropriate stage (e.g., before a commit, before a push, after a merge), depending on the hook type. You can use hooks to automate tasks such as code formatting, running tests, checking commit messages, or any other custom logic your workflow requires.

## Benefits

- **Automate repetitive or critical tasks in your Git workflow**
- **Reduce the risk of common mistakes before code enters the repository**
- **Enforce team standards and improve code quality**
- **Easily extendable: add new hooks for additional checks or automation as needed**

For more information about Git hooks and available hook types, see the [official documentation](https://git-scm.com/docs/githooks).