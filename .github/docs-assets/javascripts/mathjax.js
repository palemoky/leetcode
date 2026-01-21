window.MathJax = {
  loader: {
    load: ['[tex]/mhchem', '[tex]/extpfeil']
  },
  tex: {
    packages: {'[+]': ['mhchem', 'extpfeil']},
    inlineMath: [['$', '$'], ['\\(', '\\)']],
    displayMath: [['$$', '$$'], ['\\[', '\\]']],
    processEscapes: true,
    processEnvironments: true
  }
};

// 初始渲染和页面导航时重新渲染
document$.subscribe(function() {
  if (typeof MathJax !== 'undefined') {
    MathJax.typesetPromise();
  }
});
