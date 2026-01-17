#!/usr/bin/env python3
"""
Pre-commit hook: Convert PNG/JPG images to WebP and update Markdown references
"""
import os
import subprocess
import sys
from pathlib import Path
import re

def check_cwebp():
    """Check if cwebp is installed"""
    try:
        subprocess.run(['cwebp', '-version'], capture_output=True, check=True)
        return True
    except (subprocess.CalledProcessError, FileNotFoundError):
        return False

def convert_to_webp(image_path):
    """Convert image to WebP format"""
    webp_path = image_path.with_suffix('.webp')

    # Skip if WebP already exists and is newer
    if webp_path.exists() and webp_path.stat().st_mtime > image_path.stat().st_mtime:
        return webp_path

    try:
        subprocess.run([
            'cwebp',
            '-q', '85',  # Quality 85
            str(image_path),
            '-o', str(webp_path)
        ], check=True, capture_output=True)
        print(f"  ✓ Converted: {image_path.name} -> {webp_path.name}")
        return webp_path
    except subprocess.CalledProcessError as e:
        print(f"  ✗ Failed to convert {image_path.name}: {e}", file=sys.stderr)
        return None

def update_markdown_references(md_file, old_image, new_image):
    """Update image references in Markdown file"""
    content = md_file.read_text(encoding='utf-8')
    old_name = old_image.name
    new_name = new_image.name

    # Pattern: ![alt](image.png) or ![alt](./image.png) or <img src="image.png">
    patterns = [
        (rf'!\[([^\]]*)\]\({re.escape(old_name)}\)', rf'![\1]({new_name})'),
        (rf'!\[([^\]]*)\]\(\.\/{re.escape(old_name)}\)', rf'![\1](./{new_name})'),
        (rf'<img\s+src="\.?\/?{re.escape(old_name)}"', f'<img src="{new_name}"'),
    ]

    updated = content
    for pattern, replacement in patterns:
        updated = re.sub(pattern, replacement, updated)

    if updated != content:
        md_file.write_text(updated, encoding='utf-8')
        return True
    return False

def main():
    """Main function"""
    if not check_cwebp():
        print("⚠️  cwebp not found. Install with:")
        print("  macOS: brew install webp")
        print("  Ubuntu: sudo apt-get install webp")
        print("\nSkipping WebP conversion...")
        sys.exit(0)

    # Get staged files
    result = subprocess.run(
        ['git', 'diff', '--cached', '--name-only', '--diff-filter=ACM'],
        capture_output=True,
        text=True,
        check=True
    )

    staged_files = result.stdout.strip().split('\n')
    core_images = [
        Path(f) for f in staged_files
        if f.startswith('core/') and f.lower().endswith(('.png', '.jpg', '.jpeg'))
    ]

    if not core_images:
        sys.exit(0)

    print(f"\n🖼️  Found {len(core_images)} image(s) in core/ directory")

    converted_images = []
    for image_path in core_images:
        if not image_path.exists():
            continue

        webp_path = convert_to_webp(image_path)
        if webp_path:
            converted_images.append((image_path, webp_path))
            # Stage the WebP file
            subprocess.run(['git', 'add', str(webp_path)], check=True)

    if not converted_images:
        sys.exit(0)

    # Update Markdown references
    print("\n📝 Updating Markdown references...")
    core_dir = Path('core')
    md_files = list(core_dir.rglob('*.md'))

    updated_files = []
    for md_file in md_files:
        for old_image, new_image in converted_images:
            if update_markdown_references(md_file, old_image, new_image):
                if md_file not in updated_files:
                    updated_files.append(md_file)
                    print(f"  ✓ Updated: {md_file.relative_to(core_dir)}")

    # Stage updated Markdown files
    for md_file in updated_files:
        subprocess.run(['git', 'add', str(md_file)], check=True)

    # Remove original images
    print("\n🗑️  Removing original images...")
    for old_image, _ in converted_images:
        subprocess.run(['git', 'rm', str(old_image)], check=True)
        print(f"  ✓ Removed: {old_image.name}")

    print(f"\n✅ Converted {len(converted_images)} image(s) to WebP and updated references\n")
    sys.exit(0)

if __name__ == '__main__':
    main()
