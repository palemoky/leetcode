// 图片懒加载优化
(function() {
  // 为所有图片添加 loading="lazy" 属性
  function addLazyLoading() {
    // 获取所有 article 中的图片
    var images = document.querySelectorAll('article img:not([loading])');

    images.forEach(function(img) {
      // 跳过 logo 和小图标
      if (img.width < 50 || img.height < 50) {
        return;
      }

      // 添加 lazy loading
      img.setAttribute('loading', 'lazy');
    });
  }

  // MkDocs Material 页面加载事件
  document$.subscribe(function() {
    // 延迟执行，避免阻塞渲染
    setTimeout(addLazyLoading, 100);
  });
})();
