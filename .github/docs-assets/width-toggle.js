// 内容宽度切换功能
(function() {
  var isWide = localStorage.getItem('content-wide') === 'true';

  // 初始化宽度状态
  if (isWide) {
    document.body.setAttribute('data-md-wide', '');
  }

  // 创建切换按钮
  function createToggleButton() {
    var button = document.createElement('button');
    button.className = 'md-width-toggle md-icon';
    button.title = '切换内容宽度';
    button.setAttribute('aria-label', '切换内容宽度');
    button.type = 'button';

    // SVG 图标
    button.innerHTML = '<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24"><path d="M9 3v18H7V3h2m8 0v18h-2V3h2M3 5v14h2V5H3m16 0v14h2V5h-2z"/></svg>';

    // 点击事件
    button.addEventListener('click', function(e) {
      e.preventDefault();
      e.stopPropagation();

      var body = document.body;
      if (body.hasAttribute('data-md-wide')) {
        body.removeAttribute('data-md-wide');
        localStorage.setItem('content-wide', 'false');
      } else {
        body.setAttribute('data-md-wide', '');
        localStorage.setItem('content-wide', 'true');
      }
    });

    return button;
  }

  // 插入按钮到 header
  function insertButton() {
    // 检查是否已经插入
    if (document.querySelector('.md-width-toggle')) {
      return;
    }

    // 查找插入位置
    var header = document.querySelector('.md-header__inner');
    var source = document.querySelector('.md-header__source');

    if (header) {
      var button = createToggleButton();
      if (source) {
        // 插入到源码链接之前
        header.insertBefore(button, source);
      } else {
        // 插入到 header 末尾
        header.appendChild(button);
      }
    }
  }

  // 页面加载时插入
  if (document.readyState === 'loading') {
    document.addEventListener('DOMContentLoaded', insertButton);
  } else {
    insertButton();
  }

  // Material for MkDocs 的即时加载支持
  if (typeof document$ !== 'undefined') {
    document$.subscribe(insertButton);
  }
})();
