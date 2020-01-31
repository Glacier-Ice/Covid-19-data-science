
module.exports = {
  isMobile: () => {
    if (typeof window === 'undefined') {
      return false;
    }
    return window.innerWidth < 800;
  }
}