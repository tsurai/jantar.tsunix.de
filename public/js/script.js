window.onload = function() {
  var killswitch = document.getElementById("killswitch");
  var elements = document.getElementsByTagName("input");

  killswitch.onclick = function(e) {
    for(i = 0; i < elements.length; i++) {
      elements[i].checked = false
    }
    killswitch.style.display = "none";
  }
}