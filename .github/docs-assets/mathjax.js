window.MathJax = {
  tex: {
    inlineMath: [['$', '$'], ['\\(', '\\)']],
    displayMath: [['$$', '$$'], ['\\[', '\\]']],
    processEscapes: true,
    processEnvironments: true,
    packages: {'[+]': ['mhchem', 'extpfeil']}
  },
  loader: {
    load: ['[tex]/mhchem', '[tex]/extpfeil']
  },
  startup: {
    pageReady: () => {
      return MathJax.startup.defaultPageReady();
    }
  }
};

document$.subscribe(() => {
  if (window.MathJax && window.MathJax.typesetPromise) {
    MathJax.texReset();
    MathJax.typesetPromise();
  }
})
