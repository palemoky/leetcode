// Giscus
(function () {
  let giscusLoaded = false;

  // 创建占位符
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

  // 加载 Giscus
  function loadGiscus() {
    if (giscusLoaded) return;
    giscusLoaded = true;

    var placeholder = document.getElementById("giscus-placeholder");
    if (!placeholder) {
      placeholder = createPlaceholder();
    }

    // 移除占位符内容
    placeholder.innerHTML = "";
    placeholder.className = "giscus";

    // 加载 Giscus 脚本
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

  // 监听滚动事件，当评论区即将进入视口时加载
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
        rootMargin: "200px", // 提前 200px 开始加载
      }
    );

    observer.observe(placeholder);
  }

  // MkDocs Material 的页面切换事件
  document$.subscribe(function () {
    giscusLoaded = false;

    // 移除旧的评论区
    var existingGiscus = document.querySelector(".giscus, .giscus-placeholder");
    if (existingGiscus) {
      existingGiscus.remove();
    }

    // 设置延迟加载
    setupLazyLoad();
  });
})();
