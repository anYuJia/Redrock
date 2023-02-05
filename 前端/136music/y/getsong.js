function getSong(songs){
    let url = "http://ananqiexiafan.icu:3000/playlist/track/all?id=";
    let id = "7884263432";
    url = url + id;

    fetch(url,{
        method:"GET"
    })
    .then(res=>res.json())
    .then(json=>setSong(songs,json))
}
let setSong = function(songs,json){
    console.log(json['songs'].length)
    for(let i = 0 ; i < json['songs'].length;i++)
    {
        song={};
        song.id = json['songs'][i]['al']['id'];
        song.songName = json['songs'][i]['al']['name'];
        song.imageUrl = json['songs'][i]['al']['picUrl'];
        song.singer = json['songs'][i]['ar'][0]['name'];
        songs.push(song)
    }
    //console.log(songs)
}

function sleep (time) {
    return new Promise((resolve) => setTimeout(resolve, time));
}







let song={};
let songs = [];
getSong(songs);
sleep(200).then(() => {
      console.log(songs);
})