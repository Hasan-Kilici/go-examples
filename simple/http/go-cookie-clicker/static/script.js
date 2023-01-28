function getCookie(cname) {
  let name = cname + "=";
  let decodedCookie = decodeURIComponent(document.cookie);
  let ca = decodedCookie.split(';');
  for(let i = 0; i <ca.length; i++) {
    let c = ca[i];
    while (c.charAt(0) == ' ') {
      c = c.substring(1);
    }
    if (c.indexOf(name) == 0) {
      return c.substring(name.length, c.length);
    }
  }
  return "";
}

function setCookie(cname, cvalue, exdays) {
  const d = new Date();
  d.setTime(d.getTime() + (exdays * 24 * 60 * 60 * 1000));
  let expires = "expires="+d.toUTCString();
  document.cookie = cname + "=" + cvalue + ";" + expires + ";path=/";
}

window.onload = ()=>{
  if(!getCookie("user")){
    let username = prompt("Kullanıcı adı giriniz");
    setCookie("user", username, 365);
    setCookie("cash", 0, 365);
    setCookie("perclick", 1, 365);
    setCookie("persecond", 0, 365);
    setCookie("cursor", 0, 365);
    setCookie("grandpa", 0, 365);
    setCookie("farm", 0, 365);
    setCookie("factory", 0, 365);
  }
}

let clickElement = document.getElementById("click");

function onClick(){
  let cash = Number(getCookie("cash") + getCookie("perclick"));
  setCookie("cash", cash, 365);
}

setInterval(()=>{
  let cash = Number(getCookie("cash") + getCookie("persecond"));
  setCookie("cash", cash, 365);
},1000)
