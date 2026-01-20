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
  if (window.MathJax && MathJax.startup) {
    MathJax.startup.promise.then(() => {
      MathJax.typesetPromise().catch((err) => {
        console.log('MathJax rendering error:', err);
      });
    });
  }
});
