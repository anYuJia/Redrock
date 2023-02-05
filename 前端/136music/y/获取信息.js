const setTX = function(){
    url="http://ananqiexiafan.icu:3000/user/account";
    fetch(url,{
        method : "post",
        Authorization: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6MjAyMTIxMjY3MiwiTmFtZSI6IuadqOm5j-WuhyIsIlFRIjoiMjI1OTIwMjA0MCIsImV4cCI6MTY3MzMxNTE3OCwiaWF0IjoxNjczMzExNTc4LCJpc3MiOiJqb3NodWEifQ.t3IXSZ0YJC3oe4KWECqiZbB0-VkrbJbWio5s7a9wVNE",
        Cookie : "Cookie_1=value; MUSIC_U=bae943925ad3aa0db65486abd1b6e152eb0e6d5c8bc29bc76f9b3c363869b8e81e8907c67206e1ed03fd35b11ab07e7e6054c1a724ad28aedbc4e40dc68540b27895fe1ce8c630f0d4dbf082a8813684; NMTID=00ObHqyrDKrS-JodUEonotNs7xX07YAAAGFrsuYEQ; __csrf=c692f0a5644f297b6d89f2ac5644dd64",
    }).then(res => res.json())
    .then(json => setTx(json));
}
const setTx = function(json){
    console.log(json.profile.avatarUrl);
}
setTX();