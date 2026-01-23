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

  // Icon: material/arrow-expand-horizontal (simplified SVG for standard material icon)
  button.innerHTML = '<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24"><path d="M9 7H7v10h2v-10m6 0h2v10h-2M11 2h2v4h-2V2m0 16h2v4h-2v-4M2 13h2v7l-4-3.5L4 13v-2H2v2m20 0h-2v-2h2v7l-4-3.5 4-3.5" fill="currentColor"/></path><path d="M10 20h4V4h-4v16z" fill="currentColor"/></svg>';
  // Let's use a standard expand icon instead
  button.innerHTML = '<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24"><path d="M21 11H3c-.6 0-1 .4-1 1s.4 1 1 1h18c.6 0 1-.4 1-1s-.4-1-1-1zM3 7h18c.6 0 1-.4 1-1s-.4-1-1-1H3c-.6 0-1 .4-1 1s.4 1 1 1zM3 17h18c.6 0 1-.4 1-1s-.4-1-1-1H3c-.6 0-1 .4-1 1s.4 1 1 1z" fill="currentColor"/></svg>';
  // Better icon: arrow-expand-horizontal
  button.innerHTML = '<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24"><path d="M9 11H7v2h2v-2m4 0h-2v2h2v-2m4 0h-2v2h2v-2m2-7h-2.58l.98 1.05-1.05 1.05-1.42 1.41-1.38-1.41L12.59 5H11.41L10 3.59 8.62 5 7.21 6.41 6.16 5.36 7.14 4.38 8 3.5H5.41l1.17 1.17-1.41 1.41L4.09 4.91 2.67 6.33 1.26 4.91l1.41-1.41L3.84 4.67 5 3.5 2 3.5v17h17v-3l-1.07-1.07 1-1 1.41 1.41 1.07 1.07 1.28-1.28 1.42-1.42 1.07 1.07L22 17.5V11z" fill="currentColor"></path><path d="M21 11H3c-1.1 0-2 .9-2 2v6c0 1.1.9 2 2 2h18c1.1 0 2-.9 2-2v-6c0-1.1-.9-2-2-2zm0 8H3v-6h18v6z" fill="currentColor"/></svg>';
  // Simple "standard" icon for wide screen (arrow left-right)
  button.innerHTML = '<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24"><path d="M5 12h14v-2H5v2zm-2 4h18v-2H3v2zm2 4h14v-2H5v2z" fill="currentColor"/></svg>';
  // Let's use the material icon for aspect ratio or fullscreen
  button.innerHTML = '<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24"><path d="M4 6h16v12H4V6m18-2H2v16h20V4z" fill="currentColor"/></svg>';

  // Use a simpler SVG indicating "Previous/Next" sidebar toggle or just "Wide"
  // Icon: view-week (material) - similar to columns
  // But strictly we want "Expand"
  // Let's use < > arrows pointing out
  button.innerHTML = '<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24"><path d="M4 11h5V5H4v6m0 7h5v-6H4v6m6 0h5v-6h-5v6m6 0h5v-6h-5v6m-6-7h5V5h-5v6m6-6v6h5V5h-5z" fill="currentColor"/></svg>';

  // Let's use a generic icon that looks like a screen
  button.innerHTML = '<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24"><path d="M21 4H3c-1.1 0-2 .9-2 2v12c0 1.1.9 2 2 2h18c1.1 0 2-.9 2-2V6c0-1.1-.9-2-2-2zm0 14H3V6h18v12z" fill="currentColor"/></svg>';


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
