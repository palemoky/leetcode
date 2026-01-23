/**
 * Wide Screen Toggle Feature
 * Allows users to toggle between normal and wide screen modes.
 * Inspired by OI-wiki.
 */

document$.subscribe(() => {
  // Check if toggle button already exists to avoid duplicates
  if (document.querySelector('.md-header__option[title="切换宽屏模式"]')) {
    return;
  }

  // Determine initial state from localStorage
  const isWide = localStorage.getItem('wide-screen') === 'true';
  if (isWide) {
    document.body.setAttribute('data-wide-screen', 'true');
  }

  // Create toggle button
  const button = document.createElement('button');
  button.className = 'md-header__option md-icon';
  button.title = '切换宽屏模式';
  button.setAttribute('aria-label', '切换宽屏模式');

  // Icon: material/arrow-expand-horizontal
  button.innerHTML = '<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24"><path d="M9 11H7v2h2v-2m4 0h-2v2h2v-2m4 0h-2v2h2v-2m2-7h-2.58l.98 1.05-1.05 1.05-1.42 1.41-1.38-1.41L12.59 5H11.41L10 3.59 8.62 5 7.21 6.41 6.16 5.36 7.14 4.38 8 3.5H5.41l1.17 1.17-1.41 1.41L4.09 4.91 2.67 6.33 1.26 4.91l1.41-1.41L3.84 4.67 5 3.5 2 3.5v17h17v-3l-1.07-1.07 1-1 1.41 1.41 1.07 1.07 1.28-1.28 1.42-1.42 1.07 1.07L22 17.5V11z" fill="currentColor"></path><path d="M21 11H3c-1.1 0-2 .9-2 2v6c0 1.1.9 2 2 2h18c1.1 0 2-.9 2-2v-6c0-1.1-.9-2-2-2zm0 8H3v-6h18v6z" fill="currentColor"/></svg>';

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
