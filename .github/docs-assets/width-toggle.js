// 内容宽度切换功能
document$.subscribe(function() {
  // 创建切换按钮
  var header = document.querySelector('.md-header__inner');
  if (header && !document.querySelector('.md-width-toggle')) {
    var toggleButton = document.createElement('button');
    toggleButton.className = 'md-width-toggle md-icon';
    toggleButton.title = '切换内容宽度';
    toggleButton.setAttribute('aria-label', '切换内容宽度');

    // 使用 Material Icons
    toggleButton.innerHTML = '<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24"><path d="M9 3v18H7V3h2m8 0v18h-2V3h2M3 5v14h2V5H3m16 0v14h2V5h-2z"/></svg>';

    // 检查本地存储的宽度设置
    var isWide = localStorage.getItem('content-wide') === 'true';
    if (isWide) {
      document.body.setAttribute('data-md-wide', '');
    }

    // 点击事件
    toggleButton.addEventListener('click', function() {
      var body = document.body;
      if (body.hasAttribute('data-md-wide')) {
        body.removeAttribute('data-md-wide');
        localStorage.setItem('content-wide', 'false');
      } else {
        body.setAttribute('data-md-wide', '');
        localStorage.setItem('content-wide', 'true');
      }
    });

    // 添加到 header
    var headerTitle = header.querySelector('.md-header__title');
    if (headerTitle) {
      headerTitle.parentNode.insertBefore(toggleButton, headerTitle.nextSibling);
    }
  }
});
