window.onload = function () {
    var id = getQueryVariable('id');
    setTX();
    search();
    player(id);
    setControl();
    setInterval(timeSet,1000);
    setGeci(id);
}
const setTX = function () {
    let url = "http://ananqiexiafan.icu:3000/user/detail?uid=1777517714";
    fetch(url, {
        method: "post",
        //Authorization: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6MjAyMTIxMjY3MiwiTmFtZSI6IuadqOm5j-WuhyIsIlFRIjoiMjI1OTIwMjA0MCIsImV4cCI6MTY3MzMxNTE3OCwiaWF0IjoxNjczMzExNTc4LCJpc3MiOiJqb3NodWEifQ.t3IXSZ0YJC3oe4KWECqiZbB0-VkrbJbWio5s7a9wVNE",
        //Cookie: "Cookie_1=value; MUSIC_U=bae943925ad3aa0db65486abd1b6e152eb0e6d5c8bc29bc76f9b3c363869b8e81e8907c67206e1ed03fd35b11ab07e7e6054c1a724ad28aedbc4e40dc68540b27895fe1ce8c630f0d4dbf082a8813684; NMTID=00ObHqyrDKrS-JodUEonotNs7xX07YAAAGFrsuYEQ; __csrf=c692f0a5644f297b6d89f2ac5644dd64",
    }).then(res => res.json())
        .then(json => setTx(json));
}
const setTx = function (json) {
    document.querySelector("#Tx").setAttribute('src', json.profile.avatarUrl);
    document.querySelector('#nickname').innerHTML = json.profile.nickname;
    userGeDan(json.profile.userId);
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
    //getGD(json['playlist'][0]['id'])
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
    document.querySelector('#songImg').src = json['songs'][0]['al']['picUrl'];
    document.querySelector('#songName').innerText = json['songs'][0]['name'];
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
    document.querySelector('#singer').innerText = singers;
}
const setControl = function () {
    document.querySelector('#pause').addEventListener('click',pausePlaying,true);
    document.querySelector('#restart').addEventListener('click',restartPlay,true);
    document.querySelector('#volume').addEventListener('click',setVolume,true);
    document.querySelector('#jd').addEventListener('click',setJd,true)
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
    document.querySelector('#jd').value = timePercent*1000;
    document.querySelector('#lTime').innerText = lTime
    document.querySelector('#rTime').innerText = rTime
    document.querySelector('#nowTime').style.width=timePercent*23+"vw";
}
function fix(num, length) {
    return ('' + num).length < length ? ((new Array(length + 1)).join('0') + num).slice(-length) : '' + num;
}
const setVolume = function(){
    document.querySelector('#palyering').volume = document.querySelector('#volume').value/100;
}
let gc_times = [];
let gc_is = [];
const setGeci = function (id) {
    let url = 'http://ananqiexiafan.icu:3000/lyric?id='+id;
    fetch(url,{
        method:'get',
    }).then(res =>res.json())
    .then(json => getGeci(json['lrc']['lyric']));
}
const getGeci = function(gc){
    //let gc = ("[00:00.000] 作词 : 童子-T\n[00:01.000] 作曲 : Shingo.S/童子-T\n[00:15.05]ふたりの思い出\n[00:17.79]かき集めたなら\n[00:20.59]また泣けてきちゃう\n[00:23.50]寂しさ溢れて\n[00:26.28]最後の恋だと信じて願った\n[00:31.99]あの日々にウソはなかった\n[00:41.27]希望夢明るい未来\n[00:44.09]東京に持った大きな期待\n[00:46.83]だけど現実は甘くなくて\n[00:49.54]落ち葉見つめ深く思いつめてた\n[00:52.75]そんな時にあなたと出会って\n[00:55.72]いつもあなたに助けられて\n[00:58.50]バイトが楽しみになって\n[01:01.05]実はシフト被るように狙ってた\n[01:04.22]スタンプ使いが妙に上手くて\n[01:07.02]お化けも虫も受け付けなくて\n[01:09.81]くしゃくしゃの笑顔が可愛くて\n[01:12.66]眠れない夜は君のせいで\n[01:15.62]この気持ち今すぐに伝えたい\n[01:18.29]けどバレたくない\n[01:19.69]どうしたらいいの\n[01:21.30]迷ってるうちに\n[01:22.64]夜明けが来て\n[01:24.06]馬鹿みたいに後悔して\n[01:26.34]ふたりの想い出\n[01:28.71]かき集めたなら\n[01:31.58]また泣けてきちゃう\n[01:34.39]寂しさ溢れて\n[01:37.14]最後の恋だと信じて願った\n[01:42.96]あの日々にウソはなかった\n[02:00.42]帰り道の公園で受けた告白\n[02:03.90]ベタすぎるセリフ笑っちゃった\n[02:06.71]一生忘れられない思い出\n[02:09.56]あなたがプレゼントしてくれた\n[02:12.42]一日中ゲームやりこんで\n[02:15.12]夜ご飯は一緒に作って\n[02:18.02]贅沢なんてしなくたって\n[02:20.60]２人いればそれだけでよくて\n[02:23.76]口下手２人が本気で喧嘩\n[02:26.57]お互いブロック通じない電話\n[02:29.39]本気でぶつかり合えることが\n[02:32.19]どんな愛しいが気づけなかった\n[02:35.11]あなたが教えてくれたこと\n[02:38.24]くれたもの\n[02:39.77]胸に刻み過ごしてる今日も\n[02:42.55]だから伝えたいありがとう\n[02:45.69]ふたりの想い出\n[02:48.24]かき集めたなら\n[02:51.05]また泣けてきちゃう\n[02:53.95]寂しさ溢れて\n[02:56.71]最後の恋だと信じて願った\n[03:02.50]あの日々にウソはなかった\n[03:08.81]子供のままでいられたなら\n[03:14.24]何も怖がらず歩いて行けたかな\n[03:19.98]もっと早く大人になっていたなら\n[03:25.67]２人で乗り越えられたかな\n[03:30.64]今も君の夢夜空へ願う\n[03:37.82]今でも君は\n[03:39.83]あの頃と同じ笑顔で\n[03:43.33]今でも君は\n[03:45.44]あの頃のようにまっすぐで\n[03:48.95]今でも君は\n[03:51.13]あの頃と変わらない優しさで\n[03:54.69]今でも君は\n[03:57.50]君のままでいてほしいそう願うよ\n[04:02.48]ふたりの想い出\n[04:04.95]かき集めたなら\n[04:07.82]また泣けてきちゃう\n[04:10.68]寂しさ溢れて\n[04:13.45]最後の恋だと信じて願った\n[04:19.25]あの日々にウソはなかった\n[04:25.28]ふたりの想い出集めたら\n[04:28.15]泣き出しそうになる今夜も\n[04:30.98]寂しさ溢れて苦しくなる\n[04:36.57]最後の恋と信じ願った\n[04:39.51]あの日々に嘘はなかった\n[04:42.43]離れてもあなたの幸せ願う\n[04:48.00]ふたりの想い出集めたら\n[04:50.83]泣き出しそうになる今夜も\n[04:53.71]寂しさ溢れて苦しくなる\n[04:59.33]最後の恋と信じ願った\n[05:02.24]あの日々に嘘はなかった\n[05:05.12]離れてもあなたの幸せ願う\n");
    let gc_s = gc.split('\n');
    for(let i = 0 ; i < gc_s.length ; i++){
        let gc_time = gc_s[i].slice(2,6);
        gc_times.push(gc_time)
        index = gc_s[i].indexOf(']')
        let gc_i = gc_s[i].slice(index+1,)
        createText(gc_i);
        gc_is.push(gc_i);
    }
    setInterval(setgeci,1000)

}
setgeci = function(){
    for(let i = 0 ; i < gc_times.length ;i++){
        //console.log(document.querySelector('#lTime').innerText+"    "+gc_times[i])
        if(document.querySelector('#lTime').innerText==gc_times[i]){
            for(let j = 0 ; j < gc_times.length ;j++){
                document.querySelector('#text').childNodes[j+1].style.padding = '0px';
                document.querySelector('#text').childNodes[j+1].style.color = 'RGB(200,200,200)';
                document.querySelector('#text').childNodes[j+1].style.fontSize = '20px';
                if(i==j){
                    document.querySelector('#text').childNodes[i+1].style.padding = '15px';
                    document.querySelector('#text').childNodes[i+1].style.color = 'RGB(252,1,26)';
                    document.querySelector('#text').childNodes[i+1].style.fontSize = '32px';
                    }
                }
            document.querySelector('#text').style.top = 50*(4-i)+'px';
            break;
        }
    }
}
move = function() {

}
createText = function (gc_s) {
    let div = document.createElement('div');
    div.innerText = gc_s;
    document.querySelector('#text').appendChild(div);
}

const setJd = function () {
    document.querySelector('#palyering').currentTime = document.querySelector('#jd').value/1000*document.querySelector('#palyering').duration
}