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
  // Use label and title, letting MkDocs Material's JS handle the tooltip conversion
  const button = document.createElement('label');
  button.className = 'md-header__button md-icon width-toggle-btn';
  button.style.cursor = 'pointer';
  button.title = 'Switch to wide-screen mode'; // Theme JS will convert this to a tooltip

  // Icon container
  const iconContainer = document.createElement('span');
  iconContainer.style.display = 'inline-block';
  iconContainer.style.verticalAlign = 'middle';

  const icon = document.createElement('img');
  icon.src = '/images/wide-screen-icon.svg';
  icon.width = 24;
  icon.height = 24;
  icon.alt = 'Switch to wide-screen mode';

  iconContainer.appendChild(icon);

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
