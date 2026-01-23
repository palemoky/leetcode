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

  // Create toggle button using Material for MkDocs pattern
  const button = document.createElement('label');
  button.className = 'md-header__button md-icon';
  button.setAttribute('for', '');
  button.setAttribute('data-md-color-media', '');

  // Create tooltip span (shows immediately like theme toggle)
  const tooltip = document.createElement('span');
  tooltip.className = 'md-header__button__tooltip md-tooltip md-tooltip--grow';
  tooltip.textContent = 'Switch to wide-screen mode';

  // Create icon container
  const iconContainer = document.createElement('span');
  iconContainer.innerHTML = '<img src="/images/wide-screen-icon.svg" width="24" height="24" alt="" style="vertical-align: middle;" />';

  button.appendChild(tooltip);
  button.appendChild(iconContainer);

  // Add a specific class for precise styling
  button.classList.add('width-toggle-btn');

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
