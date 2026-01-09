// Giscus 评论系统初始化
document$.subscribe(function() {
  // 检查是否已经加载过 Giscus
  var existingGiscus = document.querySelector('.giscus');
  if (existingGiscus) {
    existingGiscus.remove();
  }

  // 在文章底部添加评论容器
  var article = document.querySelector('article');
  if (article) {
    var commentsDiv = document.createElement('div');
    commentsDiv.className = 'giscus';
    article.appendChild(commentsDiv);

    // 加载 Giscus 脚本
    var script = document.createElement('script');
    script.src = 'https://giscus.app/client.js';
    script.setAttribute('data-repo', 'palemoky/leetcode');
    script.setAttribute('data-repo-id', 'R_kgDOJ7ukCg');
    script.setAttribute('data-category', 'General');
    script.setAttribute('data-category-id', 'DIC_kwDOJ7ukCs4C0vdC');
    script.setAttribute('data-mapping', 'pathname');
    script.setAttribute('data-strict', '0');
    script.setAttribute('data-reactions-enabled', '1');
    script.setAttribute('data-emit-metadata', '0');
    script.setAttribute('data-input-position', 'top');
    script.setAttribute('data-theme', 'preferred_color_scheme');
    script.setAttribute('data-lang', 'zh-CN');
    script.setAttribute('data-loading', 'lazy');
    script.setAttribute('crossorigin', 'anonymous');
    script.async = true;

    commentsDiv.appendChild(script);
  }
});
