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
  // We use md-header__button class to inherit theme tooltip behavior
  const button = document.createElement('button');
  button.className = 'md-header__option md-header__button md-icon width-toggle-btn';
  button.setAttribute('aria-label', 'Switch to wide-screen mode');

  // Custom Tooltip (appears immediately)
  const tooltip = document.createElement('span');
  tooltip.className = 'md-tooltip';
  tooltip.textContent = 'Switch to wide-screen mode';
  tooltip.style.fontWeight = '700'; // Make tooltip bold as requested

  // Icon
  const icon = document.createElement('img');
  icon.src = '/images/wide-screen-icon.svg';
  icon.width = 24;
  icon.height = 24;
  icon.alt = 'Switch to wide-screen mode';

  button.appendChild(tooltip);
  button.appendChild(icon);

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
