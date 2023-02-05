window.onload = function () {
    id = getQueryVariable('key');
    setTX();
    getSongList();
    setControl();
    search();
    setInterval(timeSet,1000);
}
var songs = [];
var playing;
var id = "海阔天空";
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
function getSong(songs) {
    let url = "http://ananqiexiafan.icu:3000/search?keywords=";
    url = url + id;
    fetch(url, {
        method: "GET"
    })
    .then(res => res.json())
    .then(json => setSong(songs, json))
}
let setSong = function (songs, json) {
    //console.log(json['songs'].length)
    for (let i = 0; i < json['result']['songs'].length; i++) {
        let song = {};
        song.id = json['result']['songs'][i]['id'];
        song.songName = json['result']['songs'][i]['name'];
        song.singer = json['result']['songs'][i]['artists'][0]['name'];
        song.zhuanji = json['result']['songs'][i]['album']['name'];
        song.time = json['result']['songs'][i]['duration']
        songs.push(song)
    }
    //console.log(songs)
}
function sleep(time) {
    return new Promise((resolve) => setTimeout(resolve, time));
}
const getSongList = function () {
    //var songs = [];
    getSong(songs);
    sleep(1000).then(() => {
        setSongs(songs)
    })
}
const setSongs = function (songs) {
    table = document.querySelector('#song-list-body');
    let j = 0;
    for (let i = 0; i < songs.length; i++, j++) {
        let tr = document.createElement('tr');
        let song_name = document.createElement('td');
        let singer_name = document.createElement('td');
        let zhuanji = document.createElement('td');
        let time = document.createElement('td');
        let tr_class = document.createAttribute('class');
        tr_class.value = 'song-list-body-body';
        tr.setAttributeNode(tr_class);
        let song_name_class = document.createAttribute('class');
        song_name_class.value = 'song-list-body-body-song-name';
        let song_id = document.createAttribute('id');
        song_id.value = songs[i].id;
        tr.setAttributeNode(song_id);
        let singer_name_class = document.createAttribute('class');
        singer_name_class.value = 'song-list-body-body-singer-name';
        let zhuanji_class = document.createAttribute('class');
        zhuanji_class.value = 'song-list-body-body-zhuanji-name';
        let time_class = document.createAttribute('class');
        time_class.value = 'song-list-body-body-song-time';
        song_name_text = document.createTextNode(songs[i].songName);
        let singer_name_text = document.createTextNode(songs[i].singer);
        let zhuanji_text = document.createTextNode(songs[i].zhuanji);
        let time_text = document.createTextNode(fix(parseInt((songs[i].time/1000)/60),2)+":"+fix(parseInt((songs[i].time/1000)%60),2));        //console.log(songs[i].songName, songs[i].singer, songs[i].zhuanji)
        song_name.setAttributeNode(song_name_class)
        singer_name.setAttributeNode(singer_name_class)
        zhuanji.setAttributeNode(zhuanji_class)
        time.setAttributeNode(time_class)
        song_name.appendChild(song_name_text)
        singer_name.appendChild(singer_name_text)
        zhuanji.appendChild(zhuanji_text)
        time.appendChild(time_text)
        tr.appendChild(song_name)
        tr.appendChild(singer_name)
        tr.appendChild(zhuanji)
        tr.appendChild(time)
        tr.addEventListener('click',play,true)
        table.appendChild(tr)
    }
    document.querySelector('.center').style.height = 88 + (j - 6) * 4 + 'vh';
}

const getQueryVariable = function(variable)
{
    let query = window.location.search.substring(1);
    let vars = query.split("&");
    for (let i=0;i<vars.length;i++) {
        let pair = vars[i].split("=");
        if(pair[0] == variable){return pair[1];}
    }
    return(false);
}

const play = function(e){
    let Id = e.target.parentNode.getAttribute('id');
    player(Id)
}

const playAll = function(){
    player(songs[0]['id'])
}

const player = function(Id){
    let url = "http://ananqiexiafan.icu:3000/song/url?id="+Id;
    //console.log(url)
    fetch(url,{
        method:'get'
    })
    .then(res => res.json())
    .then(json => setPlayer(json['data'][0]['url'],json['data'][0]['id']))
    let url2 = 'http://ananqiexiafan.icu:3000/song/detail?ids='+Id;
    fetch(url2,{
        method:'get',
    })
    .then(res=>res.json())
    .then(json=>setPlayerInfo(json))
}

const setPlayer = function(url,id){
    document.querySelector('#palyering').src = url;
    //console.log(id)
    document.querySelector('#palyering').setAttribute('Ida',id)
}
const setPlayerInfo = function(json){
    document.querySelector('#playingCover').src = json['songs'][0]['al']['picUrl'];
    document.querySelector('#playingSongName').innerText = json['songs'][0]['name'];
    let singers = "";
    for(let i = 0 ; i < json['songs'][0]['ar'].length;i++){
        if(i==0){
            singers = singers+json['songs'][0]['ar'][i]['name']
        }
        else{
            singers = singers+'/'+json['songs'][0]['ar'][i]['name']
        }
    }
    document.querySelector('#playingSinger').innerText = singers;
}
const setControl = function () {
    document.querySelector('#playA').addEventListener('click',detai,true)
    document.querySelector('#pause').addEventListener('click',pausePlaying,true);
    document.querySelector('#restart').addEventListener('click',restartPlay,true);
    document.querySelector('#volume').addEventListener('click',setVolume,true);
    document.querySelector('#next').addEventListener('click',nextPlay,true);
    document.querySelector('#before').addEventListener('click',beforePlay,true);

}
const pausePlaying = function(){
    let audioPlayer = document.querySelector('#palyering');
    if(audioPlayer.paused){
        audioPlayer.play();
    }else{
        audioPlayer.pause();
    }
}
const restartPlay = function(){
    let audioPlayer = document.querySelector('#palyering');
    audioPlayer.load();
}
const timeSet = function(){
    let duration = document.querySelector('#palyering').duration;
    let current = document.querySelector('#palyering').currentTime;
    let rTime = fix(parseInt(duration/60),1)+':'+fix(parseInt(duration%60),2);
    let lTime = fix(parseInt(current/60),1)+':'+fix(parseInt(current%60),2);
    if(rTime=="NaN:NaN"){
        rTime = '0:00'
    }
    let timePercent = current/duration;
    document.querySelector('#lTime').innerText = lTime
    document.querySelector('#rTime').innerText = rTime
    document.querySelector('#nowTime').style.width=timePercent*23+"vw";
    if(timePercent>=1){
        nextPlay();
    }
}
function fix(num, length) {
    return ('' + num).length < length ? ((new Array(length + 1)).join('0') + num).slice(-length) : '' + num;
}
const setVolume = function(){
    document.querySelector('#palyering').volume = document.querySelector('#volume').value/100;
}
const nextPlay = function () {
    let nowId = document.querySelector('#palyering').getAttribute('Ida');
    for(let i = 0 ; i < songs.length ; i++){
        if(songs[i]['id']==nowId){
            if(i == songs.length-1){
                i = -1;
            }
            player(songs[i+1]['id']);
            break;
        }
    }
}
const beforePlay = function () {
    let nowId = document.querySelector('#palyering').getAttribute('Ida');
    for(let i = 0 ; i < songs.length ; i++){
        if(songs[i]['id']==nowId){
            if(i == 0){
                i = songs.length;
            }
            player(songs[i-1]['id']);
            break;
        }
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

const detai = function(){
    ida = document.querySelector("#palyering").getAttribute('Ida');
    document.querySelector('#playA').href = "/play/index.html?id="+ida;
}