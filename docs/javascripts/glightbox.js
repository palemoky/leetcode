document.addEventListener("DOMContentLoaded", function () {
  // Find all images within the main article/content area
  var images = document.querySelectorAll(".md-content img, article img");

  images.forEach(function (img) {
    // Skip if it's already in a link or has a skip class
    if (
      img.parentElement.tagName.toLowerCase() === "a" ||
      img.classList.contains("skip-lightbox")
    ) {
      return;
    }

    // Create lightbox wrapper
    var a = document.createElement("a");
    a.href = img.src;
    a.className = "glightbox";
    a.setAttribute("data-type", "image");

    if (img.title || img.alt) {
      a.setAttribute("data-title", img.title || img.alt);
    }

    img.parentNode.insertBefore(a, img);
    a.appendChild(img);
  });

  // Initialize GLightbox
  if (typeof GLightbox !== "undefined") {
    GLightbox({
      touchNavigation: true,
      loop: false,
      zoomable: true,
      draggable: true,
      openEffect: "zoom",
      closeEffect: "zoom",
      slideEffect: "slide",
    });
  }
});
