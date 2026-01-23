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

  // Icon: material/arrow-expand-horizontal (simplified and clearer)
  // Two arrows pointing outwards, indicating expansion
  button.innerHTML = '<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24"><path d="M4 11h5V5H4v6m0 7h5v-6H4v6m6 0h5v-6h-5v6m6 0h5v-6h-5v6m-6-7h5V5h-5v6m6-6v6h5V5h-5z" fill="currentColor"/></svg>';
  // Re-selecting a simpler "arrows out" icon for clarity
  button.innerHTML = '<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24"><path d="M21 11H3c-.6 0-1 .4-1 1s.4 1 1 1h18c.6 0 1-.4 1-1s-.4-1-1-1zM3 7h18c.6 0 1-.4 1-1s-.4-1-1-1H3c-.6 0-1 .4-1 1s.4 1 1 1zM3 17h18c.6 0 1-.4 1-1s-.4-1-1-1H3c-.6 0-1 .4-1 1s.4 1 1 1z" fill="currentColor"/></svg>';
  // Use the exact "unfold-more" or "open-in-full" style from Material Design
  button.innerHTML = '<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24"><path d="M3 5v4h2V7h4V5H3m18 0h-6v2h4v2h2V5m0 14h-4v2h6v-6h-2v4M5 19h4v2H3v-6h2v4z" fill="currentColor"/></svg>';

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
