
var intervalId;
let x = 10;
intervalId = setInterval(function (x) {
    console.log(x);
}, 200,x);
sleep(2000).then(clearInterval(intervalId));