/**
 * Wide Screen Toggle Feature
 * Allows users to toggle between normal and wide screen modes.
 * Inspired by OI-wiki.
 */

document$.subscribe(() => {
  // Check if toggle button already exists to avoid duplicates
  if (document.querySelector('.width-toggle-btn')) {
    return;
  }

  // Determine initial state from localStorage
  const isWide = localStorage.getItem('wide-screen') === 'true';
  if (isWide) {
    document.body.setAttribute('data-wide-screen', 'true');
  }

  // Create toggle button
  // Create toggle button
  // We manually construct the tooltipped structure to match the theme's rendered output
  // <button class="md-header__option md-icon">
  //   <span class="md-tooltip">Switch to wide-screen mode</span>
  //   ...icon...
  // </button>
  const button = document.createElement('button');
  button.className = 'md-header__option md-icon width-toggle-btn';
  button.style.cursor = 'pointer';
  // IMPORTANT: Do NOT set .title attribute, otherwise the browser native tooltip appears
  button.setAttribute('aria-label', 'Switch to wide-screen mode');

  // 1. Tooltip Element (theme CSS handles visibility on hover)
  // .md-tooltip--bottom is often default for header, or we rely on default positioning
  const tooltip = document.createElement('span');
  tooltip.className = 'md-tooltip';
  tooltip.textContent = 'Switch to wide-screen mode';
  tooltip.style.fontWeight = '700'; // formatting

  // 2. Icon Container
  const iconContainer = document.createElement('span');
  iconContainer.style.display = 'inline-block';
  iconContainer.style.verticalAlign = 'middle';

  const icon = document.createElement('img');
  icon.src = '/images/wide-screen-icon.svg';
  icon.width = 24;
  icon.height = 24;
  icon.alt = ''; // decorative inside button

  iconContainer.appendChild(icon);

  // Append in correct order: tooltip first (often), or just inside parent
  button.appendChild(tooltip);
  button.appendChild(iconContainer);

  button.addEventListener('click', () => {
    const body = document.body;
    const isWide = body.getAttribute('data-wide-screen') === 'true';

    if (isWide) {
      body.removeAttribute('data-wide-screen');
      localStorage.setItem('wide-screen', 'false');
    } else {
      body.setAttribute('data-wide-screen', 'true');
      localStorage.setItem('wide-screen', 'true');
    }
  });

  // Insert button into header
  const target = document.querySelector('.md-header__option[data-md-component="palette"]');
  if (target) {
    target.parentNode.insertBefore(button, target);
  } else {
    // Fallback if palette is missing
    const search = document.querySelector('.md-header__option[data-md-component="search"]');
    if (search) {
      search.parentNode.insertBefore(button, search);
    }
  }
});
