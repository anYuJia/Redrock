async function setName(obj) {
    url = "http://ananqiexiafan.icu:3000/user/account"
    fetch(url,{})
    .then(res => set(obj,res))
}
set = function(obj,res){
    obj.name = "aa";
    console.log(obj)
}

let person = {
    a:10,
};
function sleep (time) {
    return new Promise((resolve) => setTimeout(resolve, time));
}
setName(person).then(console.log(person));
function sleep (time) {
    return new Promise((resolve) => setTimeout(resolve, time));
}
  sleep(500).then(() => {
      
  })