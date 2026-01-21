/**
 * MathJax Configuration
 * Enables LaTeX math rendering with chemistry and extended arrow support
 */
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
  options: {
    ignoreHtmlClass: ".*|",
    processHtmlClass: "arithmatex"
  }
};

document$.subscribe(() => {
  if (typeof MathJax !== 'undefined' && MathJax.typesetPromise) {
    MathJax.typesetClear();
    MathJax.texReset();
    MathJax.typesetPromise().catch(err => console.log('MathJax error:', err));
  }
});
