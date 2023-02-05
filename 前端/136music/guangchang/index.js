window.onload = function () {
    setTX();
    getGD();
    search();
}
const setTX = function () {
    let url = "http://ananqiexiafan.icu:3000/user/detail?uid=1777517714";
    fetch(url, {
        method: "post",
        Authorization: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6MjAyMTIxMjY3MiwiTmFtZSI6IuadqOm5j-WuhyIsIlFRIjoiMjI1OTIwMjA0MCIsImV4cCI6MTY3MzMxNTE3OCwiaWF0IjoxNjczMzExNTc4LCJpc3MiOiJqb3NodWEifQ.t3IXSZ0YJC3oe4KWECqiZbB0-VkrbJbWio5s7a9wVNE",
        Cookie: "Cookie_1=value; MUSIC_U=bae943925ad3aa0db65486abd1b6e152eb0e6d5c8bc29bc76f9b3c363869b8e81e8907c67206e1ed03fd35b11ab07e7e6054c1a724ad28aedbc4e40dc68540b27895fe1ce8c630f0d4dbf082a8813684; NMTID=00ObHqyrDKrS-JodUEonotNs7xX07YAAAGFrsuYEQ; __csrf=c692f0a5644f297b6d89f2ac5644dd64",
    }).then(res => res.json())
        .then(json => setTx(json));
}
const setTx = function (json) {
    document.querySelector("#Tx").setAttribute('src', json.profile.avatarUrl);
    document.querySelector('#nickname').innerHTML = json.profile.nickname;
    userGeDan(json.profile.userId);
}
const getGD = function(){
    let url = "http://ananqiexiafan.icu:3000/top/playlist/highquality?limit=50"
    fetch(url,{
        method:'get',
    })
    .then(res => res.json())
    .then(json => setGD(json))
}
const setGD = function (json) {
    let body = document.querySelectorAll('.recommend-body')[0];
    for(let i = 0 ; i < json['playlists'].length;i++){
        let a = document.createElement('a')
        let a_href = document.createAttribute('href')
        a_href.value = "/gedan/index.html?id="+json['playlists'][i]['id'];
        a.setAttributeNode(a_href);
        let div = document.createElement('div');
        let div_class = document.createAttribute('class')
        let div_id = document.createAttribute('id')
        div_class.value = 'recommend-bodyly';
        div_id.value = json['playlists'][i]['id'];
        div.setAttributeNode(div_class);
        div.setAttributeNode(div_id)
        let span_p = document.createElement('span');
        let span_p_class = document.createAttribute('class')
        span_p_class.value = 'p';
        span_p.setAttributeNode(span_p_class);
        span_p.innerText = json['playlists'][i]['playCount']
        div.appendChild(span_p)
        let img = document.createElement('img');
        let img_src = document.createAttribute('src');
        img_src.value = json['playlists'][i]['coverImgUrl']
        img.setAttributeNode(img_src);
        div.appendChild(img);
        let span_title = document.createElement('span');
        let span_title_class = document.createAttribute('class');
        span_title_class.value = 't';
        span_title.setAttributeNode(span_title_class);
        span_title.innerText = json['playlists'][i]['name']
        div.appendChild(span_title);
        //div.addEventListener('click',goGeDan,true)
        a.appendChild(div);
        body.appendChild(a);
    }
}

const userGeDan = function(uid){
    url = "http://ananqiexiafan.icu:3000/user/playlist?uid="+uid;
    fetch(url,{
        method:'get',
    })
    .then(res=>res.json())
    .then(json=>setUserGD(json))
}
const setUserGD = function (json) {
    document.querySelector('.myBok').href = "/guangchang/index.html";
    document.querySelector('#left-title-a').href = "/index.html";
    document.querySelector('#mylove').href = "/gedan/index.html?id="+json['playlist'][0]['id'];
    getGD(json['playlist'][0]['id'])
    myList = document.querySelector('#myPlaylist');
    for(let i = 1 ; i < json['playlist'].length ; i++){
        let div = document.createElement('div');
        let div_class = document.createAttribute('class');
        div_class.value = 'left-my-love';
        div.setAttributeNode(div_class);
        div.innerHTML = "<svg t='1673514521947' class='icon' viewBox='0 0 1024 1024' version='1.1' xmlns='http://www.w3.org/2000/svg' p-id='1392' width='18' height='18'><path d='M958.708971 97.671507c0.00921-12.554944-5.593392-24.446785-15.255449-32.392758-9.657964-7.956206-22.350031-11.120268-34.594913-8.610098L384.945279 163.701362c-19.360953 3.955078-33.280987 21.072945-33.296337 40.941458l-0.329505 477.604632c-29.119201-16.737196-63.945381-25.944905-101.395318-25.944905-0.056282 0-0.102331 0-0.157589 0-47.749514 0-92.990904 15.038508-127.398552 43.521213-36.777625 30.447453-57.051367 67.602677-57.07695 111.974257-0.055259 88.38398 82.674954 156.37449 184.41615 156.37449 0.055259 0 0.101307 0 0.157589 0 47.759747 0 90.443895-11.95938 124.856659-40.442084 36.782741-30.443359 54.519708-69.159126 54.519708-113.514333l0 0.61603 0 0.309038 0 0.152473 2.845815-407.194002 442.041672-89.829911 0.473791 247.498884c-29.042453-16.597003-63.071478-23.818475-100.354616-23.818475-0.055259 0-0.11154 0-0.167822 0-47.749514 0-92.990904 13.145391-127.398552 41.624003-36.786835 30.442336-57.0616 68.539003-57.087183 112.914676-0.056282 88.38398 82.679048 157.331281 184.434569 157.331281 0.046049 0 0.081864 0 0.12689 0 97.022731 0 176.962388-63.573921 184.16237-146.287761 0.069585-0.984421 0.25378-5.492084 0.112564-16.951067L958.708971 97.671507zM324.453556 863.391069c-19.568684 16.198937-46.053895 21.61632-74.589812 21.61632-0.030699 0-0.066515 0-0.097214 0-54.95052 0-101.354386-31.69077-101.329826-73.255421 0.010233-18.933211 9.520841-35.062564 26.76969-49.341778 19.563568-16.194844 46.039569-22.942525 74.560136-22.942525 0.029676 0 0.066515 0 0.097214 0 54.970986 0 101.390201 33.026184 101.364619 74.600045C351.213014 832.993758 341.707522 849.111854 324.453556 863.391069zM434.70559 333.51391l0.065492-94.66708 440.756399-90.041736-0.065492 94.842065L434.70559 333.51391zM774.160815 770.655991c-0.024559 0-0.050142 0-0.075725 0-54.960753 0-101.370759-33.585933-101.344153-75.1557 0.00921-18.934235 9.520841-34.117029 26.773783-48.39215 19.563568-16.194844 46.039569-21.991874 74.570369-21.991874 0.034792 0 0.070608 0 0.106424 0 52.344159 0 102.06149 27.308973 102.06149 66.25089l0 11.796674C865.85826 742.070956 826.472228 770.655991 774.160815 770.655991z' p-id='1393' fill='#dbdbdb'></path> </svg>"+json['playlist'][i]['name'];
        let a = document.createElement('a');
        let a_href = document.createAttribute('href');
        a_href.value = "/gedan/index.html?id="+json['playlist'][i]['id'];
        a.setAttributeNode(a_href);
        a.appendChild(div);
        myList.appendChild(a)
    }
}

const search = function(){
    document.querySelector('#searchText').onchange=function(){
        document.querySelector('#search').href = "/search/index.html?key="+document.querySelector('#searchText').value;
    }
}