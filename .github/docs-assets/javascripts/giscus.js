/**
 * Giscus Comments Integration
 * Lazy-loads Giscus comments when the placeholder enters the viewport
 */
(function () {
  let giscusLoaded = false;

  // Create placeholder element
  function createPlaceholder() {
    var article = document.querySelector("article");
    if (!article) return;

    var placeholder = document.createElement("div");
    placeholder.id = "giscus-placeholder";
    placeholder.className = "giscus-placeholder";
    placeholder.innerHTML = `
      <div style="padding: 2rem; text-align: center; background: var(--md-code-bg-color); border-radius: 0.2rem; margin-top: 2rem;">
        <p style="margin: 0; color: var(--md-default-fg-color--light);">
          💬 Loading comments...
        </p>
      </div>
    `;
    article.appendChild(placeholder);

    return placeholder;
  }

  // Load Giscus script
  function loadGiscus() {
    if (giscusLoaded) return;
    giscusLoaded = true;

    var placeholder = document.getElementById("giscus-placeholder");
    if (!placeholder) {
      placeholder = createPlaceholder();
    }

    // Clear placeholder content
    placeholder.innerHTML = "";
    placeholder.className = "giscus";

    // Create and configure Giscus script element
    var script = document.createElement("script");
    script.src = "https://giscus.app/client.js";
    script.setAttribute("data-repo", "palemoky/leetcode");
    script.setAttribute("data-repo-id", "R_kgDOJ7ukCg");
    script.setAttribute("data-category", "General");
    script.setAttribute("data-category-id", "DIC_kwDOJ7ukCs4C0vdC");
    script.setAttribute("data-mapping", "pathname");
    script.setAttribute("data-strict", "0");
    script.setAttribute("data-reactions-enabled", "1");
    script.setAttribute("data-emit-metadata", "0");
    script.setAttribute("data-input-position", "top");
    script.setAttribute("data-theme", "preferred_color_scheme");
    script.setAttribute("data-lang", "zh-CN");
    script.setAttribute("data-loading", "lazy");
    script.setAttribute("crossorigin", "anonymous");
    script.async = true;

    placeholder.appendChild(script);
  }

  // Setup lazy loading with IntersectionObserver
  function setupLazyLoad() {
    var placeholder = createPlaceholder();
    if (!placeholder) return;

    var observer = new IntersectionObserver(
      function (entries) {
        entries.forEach(function (entry) {
          if (entry.isIntersecting) {
            loadGiscus();
            observer.disconnect();
          }
        });
      },
      {
        rootMargin: "200px", // Start loading 200px before entering viewport
      }
    );

    observer.observe(placeholder);
  }

  // Handle MkDocs Material page navigation
  document$.subscribe(function () {
    giscusLoaded = false;

    // Remove existing comments container
    var existingGiscus = document.querySelector(".giscus, .giscus-placeholder");
    if (existingGiscus) {
      existingGiscus.remove();
    }

    // Setup lazy loading for new page
    setupLazyLoad();
  });
})();
