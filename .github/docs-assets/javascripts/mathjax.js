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
  }
};

document$.subscribe(function() {
  if (typeof MathJax !== 'undefined' && MathJax.startup && MathJax.typesetPromise) {
    MathJax.typesetClear();
    MathJax.texReset();
    MathJax.typesetPromise().catch(function(err) {
      console.log('MathJax typeset error:', err);
    });
  }
});
