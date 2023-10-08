const $cursor = document.querySelector('#cursor');

// Listeners
document.body.addEventListener('mousemove', onMouseMove);

// Move cursor
function onMouseMove(e) {
  TweenMax.to($cursor, .1, {
    x: e.pageX - 5,
    y: e.pageY - 7
  })
}

