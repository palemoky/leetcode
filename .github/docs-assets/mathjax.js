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
  },
  startup: {
    typeset: false
  }
};

document$.subscribe(() => {
  if (window.MathJax && window.MathJax.typesetPromise) {
    if (MathJax.typesetClear) {
      MathJax.typesetClear();
    }

    if (MathJax.texReset) {
      MathJax.texReset();
    }

    MathJax.typesetPromise().catch(function (err) {
      console.log('MathJax rendering error:', err);
    });
  }
});
