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

function getTaskList(){
    let tasks;
    let userId = getCookie("token");
    fetch(`/api/tasks/${userId}`).then(async(res)=>{
        tasks = await res.json();
        for(let i = 0;i < tasks.length;i++){
            document.getElementById("tasks").innerHTML += `
              ${tasks[i].Task}<br>
            `;
        }
    })
}